package repositories

import (
	"OzonOrderService/graph/model"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *model.Product) (*model.Product, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("products").Columns("name", "price").
		Values(product.Name, product.Price).Suffix("RETURNING id")
	q, args, _ := query.ToSql()
	var id int32
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос: %v", err)
	}

	return &model.Product{ID: id, Price: product.Price, Name: product.Name}, nil
}
