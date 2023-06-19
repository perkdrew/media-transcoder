package hls

import (
	"os"
	"path/filepath"
	"time"

	"github.com/grafov/m3u8"
)

func generateHLSPlaylist(videoFilePath string, playlistFilePath string) error {
	// Create an M3U8 playlist
	playlist := &m3u8.MediaPlaylist{
		TargetDuration: 10,
	}

	// Open the video file
	videoFile, err := os.Open(videoFilePath)
	if err != nil {
		return err
	}
	defer videoFile.Close()

	// Retrieve the file information
	fileInfo, err := videoFile.Stat()
	if err != nil {
		return err
	}

	// Calculate the segment duration
	segmentDuration := 5 * time.Second

	// Read the video file in chunks and add segments to the playlist
	buffer := make([]byte, 1024*1024) // 1MB buffer size
	for {
		n, err := videoFile.Read(buffer)
		if err != nil {
			break
		}
		segment := &m3u8.MediaSegment{
			URI:         filepath.Base(videoFilePath),
			Duration:    segmentDuration.Seconds(),
			Discontinuity: true,
		}
		playlist.Segments = append(playlist.Segments, segment)
	}

	// Set the playlist end duration
	playlist.EndList = true

	// Create the playlist file
	playlistFile, err := os.Create(playlistFilePath)
	if err != nil {
		return err
	}
	defer playlistFile.Close()

	// Write the playlist to the file
	err = m3u8.Encode(playlistFile, playlist)
	if err != nil {
		return err
	}

	return nil
}
