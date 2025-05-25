package handler

type ErrorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type FileResponse struct {
	Status      string `json:"status"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Filename    string `json:"filename"`
	Hash        string `json:"hash"`
	FileSize    int64  `json:"file_size"`
	ContentType string `json:"content_type"`
}
