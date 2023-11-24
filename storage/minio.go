package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MioConfig struct {
	Host string
	AccessKey string
	SecretKey string
	Secure bool
}

var (
	MIOClient *minio.Client
)

func CreateMIOClient(config *MioConfig)(*minio.Client, error) {
	return minio.New(config.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.Secure,
	})
}
