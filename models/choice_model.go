/*
 * @Author: rui.li
 * @Date: 2024-04-24 17:26:42
 * @LastEditors: rui.li
 * @LastEditTime: 2024-04-24 17:28:29
 * @FilePath: /GoOpenAI/models/choice_model.go
 */
package models

type ChoiceModel struct {
	Index        int          `json:"index"`
	Message      MessageModel `json:"message"`
	FinishReason string       `json:"finish_reason"`
	LogProbs     *LogProbs    `json:"logprobs,omitempty"`
}

type LogProbs struct {
	// Content is a list of message content tokens with log probability information.
	Content []LogProb `json:"content"`
}

type LogProb struct {
	Token   string  `json:"token"`
	LogProb float64 `json:"logprob"`
	Bytes   []byte  `json:"bytes,omitempty"` // Omitting the field if it is null
	// TopLogProbs is a list of the most likely tokens and their log probability, at this token position.
	// In rare cases, there may be fewer than the number of requested top_logprobs returned.
	TopLogProbs []TopLogProbs `json:"top_logprobs"`
}

type TopLogProbs struct {
	Token   string  `json:"token"`
	LogProb float64 `json:"logprob"`
	Bytes   []byte  `json:"bytes,omitempty"`
}
