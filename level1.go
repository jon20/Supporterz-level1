package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

type Res_t struct {
	Ok     bool     `json:"ok"`
	Amount int      `json:"amount"`
	Items  []string `json:"items"`
}
type Res_f struct {
	Ok  bool   `json:"ok"`
	Mes string `json:"message"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/checkout", apiRequest)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func apiRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//read http request
	rf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	//encord json
	js, err := simplejson.NewJson(rf)
	arr, _ := js.Get("order").StringArray()
	//fmt.Printf("%#v\n", arr[1])
	res, flag := checkSum(arr)
	fmt.Print(flag)

	if flag == true {
		Res_t := Res_t{true, res, arr}
		red, err := json.Marshal(Res_t)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, string(red))
		fmt.Printf(string(red))
	} else {
		Res_f := Res_f{false, "item_not_found"}
		red, err := json.Marshal(Res_f)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, string(red))
	}
}

//check_items
func checkSum(arr []string) (int, bool) {
	m := map[string]int{"101": 100, "102": 130, "103": 320, "104": 320, "105": 380,
		"201": 150, "202": 270, "203": 320, "204": 280,
		"301": 100, "302": 220, "303": 250, "304": 150,
		"305": 240, "306": 270, "307": 100, "308": 150}

	sum := 0
	check_exist := false
	for res := range arr {

		for k, v := range m {
			if arr[res] == k {
				sum += v
				check_exist = true

			}
		}
		if check_exist == false {
			return sum, false
		}
	}

	return sum, true
}
