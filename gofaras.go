package gofaras

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)
type Product struct {
	ProductName    string
	Quantity int
	Price    float64
}

type Customer struct {
	Name     string 
	PhoneNumber int
	ID    string
	Email string
}
type Invoice struct {
	Key string // your api key
	Products []Product
	Customer Customer
	TestMode int // 1 is test, 0 is live
}

type FarasResponse struct {
	Error_code    int
	Error_msg     string
	InvoiceURL    string
	InvoicePDFURL string
}

var FarasURL = "https://faras.io/api_newinvoice/"

func NewInvoice(inv Invoice) (bool,string,string,string){
	//marshal json request ...
	jsonRequest, jsonerr := json.Marshal(inv)
	if jsonerr != nil {
		fmt.Println(jsonerr)
		return true,"Error in marshaling json request","",""
	}
	// creating client request
	req, err := http.NewRequest("POST", FarasURL, bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return true,"Error in creating http client request","",""
	}
	defer resp.Body.Close()
	bd, _ := ioutil.ReadAll(resp.Body)
	
	var fdata FarasResponse
	// unmarshal json response 
	if reserr := json.Unmarshal(bd, &fdata); reserr != nil {
		fmt.Println(reserr)
		return true,"Error in unmarshal json response","",""
	}
	isError := false
	if fdata.Error_code != 0{
		isError = true
	}
	return isError, fdata.Error_msg, fdata.InvoiceURL, fdata.InvoicePDFURL

}
