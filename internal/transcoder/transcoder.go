package transcoder

import (
	"fmt"
	"log"

	"github.com/perkdrew/media-transcoder/internal/models"
)

// Transcoder handles the media transcoding process
type Transcoder struct {
	FFMpegPath string      // Path to the ffmpeg executable
	Executor   models.CmdExecutor // executor to run the commands
}

// NewTranscoder creates a new Transcoder instance
func NewTranscoder(ffmpegPath string, executor models.CmdExecutor) *models.Transcoder {
	return &models.Transcoder{
		FFMpegPath: ffmpegPath,
		Executor:   executor,
	}
}

// Transcode performs media transcoding based on the provided parameters
func (t *Transcoder) Transcode(inputFile, outputFile string, parameters map[string]string) error {
	cmdArgs := []string{"-i", inputFile, "-c:v", parameters["videoCodec"], "-c:a", parameters["audioCodec"], outputFile}

	output, err := t.Executor.CombinedOutput(t.FFMpegPath, cmdArgs...)
	if err != nil {
		log.Printf("Transcoding failed: %s\n%s", err, output)
		return fmt.Errorf("transcoding failed: %s", err)
	}

	log.Println("Transcoding completed successfully")
	return nil
}
