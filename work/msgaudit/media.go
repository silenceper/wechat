package msgaudit

type MediaData struct {
	OutIndexBuf string `json:"outindexbuf,omitempty"`
	IsFinish    bool   `json:"is_finish,omitempty"`
	Data        []byte `json:"data,omitempty"`
}
