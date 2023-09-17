package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var data = "https://jsonplaceholder.typicode.com/posts"

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		// meminta permintaan GET ke API
		res, err := http.Get(data)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		//membaca data
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		// memasukkan data ke slice
		user := []User{}
		err = json.Unmarshal(resBody, &user)
		if err != nil {
			panic(err)
		}

		// mengirim data json
		result, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/getuser", getUser)

	port := 8080
	fmt.Printf("Server online at port :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
