package data

import (
	"context"
	"log"
	"time"
)

func (u *SubProject) GetAllSubProjects() ([]*PostgresSubProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, status, priority, start_date, due_date, estimated_duration, 
          notes, projects, created_at, created_by, updated_at, updated_by, invoices, incomes, expenses
          from sub_projects order by start_date`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subProjects []*PostgresSubProject

	for rows.Next() {
		var subProject PostgresSubProject
		err := rows.Scan(
			&subProject.ID,
			&subProject.Name,
			&subProject.Description,
			&subProject.Status,
			&subProject.Priority,
			&subProject.StartDate,
			&subProject.DueDate,
			&subProject.EstimatedDuration,
			&subProject.Notes,
			&subProject.Projects,
			&subProject.CreatedAt,
			&subProject.CreatedBy,
			&subProject.UpdatedAt,
			&subProject.UpdatedBy,
			&subProject.Invoices,
			&subProject.Incomes,
			&subProject.Expenses,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		subProjects = append(subProjects, &subProject)
	}

	return subProjects, nil
}

func (u *SubProject) GetSubProjectById(id string) (*PostgresSubProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, status, priority, start_date, due_date, estimated_duration, 
		notes, projects, created_at, created_by, updated_at, updated_by, invoices, incomes, expenses
        from sub_projects where id = $1`

	var subProject PostgresSubProject
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&subProject.ID,
		&subProject.Name,
		&subProject.Description,
		&subProject.Status,
		&subProject.Priority,
		&subProject.StartDate,
		&subProject.DueDate,
		&subProject.EstimatedDuration,
		&subProject.Notes,
		&subProject.Projects,
		&subProject.CreatedAt,
		&subProject.CreatedBy,
		&subProject.UpdatedAt,
		&subProject.UpdatedBy,
		&subProject.Invoices,
		&subProject.Incomes,
		&subProject.Expenses,
	)

	if err != nil {
		return nil, err
	}

	return &subProject, nil
}

func (p *PostgresSubProject) UpdateSubProject(updatedByUserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_projects set
		name = $1, 
		description = $2,
		status = $3,
		priority = $4,
		start_date = $5,
		due_date = $6,
		estimated_duration = $7,
		updated_at = $8,
		updated_by = $9
		where id = $10
	`

	_, err := db.ExecContext(ctx, stmt,
		p.Name,
		p.Description,
		p.Status,
		p.Priority,
		p.StartDate,
		p.DueDate,
		p.EstimatedDuration,
		time.Now(),
		updatedByUserId,
		p.ID,
	)

	if err != nil {
		log.Println("Error updating sub project", err)
		return err
	}

	return nil
}

func (u *PostgresSubProject) DeleteSubProject() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from sub_projects where id = $1`

	_, err := db.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *SubProject) InsertSubProject(project PostgresSubProject, createdByUserId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `insert into sub_projects (name, description, status, priority, start_date, due_date, estimated_duration, 
		notes, created_at, created_by, updated_at, updated_by, invoices, incomes, expenses)
		values ($1, $2, $3, $4, $5, $6, $7,	$8, $9, $10, $11, $12, $13, $14, $15) returning id`

	err := db.QueryRowContext(ctx, stmt,
		project.Name,
		project.Description,
		project.Status,
		project.Priority,
		project.StartDate,
		project.DueDate,
		project.EstimatedDuration,
		project.Notes,
		time.Now(),
		createdByUserId,
		time.Now(),
		createdByUserId,
		project.Invoices,
		project.Incomes,
		project.Expenses,
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func AppendSubProjectNote(subProjectId string, noteId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_projects set notes = array(SELECT DISTINCT unnest(array_append(notes, $1))) where id = $2`

	_, err := db.ExecContext(ctx, stmt, noteId, subProjectId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubProjectNote(subProjectId string, noteId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Println("Deleting note", noteId, "from sub project", subProjectId)

	stmt := `update sub_projects set notes = array_remove(notes, $1) where id = $2`

	_, err := db.ExecContext(ctx, stmt, noteId, subProjectId)
	if err != nil {
		log.Println("Error deleting note from sub project", err)
		return err
	}

	return nil
}

func AppendSubProjectInvoice(subProjectId string, invoiceId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_projects set invoices = array(SELECT DISTINCT unnest(array_append(invoices, $1))) where id = $2`

	_, err := db.ExecContext(ctx, stmt, invoiceId, subProjectId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubProjectInvoice(subProjectId string, invoiceId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_projects set invoices = array_remove(notes, $1) where id = $2`

	_, err := db.ExecContext(ctx, stmt, invoiceId, subProjectId)
	if err != nil {
		return err
	}

	return nil
}

func AppendSubProjectIncome(subProjectId string, incomeId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_projects set incomes = array(SELECT DISTINCT unnest(array_append(incomes, $1))) where id = $2`

	_, err := db.ExecContext(ctx, stmt, incomeId, subProjectId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubProjectIncome(subProjectId string, incomeId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_projects set incomes = array_remove(notes, $1) where id = $2`

	_, err := db.ExecContext(ctx, stmt, incomeId, subProjectId)
	if err != nil {
		return err
	}

	return nil
}

func AppendSubProjectExpense(subProjectId string, expenseId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_projects set expenses = array(SELECT DISTINCT unnest(array_append(expenses, $1))) where id = $2`

	_, err := db.ExecContext(ctx, stmt, expenseId, subProjectId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubProjectExpense(subProjectId string, expenseId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update sub_projects set expenses = array_remove(notes, $1) where id = $2`

	_, err := db.ExecContext(ctx, stmt, expenseId, subProjectId)
	if err != nil {
		return err
	}

	return nil
}

func GetSubProjectsByIds(ids string) ([]*PostgresSubProject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select id, name, description, status, priority, start_date, due_date, estimated_duration, 
		notes, projects, created_at, created_by, updated_at, updated_by, invoices, incomes, expenses
        from sub_projects
        where id = ANY($1)
		order by start_date desc
    `

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subProjects []*PostgresSubProject

	for rows.Next() {
		var subProject PostgresSubProject
		err := rows.Scan(
			&subProject.ID,
			&subProject.Name,
			&subProject.Description,
			&subProject.Status,
			&subProject.Priority,
			&subProject.StartDate,
			&subProject.DueDate,
			&subProject.EstimatedDuration,
			&subProject.Notes,
			&subProject.Projects,
			&subProject.CreatedAt,
			&subProject.CreatedBy,
			&subProject.UpdatedAt,
			&subProject.UpdatedBy,
			&subProject.Invoices,
			&subProject.Incomes,
			&subProject.Expenses,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		subProjects = append(subProjects, &subProject)
	}

	return subProjects, nil
}
