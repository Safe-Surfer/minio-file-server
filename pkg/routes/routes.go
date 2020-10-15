/*
	route related
*/

package routes

import (
	"net/http"
	"context"
	"fmt"

	minio "github.com/minio/minio-go/v7"

	"gitlab.com/safesurfer/minio-file-server/pkg/common"
	"gitlab.com/safesurfer/minio-file-server/pkg/types"
)

func ListObjects(minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doneCh := make(chan struct{})
		defer close(doneCh)
		for message := range minioClient.ListObjects(context.TODO(), common.GetAppMinioBucket(), minio.ListObjectsOptions{Recursive: true, Prefix: r.URL.Path}) {
			fmt.Println(message.Key)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
	}
}

func APIroot(w http.ResponseWriter, r *http.Request) {
	// root of API
	JSONresp := types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "Hit root of webserver",
		},
	}
	common.JSONResponse(r, w, 200, JSONresp)
}

func APIUnknownEndpoint(w http.ResponseWriter, r *http.Request) {
	common.JSONResponse(r, w, 404, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "This endpoint doesn't seem to exist.",
		},
	})
}
