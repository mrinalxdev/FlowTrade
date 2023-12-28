package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Trade struct {
	ID        string
	Symbol    string
	Quantity  int
	Price     float64
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
	Prices map[string]float64
}
type MarketDataSimulator struct {
	MarketData
}

func NewTradingAlgorithm() *TradingAlgorithm {
	return &TradingAlgorithm{
		MaxPortfolioSize:    20000,
		MaxOrderSize:        100,
		MaxRiskPerTrade:     0.02,
		MaxPortfolioRisk:    0.1,
		TrailingStopPercent: 0.05,
		SAMPeriod:           10,
	}
}

func (algo *TradingAlgorithm) ExecuteTrade(portfolio *Portfolio, symbol string, quantity int, marketData *MarketData) (*Trade, error) {
	if abs(quantity) > algo.MaxOrderSize {
		return nil, fmt.Errorf("order size exceeds maximum allowed")
	}

	price := marketData.GetPrice(symbol)
	trade := &Trade{
		ID:        generateTrade(),
		Symbol:    symbol,
		Quantity:  quantity,
		Price:     price,
		Timestamp: time.Now(),
	}
	tradeRisk := trade.Price * float64(abs(quantity)) / portfolio.Cash

	if tradeRisk > algo.MaxRiskPerTrade {
		return nil, fmt.Errorf("trade risk exceeds maximum allowed per trade ")
	}

	if (portfolio.TotalRisk + tradeRisk) > algo.MaxPortfolioRisk {
		return nil, fmt.Errorf("total portfolio risk exceeds maximum allowed")
	}

	if quantity > 0 {
		//Buying Process
		if portfolio.Cash < trade.Price*float64(quantity) {
			return nil, fmt.Errorf("insufficient funds to execute trade")
		}
		portfolio.Cash -= trade.Price * float64(quantity)
		portfolio.Positions[symbol] += quantity
		portfolio.TotalRisk += tradeRisk
	} else if quantity < 0 {
		//Selling the order
		if portfolio.Positions[symbol] < abs(quantity) {
			return nil, fmt.Errorf("insufficient position to sell")
		}
		portfolio.Cash += trade.Price * float64(abs(quantity))
		portfolio.Positions[symbol] += quantity
		portfolio.TotalRisk += tradeRisk
	}

	//trade to the trade history
	portfolio.History = append(portfolio.History, *trade)
	fmt.Printf("Trade Executed : %v\n", trade)

	return trade, nil
}
func generateTrade() string {
	return fmt.Sprintf("Trade-%d", time.Now().UnixNano())
}

func generateRandomPrice() float64 {
	return rand.Float64() * 100
}

// stop loss for an open positions
func (algo *TradingAlgorithm) UpdateTrailingStop(portfolio *Portfolio, symbol string, marketData *MarketData) {
	if position, ok := portfolio.Positions[symbol]; ok {
		trailingStop := algo.TrailingStopPercent * marketData.GetPrice(symbol) / float64(position)
		fmt.Printf("Trailing stop updated for %s: %.2f\n", symbol, trailingStop)
	}
}

func (simulator *MarketDataSimulator) SimulateMarketDataChange(symbol string) {
	priceChange := rand.Float64()*2 - 1
	simulator.Prices[symbol] *= (1 + priceChange)

}

func (marketData *MarketData) GetPrice(symbol string) float64 {
	if price, ok := marketData.Prices[symbol]; ok {
		return price
	}
	return 100.0
}

func RunSimulation(tradingAlgo *TradingAlgorithm, portfolio *Portfolio, marketData *MarketDataSimulator, symbols []string) {
	for i := 0; i < 100; i++ {
		//Market Data Changes
		for _, symbol := range symbols {
			marketData.SimulateMarketDataChange(symbol)
		}

		for _, symbol := range symbols {
			quantity := rand.Intn(20) - 10 // Simulating random trade Quantity
			_, err := tradingAlgo.ExecuteTrade(portfolio, symbol, quantity, &marketData.MarketData)
			if err != nil {
				fmt.Printf("Trade execution error : %v\n", err)
			}
		}

		printPortfolioInfo(portfolio)
		printMarketData(marketData.Prices)

		time.Sleep(1 * time.Second) // 1 second Time interval
	}
}

func printPortfolioInfo(portfolio *Portfolio) {
	fmt.Println("\n Portfolio:")
	fmt.Printf("Cash: $%.2f\n", portfolio.Cash)
	fmt.Printf("Total Risk: $%.2f\n", portfolio.TotalRisk)

	fmt.Println("Open Positions :")
	for symbol, quantity := range portfolio.Positions {
		fmt.Printf("%s:%d shares\n ", symbol, quantity)
	}

	fmt.Println("Trade History :")
	for _, trade := range portfolio.History {
		fmt.Printf("TradeID : %s, Symbol : %s, Quantity: %d, Price: $%.2f, Timestamp: %v\n", trade.ID, trade.Symbol, trade.Quantity, trade.Price, trade.Timestamp)
	}
}
func printMarketData(prices map[string]float64){
	fmt.Println("\nMarket Data:")
	for symbol, price := range prices {
		fmt.Printf("%s: $%.2f\n", symbol, price)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func logToFile(logEntry string) {
	file, err := os.OpenFile("trading_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(logEntry + "\n")
	if err != nil {
		fmt.Println("error writing to log file", err)
	}
}

func main() {
	tradingAlgo := NewTradingAlgorithm()
	portfolio := &Portfolio{
		Cash:      10000,
		Positions: make(map[string]int),
		TotalRisk: 0,
	}
	symbols := []string{"HDFC", "TATA", "MAHINDRA"}
	marketData := &MarketDataSimulator{
		MarketData: MarketData{
			Prices: make(map[string]float64),
		},
	}
	for _, symbol := range symbols {
		marketData.Prices[symbol] = generateRandomPrice()
	}

	//to execute the whole simulator
	RunSimulation(tradingAlgo, portfolio, marketData, symbols)
}
