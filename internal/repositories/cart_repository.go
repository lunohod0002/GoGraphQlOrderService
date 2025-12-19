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
func (r *CartRepository) Create(user_id int) (*model.Cart, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("carts").Columns("user_id").
		Values(user_id).Suffix("RETURNING id")
	q, args, _ := query.ToSql()
	var id int32
	fmt.Println(q, args)
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		fmt.Println(id)

		return nil, fmt.Errorf("не удалось выполнить запрос: %v", err)
	}
	fmt.Println(id)
	return &model.Cart{ID: int(id), UserID: user_id}, nil
}
func (r *CartRepository) Get(cart_id int) *model.Cart {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	var cart model.Cart
	query := psql.Select().
		From("carts").Where(sq.Eq{"cart_id": cart_id})

	query.RunWith(r.db).QueryRow().Scan(&cart.ID, &cart.UserID, &cart.Items, cart.TotalSum, cart.Discount)

	return &cart
}
