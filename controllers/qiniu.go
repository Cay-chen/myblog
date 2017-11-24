package controllers

import (
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"errors"
	"context"
)
// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}
func resQiNiu(localFile,fileName,bucketName string) error{
	ak := "U1Z0InLIRWomOHQ3RHBmeBLBeT2LsBsCaTZbLcRC"
	sk :="aXK10xN-DgKVOZNuk4yzTKTjtbL7WtrhtSKyOMWX"
	bucket:=bucketName
	key := fileName
	mac := qbox.NewMac(ak,sk)
	putPolicy := storage.PutPolicy{
		Scope:bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires=7200
	upToken :=putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "myblog content image",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if ret.Bucket ==bucket && fileName==ret.Key{
		return nil
	}else {
		return errors.New("七牛上传返回数据错误")
	}
}