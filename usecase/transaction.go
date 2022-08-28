package usecase

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Vin struct {
		Txid      string   `json:"txid"`
		Vout      int      `json:"vout"`
		Sequence  int64    `json:"sequence"`
		N         int      `json:"n"`
		Addresses []string `json:"addresses"`
		IsAddress bool     `json:"isAddress"`
		Value     string   `json:"value"`
		Hex       string   `json:"hex"`
}

type Vout struct {
		Value     string   `json:"value"`
		N         int      `json:"n"`
		Hex       string   `json:"hex"`
		Addresses []string `json:"addresses"`
		IsAddress bool     `json:"isAddress"`
}

type GetTransactionResponse struct {
	Txid    string `json:"txid"`
	Version int    `json:"version"`
	Vin     []Vin `json:"vin"`
	Vout []Vout `json:"vout"`
	BlockHash     string `json:"blockHash"`
	BlockHeight   int    `json:"blockHeight"`
	Confirmations int    `json:"confirmations"`
	BlockTime     int    `json:"blockTime"`
	Value         string `json:"value"`
	ValueIn       string `json:"valueIn"`
	Fees          string `json:"fees"`
	Hex           string `json:"hex"`
}

func GetTransaction(baseUri string, apiKey string, txId string) (string, error) {
	request, _ := http.NewRequest("GET", baseUri + "/tx/" + txId, nil)
	if apiKey != "" {
		request.Header.Add("api-key", apiKey)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
	c := &http.Client{}
	resp, err := c.Do(request)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		er := "request failed with status:" + resp.Status
		return "", errors.New(er)
	}
	defer resp.Body.Close()
	r := &GetTransactionResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, r)
	j, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return "", err
	}
	return string(j), nil
}