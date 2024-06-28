package handlers

import (
	"net/http"
	"yet-again-templates/rest/domain"
)

func (hh *HandlerHolder) NewGetOrderBookState() *GetOrderBookState {
	return &GetOrderBookState{logic: hh.repo}
}

// Free to embed any fileds via constructor changewhich
type GetOrderBookState struct {
	// Only handler specific logic
	logic domain.GetOrderBook
}

func (state *GetOrderBookState) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	state.logic.GetOrderBook("binance", "BTCUSDT")
}
