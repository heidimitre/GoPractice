package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main(){
	router:= mux.NewRouter()
	router.HandleFunc("/hello", hello)
	router.HandleFunc("/add", add)
	http.ListenAndServe(":8080", router)
}

func hello(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("hiiiiii"))
}

func add(res http.ResponseWriter, req *http.Request){
	qParam := req.URL.Query()
	var sum int
	for _, value := range qParam {
		valInt,_:= strconv.Atoi(value[0])
		sum+= valInt
	}
	sumString:=strconv.Itoa(sum)
	res.Write([]byte(sumString))
}