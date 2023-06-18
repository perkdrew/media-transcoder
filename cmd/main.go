package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/perkdrew/media-transcoder/internal/transcoding"
	"github.com/perkdrew/media-transcoder/internal/handlers"
	"github.com/perkdrew/media-transcoder/internal/models"
)


func main() {
	// Create a new instance of the CmdExecutor
	cmdExecutor := models.RealCmdExecutor{}

	// Create a new instance of the Transcoder
	transcoder := transcoding.NewTranscoder("/path/to/ffmpeg", cmdExecutor) 

	// Create a new instance of the TranscodingHandler
	transcodingHandler := handlers.NewTranscodingHandler(transcoder, nil) 

	// Create a new router
	router := mux.NewRouter()

	// Register API endpoints
	router.HandleFunc("/transcode", transcodingHandler.CreateTranscodeJob).Methods("POST")
	router.HandleFunc("/jobs/{jobID}", transcodingHandler.GetJobStatus).Methods("GET")

	// Start the API server
	log.Println("Media Transcoder API server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
