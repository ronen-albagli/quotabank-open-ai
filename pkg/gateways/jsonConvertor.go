package gateways

import (
	"encoding/json"
	"fmt"
)

type ChatCompletion struct {
	ID                string     `json:"id"`
	Object            string     `json:"object"`
	Created           int        `json:"created"`
	Model             string     `json:"model"`
	SystemFingerprint string     `json:"system_fingerprint"`
	Choices           []Choice   `json:"choices"`
	Usage             UsageStats `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     string  `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type UsageStats struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Openai response mock
func convert() (ChatCompletion, error) {
	jsonData := `{
		"id": "chatcmpl-123",
		"object": "chat.completion",
		"created": 1677652288,
		"model": "gpt-3.5-turbo-0125",
		"system_fingerprint": "fp_44709d6fcb",
		"choices": [{
			"index": 0,
			"message": {
				"role": "assistant",
				"content": "Hola"
			},
			"logprobs": null,
			"finish_reason": "stop"
		}],
		"usage": {
			"prompt_tokens": 9,
			"completion_tokens": 12,
			"total_tokens": 21
		}
	}`

	// Unmarshal JSON data into struct
	var completion ChatCompletion
	if err := json.Unmarshal([]byte(jsonData), &completion); err != nil {
		fmt.Println("Error:", err)
		return completion, err
	}

	return completion, nil
}
