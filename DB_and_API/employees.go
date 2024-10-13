package main

import (
    "net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "github.com/gin-gonic/gin"
	"fmt"
)

type Employee struct {
	Emp_no int `json:"emp_no"`
	Birth_date  string `json:"birth_date"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Gender  string `json:"gender"`
	Hire_date string `json:"hire_date"`
}

// func getEmployeeDataFromdb() {

// }

var db *sql.DB

func getEmployeeByID(c *gin.Context) {

	id := c.Param("id")

	query_stmt := "SELECT * FROM employees.employees where emp_no = ?;"

	println(query_stmt)

    stmt, err := db.Prepare(query_stmt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer stmt.Close()

	if err != nil {
        panic(err.Error())
    }

	var employee Employee

	err = stmt.QueryRow(id).Scan(&employee.Emp_no, &employee.Birth_date, &employee.First_name, &employee.Last_name, 
		&employee.Gender, &employee.Hire_date)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	fmt.Println(employee)



	// for results.Next() {
    //     var tag Tag
    //     err = results.Scan(&tag.emp_no, &tag.birth_date, &tag.first_name, &tag.last_name, &tag.gender, &tag.hire_date)
    //     if err != nil {
    //         panic(err.Error())
    //     }
	// 	fmt.Println(tag.emp_no, tag.birth_date, tag.first_name, tag.last_name, tag.gender, tag.hire_date)
    // }

	fmt.Println("tag", employee)

	// fmt.Println(c.JSON(tag))

    c.JSON(http.StatusOK, employee)
}

func getEmployees(c *gin.Context) {
	query_stmt := "SELECT * FROM employees.employees;"

	println(query_stmt)

    stmt, err := db.Query(query_stmt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer stmt.Close()

	if err != nil {
        panic(err.Error())
    }

	var employee Employee
	var employees []Employee

	for stmt.Next() {
		err := stmt.Scan(&employee.Emp_no, &employee.Birth_date, &employee.First_name, &employee.Last_name, 
		&employee.Gender, &employee.Hire_date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		employees = append(employees,employee)
	}

	c.JSON(http.StatusOK, employees)
}

func postEmployee(c *gin.Context) {
	var employee Employee
	
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	insert_stmt := "INSERT INTO employees (emp_no, birth_date, first_name, last_name, gender, hire_date) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(insert_stmt)

	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	_, err = stmt.Exec(employee.Emp_no, employee.Birth_date, employee.First_name, employee.Last_name, 
		employee.Gender, employee.Hire_date)

	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Employee created successfully", "employee": employee})
}

func main() {

	var err error

	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/employees")

    if err != nil {
        panic(err.Error())
    }

    router := gin.Default()

    router.GET("/Employees/id/:id", getEmployeeByID)
    router.GET("/Employees", getEmployees)
	router.POST("/Employee", postEmployee)
	// router.POST("/albums", postAlbums)
	// router.GET("/albums/:id",getAlbumByID)

    router.Run("localhost:8080")
}
