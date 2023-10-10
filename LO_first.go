package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	stan "github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
)

type DeliveryData struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type PaymentData struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type ItemData struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type ModelData struct {
	OrderUID          string       `json:"order_uid"`
	TrackNumber       string       `json:"track_number"`
	Entry             string       `json:"entry"`
	Delivery          DeliveryData `json:"delivery"`
	Payment           PaymentData  `json:"payment"`
	Items             []ItemData   `json:"items"`
	Locale            string       `json:"locale"`
	InternalSignature string       `json:"internal_signature"`
	CustomerID        string       `json:"customer_id"`
	DeliveryService   string       `json:"delivery_service"`
	Shardkey          string       `json:"shardkey"`
	SmID              int          `json:"sm_id"`
	DateCreated       string       `json:"date_created"`
	OofShard          string       `json:"oof_shard"`
}

var dataCache *cache.Cache
var ctx = context.Background()

func main() {
	//2.1 Создаем соединение с сервером NATS Streaming
	clusterID := "test-cluster"
	clientID := "sender-client"
	natsURL := "0.0.0.0:4222"

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}
	defer sc.Close()

	// Создаем канал для получения сообщений
	ch := make(chan *stan.Msg)

	// Подписываемся на канал
	sub, err := sc.Subscribe("channel", func(msg *stan.Msg) {
		ch <- msg
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Fatalf("Failed to subscribe to channel: %v", err)
	}
	defer sub.Unsubscribe()
	//2.2
	dataCache = cache.New(cache.NoExpiration, 0)
	ConnStr := "user=clown password=12345678 port=5433 dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Успешное соединение с базой данных")
	}
	defer db.Close()
	// Читаем содержимое файла model.json
	fileData, err := os.ReadFile("model.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	// Создаем структуру для декодирования JSON
	var data ModelData

	//Проверка наличия данных 2.3
	if cachedData, found := dataCache.Get("modelData"); found {
		// Данные найдены в кэше, декодируем их
		data = cachedData.(ModelData)
	} else {
		// Данных нет в кэше, читаем содержимое файла и декодируем JSON
		err = json.Unmarshal([]byte(fileData), &data)
		if err != nil {
			fmt.Println("Ошибка декодирования JSON:", err)
			return
		}
		// Сохраняем данные в кэше
		dataCache.Set("modelData", data, cache.NoExpiration)
	}

	// Декодируем данные JSON в структуру
	err = json.Unmarshal([]byte(fileData), &data)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return
	}
	// Код для соединения с базой данных -> Подготовка SQL-запроса
	delivery := "INSERT INTO delivery (order_uid,name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7,$8)"
	payment := "INSERT INTO payment(order_uid,transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11)"
	order_items := "INSERT INTO order_items(order_uid,track_number,chrt_id, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	orders := "INSERT INTO orders(order_uid,track_number, entry, locale,internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	// Выполнение SQL-запроса
	_, err = db.Exec(orders, data.OrderUID, data.TrackNumber, data.Entry, data.Locale, data.InternalSignature, data.CustomerID, data.DeliveryService, data.Shardkey, data.SmID, data.DateCreated, data.OofShard)
	if err != nil {
		fmt.Println("Ошибка при выполнении INSERT-запроса:", err)
		return
	}
	_, err = db.Exec(delivery, data.OrderUID, data.Delivery.Name, data.Delivery.Phone, data.Delivery.Zip, data.Delivery.City, data.Delivery.Address, data.Delivery.Region, data.Delivery.Email)
	if err != nil {
		fmt.Println("Ошибка при выполнении INSERT-запроса:", err)
		return
	}
	_, err = db.Exec(payment, data.OrderUID, data.Payment.Transaction, data.Payment.RequestId, data.Payment.Currency, data.Payment.Provider, data.Payment.Amount, data.Payment.PaymentDt, data.Payment.Bank, data.Payment.DeliveryCost, data.Payment.GoodsTotal, data.Payment.CustomFee)
	if err != nil {
		fmt.Println("Ошибка при выполнении INSERT-запроса:", err)
		return
	}
	_, err = db.Exec(order_items, data.OrderUID, data.TrackNumber, data.Items[0].ChrtId, data.Items[0].Price, data.Items[0].Rid, data.Items[0].Name, data.Items[0].Sale, data.Items[0].Size, data.Items[0].TotalPrice, data.Items[0].NmId, data.Items[0].Brand, data.Items[0].Status)
	if err != nil {
		fmt.Println("Ошибка при выполнении INSERT-запроса:", err)
		return
	}
	// Закрытие соединения с базой данных
	defer db.Close()

	//2.4
	restoreCacheFromDB()

	//2.5 Создаем экземпляр роутера http://localhost:8000/data/modelData
	router := mux.NewRouter()
	router.HandleFunc("/data/{id}", getDataHandler).Methods("GET")
	// Запуск HTTP-сервер на определенном порту
	log.Fatal(http.ListenAndServe(":8000", router))

}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Извлекаем данные из кэша по идентификатору
	if cachedData, found := dataCache.Get(id); found {
		// Преобразуйте данные в формат JSON
		/*jsonData, err := json.Marshal(cachedData)
		if err != nil {
			http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
			return
		}*/
		tmpl, err := template.ParseFiles("showData.html")
		if err != nil {
			// Обработайте ошибку, если не удалось загрузить шаблон
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		// Отображаем данные, используя шаблон и заполненную структуру данных
		err = tmpl.Execute(w, cachedData)
		if err != nil {
			// Обрабатываем ошибку, если не удалось отобразить данные с использованием шаблона
			http.Error(w, "Failed to render data", http.StatusInternalServerError)
			return
		}

		// Установим заголовки HTTP и отправим данные в ответ
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusOK)
		//w.Write(jsonData)
	} else {
		// Если данные не найдены в кэше, вернем ошибку "Not Found"
		http.Error(w, "Data not found", http.StatusNotFound)
	}

}

