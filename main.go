/*
	initialise the API
*/

package main

import (
	mfs "gitlab.com/safesurfer/minio-file-server/cmd/miniofileserver"
)

func main() {
	// initialise the app
	mfs.HandleWebserver()
}
