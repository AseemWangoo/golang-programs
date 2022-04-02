package main

// Request is our self-made struct to process JSON request from Client
type Request struct {
	Name string `json:"name"`
}

// ImageRequestBody is the image request body
type ImageRequestBody struct {
	FileName string `json:"filename"`
	Body     string `json:"body"`
}

// ImageUploadResponse after the image is uploaded
type ImageUploadResponse struct {
	FileName string `json:"filename"`
	Location string `json:"filelocation"`
}
