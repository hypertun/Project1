package models

type ErrResponse struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}
