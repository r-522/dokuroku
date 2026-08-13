package note

import (
	"context"
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("note not found")

type Repository interface {
	Create(ctx context.Context, title, content string, urls []string) (Note, error)
	List(ctx context.Context, query string) ([]Note, error)
	Update(ctx context.Context, id int64, title, content string, urls []string) (Note, error)
	Delete(ctx context.Context, id int64) error
}

type SQLRepository struct{ DB *sql.DB }

func (r SQLRepository) Create(ctx context.Context, title, content string, urls []string) (Note, error) {
	row := r.DB.QueryRowContext(ctx, `INSERT INTO BOOK_notes (title, content, urls) VALUES ($1, $2, $3) RETURNING id, title, content, urls, created_at`, title, content, stringArray(urls))
	return scanNote(row)
}

func (r SQLRepository) List(ctx context.Context, query string) ([]Note, error) {
	pattern := "%" + query + "%"
	rows, err := r.DB.QueryContext(ctx, `SELECT id, title, content, urls, created_at FROM BOOK_notes WHERE $1 = '' OR title ILIKE $2 OR content ILIKE $2 ORDER BY created_at DESC`, query, pattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	notes := []Note{}
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, (*stringArray)(&n.URLs), &n.CreatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	return notes, rows.Err()
}

func (r SQLRepository) Update(ctx context.Context, id int64, title, content string, urls []string) (Note, error) {
	row := r.DB.QueryRowContext(ctx, `UPDATE BOOK_notes SET title = $2, content = $3, urls = $4 WHERE id = $1 RETURNING id, title, content, urls, created_at`, id, title, content, stringArray(urls))
	return scanNote(row)
}

func (r SQLRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.DB.ExecContext(ctx, `DELETE FROM BOOK_notes WHERE id = $1`, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

type scanner interface{ Scan(dest ...any) error }

func scanNote(row scanner) (Note, error) {
	var n Note
	if err := row.Scan(&n.ID, &n.Title, &n.Content, (*stringArray)(&n.URLs), &n.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Note{}, ErrNotFound
		}
		return Note{}, err
	}
	return n, nil
}
