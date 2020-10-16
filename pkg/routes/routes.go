/*
	route related
*/

package routes

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	minio "github.com/minio/minio-go/v7"

	common "gitlab.com/safesurfer/minio-file-server/pkg/common"
	fileserverminio "gitlab.com/safesurfer/minio-file-server/pkg/minio"
	templating "gitlab.com/safesurfer/minio-file-server/pkg/templating"
)

func GetOrListObject(minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doneCh := make(chan struct{})
		defer close(doneCh)

		requestPath := r.URL.Path
		if strings.HasSuffix(requestPath, string(filepath.Separator)) {
			if requestPath == "/" {
				requestPath = ""
			}
			filesList := fileserverminio.List(minioClient, requestPath)
			err, files := templating.Template(templating.TemplateListing, templating.TemplateListingObject{
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

		err, object := fileserverminio.Get(minioClient, requestPath)
		if err != nil {
			log.Printf("%#v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("An error occurred with retrieving the requested object"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(object)
	}
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	err, index := templating.Template(templating.TemplateIndex, templating.TemplateIndexObject{
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

		err, exists := fileserverminio.BucketExists(minioClient)
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
