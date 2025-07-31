package models

import (
	"fmt"
	"log"
	"qr-code-generator/pkg/utils"
	"time"

	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

var db *gorm.DB

// Initialize the db variable
func SetDB(database *gorm.DB) {
	db = database
}

type SimpleQRCode struct {
	Content string
	Size    int
	URL     string
	Token   string
}

type QRCode struct {
	gorm.Model
	//ID           uint   `gorm:"primaryKey"`
	URL   string `gorm:"" json:"URL"`
	Size  int    `gorm:"" json:"size"`
	Token string `gorm:"" json:"token"`
	// Friends     []User `gorm:"many2many:friends;"`
}

func (code *QRCode) Generate() ([]byte, error) {
	// generate a token for the QR code
	token, err1 := utils.GenerateToken(32)

	if err1 != nil {
		log.Println("Error creating tokens:", err1)
		return nil, err1
	}
	//URL to be generated automatically and to be added as a URL to the QR code
	code.URL = "http://localhost:8080/qrcode/" + token

	code.Token = token

	//generate the QR code
	qrCode, err := qrcode.Encode(code.URL, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	//parameters needs to be passed to saveqrcode
	_, err2 := code.SaveQRcode()
	if err2 != nil {
		log.Println("Error saving QR code:", err1)
	}

	return qrCode, nil
}

func (q *QRCode) SaveQRcode() (*QRCode, error) {
	err := db.Create(&q).Error
	if err != nil {
		log.Println("Error saving qr code:", err)
		return nil, err
	}
	return q, nil
}

func ValidateORCodeByToken(token string) (*QRCode, *gorm.DB) {
	var getQRCode QRCode
	//db := db.Where("token=?", token).Find(&getQRCode)

	oneMinuteAgo := time.Now().Add(-1 * time.Minute)
	db := db.Where("token=? AND created_at >= ?", token, oneMinuteAgo).Find(&getQRCode)

	return &getQRCode, db
}
