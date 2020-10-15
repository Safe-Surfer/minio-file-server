/*
	common function calls
*/

package common

import (
	"net/http"
	"time"
	"encoding/json"
	"os"
	"log"

	"gitlab.com/safesurfer/minio-file-server/pkg/types"
)

const (
	AppBuildVersion = "0.0.0"
)

// GetAppMinioAccessKey ...
// return the accessKey for file storage
func GetAppMinioAccessKey() (output string) {
	return GetEnvOrDefault("APP_MINIO_ACCESS_KEY", "")
}

// GetAppMinioSecretKey ...
// return the secretKey for file storage
func GetAppMinioSecretKey() (output string) {
	return GetEnvOrDefault("APP_MINIO_SECRET_KEY", "")
}

// GetAppMinioBucket ...
// return the bucket for file storage
func GetAppMinioBucket() (output string) {
	return GetEnvOrDefault("APP_MINIO_BUCKET", "")
}

// GetAppMinioHost ...
// return the host for file storage
func GetAppMinioHost() (output string) {
	return GetEnvOrDefault("APP_MINIO_HOST", "")
}

// GetAppMinioUseSSL ...
// return if the file storage should use SSL
func GetAppMinioUseSSL() (output string) {
	return GetEnvOrDefault("APP_MINIO_USE_SSL", "")
}

func GetEnvOrDefault(envName string, defaultValue string) (output string) {
	output = os.Getenv(envName)
	if output == "" {
		output = defaultValue
	}
	return output
}

func GetAppPort() (output string) {
	return GetEnvOrDefault("APP_PORT", ":8080")
}

func Logging(next http.Handler) http.Handler {
	// log all requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v %v %v", r.Method, r.URL, r.Proto, r.Response, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func GetAppDistFolder() string {
	appDistFolder := GetEnvOrDefault("APP_DIST_FOLDER", "./dist")
	return appDistFolder
}

func JSONResponse(r *http.Request, w http.ResponseWriter, code int, output types.JSONMessageResponse) {
	// simpilify sending a JSON response
	output.Metadata.URL = r.RequestURI
	output.Metadata.Timestamp = time.Now().Unix()
	output.Metadata.Version = AppBuildVersion
	response, _ := json.Marshal(output)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
