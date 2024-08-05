package goseaweedfs

type UploadResult struct {
	Name  string `json:"name,omitempty"`
	Size  int64  `json:"size,omitempty"`
	Error string `json:"error,omitempty"`
}

type AssignResult struct {
	FileID    string `json:"fileID,omitempty"`
	URL       string `json:"url,omitempty"`
	PublicURL string `json:"publicURL,omitempty"`
	Count     uint64 `json:"count,omitempty"`
	Error     string `json:"error,omitempty"`
}

type SubmitResult struct {
	FileName string `json:"fileName,omitempty"`
	FileURL  string `json:"fileURL,omitempty"`
	FileID   string `json:"fileID,omitempty"`
	Size     uint64 `json:"size,omitempty"`
	Error    string `json:"error,omitempty"`
}

type ClusterStatus struct {
	IsLeader bool     `json:"isLeader,omitempty"`
	Leader   string   `json:"leader,omitempty"`
	Peers    []string `json:"peers,omitempty"`
}

type SystemStatus struct{}

type Topology struct {
	DataCenters []*DataCenter
	Free        int
	Max         int
	Layouts     []*Layout
}

type DataCenter struct {
	Free  int
	Max   int
	Racks []*Rack
}

type Rack struct {
	DataNodes []*DataNode
	Free      int
	Max       int
}

type DataNode struct {
	Free      int
	Max       int
	PublicURL string `json:"publicURL"`
	URL       string `json:"url"`
	Volumes   int
}

type Layout struct {
	Replication string
	Writables   []int64
}
