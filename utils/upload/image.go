package upload

import (
	"fmt"
	"gin-blog/utils"
	"gin-blog/utils/logging"
	"gin-blog/utils/logging/file"
	"gin-blog/utils/setting"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RunTimeRootPath + GetImagePath()
}

func CheckImageExt(filename string) bool {
	ext := file.GetExt(filename)
	for _, allowExt := range setting.AppSetting.ImageAllowTypes {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkdir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	return size <= setting.AppSetting.ImageMaxSize
}
