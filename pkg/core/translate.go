package core

import (
	"fmt"
	gateways "genaidemo/pkg/gateways"
	"time"
)

var walletId string = "SOME_WALLET_ID"
var cardId string = "SOME_CARD_ID"

func Translate(wordToTranslate string) (string, error) {
	turboApiKey := "OPENAI_API_KEY"
	quotabankApi := "QUOTABANK_API_KEY"
	turboTranslationClient := gateways.NewTurboClient(turboApiKey)
	quotaBankApi := gateways.NewQuotaBankClient("https://api.quotabank.net/api/v1", quotabankApi)

	qbBalance, err := quotaBankApi.GetCardBalance(walletId, cardId)

	if err != nil {
		return "", err
	}

	available := qbBalance.Balance.Total - qbBalance.Balance.Used
	if available <= 0 {
		return "", gateways.NewQuotaError("Not enough quota left")
	}

	messages := []gateways.Message{
		{Role: "system", Content: fmt.Sprintf("translate from english into spanish this: %s", wordToTranslate)},
	}

	translation, err := turboTranslationClient.GenerateCompletion(messages, "gpt-3.5-turbo", available)

	tokenUsed := translation.Usage.TotalTokens

	var useRequest gateways.UseRequest
	useRequest.Amount = tokenUsed
	useRequest.CardID = cardId
	useRequest.WalletID = walletId
	useRequest.Reason = "Translation usage"
	useRequest.Metadata = map[string]interface{}{
		"date": time.Now().UnixMilli(),
	}

	useResponse, err := quotaBankApi.Use(useRequest)

	if err != nil {
		fmt.Println("Error translating:", err)
		return "", err
	}

	fmt.Println("QuotaBank usage response:", useResponse)

	return translation.Choices[0].Message.Content, err
}
