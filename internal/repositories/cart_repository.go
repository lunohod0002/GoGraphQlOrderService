package repositories

import (
	"OzonOrderService/graph/model"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddItem(itemInput *model.ItemAddInput) (*model.Item, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("carts").Columns("user_id", "product_id", "quantity").
		Values(itemInput.UserID, itemInput.ProductID, itemInput.Quantity).Suffix("RETURNING id")
	q, args, _ := query.ToSql()
	var id int32
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос: %v", err)
	}

	return &model.Item{ID: id, ProductID: itemInput.ProductID, Quantity: itemInput.Quantity}, nil
}

func (r *CartRepository) Create(user_id int) (*model.Cart, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("carts").Columns("user_id").
		Values(user_id)
	q, args, _ := query.ToSql()
	var id int32
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос: %v", err)
	}

	return &model.Cart{ID: id}, nil
}
func (r *CartRepository) RemoveItem(itemInput *model.ItemRemoveInput) (*model.Item, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Delete("carts").Where("user_id", "product_id", "quantity")
	q, _, _ := query.ToSql()
	var id int32
	err := r.db.QueryRow(q).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос: %v", err)
	}
	//TODO: добавить подсчет нового кол-ва
	return &model.Item{ID: id, ProductID: itemInput.ProductID, Quantity: 1}, nil
}
