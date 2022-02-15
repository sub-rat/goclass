package exmphttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type btc struct {
	Symbol      string `json:"symbol"`
	PriceChange string `json:"priceChange"`
}

// type response struct{

// }

func InitHttpExample() {
	resp, err := http.Get("https://api2.binance.com/api/v3/ticker/24hr?symbol=ETHBTC")
	// resp, err := http.Get("https://go.dev/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// fmt.Printf("response = %+v\n", resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	// scanner := bufio.NewScanner(resp.Body)
	// for i := 0; scanner.Scan() && i < 50; i++ {
	// 	fmt.Println(scanner.Text())
	// }
	var btcData btc
	err = json.Unmarshal(body, &btcData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", btcData)
}
