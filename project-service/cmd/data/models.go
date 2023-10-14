package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Project: Project{},
	}
}

type Models struct {
	Project Project
}

type NewProject struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Notes     []string  `json:"notes"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type PostgresProject struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Notes     string    `json:"notes"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

func (u *Project) GetAll() ([]*PostgresProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, status, notes, created_at, created_by, updated_at, updated_by
	from projects order by name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*PostgresProject

	for rows.Next() {
		var project PostgresProject
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Status,
			&project.Notes,
			&project.CreatedAt,
			&project.CreatedBy,
			&project.UpdatedAt,
			&project.UpdatedBy,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		projects = append(projects, &project)
	}

	return projects, nil
}

func (u *Project) GetProjectByName(name string) (*Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, status, notes, created_at, created_by, updated_at, updated_by from projects where name = $1`

	var project Project
	row := db.QueryRowContext(ctx, query, name)

	err := row.Scan(
		&project.ID,
		&project.Name,
		&project.Status,
		&project.Notes,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (u *Project) GetProjectById(id string) (*PostgresProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, status, notes, created_at, created_by, updated_at, updated_by from projects where id = $1`

	var project PostgresProject
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&project.ID,
		&project.Name,
		&project.Status,
		&project.Notes,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (u *PostgresProject) Update(createdByUserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update projects set
		name = $1,
		status = $2,
		updated_at = $3,
		updated_by = $4
		where id = $5
	`

	_, err := db.ExecContext(ctx, stmt,
		u.Name,
		u.Status,
		time.Now(),
		createdByUserId,
		u.ID,
	)

	if err != nil {
		fmt.Println("Error updating project", err)
		return err
	}

	return nil
}

func (u *PostgresProject) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from projects where id = $1`

	_, err := db.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *Project) Insert(project NewProject, createdByUserId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `insert into projects (name, status, created_at, created_by, updated_at, updated_by)
		values ($1, $2, $3, $4 $5 $6) returning id`

	err := db.QueryRowContext(ctx, stmt,
		project.Name,
		project.Status,
		time.Now(),
		createdByUserId,
		time.Now(),
		createdByUserId,
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func AppendProjectNote(projectId string, noteId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update projects set notes = array_append(notes, $1) where id = $2`

	_, err := db.ExecContext(ctx, stmt, noteId, projectId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProjectNote(projectId string, noteId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update projects set notes = array_remove(notes, $1) where id = $2`

	_, err := db.ExecContext(ctx, stmt, noteId, projectId)
	if err != nil {
		return err
	}

	return nil
}
