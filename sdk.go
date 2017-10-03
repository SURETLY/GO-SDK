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

func NewProduction(id string, token string) Suretly {
	host := "https://api.suretly.io:3000"
	return Suretly{Id: id, Token: token, Host: host}
}

func NewDemo(id string, token string) Suretly {
	host := "https://dev.suretly.io:3000"
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
func (s Suretly) Options() (loan Loan, err Error) {
	s.get("/options", &loan, &err)
	return
}

func (s Suretly) Orders() (orders Orders, err Error) {
	s.get("/orders", &orders, &err)
	return
}

/**
*	create new order
 */
// create order and actions with orders
func (s Suretly) OrderNew(order OrderNew) (resp OrderNewResponse, err Error) {
	s.post("/order/new", order, &resp, &err)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderStatus(id string) (status OrderStatus, err Error) {
	s.get("/order/status?id="+id, &status, &err)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderStop(id string) (err Error) {
	s.post("/order/stop", map[string]string{"id": id}, nil, &err)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderIssued(id string) (err Error) {
	s.post("/order/issued", map[string]string{"id": id}, nil, &err)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderPaid(id string) (err Error) {
	s.post("/order/paid", map[string]string{"id": id}, nil, &err)
	return
}

/**
*	id - order id
	sum - paid sum
*/
func (s Suretly) OrderPartialPaid(id string, sum float32) (err Error) {
	type PartialPaid struct {
		Id  string  `json:"id"`
		Sum float32 `json:"sum"`
	}
	s.post("/order/partialpaid", PartialPaid{Id: id, Sum: sum}, nil, &err)
	return
}

/**
*	id - order id
 */
func (s Suretly) OrderUnpaid(id string) (err Error) {
	s.post("/order/unpaid", map[string]string{"id": id}, nil, &err)
	return
}

/**
*	id - order id
 */
func (s Suretly) ContractGet(id string) (text string, err Error) {
	s.get("/contract/get?id="+id, &text, &err)
	return
}

/**
*	id - order id
 */
func (s Suretly) ContractAccept(id string) (err Error) {
	s.post("/contract/accept", map[string]string{"id": id}, nil, &err)
	return
}

/**
*	list of currencies
 */
func (s Suretly) Currencies() (currencies []Currency, err Error) {
	s.get("/currencies", &currencies, &err)
	return
}

/**
*	list of countries
 */
func (s Suretly) Countries() (countries []Country, err Error) {
	s.get("/countries", &countries, &err)
	return
}

func (s Suretly) AuthKeyGen() (key string) {
	var requestId = randomId(10)
	hash := md5.New()
	hash.Write([]byte(requestId + s.Token))
	key = s.Id + "-" + requestId + "-" + hex.EncodeToString(hash.Sum(nil))
	return
}

func (s Suretly) get(uri string, target interface{}, apiError *Error) (err error) {
	req, _ := http.NewRequest("GET", s.Host+uri, nil)
	req.Header.Add("_auth", s.AuthKeyGen())

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		json.NewDecoder(res.Body).Decode(apiError)
		return
	}

	err = json.NewDecoder(res.Body).Decode(target)

	return
}

func (s Suretly) post(uri string, body interface{}, target interface{}, apiError *Error) (err error) {
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", s.Host+uri, bytes.NewReader(b))
	req.Header.Add("_auth", s.AuthKeyGen())

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		json.NewDecoder(res.Body).Decode(apiError)
		return
	}
	err = json.NewDecoder(res.Body).Decode(target)
	return
}
