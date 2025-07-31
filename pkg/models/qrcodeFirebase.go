package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"qr-code-generator/pkg/config"
	"qr-code-generator/pkg/utils"

	"github.com/skip2/go-qrcode"
)

// QRCodeNode is the Firebase model for storing QR codes
type QRCodeNode struct {
	URL       string    `json:"url"`
	Size      int       `json:"size"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (code *QRCode) GenerateFirebase(ctx context.Context) ([]byte, error) {
	// generate a token for the QR code
	token, err1 := utils.GenerateToken(32)

	if err1 != nil {
		log.Println("Error creating tokens:", err1)
		return nil, err1
	}
	//URL to be generated automatically and to be added as a URL to the QR code
	code.URL = "http://localhost:8080/firebase/qrcode/" + token

	code.Token = token

	//generate the QR code
	qrCode, err := qrcode.Encode(code.URL, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	//parameters needs to be passed to saveqrcode

	_, err2 := code.SaveQRCodeToFirebase(ctx)
	if err2 != nil {
		log.Println("Error saving QR code:", err1)
	}

	return qrCode, nil
}

func (q *QRCode) SaveQRCodeToFirebase(ctx context.Context) (*QRCode, error) {
	ref, _ := config.FirebaseDB.NewRef("qr_codes").Push(ctx, nil)

	// set timestamps
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()

	// save data to Firebase
	if err := ref.Set(ctx, q); err != nil {
		log.Println("Error saving qr code:", err)
		return nil, err
	}

	return q, nil
}

// ValidateQRCodeByTokenFirebase validates a QR code based on token and expiry (1 minute)
func ValidateQRCodeByTokenFirebase(ctx context.Context, token string) (*QRCodeNode, error) {
	var qrMap map[string]QRCodeNode
	ref := config.FirebaseDB.NewRef("qr_codes")

	// fetch all QR codes
	if err := ref.Get(ctx, &qrMap); err != nil {
		log.Println("Failed to fetch QR codes:", err)
		return nil, err
	}

	// check token validity within 1 minute
	oneMinuteAgo := time.Now().Add(-1 * time.Minute)
	for _, qr := range qrMap {
		if qr.Token == token && qr.CreatedAt.After(oneMinuteAgo) {
			return &qr, nil
		}
	}

	return nil, nil // no valid QR code found
}