func restoreCacheFromDB() {
	// Открываем соединение с базой данных
	ConnStr := "user=clown password=12345678 port=5433 dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Успешное соединение с базой данных")
	}
	defer db.Close()

	// Запросы для получения данных из таблиц
	deliveryQuery := "SELECT order_uid, name, phone, zip, city, address, region, email FROM delivery"
	paymentQuery := "SELECT order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payment"
	orderItemsQuery := "SELECT order_uid, track_number, chrt_id, price, rid, name, sale, size, total_price, nm_id, brand, status FROM order_items"
	ordersQuery := "SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders"

	// Выполнение запросов для получения данных
	deliveryRows, err := db.Query(deliveryQuery)
	if err != nil {
		fmt.Println("Ошибка при выполнении SELECT-запроса для таблицы delivery:", err)
		return
	}
	defer deliveryRows.Close()

	paymentRows, err := db.Query(paymentQuery)
	if err != nil {
		fmt.Println("Ошибка при выполнении SELECT-запроса для таблицы payment:", err)
		return
	}
	defer paymentRows.Close()

	orderItemsRows, err := db.Query(orderItemsQuery)
	if err != nil {
		fmt.Println("Ошибка при выполнении SELECT-запроса для таблицы order_items:", err)
		return
	}
	defer orderItemsRows.Close()

	ordersRows, err := db.Query(ordersQuery)
	if err != nil {
		fmt.Println("Ошибка при выполнении SELECT-запроса для таблицы orders:", err)
		return
	}
	defer ordersRows.Close()

	// Восстановление данных из таблиц
	var order_uid string
	var data ModelData
	var deliveries []DeliveryData
	for deliveryRows.Next() {
		var delivery DeliveryData
		err := deliveryRows.Scan(&order_uid, &delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
		if err != nil {
			fmt.Println("Ошибка при сканировании результата SELECT-запроса для таблицы delivery:", err)
			return
		}

		deliveries = append(deliveries, delivery)
	}
	data.Delivery = deliveries[0]

	var payments []PaymentData
	for paymentRows.Next() {
		var payment PaymentData
		err := paymentRows.Scan(&order_uid, &payment.Transaction, &payment.RequestId, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDt, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
		if err != nil {
			fmt.Println("Ошибка при сканировании результата SELECT-запроса для таблицы payment:", err)
			return
		}
		payments = append(payments, payment)
	}
	data.Payment = payments[0]

	var orderItems []ItemData
	for orderItemsRows.Next() {
		var item ItemData
		err := orderItemsRows.Scan(&order_uid, &item.TrackNumber, &item.ChrtId, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
		if err != nil {
			fmt.Println("Ошибка при сканировании результата SELECT-запроса для таблицы order_items:", err)
			return
		}

		orderItems = append(orderItems, item)
	}
	data.Items = orderItems

	var orders []ModelData
	for ordersRows.Next() {
		var order ModelData
		err := ordersRows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard)
		if err != nil {
			fmt.Println("Ошибка при сканировании результата SELECT-запроса для таблицы orders:", err)
			return
		}

		orders = append(orders, order)
	}
	// Сохранение данных в кэш
	dataCache.Set(data.OrderUID, data, cache.NoExpiration)
}
