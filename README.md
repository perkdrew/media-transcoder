# Media Transcoder

A versatile Golang service for efficient, scalable media transcoding. Convert audio/video files into multiple formats, optimize for devices/platforms. Parallel processing, robust encoding. Integrate seamlessly for high-quality conversions and optimal playback experiences.

## Project Structure

The project is structured in the following way:

- `cmd/main.go`: The entry point of the application. This is where the application starts running.
- `internal/handlers/transcoding_handler.go`: This file contains the HTTP handlers for the API endpoints of the application. It includes functionality to create a new transcoding job and fetch the status of a job.
- `internal/models/job.go`: This file defines the `Job` struct, which represents a transcoding job. It also contains a constructor function for creating a new job.
- `internal/transcoding/transcoder.go`: This file contains the main transcoding logic, where the actual conversion of media files happens.
- `pkg/utils/file_utils.go`: This file contains utility functions for dealing with files.

## Basic Code Functions

- `NewTranscodingHandler`: Initializes a new transcoding handler with a given transcoder and an empty map for jobs.
- `CreateTranscodeJob`: This handler function creates a new transcoding job, adds it to the jobs map, and starts the transcoding process in a separate goroutine.
- `GetJobStatus`: This handler function fetches the status and details of a transcoding job given a job ID.
- `NewJob`: This constructor function initializes a new job with a given ID, input file, output file, and a map of parameters.
- `Transcode`: This function in the `Transcoder` struct performs the actual media file conversion.

## Building and Running the Project

To build and run the project, use the following steps:

1. **Compile the program**: Navigate to the directory containing your main Go file (`cmd/main.go` in this project) and run the `go build` command. This compiles your program and generates an executable file.

2. **Run the program**: Execute the generated binary to start the media transcoding service. On Unix-based systems (like Linux or MacOS), you do this with `./<executable_name>`. On Windows, just type `<executable_name>.exe`.

You can also compile and run your Go program in one step using the `go run` command followed by the path to the main Go file, like `go run cmd/main.go`.

Remember to run `go mod tidy` occasionally, especially before committing your code. This command ensures your project's dependencies are up to date and removes any unused ones.

To format your Go code according to the Go standards, use the `go fmt ./...` command before committing any code. This command recursively formats all Go files in the current directory and subdirectories.

## Docker Usage

This project includes a Dockerfile for creating a Docker image of the application. To build and run the project using Docker, follow these steps:

1. **Build the Docker image**: In your terminal, navigate to the project's root directory and run `docker build -t media-transcoder .` This builds a Docker image and tags it as "media-transcoder".

2. **Run the Docker container**: Run `docker run -p 8080:8080 media-transcoder` to start a Docker container with your application. The `-p` option maps port 8000 in the container to port 8000 on your machine.

## Testing in Go

Testing is a crucial part of software development. Go includes a built-in testing tool called `go test`. To use it, you need to write test functions in your Go source files. Here's a basic example of a test function:

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
