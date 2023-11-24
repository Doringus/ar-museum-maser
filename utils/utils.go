package utils

import "github.com/skip2/go-qrcode"

func GenerateQrCode(query string) ([]byte, error) {
	return qrcode.Encode(query, qrcode.Medium, 256)
}

