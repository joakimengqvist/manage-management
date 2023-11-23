package data

import (
	"context"
	"time"
)

func InsertUserSettings(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID string
	stmt := `insert into user_settings (
			user_id,
			dark_theme,
			compact_ui,
			created_at,
			created_by,
			updated_at,
			updated_by)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err := db.QueryRowContext(ctx, stmt,
		userId,
		false,
		false,
		time.Now(),
		userId,
		time.Now(),
		userId,
	).Scan(&newID)

	if err != nil {
		return err
	}

	return nil
}

func GetUserSettingsByUserId(userId string) (*UserSettings, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select 
		id,
		user_id,
		dark_theme,
		compact_ui,
		created_at,
		created_by,
		updated_at,
		updated_by
		from user_settings where user_id = $1`

	row := db.QueryRowContext(ctx, query, userId)

	var userSettings UserSettings
	err := row.Scan(
		&userSettings.ID,
		&userSettings.UserId,
		&userSettings.DarkTheme,
		&userSettings.CompactUi,
		&userSettings.CreatedAt,
		&userSettings.CreatedBy,
		&userSettings.UpdatedAt,
		&userSettings.UpdatedBy,
	)

	if err != nil {
		return nil, err
	}

	return &userSettings, nil
}

func UpdateUserSettings(userSettings UpdateUserSettingsPayload, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update user_settings set
		dark_theme = $1,
		compact_ui = $2,
		updated_at = $3,
		updated_by = $4
		where user_id = $5`

	_, err := db.ExecContext(ctx, stmt,
		userSettings.DarkTheme,
		userSettings.CompactUi,
		time.Now(),
		userId,
		userSettings.UserId,
	)

	if err != nil {
		return err
	}

	return nil
}
