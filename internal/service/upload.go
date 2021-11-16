package service

import (
	"Blog/global"
	"Blog/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename) //获取文件名
	if !upload.CheckContainExt(fileType, fileName) {    //判断文件类型是否合法
		return nil, errors.New("file suffix is not supported")
	}
	if !upload.CheckMaxSize(fileType, file) { //检查文件大小
		return nil, errors.New("exceeded maximum file limit")
	}
	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) { //检查保存路径是否存在
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil { //创建保存路径
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) { //检查权限
		return nil, errors.New("insufficient file permissions")
	}
	dst := uploadSavePath + "/" + fileName //加密文件名
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
