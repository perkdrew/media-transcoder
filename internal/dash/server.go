package dash

import (
	"fmt"
	"log"
	"net/http"

	"github.com/perkdrew/internal/dash"
)

// Server represents the DASH server.
type Server struct {
	manifestFilePath string
	mediaDir         string
}

// NewServer creates a new instance of the DASH server.
func NewServer(manifestFilePath, mediaDir string) *Server {
	return &Server{
		manifestFilePath: manifestFilePath,
		mediaDir:         mediaDir,
	}
}

// ServeHTTP handles incoming HTTP requests for DASH manifest and media segments.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/manifest.mpd" {
		s.handleManifestRequest(w, r)
	} else if r.URL.Path != "/" {
		s.handleSegmentRequest(w, r)
	} else {
		http.NotFound(w, r)
	}
}

// handleManifestRequest handles HTTP requests for the DASH manifest file.
func (s *Server) handleManifestRequest(w http.ResponseWriter, r *http.Request) {
	manifest, err := dash.ParseManifest(s.manifestFilePath)
	if err != nil {
		log.Printf("Failed to parse DASH manifest: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/dash+xml")
	http.ServeFile(w, r, s.manifestFilePath)
}

// handleSegmentRequest handles HTTP requests for DASH media segments.
func (s *Server) handleSegmentRequest(w http.ResponseWriter, r *http.Request) {
	segmentPath := s.mediaDir + r.URL.Path
	http.ServeFile(w, r, segmentPath)
}

// Start starts the DASH server and listens for incoming requests.
func (s *Server) Start(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("DASH server listening on http://localhost%s", addr)
	err := http.ListenAndServe(addr, s)
	if err != nil {
		log.Fatalf("Failed to start DASH server: %v", err)
	}
}
