package config

type Configs struct {
	ApiKey      string `env:"API_KEY,required"`
	ApiUrl      string `env:"API_URL" envDefault:"https://www.alphavantage.co/query"`
	NDays       int    `env:"N_DAYS" envDefault:"10"`
	StockSymbol string `env:"STOCK_SYMBOL" envDefault:"SPY"`
}
