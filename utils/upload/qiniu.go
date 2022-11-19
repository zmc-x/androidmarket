package upload

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"path"
)

// 文件上传
func UploadFile(file *multipart.FileHeader) (string, error) {
	accessKey := "xQxQQsSD3WPyOzXltihXa6SKX-LcBr5HgUU6FM46"
	secretKey := "L650S5c9iJ8pUJtU44ciHcKbUSFAgZ5Fax6NVhzV"
	bucket := "hellozmc"
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	f, openerr := file.Open()
	if openerr != nil {
		return "file open failed", openerr
	}
	// 关闭文件
	defer f.Close()
	// 文件命名
	// 通过hash函数来解决上传的冲突问题
	filekey := fmt.Sprintf("androidmarket/%x%s", md5.Sum([]byte(file.Filename)), path.Ext(file.Filename))
	err := formUploader.Put(context.Background(), &ret, upToken, filekey, f, file.Size, &putExtra)
	if err != nil {
		return "upload file failed", err
	}
	return filekey, nil
}
