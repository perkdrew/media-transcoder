package models

import (
	"os/exec"
	)

type CmdExecutor interface {
	CombinedOutput(name string, arg ...string) ([]byte, error)
}

type RealCmdExecutor struct {}

func (r RealCmdExecutor) CombinedOutput(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.CombinedOutput()
}


// Transcoder handles the media transcoding process
type Transcoder struct {
	FFMpegPath string        // Path to the ffmpeg executable
	Executor   CmdExecutor   // executor to run the commands
}