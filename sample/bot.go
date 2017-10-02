package main

import (
	"fmt"
	"github.com/SURETLY/GO-SDK"
	"os"
)

func main() {
	sur := gosdk.NewSuretly("59d25e8bcea0995959de2da9", "gobot123123123", "dev")

	// получили лимиты на заявку
	loan, err := sur.Options()
	if err != nil {
		os.Exit(1)
	}
	// генерим внутренний uid заявки
	uid := gosdk.StringWithCharset(16, gosdk.Charset)
	// отправляем данные для заявкки, получаем id заявки
	newOrder := gosdk.OrderNew{
		Uid:    uid,
		Public: true,
		Borrower: gosdk.Borrower{
			Name: gosdk.Name{
				First:  "Антон",
				Middle: "Викторович",
				Last:   "Фролов",
			},
			Gender: "1",
			Birth: gosdk.Birth{
				Date:  623308357,
				Place: "г.Новосибирск",
			},
			Email:      "frolov_11123@mail.ru",
			Phone:      "+79231232766",
			Ip:         "109.226.15.42",
			ProfileUrl: "https://vk.com/frol_nsk",
			PhotoUrl:   "https://pp.userapi.com/c622420/v622420795/5368/BWdcNhJqFkc.jpg",
			Passport: gosdk.Passport{
				Series:     "4431",
				Number:     "989922",
				IssueDate:  "25.07.2007",
				IssuePlace: "Советский, отдел полиции №10, Управление МВД России по г. Новосибирску",
				IssueCode:  "554-223",
			},
			Registration: gosdk.Address{
				Country:  "Россия",
				Zip:      "630063",
				Area:     "Новосибирская область",
				City:     "Новосибирск",
				Street:   "Труженников",
				House:    "22",
				Building: "",
				Flat:     "24",
			},
			Residential: gosdk.Address{
				Country:  "Россия",
				Zip:      "630063",
				Area:     "Новосибирская область",
				City:     "Новосибирск",
				Street:   "Труженников",
				House:    "22",
				Building: "",
				Flat:     "24",
			},
		},
		UserCreditScore: 678,
		LoanSum:         loan.MaxSum / 2,
		LoanTerm:        loan.MaxTerm / 2,
		LoanRate:        38.1,
		CurrencyCode:    "RUB",
		Callback:        "callback",
	}
	id, err := sur.OrderNew(newOrder)
	fmt.Println(id, err)
	// по id заявки проверяем статус
	// и выгружаем договор по данной заявке
}
