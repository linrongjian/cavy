package xhttp

import (
	"context"
	"os"
	"testing"

	"github.com/linrongjian/cavy/common/util"
	"github.com/linrongjian/cavy/common/xlog"
)

type HttpGet struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

var ctx = context.Background()

func TestHttpGet(t *testing.T) {
	client := NewClient()
	// test
	_, bs, err := client.Req().Get("http://www.baidu.com").EndBytes(ctx)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug(string(bs))

	//rsp := new(HttpGet)
	//_, err := client.Type(TypeJSON).Get("http://api.igoogle.ink/app/v1/ping").EndStruct(ctx,rsp)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Debug(rsp)
}

func TestHttpUploadFile(t *testing.T) {
	fileContent, err := os.ReadFile("../../logo.png")
	if err != nil {
		xlog.Error(err)
		return
	}
	//xlog.Debug("fileByte：", string(fileContent))

	bm := make(BodyMap)
	bm.SetBodyMap("meta", func(bm BodyMap) {
		bm.Set("filename", "123.jpg").
			Set("sha256", "ad4465asd4fgw5q")
	}).SetFormFile("image", &util.File{Name: "logo.png", Content: fileContent})

	client := NewClient()

	rsp := new(HttpGet)
	_, err = client.Req(TypeMultipartFormData).
		Post("http://localhost:2233/admin/v1/oss/uploadImage").
		SendMultipartBodyMap(bm).
		EndStruct(ctx, rsp)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("%+v", rsp)
}
