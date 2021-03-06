package miniofileserver

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gitlab.com/safesurfer/minio-file-server/pkg/common"
	"gitlab.com/safesurfer/minio-file-server/pkg/minio"
	"gitlab.com/safesurfer/minio-file-server/pkg/routes"
	"gitlab.com/safesurfer/minio-file-server/pkg/handlers"
	"gitlab.com/safesurfer/minio-file-server/pkg/metrics"
)

// HandleWebserver ...
// manages app initialisation
func HandleWebserver() {
	// bring up the API

	envFile := common.GetAppEnvFile()
	_ = godotenv.Load(envFile)

	go metrics.Handle()

	port := common.GetAppPort()
	minioHost := common.GetAppMinioHost()
	minioAccessKey := common.GetAppMinioAccessKey()
	minioSecretKey := common.GetAppMinioSecretKey()
	minioUseSSL := common.GetAppMinioUseSSL()
	minioUseSSLBool, err := strconv.ParseBool(minioUseSSL)
	minioClient, err := minio.Open(minioHost, minioAccessKey, minioSecretKey, minioUseSSLBool)
	if err != nil {
		panic(err)
	}

	go handlers.HealthHandler(minioClient)

	router := mux.NewRouter().StrictSlash(false)
	for _, endpoint := range routes.GetEndpoints("/", minioClient) {
		router.HandleFunc(endpoint.EndpointPath, endpoint.HandlerFunc).Methods(endpoint.HTTPMethods...)
	}

	router.Use(common.Logging)

	srv := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on", port)
	log.Fatal(srv.ListenAndServe())
}
