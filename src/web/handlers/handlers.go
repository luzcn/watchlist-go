package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luzcn/watchlist-go/src/db"
	"net/http"
)

type Deps struct {
	DB db.DataAccess
}

func health(res http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(res, "{}\n")
}

// Create a product
// Example:
//
// curl -X POST 'http://localhost:5000/products' \
// -H 'Content-Type: application/json' \
// -d '{
//    "name": "product123",
//    "price": "12",
//    "notes": [
//        {
//            "note": "abc",
//            "context": "1234"
//        },
//        {
//            "note": "def",
//            "context": "78910"
//        }
//
//    ]
// }'
func createProductHandler(deps *Deps) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		product := db.Product{}
		if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
			panic(err)
		}
		deps.DB.CreateProduct(&product)
		_, _ = fmt.Fprintf(res, fmt.Sprintf(`{"id": %v}`, product.ID))
	})
}

// Get the first product record
// curl -X GET http://localhost:5000/products
func listProductsHandler(deps *Deps) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		product := deps.DB.GetProduct()

		fmt.Println(product)

		result, err := json.Marshal(product)

		if err != nil {
			panic(err)
		}

		_, _ = fmt.Fprintf(res, string(result))
	})
}

func Load(router *mux.Router, deps *Deps) {
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/products", createProductHandler(deps)).Methods("POST")
	router.HandleFunc("/products", listProductsHandler(deps)).Methods("GET")
}
