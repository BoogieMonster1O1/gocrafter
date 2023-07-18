package data

import "database/sql"

type Host struct {
	Name        string
	ID          int64
	SSHHostname sql.NullString
	SSHPort     sql.NullInt64
	IsLocal     bool
	Status      string
}
