package helper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func CreateRandomString(stringLength int) string {
	uuidBytes := make([]byte, stringLength)
	_, _ = rand.Read(uuidBytes)

	// Convert the bytes to a hexadecimal string
	requestID := hex.EncodeToString(uuidBytes)
	return requestID
}

func Chain(str ...string) string {
	for _, v := range str {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func GetEnv(key string, def ...string) string {
	resp := os.Getenv(key)
	if resp == "" {
		return Chain(def...)
	}
	return resp
}

func WrapString(payload, wrap string) (res string) {
	res = strings.TrimSpace(payload)
	if payload == "" || wrap == "" {
		return
	}
	return fmt.Sprintf("%s%s%s", wrap, payload, wrap)
}

// StandardizePhoneNumber standardizes a phone number by removing the country code and ensuring it starts with "0".
func StandardizePhoneNumber(phoneNumber string) string {
	// Remove non-digit characters from the phone number
	re := regexp.MustCompile(`[^\d]`)
	normalizedNumber := re.ReplaceAllString(phoneNumber, "")

	// Check if the number starts with the country code (e.g., +62 or 62)
	// If it does, remove the country code
	if len(normalizedNumber) > 2 && (normalizedNumber[0:2] == "62" || normalizedNumber[0:3] == "+62") {
		normalizedNumber = normalizedNumber[2:]
	}

	// Ensure the number starts with "0"
	if normalizedNumber[0] != '0' {
		normalizedNumber = "0" + normalizedNumber
	}

	return normalizedNumber
}
