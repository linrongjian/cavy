package stateless

// Redis Key
const (
	// ConfigKey 配置哈希表
	ConfigKey = "config"
	// PlayerKey 用户哈希表
	PlayerKey = "player"
	// ServerKey 服务器信息哈希表
	ServerKey = "server:"
	// RouteKey 路由哈希表
	RouteKey = "route"
	// NginxListKey nginx地址列表
	NginxListKey = "nginx"
	// WorldRank 世界排名
	WorldRank = "worldrank"
	// Online 在线集合表
	Online = "online"
	// Announce 公告记录
	Announce = "announce"
	// AccountKey 账号表Key
	AccountKey = "account"
	// AccountIDKey 账号表Key-id索引
	AccountIDKey = "accountid"
	// AccountUnionIDKey 账号表Key-unionid索引
	AccountUnionIDKey = "accountunionid"
	// AccountOpenidBindKey 账号表openid-account索引
	AccountOpenidBindKey = "accountbindopenid"
	// StatisticsKey 统计key
	StatisticsKey = "statistics"
	// BlackListBak
	BlackListBak = "blacklistbak"
	// BlackList 黑名单
	BlackList = "blacklist:"
	// RequestNum 请求次数
	RequestNum = "requestnum:"
	// RechargeSwtich 支付开关
	RechargeSwtich = "rechargeswtich:"
	// SyncDatabase 同步数据
	SyncDatabase = "syncDatabase"
	// ChannelKey 渠道
	ChannelKey = "channel"
	// AccountOpenidBindQQ 账号表绑定QQ
	AccountOpenidBindQQ = "accountbindQQ"
	// ListUserNic 玩家名字
	SetUserNicK = "userName"
)

// ServerHashKey 服务器信息哈希表key
func ServerHashKey(name string) string {
	return ServerKey + name
}
