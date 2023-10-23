package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Project:    Project{},
		SubProject: SubProject{},
	}
}

type Models struct {
	Project    Project
	SubProject SubProject
}

type NewProject struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Notes       []string  `json:"notes"`
	SubProjects []string  `json:"sub_projects"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}

type PostgresProject struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`
	SubProjects string    `json:"sub_projects"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
}

type NewSubProject struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             []string  `json:"notes"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
	Invoices          []string  `json:"invoices"`
	Incomes           []string  `json:"incomes"`
	Expenses          []string  `json:"expenses"`
}

type SubProject struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             []string  `json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         string    `json:"updated_by"`
	Projects          []string  `json:"projects"`
	Invoices          []string  `json:"invoices"`
	Incomes           []string  `json:"incomes"`
	Expenses          []string  `json:"expenses"`
}

type PostgresSubProject struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	Priority          int       `json:"priority"`
	StartDate         time.Time `json:"start_date"`
	DueDate           time.Time `json:"due_date"`
	EstimatedDuration int       `json:"estimated_duration"`
	Notes             string    `json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         string    `json:"updated_by"`
	Projects          string    `json:"projects"`
	Invoices          string    `json:"invoices"`
	Incomes           string    `json:"incomes"`
	Expenses          string    `json:"expenses"`
}

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
			fmt.Println("Error scanning", err)
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
			fmt.Println("Error scanning", err)
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
		&project.UpdatedAt,
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
		fmt.Println("Error updating project", err)
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

// --------------------
// SUB PROJECTS -------
// --------------------

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
			fmt.Println("Error scanning", err)
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
		fmt.Println("Error updating sub project", err)
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

	stmt := `update sub_projects set notes = array_remove(notes, $1) where id = $2`

	_, err := db.ExecContext(ctx, stmt, noteId, subProjectId)
	if err != nil {
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
			fmt.Println("Error scanning", err)
			return nil, err
		}

		subProjects = append(subProjects, &subProject)
	}

	return subProjects, nil
}

// ------------------------
// -- COMMON (entangled) --
// ------------------------

func AppendProjectsToSubProject(projectIds []string, subProjectId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		with new_projects as (
			select distinct unnest(projects || $1) as value
			from sub_projects
			where id = $2
		)
		update sub_projects
		set projects = array(select value from new_projects)
		where id = $2
	`

	_, err := db.ExecContext(ctx, stmt, convertToPostgresArray(projectIds), subProjectId)
	if err != nil {
		return err
	}

	for _, projectId := range projectIds {
		stmt := `update projects set sub_projects = array(SELECT DISTINCT unnest(array_append(sub_projects, $1))) where id = $2`

		_, err := db.ExecContext(ctx, stmt, subProjectId, projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteProjectsFromSubProject(projectIds []string, subProjectId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	for _, projectId := range projectIds {

		stmtUpdateSubProjects := `update sub_projects set projects = array_remove(projects, $1) where id = $2`

		_, err := db.ExecContext(ctx, stmtUpdateSubProjects, projectId, subProjectId)
		if err != nil {
			return err
		}

		stmtUpdateProjects := `update projects set sub_projects = array_remove(sub_projects, $1) where id = $2`

		_, err = db.ExecContext(ctx, stmtUpdateProjects, subProjectId, projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func AppendSubProjectsToProject(projectId string, subProjectIds []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		with new_sub_projects as (
			select distinct unnest(sub_projects || $1) as value
			from projects
			where id = $2
		)
		update projects
		set sub_projects = array(select value from new_sub_projects)
		where id = $2
	`

	_, err := db.ExecContext(ctx, stmt, convertToPostgresArray(subProjectIds), projectId)
	if err != nil {
		return err
	}

	for _, subProjectId := range subProjectIds {
		stmt := `update sub_projects set projects = array(SELECT DISTINCT unnest(array_append(projects, $1))) where id = $2`

		_, err := db.ExecContext(ctx, stmt, projectId, subProjectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteSubProjectsFromProject(projectId string, subProjectIds []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
        update projects
        set sub_projects = array(
            select distinct unnest(sub_projects) except all $1
        )
        where id = $2
    `

	_, err := db.ExecContext(ctx, stmt, convertToPostgresArray(subProjectIds), projectId)
	if err != nil {
		return err
	}

	for _, subProjectId := range subProjectIds {
		stmt := `update sub_projects set projects = array_remove(projects, $1) where id = $2`

		_, err := db.ExecContext(ctx, stmt, projectId, subProjectId)
		if err != nil {
			return err
		}
	}

	return nil
}
