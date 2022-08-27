package usecase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Backend struct {
	Chain           string `json:"chain"`
	Blocks          int    `json:"blocks"`
	Headers         int    `json:"headers"`
	BestBlockHash   string `json:"bestBlockHash"`
	Difficulty      string `json:"difficulty"`
	SizeOnDisk      int64  `json:"sizeOnDisk"`
	Version         string `json:"version"`
	Subversion      string `json:"subversion"`
	ProtocolVersion string `json:"protocolVersion"`
}

type Blockbook struct {
	Coin            string    `json:"coin"`
	Host            string    `json:"host"`
	Version         string    `json:"version"`
	GitCommit       string    `json:"gitCommit"`
	BuildTime       time.Time `json:"buildTime"`
	SyncMode        bool      `json:"syncMode"`
	InitialSync     bool      `json:"initialSync"`
	InSync          bool      `json:"inSync"`
	BestHeight      int       `json:"bestHeight"`
	LastBlockTime   time.Time `json:"lastBlockTime"`
	InSyncMempool   bool      `json:"inSyncMempool"`
	LastMempoolTime time.Time `json:"lastMempoolTime"`
	MempoolSize     int       `json:"mempoolSize"`
	Decimals        int       `json:"decimals"`
	DbSize          int64     `json:"dbSize"`
	About           string    `json:"about"`
} 

type StatusResponse struct {
	Blockbook Blockbook `json:"blockbook"`
	Backend  Backend `json:"backend"`
}

func GetStatus(baseUri string, apiKey string) (string, error) {
	request, _ := http.NewRequest("GET", baseUri, nil)
	if apiKey != "" {
		request.Header.Add("api-key", apiKey)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
	c := &http.Client{}
	resp, err := c.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	r := &StatusResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, r)
	j, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return "", err
	}
	return string(j), nil
}