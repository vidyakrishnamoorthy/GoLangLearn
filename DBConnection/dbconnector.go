package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	emp_no string `json:"emp_no"`
	birth_date  string `json:"birth_date"`
	first_name string `json:"first_name"`
	last_name string `json:"last_name"`
	gender  string `json:"gender"`
	hire_date string `json:"hire_date"`
}

func main() {
    fmt.Println("Go MySQL Tutorial")

    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/employees")

    if err != nil {
        panic(err.Error())
    }

	err = db.Ping()
    if err != nil {
        panic(err)
    }

	println(db)
	println(db.Query("SELECT * FROM employees.employees limit 10;"))

	results, err := db.Query("SELECT * FROM employees.employees limit 10;")

    // if err != nil {
    //     panic(err.Error())
    // }

	for results.Next() {
        var tag Tag
        err = results.Scan(&tag.emp_no, &tag.birth_date, &tag.first_name, &tag.last_name, &tag.gender, 
			&tag.hire_date)
        if err != nil {
            panic(err.Error())
        }
		fmt.Println(tag.emp_no, tag.birth_date, tag.first_name, tag.last_name, tag.gender, tag.hire_date)
    }

    defer db.Close()

}