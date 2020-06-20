package models

import "time"

type FixedPoint int // Represents the actual value multiplied by 100

type AssetInfo struct {
	TipReg                 int
	DataCollectionDate     time.Time
	BDICode                int
	Ticker                 string
	MarketType             int
	CompanyName            string
	SecurityType           string
	FutureMarketExpiration string // Need to check this against an example that actually has a value
	Currency               string
	PriceOpen              FixedPoint
	PriceMax               FixedPoint
	PriceMin               FixedPoint
	PriceMean              FixedPoint
	PriceClose             FixedPoint
	PriceBid               FixedPoint
	PriceAsk               FixedPoint
	TotalTrades            int
	TotalQuantity          int
	TotalVolume            FixedPoint
	PreExe                 FixedPoint // Needs further investigation
	IndOpc                 int        // Needs further investigation
	ExpirationDate         time.Time
	FatCot                 int // Needs further investigation
	PtoExe                 int // Needs further investigation
	ISINCode               string
	DistributionNumber     int
}
