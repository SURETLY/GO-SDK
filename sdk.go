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
}

var client = &http.Client{Timeout: 10 * time.Second}
const Host = "api.suretly.io"

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

func (s Suretly) get(uri string, target interface{}) error {
	var Header = map[string][]string{
		"_auth": {s.authKeyGen()},
	}
	req := http.Request{
		Header: Header,
		Method: "GET",
		Host: Host,
		RequestURI: uri,
	}
	res, err := client.Do(&req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}