package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"strings"
	"telegram-door-bell/internal/utils/base64"
)

const (
	AlgHS256 = "HS256"
	TypJWT   = "JWT"
)

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type Payload struct {
	Subject  string `json:"sub"`
	IssuedAt int64  `json:"iat"`
}

func GenerateToken(payload Payload, secret string) string {
	headerBytes, _ := json.Marshal(Header{
		Algorithm: AlgHS256,
		Type:      TypJWT,
	})
	headerEncoded := base64.Encode(headerBytes)

	payloadBytes, _ := json.Marshal(payload)
	payloadEncoded := base64.Encode(payloadBytes)

	signature := generateSignature(headerEncoded, payloadEncoded, secret)

	return headerEncoded + "." + payloadEncoded + "." + signature
}

func generateSignature(header string, payload string, secret string) string {
	secret = base64.DecodeString(secret)

	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(header + "." + payload))

	signature := base64.Encode(h.Sum(nil))
	return signature
}

func Verify(token string, secret string) bool {
	tokenParts := strings.Split(token, ".")

	if len(tokenParts) != 3 {
		return false
	}

	signatureGenerated := generateSignature(tokenParts[0], tokenParts[1], secret)

	return signatureGenerated == tokenParts[2]
}
func GetPayload(token string) (*Payload, error) {
	tokenParts := strings.Split(token, ".")

	if len(tokenParts) != 3 {
		return nil, errors.New("invalid token format")
	}

	var payload Payload
	err := json.Unmarshal(base64.Decode(tokenParts[1]), &payload)

	return &payload, err
}
