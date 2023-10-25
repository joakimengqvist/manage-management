package data

import (
	"context"
	"log"
	"time"
)

func (p *Privilege) GetAllPrivileges() ([]*Privilege, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, created_at, updated_at from privileges order by name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privileges []*Privilege

	for rows.Next() {
		var privilege Privilege
		err := rows.Scan(
			&privilege.ID,
			&privilege.Name,
			&privilege.Description,
			&privilege.CreatedAt,
			&privilege.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		privileges = append(privileges, &privilege)
	}

	return privileges, nil
}

func (p *Privilege) InsertPrivilege(privilege Privilege) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `insert into privileges (name, description, created_at, updated_at)
		values ($1, $2, $3, $4) returning id`

	err := db.QueryRowContext(ctx, stmt,
		privilege.Name,
		privilege.Description,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return "", err
	}

	return newID, nil
}

func (p *Privilege) GetPrivilegeById(id string) (*Privilege, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, created_at, updated_at from privileges where id = $1`

	var privilege Privilege
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&privilege.ID,
		&privilege.Name,
		&privilege.Description,
		&privilege.CreatedAt,
		&privilege.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &privilege, nil
}

func (p *Privilege) UpdatePrivilege() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update privileges set
		name = $1,
		description = $2,
		updated_at = $3
		where id = $4
	`

	_, err := db.ExecContext(ctx, stmt,
		p.Name,
		p.Description,
		time.Now(),
		p.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *Privilege) DeletePrivilege() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from privileges where id = $1`

	_, err := db.ExecContext(ctx, stmt, p.ID)
	if err != nil {
		return err
	}

	return nil
}
