package suretly

type Orders struct {
	Total int     `json:"total"`
	List  []Order `json:"list"`
}

type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Country struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	CurrencyCode string `json:"currency_code"`
}

type Loan struct {
	MinTerm    int     `json:"min_term"`
	MaxTerm    int     `json:"max_term"`
	MinSum     float32 `json:"min_sum"`
	MaxSum     float32 `json:"max_sum"`
	ServerTime int     `json:"server_time"`
}

type OrderPaymentStatusCode int

const (
	ORDER_PAY_STATUS_NEW          OrderPaymentStatusCode = iota
	ORDER_PAY_STATUS_NO                                  // 1 - займ не выплачен
	ORDER_PAY_STATUS_PARTIAL                             // 2 - займ выплчен частично
	ORDER_PAY_STATUS_PAID                                // 3 - займ выплачен
	ORDER_PAY_STATUS_PAID_SURETLY                        // 4 - займ выплачен на сайте suretly
	ORDER_PAY_STATUS_DEFOLT                              // 5 - займ не отдан
)

type OrderStatusCode int

const (
	ORDER_STATUS_NEW      OrderStatusCode = iota // 0 - Новая заявка, ждем акцепта договора
	ORDER_STATUS_OPEN                            // 1 - Договр акцептован идет поиск поручителей
	ORDER_STATUS_CANCELED                        // 2 - Заявка анулирована по неизвестной причине
	ORDER_STATUS_TIMEOUT                         // 3 - Заявка остановлена, по истечению времени, сумма не набрана
	ORDER_STATUS_DONE                            // 4 - Заявка успешно завершена, сумма набрана
	ORDER_STATUS_ISSUED                          // 5 - Заявка оплачена и выдана
)

type Order struct {
	Id              string                 `json:"id"`
	Uid             string                 `json:"uid"`
	Status          OrderStatusCode        `json:"status"`
	PaymentStatus   OrderPaymentStatusCode `json:"payment_status"`
	Borrower        Borrower               `json:"borrower"`
	UserCreditScore int                    `json:"user_credit_score"`
	Cost            float32                `json:"cost"`      // стоимость поручительства
	LoanSum         float32                `json:"loan_sum"`  // сумма займа
	LoanTerm        int                    `json:"loan_term"` // срок в днях
	LoanRate        float32                `json:"loan_rate"` // процентная ставка
	CurrencyCode    string                 `json:"currency_code"`
	MaxWaitTime     int                    `json:"max_wait_time"` // сколько по времени ищем заемщиков (сек)
	CreatedAt       int                    `json:"created_at"`
	ModifyAt        int                    `json:"modify_at"`
	ClosedAt        int                    `json:"closed_at"`
	BidsCount       int                    `json:"bids_count"`
	BidsSum         float32                `json:"bids_sum"`
	Callback        string                 `json:"callback"`
}

type OrderNew struct {
	Uid             string   `json:"uid"`
	Public          bool     `json:"is_public"`
	Borrower        Borrower `json:"borrower"`
	UserCreditScore int      `json:"user_credit_score"`
	LoanSum         float32  `json:"loan_sum"`  // сумма займа
	LoanTerm        int      `json:"loan_term"` // срок в днях
	LoanRate        float32  `json:"loan_rate"` // процентная ставка
	CurrencyCode    string   `json:"currency_code"`
	Callback        string   `json:"callback"`
}

type OrderNewResponse struct {
	Id string `json:"id"`
}

type OrderStatus struct {
	Id            string                 `json:"id"`
	Status        OrderStatusCode        `json:"status"`
	PaymentStatus OrderPaymentStatusCode `json:"payment_status"`
	Public        bool                   `json:"public"`
	Cost          float32                `json:"cost"`
	Sum           float32                `json:"sum"`
	BidsCount     int                    `json:"bids_count"`
	BidsSum       float32                `json:"bids_sum"`
	StopTime      int                    `json:"stop_time"`
}

type Borrower struct {
	Name         Name     `json:"name"`
	Gender       string   `json:"gender"`
	Birth        Birth    `json:"birth"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Ip           string   `json:"ip"`
	ProfileUrl   string   `json:"profile_url"`
	PhotoUrl     string   `json:"photo_url"`
	Passport     Passport `json:"passport"`
	Registration Address  `json:"registration"`
	Residential  Address  `json:"residential"`
}

// full name
type Name struct {
	First  string `json:"first"`
	Middle string `json:"middle"`
	Last   string `json:"last"`
	Maiden string `json:"maiden"`
}

// birth date and place
type Birth struct {
	Date  int    `json:"date"`
	Place string `json:"place"`
}

// passport
type Passport struct {
	Series       string `json:"series"`
	Number       string `json:"number"`
	IssueDate    string `json:"issue_date"`
	IssuePlace   string `json:"issue_place"`
	IssueCode    string `json:"issue_code"`
	Registration string `json:"registration"`
}

// address
type Address struct {
	Country  string `json:"country"`
	Zip      string `json:"zip"`
	Area     string `json:"area"`
	City     string `json:"city"`
	Street   string `json:"street"`
	House    string `json:"house"`
	Building string `json:"building"`
	Flat     string `json:"flat"`
}

type Error struct {
	Code int
	Msg  string
}
