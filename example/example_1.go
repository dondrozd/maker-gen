package example

import (
	"database/sql"
	"time"
)

type ExampleOne struct {
	PublicString     string
	PublicInt        int64
	PublicTime       time.Time
	PublicNullString sql.NullString
}
