package gosdk

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"
	"crypto/tls"
)

type Suretly struct {
	Id    string
	Token string
	Host string
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
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
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

// create order and actions with orders
func (s Suretly) OrderNew(order OrderNew) (err error) {
	err = s.post("/order/new", order, nil)
	return
}

func (s Suretly) OrderStatus(id string) (status OrderStatus, err error){
	err = s.get("/order/status?id="+id, status)
	return
}

func (s Suretly) OrderStop(id string) (err error){
	err = s.post("/order/stop", map[string]string{"id": id}, nil)
	return
}

func (s Suretly) OrderIssued(id string) {
	err = s.post("/order/issued", map[string]string{"id": id}, nil)
	return
}

func (s Suretly) OrderPaid(id string) {
	err = s.post("/order/paid", map[string]string{"id": id}, nil)
	return
}

func (s Suretly) OrderPartialPaid(id string, sum float32) (err error){
	err = s.post("/order/partialpaid", map[string]string{"id": id, "sum": string(sum)}, nil)
	return
}

func (s Suretly) OrderUnpaid(id string) (err error){
	err = s.post("/order/unpaid", map[string]string{"id": id}, nil)
	return
}

func (s Suretly) ContractGet(id string) (text string, err error) {
	err = s.get("/order/contract/get?id="+id, text)
	return
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

func (s Suretly) get(uri string, target interface{}) (err error) {
	req, _ := http.NewRequest("GET", s.Host + uri, nil)
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
	req, _ := http.NewRequest("POST", s.Host + uri, b)
	req.Header.Add("_auth", s.authKeyGen())

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(target)

	return
}
