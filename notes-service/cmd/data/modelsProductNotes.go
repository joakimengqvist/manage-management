package data

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (n *ProductNote) GetProductNoteById(id string) (*ProductNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, product_id, title, note, created_at, updated_at from product_notes where id = $1`

	var note ProductNote
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&note.ID,
		&note.AuthorId,
		&note.AuthorName,
		&note.AuthorEmail,
		&note.ProductId,
		&note.Title,
		&note.Note,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (n *ProductNote) GetProductNotesByProductId(id string) ([]*ProductNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, product_id, title, note, created_at, updated_at from product_notes where product_id = $1`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Print("rows", err)
		return nil, err
	}
	defer rows.Close()

	var notes []*ProductNote

	for rows.Next() {
		var note ProductNote
		err := rows.Scan(
			&note.ID,
			&note.AuthorId,
			&note.AuthorName,
			&note.AuthorEmail,
			&note.ProductId,
			&note.Title,
			&note.Note,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		notes = append(notes, &note)
	}

	log.Println("notes appended", notes)

	return notes, nil
}

func (n *ProductNote) GetProductNotesByAuthorId(id string) ([]*ProductNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, product_id, title, note, created_at, updated_at from product_notes where author_id = $1 order by updated_at desc`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*ProductNote

	for rows.Next() {
		var note ProductNote
		err := rows.Scan(
			&note.ID,
			&note.AuthorId,
			&note.AuthorName,
			&note.AuthorEmail,
			&note.ProductId,
			&note.Title,
			&note.Note,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		notes = append(notes, &note)
	}

	return notes, nil
}

func (n *ProductNote) UpdateProductNote() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update product_notes set
		author_id = $1,
		author_name = $2,
		author_email = $3,
		product_id = $4,
		title = $5,
		note = $6,
		updated_at = $7
		where id = $8
	`

	_, err := db.ExecContext(ctx, stmt,
		n.AuthorId,
		n.AuthorName,
		n.AuthorEmail,
		n.ProductId,
		n.Title,
		n.Note,
		time.Now(),
		n.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteProductNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from product_notes where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductNote) InsertProductNote(note ProductNote) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `insert into product_notes (author_id, author_name, author_email, product_id, title, note, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	err := db.QueryRowContext(ctx, stmt,
		note.AuthorId,
		note.AuthorName,
		note.AuthorEmail,
		note.ProductId,
		note.Title,
		note.Note,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}
