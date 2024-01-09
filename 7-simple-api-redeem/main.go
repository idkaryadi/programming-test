package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// reference https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go#reading-a-request-body

type User struct {
	ID   string // uuid
	Name string
}

type UserRequest struct {
	Name string `json:"name"`
}

type Invoice struct {
	ID     string // uuid
	UserID string
	Amount int
	Redeem bool
}

type InvoiceRequest struct {
	UserID string `json:"userId"`
	Amount int    `json:"amount"`
}

type Voucher struct {
	ID          string
	InvoiceID   string
	ExpiredDate time.Time
}

type VoucherRequest struct {
	InvoiceID string `json:"invoiceId"`
}

// TODO: update to db
var UserDB []User
var InvoiceDB []Invoice
var VoucherDB []Voucher

func main() {
	http.HandleFunc("/register", RegisterUser)
	http.HandleFunc("/user", GetUser)
	http.HandleFunc("/invoice", CreateInvoice)
	http.HandleFunc("/invoices", GetInvoice)
	http.HandleFunc("/generate-vouchers", GetInvoice)

	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// handle body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err", err)
		io.WriteString(w, "error body request")
		return
	}

	var userReq UserRequest
	json.Unmarshal(body, &userReq)
	// fmt.Println("body", userReq)

	// generate new user
	id := uuid.New().String()
	newUser := User{
		ID:   id,
		Name: userReq.Name,
	}
	// TODO: save db
	// fmt.Println("new User", newUser)
	UserDB = append(UserDB, newUser)

	// fmt.Println("UserDB", UserDB)
	// io.WriteString(w, "success")
	w.Header().Set("Content-Type", "application/json")
	p := map[string]interface{}{
		"status": "success",
		"payload": map[string]interface{}{
			"userId": id,
		},
	}
	json.NewEncoder(w).Encode(p)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := map[string]interface{}{
		"status":  "success",
		"payload": UserDB,
	}
	json.NewEncoder(w).Encode(p)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	// handle body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err", err)
		io.WriteString(w, "error body request")
		return
	}

	var invoiceReq InvoiceRequest
	json.Unmarshal(body, &invoiceReq)

	// input data
	id := uuid.New().String()
	newInvoice := Invoice{
		ID:     id,
		UserID: invoiceReq.UserID,
		Amount: invoiceReq.Amount,
		Redeem: false,
	}
	// TODO: save db
	fmt.Println("new Invoice", newInvoice)
	InvoiceDB = append(InvoiceDB, newInvoice)

	w.Header().Set("Content-Type", "application/json")
	p := map[string]interface{}{
		"status": "success",
		"payload": map[string]interface{}{
			"invoice": id,
		},
	}
	json.NewEncoder(w).Encode(p)
}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := map[string]interface{}{
		"status":  "success",
		"payload": InvoiceDB,
	}
	json.NewEncoder(w).Encode(p)
}

func getInvoiceById(id string) (Invoice, error) {
	for _, inv := range InvoiceDB {
		if inv.ID == id {
			return inv, nil
		}
	}
	return Invoice{}, errors.New("not found")
}

func insertBulkVoucher(listOfVoucher []Voucher) {
	VoucherDB = append(VoucherDB, listOfVoucher...)
}

func updateInvoiceById(id string, updatedInvoice Invoice) {
	for i, inv := range InvoiceDB {
		if inv.ID == id {
			InvoiceDB[i] = updatedInvoice
		}
	}
}

func GenerateVoucher(w http.ResponseWriter, r *http.Request) {
	// handle body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err", err)
		io.WriteString(w, "error body request")
		return
	}

	var voucherReq VoucherRequest
	json.Unmarshal(body, &voucherReq)

	// get invoice by id
	inv, err := getInvoiceById(voucherReq.InvoiceID)
	if err != nil {
		io.WriteString(w, "error data not found")
		return
	}

	if inv.UserID == "" {
		io.WriteString(w, "error unregistered user")
		return
	}
	// check is redeem
	if inv.Redeem {
		io.WriteString(w, "already redeem")
		return
	}

	// cek Amount
	// if < 1jt skip
	// di floor 1 juta
	countVoucher := math.Floor(float64(inv.Amount) / float64(1000000))
	if countVoucher < 1 {
		io.WriteString(w, "kurang dari batas minimal")
		return
	}

	// generate voucher
	var listOfVoucher []Voucher
	expDate := time.Now().AddDate(0, 3, 0)
	for i := 0; i < int(countVoucher); i++ {
		id := uuid.New().String()
		newVoucher := Voucher{
			ID:          id,
			InvoiceID:   inv.ID,
			ExpiredDate: expDate,
		}
		listOfVoucher = append(listOfVoucher, newVoucher)
	}

	// transaksional
	// save voucher
	insertBulkVoucher(listOfVoucher)

	// update invoice
	updateInvoice := Invoice{
		ID:     inv.ID,
		UserID: inv.UserID,
		Amount: inv.Amount,
		Redeem: true,
	}
	updateInvoiceById(inv.ID, updateInvoice)

	w.Header().Set("Content-Type", "application/json")
	p := map[string]interface{}{
		"status":  "success",
		"payload": listOfVoucher,
	}
	json.NewEncoder(w).Encode(p)
}
