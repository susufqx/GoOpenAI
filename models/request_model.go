package models

type RequestModel struct {
	Model    string         `json:"model"`
	Messages []MessageModel `json:"messages"`
}
