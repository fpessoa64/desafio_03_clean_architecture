package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/fpessoa64/desafio03_clean_arch/internal/entity"
)

type OrderRepositoryMySQL struct {
	db *sql.DB
}

func NewOrderRepositoryMySQL(db *sql.DB) *OrderRepositoryMySQL {
	return &OrderRepositoryMySQL{db: db}
}

func (r *OrderRepositoryMySQL) Create(ctx context.Context, o *entity.Order) error {
	query := `INSERT INTO orders (name, amount, status, created_at) VALUES (?, ?, ?, ?)`
	now := time.Now()
	result, err := r.db.ExecContext(ctx, query, o.Name, o.Amount, o.Status, now)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	o.ID = id
	o.CreatedAt = now
	return nil
}

func (r *OrderRepositoryMySQL) List(ctx context.Context) ([]entity.Order, error) {
	query := `SELECT id, name, amount, status, created_at FROM orders ORDER BY id DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]entity.Order, 0)
	for rows.Next() {
		var o entity.Order
		if err := rows.Scan(&o.ID, &o.Name, &o.Amount, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, rows.Err()
}
