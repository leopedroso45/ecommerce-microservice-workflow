package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		log.Fatalf("Error during jsonFile opening: %v", err)
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	return data
}

/*ListProducts it gives all products*/
func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	data := loadData()
	vars := mux.Vars(r)

	var products Products
	json.Unmarshal(data, &products)

	for _, v := range products.Products {
		if v.Uuid == vars["id"] {
			product, _ := json.Marshal(v)
			w.Write(product)
		}
	}
}

func main() {
	/*Create router */
	router := mux.NewRouter()
	/*Create HandleFunc with routes */
	router.HandleFunc("/products", ListProducts)
	router.HandleFunc("/product/{id}", GetProductByID)

	http.ListenAndServe(":8081", router)
}
