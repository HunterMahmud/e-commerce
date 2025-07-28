package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productsList []Product

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
func aboutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'am Mahmud. I am a software engineer.")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}

	if r.Method != "GET" {
		http.Error(w, "Please give me get request", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productsList)
}

func addProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}

	if r.Method != "POST" {
		http.Error(w, "Please give me post request with data.", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding body."+err.Error(), 400)
		return
	}

	// fmt.Println(newProduct)
	newProduct.Id = len(productsList) + 1

	productsList = append(productsList, newProduct)

	encoder := json.NewEncoder(w)

	encoder.Encode(productsList)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloFunc)

	mux.HandleFunc("/about", aboutFunc)

	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/product", addProducts)

	fmt.Println("Server is running on port: 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("something went wrong:", err)
	}
}

func init() {
	p1 := Product{
		Id:          1,
		Title:       "Apple",
		Description: "This is an apple.",
		Price:       14.99,
		ImgUrl:      "https://assets.clevelandclinic.org/transform/cd71f4bd-81d4-45d8-a450-74df78e4477a/Apples-184940975-770x533-1_jpg",
	}

	p2 := Product{
		Id:          2,
		Title:       "Banana",
		Description: "Fresh yellow banana, full of energy.",
		Price:       9.99,
		ImgUrl:      "https://www.rainforest-alliance.org/wp-content/uploads/2021/06/bananas-1-e1624909301641.jpg.optimal.jpg",
	}

	p3 := Product{
		Id:          3,
		Title:       "Mango",
		Description: "Sweet and juicy mango from Rajshahi.",
		Price:       19.99,
		ImgUrl:      "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	}

	p4 := Product{
		Id:          4,
		Title:       "Orange",
		Description: "Citrus fruit rich in vitamin C.",
		Price:       12.50,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQR1d55KRpVZtNZDCGxi193hqjtIhz1Nb-OBQ&s",
	}

	p5 := Product{
		Id:          5,
		Title:       "Pineapple",
		Description: "Tropical fruit with a tangy taste.",
		Price:       24.99,
		ImgUrl:      "https://snaped.fns.usda.gov/sites/default/files/styles/crop_ratio_7_5/public/seasonal-produce/2018-05/pineapple.jpg.webp?itok=g7zYmLlW",
	}

	p6 := Product{
		Id:          6,
		Title:       "Watermelon",
		Description: "Large juicy fruit, perfect for summer.",
		Price:       29.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQoHLnenQKdhz19ApHVgb0SSfziDN3YDBmLQg&s",
	}

	p7 := Product{
		Id:          7,
		Title:       "Strawberry",
		Description: "Sweet red berries, rich in antioxidants.",
		Price:       19.50,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR0gRT71EwtApQ8ACXCGyfvsCkuxwIXsx2z8Q&s",
	}

	p8 := Product{
		Id:          8,
		Title:       "Blueberry",
		Description: "Small berries full of flavor and nutrients.",
		Price:       22.00,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRe2zLJUrzmTBzPSuMJSJNFFq5EC38TFJk3-A&s",
	}

	p9 := Product{
		Id:          9,
		Title:       "Papaya",
		Description: "Soft tropical fruit good for digestion.",
		Price:       15.75,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSZYT1tOS5o7hV4hje57tDhYLYd-f6WwZdDpg&s",
	}

	p10 := Product{
		Id:          10,
		Title:       "Grapes",
		Description: "Seedless green grapes, perfect as snacks.",
		Price:       17.80,
		ImgUrl:      "https://snaped.fns.usda.gov/sites/default/files/styles/crop_ratio_7_5/public/seasonal-produce/2018-05/grapes_0.jpg.webp?itok=ZiqbgHzZ",
	}

	productsList = append(productsList, p1)
	productsList = append(productsList, p2)
	productsList = append(productsList, p3)
	productsList = append(productsList, p4)
	productsList = append(productsList, p5)
	productsList = append(productsList, p6)
	productsList = append(productsList, p7)
	productsList = append(productsList, p8)
	productsList = append(productsList, p9)
	productsList = append(productsList, p10)

}
