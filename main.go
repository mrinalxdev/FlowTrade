package main

import (
	"time"
)

type Trade struct {
	ID        string
	Sysmbol   string
	Quantity  string
	Price     int
	Timestamp time.Time
}

type TradingAlgorithm struct {
	MaxPortfolioSize    int
	MaxOrderSize        int
	MaxRiskPerTrade     float64
	MaxPortfolioRisk    float64
	TrailingStopPercent float64
	SAMPeriod           int
}
type Portfolio struct {
	Cash      float64
	Positions map[string]int
	TotalRisk float64
	History   []Trade
}
type MarketData struct {
	Prices map[string][]float64
}
type MaketDataSimulator struct {
	MarketData
}

func NewTradingAlgorithm() *TradingAlgorithm {
	return &TradingAlgorithm{
		MaxPortfolioSize: 20000,
		MaxOrderSize:     100,
		MaxRiskPerTrade:  0.02,
		MaxPortfolioRisk: 0.1,
		TrailingStopPercent: 0.05,
		SAMPeriod: 10,
	}
}

