package http_manager

import (
	"encoding/json"
	"fmt"
	http_model "go_workspace/hello/models"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

func BinaceMarketCall() {
	resp, err := http.Get("https://api.binance.com/api/v3/exchangeInfo")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var exchangeInforData map[string]interface{}
	json.Unmarshal(data, &exchangeInforData)

	var result []http_model.Symbol
	mapstructure.Decode(exchangeInforData["symbols"], &result)

	fmt.Println(result)
}
