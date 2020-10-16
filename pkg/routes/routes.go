/*
	route related
*/

package routes

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	minio "github.com/minio/minio-go/v7"

	"gitlab.com/safesurfer/minio-file-server/pkg/common"
)

func GetOrListObject(minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doneCh := make(chan struct{})
		defer close(doneCh)

		if strings.HasSuffix(r.URL.Path, "/") {
			filesList := []string{}
			for message := range minioClient.ListObjects(context.TODO(), common.GetAppMinioBucket(), minio.ListObjectsOptions{Recursive: false, Prefix: r.URL.Path}) {
				filesList = append(filesList, message.Key)
			}
			files := ""
			for _, fileName := range filesList {
				files += fileName + "\n"
			}
			w.WriteHeader(200)
			w.Write([]byte(files))
			return

		}

		object, err := minioClient.GetObject(context.TODO(), common.GetAppMinioBucket(), r.URL.Path, minio.GetObjectOptions{})
		if err != nil {
			log.Printf("%#v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer object.Close()
		objectBytes, err := ioutil.ReadAll(object)
		w.WriteHeader(http.StatusOK)
		w.Write(objectBytes)
	}
}
