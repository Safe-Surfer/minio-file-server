package handlers

import (
	"log"
	"net/http"

	minio "github.com/minio/minio-go/v7"

	common "gitlab.com/safesurfer/minio-file-server/pkg/common"
	routes "gitlab.com/safesurfer/minio-file-server/pkg/routes"
)

// HealthHandler ...
func HealthHandler(minioClient *minio.Client) {
	if common.GetAppHealthPortEnabled() != "true" {
		return
	}

	port := common.GetAppHealthPort()
	http.Handle("/", routes.Healthz(minioClient))
	log.Printf("Health listening on %v", port)
	http.ListenAndServe(port, nil)
}
