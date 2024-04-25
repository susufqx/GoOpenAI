/*
 * @Author: rui.li
 * @Date: 2024-04-24 18:49:15
 * @LastEditors: rui.li
 * @LastEditTime: 2024-04-25 09:27:46
 * @FilePath: /GoOpenAI/tests/chat_test.go
 */
package tests

import (
	"context"
	"os"
	"testing"

	"github.com/susufqx/GoOpenAI/models"
	"github.com/susufqx/GoOpenAI/services"
)

func TestChat(t *testing.T) {
	chatService := services.NewChatService(os.Getenv("OPENAI_API_KEY"))
	resp, err := chatService.Send(context.Background(), services.NewChatRequest("gpt-3.5-turbo", []models.MessageModel{
		{
			Role:    "system",
			Content: "You are a player assistant.",
		},
		{
			Role:    "user",
			Content: "用Go语言实现一个归并排序的代码",
		},
	}))

	if err != nil {
		t.Error(err)
	}

	t.Log(resp.Choices[0].Message.Content)
}
