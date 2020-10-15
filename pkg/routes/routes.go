/*
	route related
*/

package routes

import (
	"net/http"

	"gitlab.com/safesurfer/minio-file-server/pkg/common"
	"gitlab.com/safesurfer/minio-file-server/pkg/types"
)

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
