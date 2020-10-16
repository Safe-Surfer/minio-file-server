package minio

import (
	"context"
	"io/ioutil"
	"log"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"gitlab.com/safesurfer/minio-file-server/pkg/common"
)

// Open ...
// open a Minio client
func Open(endpoint string, accessKey string, secretKey string, useSSL bool) (*minio.Client, error) {
	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
}

func Get(minioClient *minio.Client, filePath string) (err error, objectBytes []byte) {
	object, err := minioClient.GetObject(context.TODO(), common.GetAppMinioBucket(), filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("%#v\n", err)
		return err, []byte{}
	}
	defer object.Close()
	objectBytes, err = ioutil.ReadAll(object)
	if err != nil {
		log.Printf("%#v\n", err)
		return err, []byte{}
	}
	return err, objectBytes
}

func List(minioClient *minio.Client, path string) []minio.ObjectInfo {
	filesList := []minio.ObjectInfo{}
	for message := range minioClient.ListObjects(context.TODO(), common.GetAppMinioBucket(), minio.ListObjectsOptions{Recursive: false, Prefix: path}) {
		filesList = append(filesList, message)
	}
	return filesList
}

func ListBuckets(minioClient *minio.Client) (err error, bucketList []minio.BucketInfo) {
	bucketList, err = minioClient.ListBuckets(context.TODO())
	return err, bucketList
}

func BucketExists(minioClient *minio.Client) (err error, exists bool) {
	exists, err = minioClient.BucketExists(context.TODO(), common.GetAppMinioBucket())
	return err, exists
}
