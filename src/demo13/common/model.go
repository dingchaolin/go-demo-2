package common

type Metric struct {
	Metric    string   `json:"metric"`
	EndPoint  string   `json:"endpoint"`
	Tag       []string `json:"tag"`
	Value     float64  `json:"value"`
	Timestamp int64    `json:"timestamp"`
}
