package api

import (
	"encoding/base64"
	"encoding/json"
)

// Cookie 请求头信息
type Cookie struct {
	Token string `json:"Token"`
}

// Encode 编码
func (c *Cookie) Encode() string {
	data, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

// Decode 解码
func (c *Cookie) Decode(str string) bool {
	if str == "" {
		//fmt.Println("-------------11 ")
		return false
	}

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		//fmt.Println("-------------22 ", err)
		return false
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		//fmt.Println("-------------33 ", err)
		return false
	}

	return true
}
