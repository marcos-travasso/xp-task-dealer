package openai_dealer

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
)

func generateHash(slice []openai.ChatCompletionMessage) string {
	jsonData, err := json.Marshal(slice)
	if err != nil {
		panic("error hashing messages: " + err.Error())
	}

	hash := sha256.Sum256(jsonData)
	return hex.EncodeToString(hash[:])
}
