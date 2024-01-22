package data

import (
	"context"
	"log"
	"time"
)

func (u *Project) GetAllProjects() ([]*PostgresProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, status, notes, sub_projects, created_at, created_by, updated_at, updated_by
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
			&project.SubProjects,
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

func GetProjectsByIds(ids string) ([]*PostgresProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select id, name, status, notes, sub_projects, created_at, created_by, updated_at, updated_by
        from projects
        where id = ANY($1)
		order by name
    `

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
			&project.SubProjects,
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

func (u *Project) GetProjectById(id string) (*PostgresProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, status, notes, sub_projects, created_at, created_by, updated_at, updated_by from projects where id = $1`

	var project PostgresProject
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&project.ID,
		&project.Name,
		&project.Status,
		&project.Notes,
		&project.SubProjects,
		&project.CreatedAt,
		&project.CreatedBy,
		&project.UpdatedAt,
		&project.UpdatedBy,
	)

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (u *PostgresProject) UpdateProject(updatedByUserId string) error {
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
		updatedByUserId,
		u.ID,
	)

	if err != nil {
		log.Println("Error updating project", err)
		return err
	}

	return nil
}

func (u *PostgresProject) DeleteProject() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from projects where id = $1`

	_, err := db.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *Project) InsertProject(project NewProject, createdByUserId string) (string, error) {
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

	stmt := `update projects set notes = array(SELECT DISTINCT unnest(array_append(notes, $1))) where id = $2`

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
