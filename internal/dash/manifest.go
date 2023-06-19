package dash

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Manifest represents a DASH manifest.
type Manifest struct {
	XMLName   xml.Name   `xml:"MPD"`
	MediaInfo []MediaInfo `xml:"Period>AdaptationSet>Representation"`
}

// MediaInfo represents media information within a DASH manifest.
type MediaInfo struct {
	ID           string `xml:"id,attr"`
	Bandwidth    uint   `xml:"bandwidth,attr"`
	SegmentURL   string `xml:"SegmentTemplate>SegmentURL,attr"`
}

// ParseManifest parses a DASH manifest from the provided file path.
func ParseManifest(filepath string) (*Manifest, error) {
	xmlFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open manifest file: %v", err)
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest file: %v", err)
	}

	var manifest Manifest
	err = xml.Unmarshal(byteValue, &manifest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal manifest XML: %v", err)
	}

	return &manifest, nil
}

// GetMediaInfoByID retrieves the media information for a given ID from the DASH manifest.
func (m *Manifest) GetMediaInfoByID(id string) (*MediaInfo, error) {
	for _, mediaInfo := range m.MediaInfo {
		if mediaInfo.ID == id {
			return &mediaInfo, nil
		}
	}

	return nil, fmt.Errorf("media with ID '%s' not found in the manifest", id)
}
