package main

import (
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

type Products struct {
	Products []Product
}

var productsUrl string

func init() {
	productsUrl = os.Getenv("PRODUCT_URL")
}

func loadProducts() []Product {
	response, err := http.Get(productsUrl + "/products")
	if err != nil {
		fmt.Printf("Catalog - Error requesting products: %v", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	var products Products
	json.Unmarshal(data, &products)

	return products.Products
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()
	t := template.Must(template.ParseFiles("template/catalog.html"))
	t.Execute(w, products)
}

func ShowProducts(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	response, err := http.Get(productsUrl + "/product/" + vars["id"])
	if err != nil{
		log.Fatalf("The http request failed with error: %v", err)
	}
	data,_ := ioutil.ReadAll(response.Body)
	var product Product
	json.Unmarshal(data, &product)

	t := template.Must(template.ParseFiles("template/view.html"))
	t.Execute(w, product)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", ListProducts)
	router.HandleFunc("/product/{id}", ShowProducts)

	http.ListenAndServe(":8080", router)

}
