package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"
)

func main() {

}

type Suretly struct {
	Id    string
	Token string
	Mode  string
}

var client = &http.Client{Timeout: 10 * time.Second}
var Endpoint = "https://api.suretly.io/"

// Public API methods
// common
func (s Suretly) Options() {

}

func (s Suretly) Orders() {

}

// create order and actions with orders
func (s Suretly) OrderNew() {

}

func (s Suretly) OrderStatus() {

}

func (s Suretly) OrderStop() {

}

func (s Suretly) OrderIssued() {

}

func (s Suretly) OrderPaid() {

}

func (s Suretly) OrderPartialPaid() {

}

func (s Suretly) OrderUnpaid() {

}

func (s Suretly) ContractGet() {

}

func (s Suretly) ContractAccept() {

}

func (s Suretly) Currencies() {

}

func (s Suretly) Countries() {

}

func (s Suretly) authKeyGen() (key string) {
	var requestId = randomId(10)
	hash := md5.New()
	hash.Write([]byte(requestId + s.Token))
	key = s.Id + "-" + requestId + "-" + hex.EncodeToString(hash.Sum(nil))
	return
}

func (s Suretly) get(url string, target interface{}) error {
	var Header = map[string][]string{
		"_auth": {s.authKeyGen()},
	}
	req := http.Request{
		Header: Header,
		Method: "GET",
	}
	r, err := client.Do(&req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
