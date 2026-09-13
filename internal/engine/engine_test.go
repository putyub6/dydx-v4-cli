package engine

import "testing"

func TestTickerPositive(t *testing.T) {
	if ticker("BTCUSDT", 1) <= 0 {
		t.Fatal("ticker")
	}
}

func TestAllowRejects(t *testing.T) {
	if allow(50, 100, 0.1) {
		t.Fatal("should reject")
	}
}

func TestAllowSmall(t *testing.T) {
	if !allow(10, 100, 0.5) {
		t.Fatal("should allow")
	}
}

func TestBacktest(t *testing.T) {
	r := Backtest(16)
	if r.Bars != 16 || r.Equity <= 0 {
		t.Fatalf("%+v", r)
	}
}

func TestDefault(t *testing.T) {
	if Default().Symbol == "" {
		t.Fatal("empty symbol")
	}
}
