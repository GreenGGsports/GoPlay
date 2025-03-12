package models

import (
	"fmt"
)

func CreateTodo() error {
	// Use Exec() instead of Query()
	_, err := con.Exec("INSERT INTO TODO VALUES(?, ?)", "Evés ivás", "Evés isvás dinom dánom")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}
	return nil
}
