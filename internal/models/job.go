package models

// Job represents a transcoding job
type Job struct {
	ID          string            `json:"jobID"`
	Status      string            `json:"status"`
	InputFile   string            `json:"inputFile"`
	OutputFile  string            `json:"outputFile"`
	Progress    int               `json:"progress"`
	Parameters  map[string]string `json:"parameters"`
}

// NewJob creates a new Job instance with default values
func NewJob(id, inputFile, outputFile string, parameters map[string]string) *Job { 
	return &Job{
		ID:         id,
		Status:     "pending",
		InputFile:  inputFile,
		OutputFile: outputFile,
		Progress:   0,
		Parameters: parameters, // set the parameters here
	}
}
