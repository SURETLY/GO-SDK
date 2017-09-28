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
	loan = Loan{}
	err = s.get("/options", &loan)
	return
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
	req := http.Request{}

	req.Header = map[string][]string{
		"_auth": {s.authKeyGen()},
	}
	req.Method = "POST"
	req.Host = s.Host
	req.RequestURI = uri

	res, err := client.Do(&req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(target)

	return
}
