package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/address_book/{username}/entry", addEntry)
  r.HandleFunc("/address_book", AddressBookHandler)
//http.HandleFunc("/address_book", AddressBookHandler)
  http.ListenAndServe(":8080", r)
}



// Create address book
// Update address book (add or update person)
// Get address book (all contacts) OR search by person name

func AddressBookHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
    fetchData(r)
  case "POST":
    postData(r)
  default:
    fmt.Println("No match to route")
  }

}

func fetchData(r *http.Request) {
 // curl -X GET http://localhost:8080/address_book?username=christine
  username := r.FormValue("username")
  fmt.Println("in get", username)
  fmt.Println(username)
}

func postData(r *http.Request) {
  //curl -X POST http://localhost:8080/address_book -d "username=christine"
  fmt.Println("in post")
  username := r.FormValue("username")
  fmt.Println("username: ", username)
}

func addEntry(w http.ResponseWriter, r *http.Request) {
  //curl -X POST http://localhost:8080/address_book/christine/entry -d "first_name=Test&last_name="Testerson"
  vars := mux.Vars(r)
  username := vars["username"]
  first_name := r.FormValue("first_name")
  fmt.Println("in add entry", first_name)
  fmt.Println(vars)
  fmt.Println(username)
}

