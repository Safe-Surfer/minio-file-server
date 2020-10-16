/*
	route related
*/

package routes

import (
	"net/http"
	"strings"
	"path/filepath"

	minio "github.com/minio/minio-go/v7"

	fileserverminio "gitlab.com/safesurfer/minio-file-server/pkg/minio"
)

func GetOrListObject(minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doneCh := make(chan struct{})
		defer close(doneCh)

		requestPath := r.URL.Path
		if strings.HasSuffix(requestPath, string(filepath.Separator)) {
			filesList := fileserverminio.List(minioClient, requestPath)
			files := ""
			for _, fileName := range filesList {
				files += fileName + "\n"
			}
			w.WriteHeader(200)
			w.Write([]byte(files))
			return

		}

		err, object := fileserverminio.Get(minioClient, requestPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An error occurred"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(object)
	}
}
