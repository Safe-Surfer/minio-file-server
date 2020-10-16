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

// Get ...
// retrieves a given object
func Get(minioClient *minio.Client, filePath string) (objectBytes []byte, err error) {
	object, err := minioClient.GetObject(context.TODO(), common.GetAppMinioBucket(), filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("%#v\n", err)
		return []byte{}, err
	}
	defer object.Close()
	objectBytes, err = ioutil.ReadAll(object)
	if err != nil {
		log.Printf("%#v\n", err)
		return []byte{}, err
	}
	return objectBytes, err
}

// List ...
// returns a list of files in a given path
func List(minioClient *minio.Client, path string) []minio.ObjectInfo {
	filesList := []minio.ObjectInfo{}
	for message := range minioClient.ListObjects(context.TODO(), common.GetAppMinioBucket(), minio.ListObjectsOptions{Recursive: false, Prefix: path}) {
		filesList = append(filesList, message)
	}
	return filesList
}

// ListBuckets ...
// returns a list of the available buckets
func ListBuckets(minioClient *minio.Client) ([]minio.BucketInfo, error) {
	return minioClient.ListBuckets(context.TODO())
}

// BucketExists ...
// returns if a bucket exists
func BucketExists(minioClient *minio.Client) (exists bool, err error) {
	exists, err = minioClient.BucketExists(context.TODO(), common.GetAppMinioBucket())
	return exists, err
}
