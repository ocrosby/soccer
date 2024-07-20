// internal/location-service/database/repository/country_repository.go
package repository

import (
	"context"
	"database/sql"
	"github.com/ocrosby/soccer/internal/location-service/database/models"
	"log"
)

type CountryRepositoryInterface interface {
	Create(ctx context.Context, country models.Country) error
	FindById(ctx context.Context, id string) (*models.Country, error)
	FindByCode(ctx context.Context, code string) (*models.Country, error)
	FindAll(ctx context.Context) ([]models.Country, error)
	Update(ctx context.Context, country models.Country) error
	Delete(ctx context.Context, code string) error
}

type CountryRepository struct {
	db *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (r *CountryRepository) Create(ctx context.Context, country models.Country) error {
	query := `INSERT INTO countries (name, code) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, country.Name, country.Code)
	return err
}

func (r *CountryRepository) FindById(ctx context.Context, id string) (*models.Country, error) {
	query := `SELECT id, name, code FROM countries WHERE id = $1`
	var country models.Country
	err := r.db.QueryRowContext(ctx, query, id).Scan(&country.ID, &country.Name, &country.Code)
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *CountryRepository) FindByCode(ctx context.Context, code string) (*models.Country, error) {
	query := `SELECT id, name, code FROM countries WHERE code = $1`
	var country models.Country
	err := r.db.QueryRowContext(ctx, query, code).Scan(&country.ID, &country.Name, &country.Code)
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *CountryRepository) FindAll(ctx context.Context) ([]models.Country, error) {
	var (
		err       error
		countries []models.Country
	)

	query := `SELECT id, name, code FROM countries`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}(rows)

	for rows.Next() {
		var country models.Country
		if err = rows.Scan(&country.ID, &country.Name, &country.Code); err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}
	return countries, nil
}

func (r *CountryRepository) Update(ctx context.Context, country models.Country) error {
	query := `UPDATE countries SET name = $1 WHERE code = $2`
	_, err := r.db.ExecContext(ctx, query, country.Name, country.Code)
	return err
}

func (r *CountryRepository) Delete(ctx context.Context, code string) error {
	query := `DELETE FROM countries WHERE code = $1`
	_, err := r.db.ExecContext(ctx, query, code)
	return err
}
