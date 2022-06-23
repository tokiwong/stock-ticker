package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	defaultFunction = "TIME_SERIES_DAILY"
)

// Service is the exposed interface for the stocks API backend
type Service interface {
	// GetStockData returns the Stock history data given a Stock Symbol and nDays window
	GetStockData(stockSymbol string, nDays int) (*stockData, error)
}

type alpha struct {
	apiKey   string
	apiUrl   string
	function string

	cache map[string]cacheData
}

// NewStock returns an instance of alpha
func NewStock(apiUrl, apiKey string) Service {
	return &alpha{
		apiKey:   apiKey,
		apiUrl:   apiUrl,
		cache:    map[string]cacheData{},
		function: defaultFunction,
	}
}

// GetStockData returns daily historical stock data given a Stock Symbol and number of days
func (s *alpha) GetStockData(stockSymbol string, nDays int) (*stockData, error) {
	queryStr := fmt.Sprintf("apikey=%s&function=%s&symbol=%s", s.apiKey, s.function, stockSymbol)
	u, err := url.Parse(s.apiUrl)
	if err != nil {
		return nil, err
	}

	u.RawQuery = queryStr

	var jsonData []byte
	if ok := checkCache(s.cache, stockSymbol); ok {
		jsonData = s.cache[stockSymbol].data
	} else {
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		jsonData = body
		s.cache[stockSymbol] = cacheData{
			data:        jsonData,
			lastUpdated: time.Now(),
		}
	}

	var data stockData
	json.Unmarshal(jsonData, &data)

	return computeStockData(&data, nDays)
}

// computeStockData filters the StockData for the desired dates and computes the closing average
func computeStockData(stock *stockData, nDays int) (*stockData, error) {
	output := stockData{}
	output.Metadata = stock.Metadata
	output.OutputData.NDays = nDays

	dates := getDates(stock.Timeseries, nDays)

	if stock.Timeseries != nil {
		output.Timeseries = map[string]dailyData{}
		var sum float64
		for match, data := range stock.Timeseries {
			if contains(dates, match) {
				output.Timeseries[match] = data
				c, err := strconv.ParseFloat(data.Close, 32)
				if err != nil {
					return nil, err
				}
				sum += c
			}
		}
		output.OutputData.AvgClose = fmt.Sprintf("%f", sum/float64(nDays))
	}

	return &output, nil
}

// getDates returns a slice of trading dates given timeseries data and number of days
func getDates(timeseries map[string]dailyData, nDays int) []string {
	var dailies, dates []string

	for k := range timeseries {
		dailies = append(dailies, k)
	}
	sort.Strings(dailies)

	for i := 1; i <= nDays; i++ {
		dates = append(dates, dailies[len(dailies)-i])
	}

	return dates
}
