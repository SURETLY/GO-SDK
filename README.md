# GO-SDK

SDK for Suretly Lender API
## Установка
Для подключения SDK необходимо скопировать  папку с SDK в проект и подключить файл Suretly.java в месте использования SDK.
## Подключение
    suretly := suretly.NewDemo(id, token)
    or
    suretly := suretly.NewProduction(id, token)

## Вызов методов API через SDK

### #1 Общие методы

#### #1.1 Получение параметров для поиска поручителей

    loan, err := suretly.Options()

#### #1.2 Список заявок

    orders, err := suretly.Orders()

### #2 Создание и работа с заявками

#### #2.2 Создать заявку на поручительство

    orderNewResponse, err := suretly.OrderNew(OrderNew)

#### #2.3 Получить статус заявки

    orderStatus, err := suretly.OrderStatus(orderId)

#### #2.4 Отменить заявку

    err := suretly.OrderStop(orderId)

#### #2.9 Получить контракт для заемщика

    htmlText, err := suretly.ContractGet(orderId)

#### #2.10 Подтвердить что договор по заявке подписан заемщиком

    err := suretly.ContractAccept(orderId)

#### #2.11 Подтвердить что заявка оплачена и выдана

    err := suretly.OrderIssued(orderId)

### #3 Работа с оплатой заявки

#### #3.5 Отметить займ как выплаченный

    err := suretly.OrderPaid(orderId)

#### #3.6 Отметить займ как выплаченный частично

    err := suretly.OrderParialPaid(orderId, sum)

#### #3.7 Отметить займ как просроченный

    err := suretly.OrderUnPaid(orderID)

### Справочники

#### Валюты

    suretly.Currencies()

#### Страны

    suretly.Countries()