package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func merchantMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var autor = r.Header.Get("Authorization")

		if autor != "merchant" {
			w.Write([]byte("anda tidak punya akses"))
			return
		}
		next.ServeHTTP(w,r)
	}
}

func SuperMidlle(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var autor = r.Header.Get("Authorization")

		if autor != "su-admin" {
			w.Write([]byte("anda tidak punya akses"))
			return
		}
		next.ServeHTTP(w,r)
	}
}

type Merchant struct {
	Id     			int64  `json:"id"`
	Nama_toko		string `json:"nama_toko"`
	Deskripsi  		string `json:"Deskripsi"`
}
type MerchantData struct {
	MerchantToko []Merchant `json:"toko"`
}
type Toko struct {
	ID          int    `json:"id"`
	NamaToko    string `json:"nama_toko"`
	Deskripsi   string `json:"Deskripsi"`
	JumlahProduk int    `json:"jumlah_produk"`
  }
// Define a struct to hold the decoded JSON data
type Data struct {
	Toko []Toko `json:"toko"`
  }

func getMerchant(w http.ResponseWriter, r *http.Request){
	resp,_ := http.Get("http://127.0.0.1:8000/api/merchants")
	output,_ := ioutil.ReadAll(resp.Body)
	data2 := []byte(output)
	var data MerchantData
	json.Unmarshal(data2, &data)
	json.NewEncoder(w).Encode(data)
}

func getAllToko(w http.ResponseWriter, r *http.Request){
	resp,err := http.Get("http://127.0.0.1:9000/api/merchants")
	success := (err == nil)
	if success {
		output,_ := ioutil.ReadAll(resp.Body)
		data2 := []byte(output)
		var data Data
		json.Unmarshal(data2, &data)
		json.NewEncoder(w).Encode(data)
	}else{
		log.Println("error",err)
		panic("kesahalan di sisi get api")
	}
	
	
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/merchants",merchantMiddle(getMerchant))
	mux.HandleFunc("/toko",SuperMidlle(getAllToko))
	fmt.Println("server running")
	http.ListenAndServe(":6000",mux)
}


