/*
	route related
*/

package routes

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	minio "github.com/minio/minio-go/v7"

	common "gitlab.com/safesurfer/minio-file-server/pkg/common"
	fileserverminio "gitlab.com/safesurfer/minio-file-server/pkg/minio"
	templating "gitlab.com/safesurfer/minio-file-server/pkg/templating"
)

// GetOrListObject ...
// returns a list or path depending of the request
func GetOrListObject(minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doneCh := make(chan struct{})
		defer close(doneCh)

		requestPath := r.URL.Path
		if strings.HasSuffix(requestPath, string(filepath.Separator)) {
			tidyRequestPath := requestPath
			if requestPath == "/" {
				tidyRequestPath = ""
			}
			filesList := fileserverminio.List(minioClient, tidyRequestPath)
			files, err := templating.Template(templating.TemplateListing, templating.TemplateListingObject{
				SiteTitle: common.GetAppSiteTitle(),
				Path:      requestPath,
				Items:     filesList,
			})
			if err != nil {
				log.Printf("%#v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("An error occurred with listing objects"))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(files))
			return
		}

		object, objectInfo, err := fileserverminio.Get(minioClient, requestPath)
		if err != nil {
			log.Printf("%#v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An error occurred with retrieving the requested object"))
			return
		}
		log.Println(objectInfo.Key, objectInfo.Size, objectInfo.ContentType)
		w.Header().Set("content-length", fmt.Sprintf("%d", objectInfo.Size))
		w.Header().Set("content-type", objectInfo.ContentType)
		w.Header().Set("accept-ranges", "bytes")
		w.WriteHeader(http.StatusOK)
		w.Write(object)
	}
}

// GetRoot ...
// returns an index with the site title and path
func GetRoot(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	index, err := templating.Template(templating.TemplateIndex, templating.TemplateIndexObject{
		SiteTitle: common.GetAppSiteTitle(),
		Path:      requestPath,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(index))
}

// Healthz ...
// HTTP handler for health checks
func Healthz(minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := "App unhealthy"
		code := http.StatusInternalServerError

		exists, err := fileserverminio.BucketExists(minioClient)
		if err == nil && exists == true {
			response = "App healthy"
			code = http.StatusOK
		} else {
			log.Printf("Bucket '%s' exists: %#v\n", common.GetAppMinioBucket(), err.Error())
		}
		w.WriteHeader(code)
		w.Write([]byte(response))
	}
}
