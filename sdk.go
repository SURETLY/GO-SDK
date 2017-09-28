package gosdk

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"
)

type Suretly struct {
	Id    string
	Token string
	Host  string
}

func NewSuretly(id string, token string, mode string) Suretly {
	host := "https://api.suretly.io:3000"
	if mode == "demo" {
		host = "https://demo.suretly.io:3000"
	} else if mode == "dev" {
		host = "https://dev.suretly.io:3000"
	}
	return Suretly{Id: id, Token: token, Host: host}
}

var client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

// Public API methods
// common
func (s Suretly) Options() (loan Loan, err error) {
	err = s.get("/options", &loan)
	return
}

func (s Suretly) Orders() (orders Orders, err error) {
	err = s.get("/orders", &orders)
	return
}

/**
*	create new order
 */
// create order and actions with orders
func (s Suretly) OrderNew(order OrderNew) (err error) {
	err = s.post("/order/new", order, nil)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderStatus(id string) (status OrderStatus, err error) {
	err = s.get("/order/status?id="+id, status)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderStop(id string) (err error) {
	err = s.post("/order/stop", map[string]string{"id": id}, nil)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderIssued(id string) (err error) {
	err = s.post("/order/issued", map[string]string{"id": id}, nil)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderPaid(id string) (err error) {
	err = s.post("/order/paid", map[string]string{"id": id}, nil)
	return
}

/**
*	id - order id
	sum - paid sum
 */
func (s Suretly) OrderPartialPaid(id string, sum float32) (err error) {
	type PartialPaid struct {
		Id  string  `json:"id"`
		Sum float32 `json:"sum"`
	}
	err = s.post("/order/partialpaid", PartialPaid{Id: id, Sum: sum}, nil)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderUnpaid(id string) (err error) {
	err = s.post("/order/unpaid", map[string]string{"id": id}, nil)
	return
}

/**
*	id - order id
 */
func (s Suretly) ContractGet(id string) (text string, err error) {
	err = s.get("/contract/get?id="+id, text)
	return
}

/**
*	id - order id
 */
func (s Suretly) ContractAccept(id string) (err error) {
	err = s.post("/contract/accept", map[string]string{"id": id}, nil)
	return
}

/**
*	list of currencies
 */
func (s Suretly) Currencies() (currencies []Currency, err error) {
	err = s.get("/currencies", currencies)
	return
}

/**
*	list of countries
 */
func (s Suretly) Countries() (countries []Country, err error) {
	err = s.get("/countries", countries)
	return
}

func (s Suretly) authKeyGen() (key string) {
	var requestId = randomId(10)
	hash := md5.New()
	hash.Write([]byte(requestId + s.Token))
	key = s.Id + "-" + requestId + "-" + hex.EncodeToString(hash.Sum(nil))
	return
}

func (s Suretly) get(uri string, target interface{}) (err error) {
	req, _ := http.NewRequest("GET", s.Host+uri, nil)
	req.Header.Add("_auth", s.authKeyGen())

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(target)

	return
}

func (s Suretly) post(uri string, body interface{}, target interface{}) (err error) {
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", s.Host+uri, bytes.NewReader(b))
	req.Header.Add("_auth", s.authKeyGen())

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(target)

	return
}
