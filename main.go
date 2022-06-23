package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/tokiwong/stock-ticker/pkg/config"
	"github.com/tokiwong/stock-ticker/pkg/stocks"
)

type stocksHandler struct {
	service     stocks.Service
	ndays       int
	stockSymbol string
}

func main() {
	cfg := config.Configs{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("unexpected error while initializing the config: %v", err)
	}

	stockSvc := stocks.NewStock(cfg.ApiUrl, cfg.ApiKey)

	http.Handle("/api/daily", &stocksHandler{
		stockSvc,
		cfg.NDays,
		cfg.StockSymbol,
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr: ":8080",
	}

	start(server)
}

func (s *stocksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p, err := s.service.GetStockData(s.stockSymbol, s.ndays)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := make(map[string]string)
		resp["message"] = fmt.Sprintf("Unexpected error occurred: %v", err)
	}
	jsonResp, _ := json.MarshalIndent(p, "", "   ")
	w.Write(jsonResp)
	return
}

func start(server *http.Server) {
	log.Printf("Server listening on %s", server.Addr)
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err.Error())
	}
}
