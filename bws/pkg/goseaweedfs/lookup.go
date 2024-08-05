package goseaweedfs

import "math/rand"

type VolumeLocation struct {
	URL       string `json:"url,omitempty"`
	PublicURL string `json:"publicURL,omitempty"`
}

type VolumeLocations []*VolumeLocation

func (c VolumeLocations) Head() *VolumeLocation {
	if len(c) == 0 {
		return nil
	}
	return c[0]
}

func (c VolumeLocations) RandomPickForRead() *VolumeLocation {
	if len(c) == 0 {
		return nil
	}
	return c[rand.Intn(len(c))]
}

type LookupResult struct {
	VolumeLocations VolumeLocations `json:"volumeLocations,omitempty"`
	Error           string          `json:"error,omitempty"`
}
