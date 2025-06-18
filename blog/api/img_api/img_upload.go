package img_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strings"
)

var (
	WhitelistSuffix = []string{
		"jpg", "png", "gif", "psd", "tif", "bmp",
	}
)

// 图片格式白名单
type FileUploadResponse struct {
	Filename  string `json:"filename"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// 上传图片，返回图片uri
func (ImgApi) ImgUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fileList, ok := form.File["image"]
	if !ok {
		res.FailWithMsg("不存在文件", c)
		return
	}
	//判断路径是否存在
	//不存在则创建
	defaultPath := global.Config.Upload.Path
	_, err = os.ReadDir(defaultPath)
	if err != nil {
		err = os.Mkdir(defaultPath, os.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []FileUploadResponse
	for _, file := range fileList {
		//文件后缀判断
		filename := file.Filename
		filesuffix := strings.Replace(path.Ext(filename), ".", "", -1)
		flag := false
		lenSuffix := len(WhitelistSuffix)
		for key, suffix := range WhitelistSuffix {
			if strings.EqualFold(suffix, filesuffix) {
				flag = true
			}
			if key+1 == lenSuffix && !flag {
				resList = append(resList, FileUploadResponse{
					file.Filename,
					false,
					fmt.Sprintf("当前图片格式%s，需要上传图片格式为%s", filesuffix, WhitelistSuffix),
				})
			}

		}
		if !flag {
			continue
		}

		filePath := path.Join(defaultPath, file.Filename)
		//判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				file.Filename,
				false,
				fmt.Sprintf("当前图片大小：%.2fMB,图片过大，需上传%dMB大小的图片", size, global.Config.Upload.Size),
			})
			continue
		}

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imgHash := utils.Md5(byteData)
		//图片重复
		var bannerModel models.BannerModel

		err = global.DB.Take(&bannerModel, "hash = ?", imgHash).Error
		fmt.Println(err)
		if err == nil {
			resList = append(resList, FileUploadResponse{
				bannerModel.Path,
				false,
				fmt.Sprintf("图片已存在"),
			})
			continue
		}

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			resList = append(resList, FileUploadResponse{
				file.Filename,
				false,
				err.Error(),
			})
		}
		resList = append(resList, FileUploadResponse{
			filePath,
			true,
			fmt.Sprintf("图片上传成功"),
		})
		//入库
		fmt.Println(filePath, imgHash, filename)
		global.DB.Create(&models.BannerModel{
			Path: filePath,
			Hash: imgHash,
			Name: filename,
		})
	}
	res.OkWithData(resList, c)
}
