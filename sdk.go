package main

import (
	"net/http"
	"time"
)

func init() {

}

type Suretly struct {
	Id     string
	Token  string
}

var client = &http.Client{ Timeout: 10 * time.Second }

func (s Suretly) authKeyGen(id string, token string)  {

}

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
