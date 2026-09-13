package engine

import (
	"crypto/sha256"
	"encoding/binary"
)

type Config struct {
	Exchange string
	Symbol   string
	Strategy string
	MaxPos   float64
}

func Default() Config {
	return Config{
		Exchange: "dydx",
		Symbol:   "BTC-USD",
		Strategy: "perp",
		MaxPos:   0.25,
	}
}

type Report struct {
	Equity float64
	Fills  int
	Bars   int
}

func ticker(symbol string, i int) float64 {
	sum := sha256.Sum256([]byte(symbol + string(rune(i))))
	n := binary.BigEndian.Uint16(sum[:2])
	return 100 + float64(n)/1000
}

func allow(notional, equity, maxPos float64) bool {
	if equity <= 0 {
		return false
	}
	return notional/equity <= maxPos
}

func Backtest(bars int) Report {
	cfg := Default()
	equity := 10000.0
	fills := 0
	for i := 0; i < bars; i++ {
		price := ticker(cfg.Symbol, i)
		qty := 0.01
		if !allow(qty*price, equity, cfg.MaxPos) {
			continue
		}
		equity -= qty * price * 0.0008
		fills++
	}
	return Report{Equity: equity, Fills: fills, Bars: bars}
}
