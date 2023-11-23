package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User:         User{},
		UserSettings: UserSettings{},
		Privilege:    Privilege{},
	}
}

type Models struct {
	User         User
	UserSettings UserSettings
	Privilege    Privilege
}

type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Privileges string    `json:"privileges"`
	Projects   string    `json:"projects"`
	Password   string    `json:"-"`
	Active     int       `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ReturnedUser struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Privileges []string  `json:"privileges"`
	Projects   []string  `json:"projects"`
	Password   string    `json:"-"`
	Active     int       `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserSettings struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	DarkTheme bool      `json:"dark_theme"`
	CompactUi bool      `json:"compact_ui"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type UpdateUserSettingsPayload struct {
	UserId    string `json:"user_id"`
	DarkTheme bool   `json:"dark_theme"`
	CompactUi bool   `json:"compact_ui"`
}

type Privilege struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PrivilegePayload struct {
	UserId string `json:"userId"`
	Action string `json:"action"`
}
