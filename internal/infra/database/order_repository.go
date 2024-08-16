package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/allurco/desafio-cleanarch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) List(page, limit int, sort string) []entity.Order {
	results := []entity.Order{}

	stmt := fmt.Sprintf("SELECT id, price, tax, final_price FROM orders ORDER BY id %s LIMIT ? OFFSET ?", sort)

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	rows, err := r.Db.Query(stmt, limit, offset)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var result entity.Order
		err = rows.Scan(&result.ID, &result.Price, &result.Tax, &result.FinalPrice)
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}

	return results
}
