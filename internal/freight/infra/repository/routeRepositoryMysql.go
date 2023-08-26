package repository

import (
	"database/sql"
	"fmt"

	"github.com/joaomarcosbc/imersaoFC-ms/internal/freight/entity"
)

type RouteRepositoryMySql struct {
	db *sql.DB
}

func NewRouteRepositoryMySql(db *sql.DB) *RouteRepositoryMySql {
	return &RouteRepositoryMySql{
		db: db,
	}
}

func (r *RouteRepositoryMySql) Create(route *entity.Route) error {
	sqlSmt := "INSERT INTO routes (id, name, distance, status, freight_price) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(sqlSmt, route.ID, route.Name, route.Distance, route.Status, route.FreightPrice)
	if err != nil {
		return err
	}

	return nil
}

func (r *RouteRepositoryMySql) FindByID(id string) (*entity.Route, error) {
	sqlSmt := "SELECT * FROM routes WHERE id = ?"
	row := r.db.QueryRow(sqlSmt, id)

	var starteAt, finishedAt sql.NullTime

	var route entity.Route

	err := row.Scan(
		&route.ID,
		&route.Name,
		&route.Distance,
		&route.Status,
		&route.FreightPrice,
		&starteAt,
		&finishedAt,
	)

	if err != nil {
		return nil, err
	}

	if starteAt.Valid {
		route.StartedAt = starteAt.Time
	}

	if finishedAt.Valid {
		route.FinishedAt = finishedAt.Time
	}

	return &route, nil
}

func (r *RouteRepositoryMySql) Update(route *entity.Route) error {
	startedAt := route.StartedAt.Format("2006-01-02 15:04:05")
	fmt.Println(startedAt)
	finishedAt := route.FinishedAt.Format("2006-01-02 15:04:05")

	sql := "UPDATE routes SET status = ?, freight_price = ?, started_at=?, finished_at = ? WHERE id = ?"

	_, err := r.db.Exec(sql, route.Status, route.FreightPrice, startedAt, finishedAt, route.ID)
	if err != nil {
		return err
	}
	return nil
}
