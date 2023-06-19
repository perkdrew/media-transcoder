package transcoder

import (
	"errors"
	"testing"
)

type MockExecutorSuccess struct{}

func (m MockExecutorSuccess) CombinedOutput(name string, arg ...string) ([]byte, error) {
	return []byte("Success"), nil
}

type MockExecutorFail struct{}

func (m MockExecutorFail) CombinedOutput(name string, arg ...string) ([]byte, error) {
	return []byte("Failed"), errors.New("command failed")
}

func TestTranscode(t *testing.T) {
	transcoder := NewTranscoder("/path/to/ffmpeg", MockExecutorSuccess{})
	err := transcoder.Transcode("input", "output", map[string]string{
		"videoCodec": "libx264",
		"audioCodec": "aac",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	transcoder = NewTranscoder("/path/to/ffmpeg", MockExecutorFail{})
	err = transcoder.Transcode("input", "output", map[string]string{
		"videoCodec": "libx264",
		"audioCodec": "aac",
	})

	if err == nil {
		t.Errorf("Expected error, got none")
	}
}
