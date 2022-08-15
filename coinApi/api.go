package coinApi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    coinAPI, err := UnmarshalCoinAPI(bytes)
//    bytes, err = coinAPI.Marshal()


func UnmarshalCoinAPI(data []byte) (CoinAPI, error) {
	var r CoinAPI
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CoinAPI) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CoinAPI struct {
	Status string `json:"status"`
	CoinData   Data   `json:"data"`  
}

type Data struct {
	Stats Stats  `json:"stats"`
	Coins []Coin `json:"coins"`
}

type Coin struct {
	UUID           string   `json:"uuid"`          
	Symbol         string   `json:"symbol"`        
	Name           string   `json:"name"`          
	Color          string   `json:"color"`         
	IconURL        string   `json:"iconUrl"`       
	MarketCap      string   `json:"marketCap"`     
	Price          string   `json:"price"`         
	ListedAt       int64    `json:"listedAt"`      
	Tier           int64    `json:"tier"`          
	Change         string   `json:"change"`        
	Rank           int64    `json:"rank"`          
	Sparkline      []string `json:"sparkline"`     
	LowVolume      bool     `json:"lowVolume"`     
	CoinrankingURL string   `json:"coinrankingUrl"`
	The24HVolume   string   `json:"24hVolume"`     
	BtcPrice       string   `json:"btcPrice"`      
}

type Stats struct {
	Total          int64  `json:"total"`         
	TotalCoins     int64  `json:"totalCoins"`    
	TotalMarkets   int64  `json:"totalMarkets"`  
	TotalExchanges int64  `json:"totalExchanges"`
	TotalMarketCap string `json:"totalMarketCap"`
	Total24HVolume string `json:"total24hVolume"`
}



func GetAllCoins() []Coin {
	url := "https://coinranking1.p.rapidapi.com/coins?referenceCurrencyUuid=yhjMzLPhuIDl&timePeriod=24h&tiers%5B0%5D=1&orderBy=marketCap&orderDirection=desc&limit=100&offset=0"

	req, _ := http.NewRequest("GET", url, nil);

	req.Header.Add("X-RapidAPI-Key", "81b82b588fmshf1cd64975bd20acp10f561jsn0606cfe22d88")
	req.Header.Add("X-RapidAPI-Host", "coinranking1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req);

	defer res.Body.Close();
	responseData, err := ioutil.ReadAll(res.Body);
	if err != nil {
		log.Fatalln("Error while reading response body", err.Error());
	}
	responseCoinData, err := UnmarshalCoinAPI(responseData);
	if err != nil{
		log.Fatalln("Error while unmarshalling coin api response", err.Error());
	}
	
	return responseCoinData.CoinData.Coins;
}