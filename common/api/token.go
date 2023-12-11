package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"cavy/common/mlog"
)

const (
	// myTimeExpired Token过期时间
	myTimeExpired = (60 * 60 * 24)
	// UpdateTokenTime 更新token时间
	UpdateTokenTime = (60 * 60 * 24 / 2)
)

var (
	// myKey token Key
	mykey = []byte("@#$yymmxxkkyzilm")
)

// Token 结构
type Token struct {
	PlayerID   int64
	CreateTime int64
}

func (t *Token) encode() string {
	return fmt.Sprintf("%d@%d", t.PlayerID, t.CreateTime)
}

func (t *Token) decode(data string) error {
	str := strings.Split(data, "@")
	var err error
	t.PlayerID, err = strconv.ParseInt(str[0], 10, 64)
	if err != nil {
		return err
	}
	t.CreateTime, err = strconv.ParseInt(str[1], 10, 64)
	if err != nil {
		return err
	}
	return nil
}

// IsExpire 是否过期
func (t *Token) IsExpire() bool {
	if t.RemainingTime() < 0 {
		log := mlog.NewLogger(mlog.Fields{"PlayerID": t.PlayerID, "CreateTime": t.CreateTime})
		log.Warn("Token Is Expire!!!")
		return true
	}
	return false
}

// RemainingTime 剩余过期时间单位秒
func (t *Token) RemainingTime() int64 {
	now := time.Now().Unix()
	pastTime := now - t.CreateTime
	return myTimeExpired - pastTime
}

// GenToken 生成token
func GenToken(playerID int64) string {
	token := Token{
		PlayerID:   playerID,
		CreateTime: time.Now().Unix(),
	}
	log.Println(token)
	return encrypt(mykey, token.encode())
}

// ParseToken 解析Token
func ParseToken(token string) (*Token, error) {
	plainToken, err := decrypt(mykey, token)
	if err != nil {
		return nil, err
	}

	info := &Token{}
	if err := info.decode(plainToken); err != nil {
		return nil, err
	}

	if info.PlayerID == 0 {
		return nil, fmt.Errorf("Parse Player Is 0")
	}

	return info, nil
}
