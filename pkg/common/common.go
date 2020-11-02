/*
	common function calls
*/

package common

import (
	"net/http"
	"os"
	"log"
)

// AppBuild metadata
const (
	AppBuildVersion = "0.0.0"
	AppBuildHash    = "???"
	AppBuildDate    = "???"
	AppBuildMode    = "???"
)

// GetAppSiteTitle ...
// site title to template in
func GetAppSiteTitle() (output string) {
	return GetEnvOrDefault("APP_SITE_TITLE", "Minio-File-Server")
}

// GetAppEnvFile ...
// location of an env file to load
func GetAppEnvFile() (output string) {
	return GetEnvOrDefault("APP_ENV_FILE", ".env")
}

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

// GetAppHealthPortEnabled ...
// enable the binding of a health port
func GetAppHealthPortEnabled() (output string) {
	return GetEnvOrDefault("APP_HEALTH_PORT_ENABLED", "true")
}

// GetAppHealthPort ...
// the port to bind the health service to
func GetAppHealthPort() (output string) {
	return GetEnvOrDefault("APP_HEALTH_PORT", ":8081")
}

// GetAppPort ...
// the port to serve web traffic on
func GetAppPort() (output string) {
	return GetEnvOrDefault("APP_PORT", ":8080")
}

// GetAppMetricsPort ...
// return the port which the app should serve metrics on
func GetAppMetricsPort() (output string) {
	return GetEnvOrDefault("APP_PORT_METRICS", ":2112")
}

// GetAppMetricsEnabled ...
// serve metrics endpoint
func GetAppMetricsEnabled() (output string) {
	return GetEnvOrDefault("APP_METRICS_ENABLED", "true")
}

// GetAppRealIPHeader ...
// the header to use instead of r.RemoteAddr
func GetAppRealIPHeader() (output string) {
	return GetEnvOrDefault("APP_HTTP_REAL_IP_HEADER", "")
}

// GetEnvOrDefault ...
// given an env var return it's value, else return a default
func GetEnvOrDefault(envName string, defaultValue string) (output string) {
	output = os.Getenv(envName)
	if output == "" {
		output = defaultValue
	}
	return output
}

// GetRequestIP ...
// returns r.RemoteAddr unless RealIPHeader is set
func GetRequestIP(r *http.Request) (requestIP string) {
	realIPHeader := GetAppRealIPHeader()
	headerValue := r.Header.Get(realIPHeader)
	if realIPHeader == "" || headerValue == "" {
		return r.RemoteAddr
	}
	return headerValue
}

// Logging ...
// a basic middleware for logging
func Logging(next http.Handler) http.Handler {
	// log all requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestIP := GetRequestIP(r)
		log.Printf("%v %v %v %v %v %v %#v", r.Method, r.URL, r.Proto, r.Response, requestIP, r.RemoteAddr, r.Header)
		next.ServeHTTP(w, r)
	})
}
