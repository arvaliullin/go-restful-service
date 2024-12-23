package repository

import (
	"database/sql"
	"github.com/arvaliullin/go-restful-service/internal/domain"
)

type GlossaryRepository struct {
	Db *sql.DB
}

func NewGlossaryRepository(db *sql.DB) *GlossaryRepository {
	return &GlossaryRepository{Db: db}
}

func (r *GlossaryRepository) GetAll() ([]domain.GlossaryTerm, error) {
	query := "SELECT id, term, description FROM glossary"
	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var glossaryTerms []domain.GlossaryTerm
	for rows.Next() {
		var term domain.GlossaryTerm
		if err := rows.Scan(&term.ID, &term.Term, &term.Description); err != nil {
			return nil, err
		}
		glossaryTerms = append(glossaryTerms, term)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return glossaryTerms, nil
}

func (r *GlossaryRepository) GetByTerm(term string) (*domain.GlossaryTerm, error) {
	query := "SELECT id, term, description FROM glossary WHERE term = $1"
	row := r.Db.QueryRow(query, term)

	var glossaryTerm domain.GlossaryTerm
	if err := row.Scan(&glossaryTerm.ID, &glossaryTerm.Term, &glossaryTerm.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // return nil if no rows are found
		}
		return nil, err
	}

	return &glossaryTerm, nil
}

func (r *GlossaryRepository) Create(term domain.GlossaryTerm) error {
	query := "INSERT INTO glossary (term, description) VALUES ($1, $2)"
	_, err := r.Db.Exec(query, term.Term, term.Description)
	return err
}

func (r *GlossaryRepository) Update(term domain.GlossaryTerm) error {
	query := "UPDATE glossary SET description = $1 WHERE term = $2"
	result, err := r.Db.Exec(query, term.Description, term.Term)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *GlossaryRepository) Delete(term string) error {
	query := "DELETE FROM glossary WHERE term = $1"
	result, err := r.Db.Exec(query, term)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
