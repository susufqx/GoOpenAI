/*
 * @Author: rui.li
 * @Date: 2024-04-24 17:44:25
 * @LastEditors: rui.li
 * @LastEditTime: 2024-04-24 18:16:45
 * @FilePath: /GoOpenAI/services/chat_serice.go
 */
package services

import (
	"context"

	"github.com/susufqx/GoOpenAI/models"
	"github.com/susufqx/GoOpenAI/utils"
)

type ChatService struct {
	authToken     string
	orgnizationId *string
}

func NewChatService(authToken string) *ChatService {
	return &ChatService{
		authToken: authToken,
	}
}

func NewChatRequest(modelName string, messages []models.MessageModel) models.RequestModel {
	return models.RequestModel{
		Model:    modelName,
		Messages: messages,
	}
}

func (s *ChatService) SetOrgnizationId(orgId string) {
	*s.orgnizationId = orgId
}

func (s *ChatService) Send(ctx context.Context, req models.RequestModel) (*models.CompleteResponseModel, error) {
	url := "https://api.openai.com/v1/chat/completions"
	headers := map[string]string{
		"Authorization": "Bearer " + s.authToken,
		"Content-Type":  "application/json",
	}

	if s.orgnizationId != nil {
		headers["OpenAI-Organization"] = *s.orgnizationId
	}

	reqBytes, err := utils.Json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := postRawData(ctx, url, reqBytes, headers)
	if err != nil {
		return nil, err
	}

	completeResponseModel := &models.CompleteResponseModel{}
	err = utils.Json.Unmarshal(data, completeResponseModel)
	if err != nil {
		return nil, err
	}

	return completeResponseModel, nil
}
