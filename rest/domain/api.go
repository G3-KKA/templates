package domain

type GetOrderHistory interface {
	GetOrderHistory(client *Client) ([]*HistoryOrder, error)
}
type SaveOrder interface {
	SaveOrder(client *Client, order *HistoryOrder) error
}
type SaveOrderBook interface {
	SaveOrderBook(exchange_name, pair string, orderBook []*DepthOrder) error
}
type GetOrderBook interface {
	GetOrderBook(exchange_name, pair string) ([]*DepthOrder, error)
}
type Repository interface {
	GetOrderBook
	GetOrderHistory
	SaveOrder
	SaveOrderBook
}
