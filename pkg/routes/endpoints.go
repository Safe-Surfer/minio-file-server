package routes

import (
	"net/http"

	minio "github.com/minio/minio-go/v7"

	"gitlab.com/safesurfer/minio-file-server/pkg/types"
)

// GetEndpoints ...
// returns array of endpoints
func GetEndpoints(endpointPrefix string, minioClient *minio.Client) types.Endpoints {
	return types.Endpoints{
		{
			EndpointPath: "/{*.}/{*.}/{*.}/{*.}",
			HandlerFunc: GetOrListObject(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
		{
			EndpointPath: "/{*.}/{*.}/{*.}/",
			HandlerFunc: GetOrListObject(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
		{
			EndpointPath: "/{*.}/{*.}/{*.}",
			HandlerFunc: GetOrListObject(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
		{
			EndpointPath: "/{*.}/{*.}/",
			HandlerFunc: GetOrListObject(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
		{
			EndpointPath: "/{*.}/{*.}",
			HandlerFunc: GetOrListObject(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
		{
			EndpointPath: "/{*.}/",
			HandlerFunc: GetOrListObject(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
		{
			EndpointPath: "/{*.}",
			HandlerFunc: GetOrListObject(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
		{
			EndpointPath: "/",
			HandlerFunc: GetOrListObject(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
	}
}
