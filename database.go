package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func InitDB(cfg Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("БД не відповідає:", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("✅ БД підключена, міграції виконані")
	return db
}

type PostgresRepo struct{ db *sql.DB }

func (r *PostgresRepo) Create(g Guitar) error {
	_, err := r.db.Exec(`INSERT INTO guitars (id, brand, model, strings_count, is_electric) VALUES ($1, $2, $3, $4, $5)`,
		g.ID, g.Brand, g.Model, g.StringsCount, g.IsElectric)
	return err
}

func (r *PostgresRepo) GetAll() ([]Guitar, error) {
	rows, err := r.db.Query(`SELECT id, brand, model, strings_count, is_electric FROM guitars`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guitars []Guitar
	for rows.Next() {
		var g Guitar
		if err := rows.Scan(&g.ID, &g.Brand, &g.Model, &g.StringsCount, &g.IsElectric); err != nil {
			return nil, err
		}
		guitars = append(guitars, g)
	}
	return guitars, nil
}

func (r *PostgresRepo) GetByID(id string) (Guitar, error) {
	var g Guitar
	err := r.db.QueryRow(`SELECT id, brand, model, strings_count, is_electric FROM guitars WHERE id=$1`, id).
		Scan(&g.ID, &g.Brand, &g.Model, &g.StringsCount, &g.IsElectric)
	if err == sql.ErrNoRows {
		return g, fmt.Errorf("не знайдено")
	}
	return g, err
}

func (r *PostgresRepo) Update(id string, g Guitar) error {
	_, err := r.db.Exec(`UPDATE guitars SET brand=$1, model=$2, strings_count=$3, is_electric=$4 WHERE id=$5`,
		g.Brand, g.Model, g.StringsCount, g.IsElectric, id)
	return err
}

func (r *PostgresRepo) PatchElectric(id string, isElectric bool) error {
	_, err := r.db.Exec(`UPDATE guitars SET is_electric=$1 WHERE id=$2`, isElectric, id)
	return err
}

func (r *PostgresRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM guitars WHERE id=$1`, id)
	return err
}
