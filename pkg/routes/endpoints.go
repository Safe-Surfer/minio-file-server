package routes

import (
	"net/http"

	minio "github.com/minio/minio-go/v7"

	"gitlab.com/safesurfer/minio-file-server/pkg/types"
)

func GetEndpoints(endpointPrefix string, minioClient *minio.Client) types.Endpoints {
	return types.Endpoints{
		{
			EndpointPath: "/",
			HandlerFunc: ListObjects(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
		{
			EndpointPath: "/{*.}",
			HandlerFunc: ListObjects(minioClient),
			HTTPMethods: []string{http.MethodGet},
		},
	}
}
