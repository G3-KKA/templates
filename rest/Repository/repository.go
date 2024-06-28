package clientrepo

import (
	"database/sql"
	"yet-again-templates/rest/domain"
)

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

type ClientRepository struct {
	db *sql.DB
}

func (repo *ClientRepository) GetOrderBook(
	exchange_name, pair string) ([]*domain.DepthOrder, error) {
	return nil, nil
}
func (repo *ClientRepository) GetOrderHistory(
	client *domain.Client) ([]*domain.HistoryOrder, error) {
	return nil, nil
}
func (repo *ClientRepository) SaveOrder(
	client *domain.Client, order *domain.HistoryOrder) error {
	return nil
}
func (repo *ClientRepository) SaveOrderBook(
	exchange_name, pair string, orderBook []*domain.DepthOrder) error {
	return nil
}
