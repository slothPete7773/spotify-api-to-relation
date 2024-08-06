package repository

type StorageRepository interface {
	UploadFile(srcPath string, objectName string) error
}
