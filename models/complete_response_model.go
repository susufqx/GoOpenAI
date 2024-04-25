/*
 * @Author: rui.li
 * @Date: 2024-04-24 17:35:59
 * @LastEditors: rui.li
 * @LastEditTime: 2024-04-24 17:37:14
 * @FilePath: /GoOpenAI/models/complete_response_model.go
 */
package models

type CompleteResponseModel struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`

	Choices           []ChoiceModel `json:"choices"`
	Usage             UsageModel    `json:"usage"`
	SystemFingerPrint string        `json:"system_fingerprint"`
}

type UsageModel struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
