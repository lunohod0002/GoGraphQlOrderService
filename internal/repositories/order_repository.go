package repositories

import (
	"OzonOrderService/graph/model"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(cart *model.Cart, name string) (*model.Order, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("orders").Columns("cart_id", "name").
		Values(cart.ID, name)
	q, args, _ := query.ToSql()
	var id int32
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос: %v", err)
	}

	return &model.Order{ID: int(id), Cart: cart, Status: "CREATED", CreatedAt: time.Now().String()}, nil
}

func (r *OrderRepository) GetAll(user_id int) ([]*model.Order, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	orders := psql.Select("*").
		From("orders").Where(sq.Eq{"user_id": user_id})

	rows, sqlerr := orders.RunWith(r.db).Query()
	if sqlerr != nil {
		panic(fmt.Sprintf("QueryRow failed: %v", sqlerr))
	}
	var ordersSl []*model.Order

	for rows.Next() {
		var order model.Order
		rows.Scan(&order.Cart, &order.Status, &order.CreatedAt, &order.ID)
		ordersSl = append(ordersSl, &order)
	}
	return ordersSl, nil
}
