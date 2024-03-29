package mongodb

import (
	"context"
	"strings"
	"time"

	"github.com/linrongjian/cavy/core/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB 数据存储
type Store struct {
	opts   *Options
	ctx    context.Context
	cancel context.CancelFunc
	client *mongo.Client
}

func (s *Store) Init(opts ...Option) {
	for _, o := range opts {
		o(s.opts)
	}
}

func (s *Store) DbName() string {
	return s.opts.DbName
}

func (s *Store) Opts() *Options {
	return s.opts
}

func (s *Store) Connect() error {
	if s.client != nil {
		return nil
	}

	if !strings.HasPrefix(s.opts.RawUrl, "mongodb://") {
		s.opts.RawUrl = "mongodb://" + s.opts.RawUrl
	}
	logger.Infof("mongo url:%s", s.opts.RawUrl)

	opts := options.Client().
		SetMinPoolSize(s.opts.MinPoolSize).
		SetMaxPoolSize(s.opts.MaxPoolSize).
		SetConnectTimeout(s.opts.ConnectTimeout).
		SetSocketTimeout(s.opts.SocketTimeout).
		SetMaxConnIdleTime(s.opts.MaxConnIdleTime).
		SetRetryWrites(true).
		SetRetryReads(true).
		ApplyURI(s.opts.RawUrl)

	//if opts.ReplicaSet == nil || *opts.ReplicaSet == "" {
	//	return errors.Server("this system only supports replica sets. example: mongodb://0.0.0.0:27017,0.0.0.0:27018,0.0.0.0:27019/?replicaSet=rs1")
	//}

	s.ctx, s.cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer s.cancel()

	if mc, err := mongo.Connect(s.ctx, opts); err != nil {
		return err
	} else {
		s.client = mc
	}

	if err := s.client.Ping(s.ctx, readpref.Primary()); err != nil {
		return err
	}

	logger.Debug("Store [mongodb] Connect to %s", s.opts.RawUrl)

	return nil
}

func (s *Store) Disconnect() error {
	if s.client == nil {
		return nil
	}

	if err := s.client.Disconnect(s.ctx); err != nil {
		return err
	}

	s.client = nil

	return nil
}

// Client 获取客户端
func (s *Store) Client() *mongo.Client {
	return s.client
}

// Database 获取数据库对象
func (s *Store) D(dbname ...string) *mongo.Database {
	if len(dbname) > 0 && dbname[0] != "" {
		return s.client.Database(dbname[0])
	}
	return s.client.Database(s.opts.DbName)
}

// Collection 获取集合对象
func (s *Store) C(name string, dbname ...string) *mongo.Collection {
	if len(dbname) > 0 && dbname[0] != "" {
		return s.client.Database(dbname[0]).Collection(name)
	}
	return s.client.Database(s.opts.DbName).Collection(name)
}

// CloneCollection 克隆集合对象
func (s *Store) CloneC(name string, dbname ...string) (*mongo.Collection, error) {
	return s.C(name, dbname...).Clone()
}

func (s *Store) GetIncId(id string) (int64, error) {
	return GetIncId(context.Background(), s.D(), id)
}

// 获取集合列表
func (s *Store) ListCollectionNames(dbname ...string) ([]string, error) {
	return s.D(dbname...).ListCollectionNames(context.Background(), bson.M{})
}

func (s *Store) Scan(dbName, tabName string, cur, size int64, filter interface{}, result interface{}, fn ...func(opts *options.FindOptions) *options.FindOptions) *Scan {
	var scan *Scan
	_ = s.Client().UseSession(context.Background(), func(sctx mongo.SessionContext) error {
		col := sctx.Client().Database(dbName).Collection(tabName)

		count, _ := col.CountDocuments(sctx, filter)
		scan = NewScan(count, cur, size)

		if count > 0 {
			opts := scan.FindOptions()
			if len(fn) > 0 {
				opts = fn[0](opts)
			}
			cur, err := col.Find(sctx, filter, opts)
			if err != nil {
				return err
			}
			if err := cur.All(nil, result); err != nil {
				return err
			}
		}

		return nil
	})
	return scan
}

func NewStore(opts ...Option) *Store {
	ms := &Store{
		opts: newOptions(opts...),
	}
	return ms
}
