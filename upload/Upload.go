package upload

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"mime/multipart"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// AccessKey 公钥
var AccessKey = utils.AccessKey

// SecretKey 私钥
var SecretKey = utils.SecretKey

// Bucket Bucket名
var Bucket = utils.Bucket

// ImgURL 地址
var ImgURL = utils.QiniuServer

// File 上传文件
func File(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}

	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.Error
	}

	url := ImgURL + ret.Key
	return url, errmsg.Success
}
