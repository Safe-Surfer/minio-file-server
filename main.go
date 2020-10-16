/*
	initialise the API
*/

package main

import (
	"log"

	mfs "gitlab.com/safesurfer/minio-file-server/cmd/miniofileserver"
	common "gitlab.com/safesurfer/minio-file-server/pkg/common"
)

func main() {
	// initialise the app
	log.Printf("Minio-File-Server (%v, %v, %v, %v)\n", common.AppBuildVersion, common.AppBuildHash, common.AppBuildMode, common.AppBuildDate)
	mfs.HandleWebserver()
}
