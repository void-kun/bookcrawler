package goseaweedfs

import "encoding/json"

type ChunkInfo struct {
	Fid    string `json:"fid"`
	Offset int64  `json:"offset"`
	Size   int64  `json:"size"`
}

type ChunkManifest struct {
	Name   string       `json:"name,omitempty"`
	Mime   string       `json:"mine,omitempty"`
	Size   int64        `json:"size,omitempty"`
	Chunks []*ChunkInfo `json:"chunks,omitempty"`
}

func (c *ChunkManifest) Marshal() ([]byte, error) {
	return json.Marshal(c)
}
