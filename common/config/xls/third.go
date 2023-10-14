package xls

// ThirdConfig 游戏渠道配置 对应 Third.xlsx
type ThirdConfig struct {
	GameId         int    `xlsx:"column:0"` // 游戏ID
	Third          string `xlsx:"column:1"` // 第三方渠道ID
	AppId          string `xlsx:"column:2"` // AppId
	AppKey         string `xlsx:"column:3"` // AppKey
	LoginUrl       string `xlsx:"column:4"` // 登录地址
	PayCallBackUrl string `xlsx:"column:5"` // 支付回调地址
	Token          string `xlsx:"column:6"` // token

}

func (t *ThirdConfig) Head() int {
	return 2
}

func (t *ThirdConfig) Obj() interface{} {
	return t
}

func (t *ThirdConfig) Data() interface{} {
	return &Config.Third
}
