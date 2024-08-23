package model

import (
	"io"
	"mime/multipart"
	"pichub/utils"
	"pichub/utils/errmsg"

	"gorm.io/gorm"
)

type ImageInfo struct {
	gorm.Model
	MD5  string `gorm:"type:varchat(50);unique"`
	Type string `gorm:"type:varchat(20)"`
	Size int64  `gorm:"type:varchat(20)"`
	Data []byte `gorm:"type:blob"`
}

func UploadImage(file *multipart.FileHeader) (int, string) {
	src, _ := file.Open()
	defer src.Close()
	// 计算MD5
	md5 := utils.MathMD5(src)
	// 获取二进制数据
	src.Seek(0, io.SeekStart)
	fileBytes, _ := io.ReadAll(src)
	imageinfo := ImageInfo{
		MD5:  md5,
		Type: file.Header.Get("Content-Type"),
		Data: fileBytes,
		Size: file.Size,
	}
	if err := db.Create(&imageinfo).Error; err != nil {
		return errmsg.ERROR_IMAGE_EXIST, imageinfo.MD5
	}
	return errmsg.SUCCESS, imageinfo.MD5
}
func GetImageData(md5 string) ImageInfo {
	var imageinfo ImageInfo
	_ = db.First(&imageinfo, "md5 = ?", md5).Error
	return imageinfo
}
