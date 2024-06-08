package clients

import (
	"database/sql"
	"fmt"
)

func NewClient(sqlDB *sql.DB) {
	fmt.Println("Hello, World!")
}
