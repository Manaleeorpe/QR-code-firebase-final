package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qr-code-generator/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

//var QRCode models.SimpleQRCode

func QrcodeGenerate(w http.ResponseWriter, r *http.Request) {

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
	codeData, err = qrCode.Generate()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			fmt.Sprintf("Could not generate QR code. %v", err),
		)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(codeData)
}

func ValidateQRCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Token := vars["token"]

	if Token == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Token is required to validate the QR code.")
		return
	}

	q, db := models.ValidateORCodeByToken(Token)

	if db.Error != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Database error: " + db.Error.Error())
		return
	}

	if db.RowsAffected == 0 { // how many records were found by the query—not how many were changed.
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("QR code not validated or not found.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var res string
	if q == nil {
		res = "QR code is not valid"
	} else {
		res = "QR code is valid"
	}

	json.NewEncoder(w).Encode(res)

}
