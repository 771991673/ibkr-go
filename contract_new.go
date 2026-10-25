package ibkr

import (
	"github.com/shopspring/decimal"
)

// Stock returns a SMART-routed, USD-denominated stock or ETF [Contract] for
// symbol. It fills only Symbol, SecType, Exchange, and Currency; set
// PrimaryExchange or Currency explicitly for non-US listings.
func Stock(symbol string) Contract {
	return Contract{
		Symbol:   symbol,
		SecType:  SecTypeStock,
		Exchange: "SMART",
		Currency: "USD",
	}
}

// StockHK returns a SMART-routed Hong Kong stock [Contract] (SEHK / HKD).
// Numeric tickers such as "700" or "00700" are accepted; leading zeros are
// not required by the Gateway when PrimaryExchange is SEHK.
func StockHK(symbol string) Contract {
	return Contract{
		Symbol:          symbol,
		SecType:         SecTypeStock,
		Exchange:        "SMART",
		PrimaryExchange: "SEHK",
		Currency:        "HKD",
	}
}

// StockCN returns a SMART-routed mainland China A-share [Contract] (CNH).
// primaryExchange should be "SSE" (Shanghai) or "SZSE" (Shenzhen).
func StockCN(symbol, primaryExchange string) Contract {
	if primaryExchange == "" {
		primaryExchange = "SSE"
	}
	return Contract{
		Symbol:          symbol,
		SecType:         SecTypeStock,
		Exchange:        "SMART",
		PrimaryExchange: primaryExchange,
		Currency:        "CNH",
	}
}

// Crypto returns a PAXOS-routed cryptocurrency [Contract] for symbol
// (e.g. "BTC", "ETH") quoted in USD.
func Crypto(symbol string) Contract {
	return Contract{
		Symbol:   symbol,
		SecType:  SecTypeCrypto,
		Exchange: "PAXOS",
		Currency: "USD",
	}
}

// Forex returns an IDEALPRO forex pair [Contract] from a six-letter pair code
// such as "EURUSD" (base EUR, quote/currency USD). pair must be exactly six
// uppercase ASCII letters; invalid pairs return a [*ValidationError].
func Forex(pair string) (Contract, error) {
	if len(pair) != 6 {
		return Contract{}, &ValidationError{Field: "Pair", Value: pair, Message: "must be exactly six uppercase ASCII letters"}
	}
	for _, c := range pair {
		if c < 'A' || c > 'Z' {
			return Contract{}, &ValidationError{Field: "Pair", Value: pair, Message: "must be exactly six uppercase ASCII letters"}
		}
	}
	return Contract{
		Symbol:   pair[:3],
		Currency: pair[3:],
		SecType:  SecTypeForex,
		Exchange: "IDEALPRO",
	}, nil
}

// OptionContract returns a SMART-routed, USD-denominated, 100-multiplier
// option [Contract]. lastTradeDateOrContractMonth follows the IBKR YYYYMMDD
// (or YYYYMM) convention.
//
// Named OptionContract, not Option, because [Option] is already the
// functional-options type passed to [DialContext].
func OptionContract(symbol, lastTradeDateOrContractMonth string, strike decimal.Decimal, right Right) Contract {
	return Contract{
		Symbol:     symbol,
		SecType:    SecTypeOption,
		Expiry:     lastTradeDateOrContractMonth,
		Strike:     new(strike),
		Right:      right,
		Multiplier: "100",
		Exchange:   "SMART",
		Currency:   "USD",
	}
}

// Future returns a USD-denominated future [Contract] on the given exchange.
// lastTradeDateOrContractMonth follows the IBKR YYYYMMDD (or YYYYMM)
// convention.
func Future(symbol, lastTradeDateOrContractMonth, exchange string) Contract {
	return Contract{
		Symbol:   symbol,
		SecType:  SecTypeFuture,
		Expiry:   lastTradeDateOrContractMonth,
		Exchange: exchange,
		Currency: "USD",
	}
}
