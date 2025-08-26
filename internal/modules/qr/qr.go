package qr

import (
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(data string) ([]byte, error) {
	png, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	_ = os.WriteFile("qrcode.png", png, 0644)
	fmt.Println("Qrcode PNG создался ")
	return png, nil
}
