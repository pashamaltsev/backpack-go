package backpackgo

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"maps"
	"sort"
	"strings"

	"github.com/pashamaltsev/backpack-go/utils"
)

func auth(apikey, secret string, payload map[string]any, instruction string, timestamp, window int64) (map[string]string, error) {
	signature, err := sign(secret, payload, instruction, timestamp, window)
	if err == nil {
		head := map[string]string{
			"X-API-Key":   apikey,
			"X-Signature": signature,
			"X-Timestamp": fmt.Sprintf("%d", timestamp),
			"X-Window":    fmt.Sprintf("%d", window),
		}
		return head, nil
	} else {
		return nil, err
	}
}

func sign(secret string, payload map[string]any, instruction string, timestamp, window int64) (string, error) {
	sign := fmt.Sprintf("instruction=%s&{payload}timestamp=%d&window=%d", instruction, timestamp, window)
	sign = strings.ReplaceAll(sign, "{payload}", sortPayload(payload))

	// Sign the string using the private key
	privateKeyBytes, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}
	privateKey := ed25519.NewKeyFromSeed(privateKeyBytes)
	signature := ed25519.Sign(privateKey, utils.StringToBytes(sign))
	return base64.StdEncoding.EncodeToString(signature), nil
}

func sortPayload(payload map[string]any) string {
	keys := make([]string, 0, len(payload))
	for k := range maps.Keys(payload) {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("%s=%v&", k, payload[k]))
	}
	return builder.String()
}
