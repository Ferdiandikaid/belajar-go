package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type DataCustomer struct {
	Id    int    `json:"customer_id"`
	Nama  string `json:"customer_nama"`
	Email string `json:"customer_email"`
}

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/getAllData", api_getAllData)
	http.ListenAndServe(":5050", mux)
}

func test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string("wow"))
}

func connect() *sql.DB {
	connection, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/customer_go")
	if err != nil {
		panic(err)
	} else {
		return connection
	}
}

func getData(connection *sql.DB) []DataCustomer {
	var dataCustomer = DataCustomer{}
	var dataCustomerArray []DataCustomer
	data, err := connection.Query("SELECT * FROM customer LIMIT 1")
	if err != nil {
		panic(err)
	} else {
		for data.Next() {
			err = data.Scan(&dataCustomer.Id, &dataCustomer.Nama, &dataCustomer.Email)
			if err != nil {
				panic(err)
			} else {
				dataCustomerArray = append(dataCustomerArray, DataCustomer{
					Id:    dataCustomer.Id,
					Nama:  dataCustomer.Nama,
					Email: dataCustomer.Email,
				})
			}
		}
	}
	return dataCustomerArray
}

func api_getAllData(w http.ResponseWriter, r *http.Request) {
	var dataCustomer = getData(connect())
	dataJson, err := json.Marshal(dataCustomer)
	if err != nil {
		panic(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(dataJson))
	}
}
