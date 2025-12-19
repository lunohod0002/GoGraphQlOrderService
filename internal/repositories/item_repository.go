package repositories

import (
	"OzonOrderService/graph/model"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) AddItem(itemInput *model.ItemUpdateInput) (*model.Item, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("items").Columns("cart_id", "product_id", "quantity").
		Values(itemInput.CartID, itemInput.ProductID, itemInput.Quantity).Suffix("RETURNING id")
	q, args, _ := query.ToSql()
	var id int32
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос: %v", err)
	}

	return &model.Item{ID: int(id), ProductID: itemInput.ProductID, Quantity: itemInput.Quantity}, nil
}
func (r *ItemRepository) GetItem(cart_id int, product_id int) (*model.Item, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	var item model.Item
	query := psql.Select().
		From("items").Where(sq.And{sq.Eq{"cart_id": cart_id}, sq.Eq{"product_id": product_id}})

	query.RunWith(r.db).QueryRow().Scan(&item.ID, &item.ProductID, &item.CartID, item.Quantity)

	return &item, nil
}
func (r *ItemRepository) DeleteItem(item_id int) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Delete("items").Where(sq.Eq{"item_id": item_id})

	query.RunWith(r.db).Exec()

}
func (r *ItemRepository) UpdateItem(item_id int, quantity int) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Update("items").Set("quantity", quantity).Where(sq.Eq{"item_id": item_id})
	query.RunWith(r.db).Exec()

}
