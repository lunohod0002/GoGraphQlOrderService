package repositories

import (
	"OzonOrderService/graph/model"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.
		Insert("users").Columns("fio", "balance").
		Values(user.Fio, user.Balance).Suffix("RETURNING id")
	q, args, _ := query.ToSql()
	var id int32
	fmt.Println(q, args)
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		fmt.Println(id)

		return nil, fmt.Errorf("не удалось выполнить запрос: %v", err)
	}
	fmt.Println(id)
	return &model.User{ID: id, Fio: user.Fio, Balance: user.Balance}, nil
}
