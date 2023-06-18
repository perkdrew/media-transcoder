package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/perkdrew/media-transcoder/internal/handlers"
	"github.com/perkdrew/media-transcoder/internal/models"
	"github.com/perkdrew/media-transcoder/internal/transcoding"
	"github.com/stretchr/testify/assert"
)

type mockJobIDGenerator struct{}

func (m mockJobIDGenerator) Generate() string {
	return "mockJobID"
}

func TestCreateTranscodeJob(t *testing.T) {
	transcoder := transcoding.NewTranscoder("ffmpeg", models.RealCmdExecutor{})
	handler := handlers.NewTranscodingHandler(transcoder, mockJobIDGenerator{})

	requestBody := strings.NewReader(`{"InputFile":"input.mp4","OutputFile":"output.mp4","Parameters":{"videoCodec":"libx264","audioCodec":"aac"}}`)
	request := httptest.NewRequest(http.MethodPost, "/jobs", requestBody)
	responseRecorder := httptest.NewRecorder()

	handler.CreateTranscodeJob(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	var job models.Job
	_ = json.NewDecoder(responseRecorder.Body).Decode(&job)

	assert.Equal(t, "mockJobID", job.ID)
	assert.Equal(t, "input.mp4", job.InputFile)
	assert.Equal(t, "output.mp4", job.OutputFile)
	assert.Equal(t, map[string]string{"videoCodec": "libx264", "audioCodec": "aac"}, job.Parameters)
}
