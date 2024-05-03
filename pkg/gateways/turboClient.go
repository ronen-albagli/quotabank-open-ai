package gateways

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TurboClient struct {
	apiKey  string
	baseURL string
}

type CompletionRequest struct {
	Messages  []Message `json:"messages"`
	Model     string    `json:"model"`
	MaxTokens int64     `json:"maxTokens"`
}

type ChatGptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseBody struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func NewTurboClient(apiKey string) *TurboClient {
	return &TurboClient{
		apiKey:  apiKey,
		baseURL: "https://api.openai.com/v1/",
	}
}

func (c *TurboClient) sendRequest(method, endpoint string, requestBody interface{}) (*http.Response, error) {
	requestBodyJSON, _ := json.Marshal(requestBody)
	req, err := http.NewRequest(method, c.baseURL+endpoint, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	client := &http.Client{}
	return client.Do(req)
}

func (c *TurboClient) GenerateCompletion(messages []Message, model string, maxTokens int64) (ChatCompletion, error) {
	requestBody := CompletionRequest{
		Messages:  messages,
		Model:     model,
		MaxTokens: maxTokens,
	}
	resp, err := c.sendRequest("POST", "/chat/completions", requestBody)
	// if err != nil {
	// 	return nil, err
	// }
	defer resp.Body.Close()
	var responseBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	// if err != nil {
	// 	return , err
	// }
	var completions []string
	for _, choice := range responseBody.Choices {
		completions = append(completions, choice.Text)
	}

	stubResponse, err := convert()

	return stubResponse, err
	// return completions, nil
}
