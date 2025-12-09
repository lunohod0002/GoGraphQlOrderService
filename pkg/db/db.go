package main

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Animal struct {
	Id   uint
	Name string
	Age  uint
}

func initDb() (*sql.DB, error) {
	connString :=
		"postgresql://postgres:postgres@localhost:5432/postgres"
	db, err := sql.Open("pgx", connString)
	if err != nil {
		panic("Unable to connect to database")
	}
	defer db.Close()
	return db, nil
	//// a little magic to tell squirrel it's postgres
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	animals := psql.Select("name", "age").
		From("animals").Where(sq.Gt{"age": 4}).Limit(2)

	rows, sqlerr := animals.RunWith(db).Query()
	if sqlerr != nil {
		panic(fmt.Sprintf("QueryRow failed: %v", sqlerr))
	}
	var animalsSl []Animal

	for rows.Next() {
		var animal Animal
		rows.Scan(&animal.Name, &animal.Age)
		animalsSl = append(animalsSl, animal)
	}
	fmt.Println(animalsSl)
	query := psql.
		Insert("animals").Columns("name", "age", "id").
		Values("moeascascssssc2", 13, 14)
	query.RunWith(db).Exec()

	////query2 := psql.
	////	Delete("animals").
	////	Where(sq.Eq{"name": "Tima"})
	////query2.RunWith(db).Exec()
	//query3 := psql.
	//	Update("animals").Set("age", 14).Where(sq.Eq{"name": "Pushok"})
	//query3.RunWith(db).Exec()
	//var newAnimal Animal
	//animal := psql.Select("id", "name", "age").
	//	From("animals").Where(sq.Eq{"name": "Pushok"})
	//
	//animal.RunWith(db).QueryRow().Scan(&newAnimal.Id, &newAnimal.Name, &newAnimal.Age)
	//if sqlerr != nil {
	//	panic(fmt.Sprintf("QueryRow failed: %v", sqlerr))
	//}
	//fmt.Println(newAnimal)
	////query2 := psql.
	////	Delete("animals").
	////	Where(sq.Eq{"name": "Tima"})
	////query2.RunWith(db).Exec()
	////tx, err := dbpool.Begin(context.Background())
	////if err != nil {
	////	log.Fatal("Qwqwdw")
	////}
	////defer tx.Rollback(context.Background())
	////
	////_, err = dbpool.Exec(context.Background(), "Insert into animals (id,name,age) VALUES ($1,$2,$3) ON CONFLICT DO NOTHING", 4, "Duke", 3)
	////if err != nil {
	////	log.Fatal("Не удалось выполнить запрос: %v\n", err)
	////}
	////_, err = dbpool.Exec(context.Background(), "Insert into animals (id,name,age) VALUES ($1,$2,$3) ON CONFLICT DO NOTHING", 5, "Pushok", 3)
	////if err != nil {
	////	log.Fatal("Не удалось выполнить запрос: %v\n", err)
	////}
	////_, err = dbpool.Exec(context.Background(), "Insert into animals (id,name,age) VALUES ($1,$2,$3) ON CONFLICT DO NOTHING", 6, "Bobik", 3)
	////if err != nil {
	////	log.Fatal("Не удалось выполнить запрос: %v\n", err)
	////}
	//
	////var animal Animal
	////err = dbpool.QueryRow(context.Background(), "SELECT Animals.id,Animals.age,Animals.name FROM Animals WHERE Animals.id = $1", 1).Scan(&animal.Id, &animal.Age, &animal.Name)
	////if err != nil {
	////	log.Fatal("Не удалось выполнить запрос: %v\n", err)
	////}
	////fmt.Println(animal)
	////
	////animalId := 1
	////newAge := 10
	////_, err = dbpool.Exec(context.Background(), "Update Animals SET age = $1 WHERE id = $2", newAge, animalId)
	////if err != nil {
	////	log.Fatal("Не удалось выполнить запрос: %v\n", err)
	////}
	////fmt.Println("Updated")
	////
	////animalId = 1
	////_, err = dbpool.Exec(context.Background(), "DELETE FROM Animals  WHERE id = $1", animalId)
	////if err != nil {
	////	log.Fatal("Не удалось выполнить запрос: %v\n", err)
	////}
	////fmt.Println("Deleted")
	return db, nil
}
