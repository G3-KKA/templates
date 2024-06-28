package handlers

import (
	"net/http"
	"yet-again-templates/rest/domain"
)

type SaveOrderBookState struct {
	// Only handler specific logic
	logic domain.GetOrderBook
}

func (state *SaveOrderBookState) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	state.logic.GetOrderBook("binance", "BTCUSDT")
}
