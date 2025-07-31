package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"qr-code-generator/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

//var QRCode models.SimpleQRCode

func QrcodeGenerateFirebase(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)
	//var size, content string = r.FormValue("size"), r.FormValue("content")
	var size = r.FormValue("size")
	var codeData []byte

	w.Header().Set("Content-Type", "application/json")

	/*
		if content == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(
				"Could not determine the desired QR code content.",
			)
			return
		}*/

	qrCodeSize, err := strconv.Atoi(size) //string to int conversion
	if err != nil || size == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Could not determine the desired QR code size.")
		return
	}

	qrCode := models.QRCode{Size: qrCodeSize}

	ctx := context.Background()

	codeData, err1 := qrCode.GenerateFirebase(ctx)

	//codeData, err = qrCode.Generate()
	if err1 != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Could not generate QR code. %v", err1),
		)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(codeData)
}

func ValidateQRCodeFirebase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Token := vars["token"]

	if Token == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Token is required to validate the QR code.")
		return
	}
	ctx := context.Background()
	q, err := models.ValidateQRCodeByTokenFirebase(ctx, Token)

	//q, db := models.ValidateORCodeByToken(Token)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Could not Validate QR code. %v", err),
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//log.Println("QR code validation result:", q)
	var res string
	if q == nil {
		res = "QR code is not valid"
	} else {
		res = "QR code is valid"
	}

	json.NewEncoder(w).Encode(res)

}
