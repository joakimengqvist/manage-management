package data

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (n *InvoiceItemNote) GetInvoiceItemNoteById(id string) (*InvoiceItemNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, invoice_item_id, title, note, created_at, updated_at from invoice_item_notes where id = $1`

	var note InvoiceItemNote
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&note.ID,
		&note.AuthorId,
		&note.AuthorName,
		&note.AuthorEmail,
		&note.InvoiceItemId,
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

func (n *InvoiceItemNote) GetInvoiceItemNotesByInvoiceItemId(id string) ([]*InvoiceItemNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, invoice_item_id, title, note, created_at, updated_at from invoice_item_notes where invoice_item_id = $1`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Print("rows", err)
		return nil, err
	}
	defer rows.Close()

	var notes []*InvoiceItemNote

	for rows.Next() {
		var note InvoiceItemNote
		err := rows.Scan(
			&note.ID,
			&note.AuthorId,
			&note.AuthorName,
			&note.AuthorEmail,
			&note.InvoiceItemId,
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

	fmt.Print("notes appended", notes)

	return notes, nil
}

func (n *InvoiceItemNote) GetInvoiceItemNotesByAuthorId(id string) ([]*InvoiceItemNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, invoice_item_id, title, note, created_at, updated_at from invoice_item_notes where author_id = $1 order by updated_at desc`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*InvoiceItemNote

	for rows.Next() {
		var note InvoiceItemNote
		err := rows.Scan(
			&note.ID,
			&note.AuthorId,
			&note.AuthorName,
			&note.AuthorEmail,
			&note.InvoiceItemId,
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

func (n *InvoiceItemNote) UpdateInvoiceItemNote() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update invoice_item_notes set
		author_id = $1,
		author_name = $2,
		author_email = $3,
		invoice_item_id = $4,
		title = $5,
		note = $6,
		updated_at = $7
		where id = $8
	`

	_, err := db.ExecContext(ctx, stmt,
		n.AuthorId,
		n.AuthorName,
		n.AuthorEmail,
		n.InvoiceItemId,
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

func DeleteInvoiceItemNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from invoice_item_notes where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *InvoiceItemNote) InsertInvoiceItemNote(note InvoiceItemNote) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `insert into invoice_item_notes (author_id, author_name, author_email, invoice_item_id, title, note, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	err := db.QueryRowContext(ctx, stmt,
		note.AuthorId,
		note.AuthorName,
		note.AuthorEmail,
		note.InvoiceItemId,
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
