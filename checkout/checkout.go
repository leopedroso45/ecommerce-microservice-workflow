package main

import (
	"checkout/queue"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Order struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductId string `json:"product_id"`
}

var productsUrl string

//SET PRODUCT_URL=http://localhost:8081
func init() {
	productsUrl = os.Getenv("PRODUCT_URL")
	//productsUrl = "http://localhost:8081"
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := http.Get(productsUrl + "/product/" + vars["id"])
	if err != nil {
		log.Fatalf("The http request failed with error: %v", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))
	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("template/checkout.html"))
	t.Execute(w, product)
}

func finish(w http.ResponseWriter, r *http.Request) {
	var order Order
	order.Name = r.FormValue("name")
	order.Email = r.FormValue("email")
	order.Phone = r.FormValue("phone")
	order.ProductId = r.FormValue("product_id")

	data, _ := json.Marshal(order)

	fmt.Println(string(data))

	connection := queue.Connect()
	queue.Notify(data, "checkout_ex", "", connection)

	w.Write([]byte("Processou!"))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/finish", finish)
	router.HandleFunc("/{id}", displayCheckout)

	http.ListenAndServe(":8082", router)
}
