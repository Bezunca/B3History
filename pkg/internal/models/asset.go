package models

import "time"

type FixedPoint int // Represents the actual value multiplied by 100

type AssetInfo struct {
	Year                   int         	`json:"year"`
	TipReg                 int         	`json:"registration_type"`
	DataCollectionDate     time.Time   	`json:"data"`
	BDICode                int         	`json:"bdi_code"`
	Ticker                 string		`json:"ticker"`
	MarketType             int			`json:"market_type"`
	CompanyName            string		`json:"shot_company_name"`
	SecurityType           string		`json:"security_type"`
	FutureMarketExpiration string       `json:" future_market_expiration"` // Need to check this against an example that actually has a value
	Currency               string		`json:"currency"`
	PriceOpen              FixedPoint	`json:"price_open"`
	PriceMax               FixedPoint	`json:"price_max"`
	PriceMin               FixedPoint	`json:"price_min"`
	PriceMean              FixedPoint	`json:"price_mean"`
	PriceClose             FixedPoint	`json:"price_close"`
	PriceBid               FixedPoint	`json:"price_bid"`
	PriceAsk               FixedPoint	`json:"price_ask"`
	TotalTrades            int			`json:"total_trades"`
	TotalQuantity          int			`json:"total_quantity"`
	TotalVolume            FixedPoint	`json:"total_volume"`
	PreExe                 FixedPoint   `json:"execution_price"` // Needs further investigation
	IndOpc                 int        	`json:"options_indicator"`	// Needs further investigation
	ExpirationDate         time.Time	`json:"expiration_date"`
	FatCot                 int 			`json:"amount_factor"` // Needs further investigation
	PtoExe                 int 			`json:"execution_points"` // Needs further investigation
	ISINCode               string		`json:"isin_code"`
	DistributionNumber     int			`json:"distribution_number"`
}
