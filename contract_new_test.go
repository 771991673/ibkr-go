package ibkr_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ThomasMarcelis/ibkr-go/v2"
)

func TestForex(t *testing.T) {
	t.Parallel()

	contract, err := ibkr.Forex("EURUSD")
	if err != nil {
		t.Fatalf("Forex(EURUSD) error = %v", err)
	}
	want := ibkr.Contract{
		Symbol:   "EUR",
		SecType:  ibkr.SecTypeForex,
		Exchange: "IDEALPRO",
		Currency: "USD",
	}
	if !reflect.DeepEqual(contract, want) {
		t.Fatalf("Forex(EURUSD) = %+v, want %+v", contract, want)
	}
}

func TestStockHK(t *testing.T) {
	t.Parallel()
	got := ibkr.StockHK("700")
	want := ibkr.Contract{
		Symbol:          "700",
		SecType:         ibkr.SecTypeStock,
		Exchange:        "SMART",
		PrimaryExchange: "SEHK",
		Currency:        "HKD",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("StockHK(700) = %+v, want %+v", got, want)
	}
}

func TestStockCN(t *testing.T) {
	t.Parallel()
	got := ibkr.StockCN("600519", "SSE")
	if got.Currency != "CNH" || got.PrimaryExchange != "SSE" {
		t.Fatalf("StockCN = %+v", got)
	}
}

func TestCrypto(t *testing.T) {
	t.Parallel()
	got := ibkr.Crypto("BTC")
	want := ibkr.Contract{
		Symbol:   "BTC",
		SecType:  ibkr.SecTypeCrypto,
		Exchange: "PAXOS",
		Currency: "USD",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Crypto(BTC) = %+v, want %+v", got, want)
	}
}

func TestForexRejectsInvalidPair(t *testing.T) {
	t.Parallel()

	for _, pair := range []string{
		"EURUS",
		"EURUSDX",
		"eurusd",
		"EUR_SD",
		"EURÉD",
	} {
		t.Run(pair, func(t *testing.T) {
			t.Parallel()

			contract, err := ibkr.Forex(pair)
			if !reflect.DeepEqual(contract, ibkr.Contract{}) {
				t.Fatalf("Forex(%q) contract = %+v, want zero Contract", pair, contract)
			}
			validation, ok := errors.AsType[*ibkr.ValidationError](err)
			if !ok {
				t.Fatalf("Forex(%q) error = %v, want *ValidationError", pair, err)
			}
			if validation.Field != "Pair" || validation.Value != pair ||
				validation.Message != "must be exactly six uppercase ASCII letters" {
				t.Fatalf("Forex(%q) error = %+v, want exact Pair validation", pair, validation)
			}
		})
	}
}
