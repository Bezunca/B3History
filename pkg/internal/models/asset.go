package models

import "time"

type FixedPoint int // Represents the actual value multiplied by 100

type AssetInfo struct {
	Year                   int         	`json:"year"`
	TipReg                 int         	`json:"-"`
	DataCollectionDate     time.Time   	`json:"data"`
	BDICode                int         	`json:"-"`
	Ticker                 string		`json:"ticker"`
	MarketType             int			`json:"market_type"`
	CompanyName            string		`json:"shot_company_name"`
	SecurityType           string		`json:"security_type"`
	FutureMarketExpiration string       `json:" future_market_expiration"` // Need to check this against an example that actually has a value
	Currency               string		`json:"currency"`
	PriceOpen              FixedPoint	`json:"-"`
	PriceMax               FixedPoint	`json:"-"`
	PriceMin               FixedPoint	`json:"-"`
	PriceMean              FixedPoint	`json:"-"`
	PriceClose             FixedPoint	`json:"price_close"`
	PriceBid               FixedPoint	`json:"-"`
	PriceAsk               FixedPoint	`json:"-"`
	TotalTrades            int			`json:"-"`
	TotalQuantity          int			`json:"-"`
	TotalVolume            FixedPoint	`json:"total_volume"`
	PreExe                 FixedPoint   `json:"execution_price"` // Needs further investigation
	IndOpc                 int        	`json:"-"`	// Needs further investigation
	ExpirationDate         time.Time	`json:"expiration_date"`
	FatCot                 int 			`json:"-"` // Needs further investigation
	PtoExe                 int 			`json:"-"` // Needs further investigation
	ISINCode               string		`json:"-"`
	DistributionNumber     int			`json:"-"`
}
