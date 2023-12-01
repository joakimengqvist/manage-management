package data

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (n *SubProjectNote) GetSubProjectNoteById(id string) (*SubProjectNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, sub_project_id, title, note, created_at, updated_at from sub_project_notes where id = $1`

	var note SubProjectNote
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&note.ID,
		&note.AuthorId,
		&note.AuthorName,
		&note.AuthorEmail,
		&note.SubProjectId,
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

func (n *SubProjectNote) GetSubProjectNotesBySubProjectId(id string) ([]*SubProjectNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, sub_project_id, title, note, created_at, updated_at from sub_project_notes where sub_project_id = $1`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Print("rows", err)
		return nil, err
	}
	defer rows.Close()

	var notes []*SubProjectNote

	for rows.Next() {
		var note SubProjectNote
		err := rows.Scan(
			&note.ID,
			&note.AuthorId,
			&note.AuthorName,
			&note.AuthorEmail,
			&note.SubProjectId,
			&note.Title,
			&note.Note,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			fmt.Println("Error scanning", err)
			return nil, err
		}

		notes = append(notes, &note)
	}

	return notes, nil
}

func (n *SubProjectNote) GetSubProjectNotesByAuthorId(id string) ([]*SubProjectNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, author_id, author_name, author_email, sub_project_id, title, note, created_at, updated_at from sub_project_notes where author_id = $1 order by updated_at desc`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*SubProjectNote

	for rows.Next() {
		var note SubProjectNote
		err := rows.Scan(
			&note.ID,
			&note.AuthorId,
			&note.AuthorName,
			&note.AuthorEmail,
			&note.SubProjectId,
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

func (n *SubProjectNote) UpdateSubProjectNote() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_project_notes set
		author_id = $1,
		author_name = $2,
		author_email = $3,
		sub_project_id = $4,
		title = $5,
		note = $6,
		updated_at = $7
		where id = $8
	`

	_, err := db.ExecContext(ctx, stmt,
		n.AuthorId,
		n.AuthorName,
		n.AuthorEmail,
		n.SubProjectId,
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

func DeleteSubProjectNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from sub_project_notes where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *SubProjectNote) InsertSubProjectNote(note SubProjectNote) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `insert into sub_project_notes (author_id, author_name, author_email, sub_project_id, title, note, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	err := db.QueryRowContext(ctx, stmt,
		note.AuthorId,
		note.AuthorName,
		note.AuthorEmail,
		note.SubProjectId,
		note.Title,
		note.Note,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		fmt.Println("Error inserting sub project note", err)
		return "", err
	}

	return newID, nil
}
