package gateways

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type QuotaBankClient struct {
	baseURL string
	apiKey  string
}

type UseRequest struct {
	Amount   int                    `json:"amount"`
	CardID   string                 `json:"cardId"`
	Metadata map[string]interface{} `json:"metadata"`
	Reason   string                 `json:"reason"`
	WalletID string                 `json:"walletId"`
}

type UseResponse struct {
	// Define response fields here
	// Adjust based on the actual response structure from the API
}

type QuotaError struct {
	message string
}

// Error returns the error message.
func (e *QuotaError) Error() string {
	return e.message
}

// NewQuotaError creates a new QuotaError instance with the given message.
func NewQuotaError(message string) *QuotaError {
	return &QuotaError{message: message}
}

type BalanceQuota struct {
	Total      int64 `json:"total" bson:"total"`
	Used       int64 `json:"used" bson:"used"`
	LowerLimit int64 `json:"lowerLimit" bson:"lowerLimit"`
}

type CardBalanceResponse struct {
	Balance BalanceQuota `json:"balance"`
}

func NewQuotaBankClient(baseURL string, apiKey string) *QuotaBankClient {
	return &QuotaBankClient{
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (c *QuotaBankClient) Use(reqBody UseRequest) (*UseResponse, error) {
	url := c.baseURL + "/wallet/quota/use"
	reqBodyJSON, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("qbjwt", c.apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var responseBody UseResponse
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	return &responseBody, nil
}

func (c *QuotaBankClient) GetCardBalance(walletID, cardID string) (*CardBalanceResponse, error) {
	url := fmt.Sprintf("%s/wallet/%s/cards/%s/balance", c.baseURL, walletID, cardID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("qbjwt", c.apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var responseBody CardBalanceResponse
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	return &responseBody, nil
}
