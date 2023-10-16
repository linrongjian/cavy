package api

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
)

// defaultClient 默认的httpClient
var defaultClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

// HTTPPost 发送httpPost请求(不认证服务端证书)
func HTTPPost(url string, contentType string, body io.Reader) ([]byte, error) {
	resp, err := defaultClient.Post(url, contentType, body)
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp == nil {
		return []byte{}, nil
	}
	data, err := ioutil.ReadAll(resp.Body)

	return data, err
}

// HTTPPostToken 发送httpPost请求(不认证服务端证书)
func HTTPPostToken(url string, playerID int64, contentType string, body io.Reader) ([]byte, error) {
	ck := &Cookie{
		Token: GenToken(playerID),
	}
	//提交请求
	req, err := http.NewRequest("POST", url, body)
	//增加header选项
	req.Header.Add("Session", ck.Encode())
	req.Header.Set("Content-Type", contentType)

	resp, err := defaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp == nil {
		return []byte{}, nil
	}
	data, err := ioutil.ReadAll(resp.Body)

	return data, err
}

// HTTPResponse data数据解析成pb协议
// func HTTPResponse(data []byte) *protobuf.HTTPResponse {
// 	result := &protobuf.HTTPResponse{
// 		Result: proto.Int32(int32(gerrors.ParseErr)),
// 	}
// 	err := proto.Unmarshal(data, result)
// 	if err != nil {
// 		result.Msg = proto.String(fmt.Sprintf("protobuf解析失败,err:%v", err.Error()))
// 	}
// 	return result

// }
