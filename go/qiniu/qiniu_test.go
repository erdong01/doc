package qiniu_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func TestXxx(t *testing.T) {
	bucket := "fqykt"
	key := "script/test.js"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac("VzE24dWaAOmPCYw4tHW7umZQsogg9GtYjGlcGDX1", "AguEr0-DFOZZNWVwzP_wAjZ3nbF34W0Y0vKkyW1V")
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneXinjiapo
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	cfg.CentralRsHost = "rxvob9d3a.sabkt.gdipper.com"
	// deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	// privateAccessURL := storage.MakePrivateURL(mac, "rxvob9d3a.sabkt.gdipper.com", key, deadline)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
