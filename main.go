package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

type Customers struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type Employee struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price int `json:"price"`
	Unit     string `json:"unit"`
}

var db = config.ConnectDB()

func main() {
	// Tulis kode kamu disini
	router := gin.Default()
	// router.Use(LoggerMiddleware)

	Customer := router.Group("/Customers")
	{
		Customer.GET("/:id", getCustomers)
		Customer.POST("/", create_Customer)
		Customer.PUT("/:id", UpdateCustomer)
		Customer.DELETE("/:id", DeleteCustomer)
	}

	Employee := router.Group("/employees")
	{
		Employee.GET("/:id", getEmployee)
		Employee.POST("/", createEmployee)
		Employee.PUT("/:id", updateEmployee)
		Employee.DELETE("/:id", deleteEmployee)
	}

	Product := router.Group("/products")
	{
		Product.POST("/", createProduct)
	}


	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func create_Customer(c *gin.Context) {
	var customer Customers
	err := c.ShouldBind(&customer)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO Customers (name, phoneNumber, address) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err = db.QueryRow(query, customer.Name, customer.PhoneNumber, customer.Address).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	customer.Id = id
	c.JSON(http.StatusCreated, gin.H{"message": "Customer created", "data": customer})
}

func getCustomers(c *gin.Context) {
	id := c.Param("id")

	if id != "" {
		if _, err := strconv.Atoi(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
			return
		}
	}

	query := "SELECT id, name, phonenumber, address FROM Customers"
	var rows *sql.Rows
	var err error

	if id != "" {
		query += " WHERE id = $1"
		rows, err = db.Query(query, id)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customer"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customers
		if err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse customer data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Customers retrieved", "data": customer})
	}
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer Customers

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	query := `UPDATE Customers SET name = $1, phoneNumber = $2, address = $3 WHERE id = $4`
	result, err := db.Exec(query, customer.Name, customer.PhoneNumber, customer.Address, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully", "data": customer})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM Customers WHERE id = $1"
	result, err := db.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm deletion"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully", "data": "OK"})
}

func createEmployee(c *gin.Context) {
	var worker Employee
	err := c.ShouldBind(&worker)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO Employees (name, phoneNumber, address) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err = db.QueryRow(query, worker.Name, worker.PhoneNumber, worker.Address).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Employees"})
		return
	}

	worker.Id = id
	c.JSON(http.StatusCreated, gin.H{"message": "Employees created", "data": worker})
}

func getEmployee(c *gin.Context) {
	id := c.Param("id")

	if id != "" {
		if _, err := strconv.Atoi(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}
	}

	query := "SELECT id, name, phonenumber, address FROM Employees"
	var rows *sql.Rows
	var err error

	if id != "" {
		query += " WHERE id = $1"
		rows, err = db.Query(query, id)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve worker"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var employee Employee
		if err := rows.Scan(&employee.Id, &employee.Name, &employee.PhoneNumber, &employee.Address); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse employee data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Employee retrieved", "data": employee})
	}
}

func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var worker Employee

	if err := c.ShouldBindJSON(&worker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	query := `UPDATE Employees SET name = $1, phoneNumber = $2, address = $3 WHERE id = $4`
	result, err := db.Exec(query, worker.Name, worker.PhoneNumber, worker.Address, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully", "data": worker})
}

func deleteEmployee(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM Employees WHERE id = $1"
	result, err := db.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm deletion"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully", "data": "OK"})
}

func createProduct(c *gin.Context) {
	var product Product
	err := c.ShouldBind(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO Products (name, price, unit) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err = db.QueryRow(query, product.Name, product.Price, product.Unit).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	product.Id = id
	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "data": product})
}



