package service

import (
	"errors"
	"github.com/catherine.li/go_blog/global"
	"github.com/catherine.li/go_blog/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	//在UploadFile Service方法中，首先获取文件所需的基本信息
	//接着对文件进行业务所检查（文件大小是否符合需求、文件后缀是否达到要求）
	//并且判断其是否具备写入条件（目录是否存在、权限是否足够），最后进行真正的写入文件操作。

	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
