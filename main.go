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
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Unit  string `json:"unit"`
}

type TransactionRequest struct {
	BillDate    string       `json:"billDate"`
	EntryDate   string       `json:"entryDate"`
	FinishDate  string       `json:"finishDate"`
	EmployeeId  string       `json:"employeeId"`
	CustomerId  string       `json:"customerId"`
	BillDetails []BillDetail `json:"billDetails"`
}

type BillDetail struct {
	ID           string `json:"id"`
	ProductId    string `json:"productId"`
	ProductPrice int    `json:"productPrice"`
	Qty          int    `json:"qty"`
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
		Product.GET("/", getAllProduct)
		Product.GET("/:id", getProductByID)
		Product.PUT("/:id", updateProduct)
		Product.DELETE("/:id", deleteProduct)
	}

	Transaction := router.Group("/transaction")
	{
		Transaction.POST("/", createTransaction)
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

func getAllProduct(c *gin.Context) {
	searchName := c.Query("name")

	query := "SELECT id, name, price, unit FROM Products"

	var rows *sql.Rows
	var err error

	if searchName != "" {
		query += " WHERE name ILIKE '%' || $1 || '%'"
		rows, err = db.Query(query, searchName)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	defer rows.Close()

	var matchedProduct []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Unit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse product data"})
			return
		}
		matchedProduct = append(matchedProduct, product)
	}

	if len(matchedProduct) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Product retrieved", "data": matchedProduct})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	}
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	// Validasi ID agar memastikan hanya berupa angka
	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Query untuk mendapatkan produk berdasarkan ID
	query := "SELECT id, name, price, unit FROM Products WHERE id = $1"
	var product Product

	// Menggunakan QueryRow karena kita hanya mengharapkan satu hasil
	err := db.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Price, &product.Unit)
	if err != nil {
		// Mengembalikan pesan error yang sesuai jika produk tidak ditemukan
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
		}
		return
	}

	// Mengembalikan respons sukses jika produk ditemukan
	c.JSON(http.StatusOK, gin.H{"message": "Product retrieved", "data": product})
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	query := `UPDATE Products SET name = $1, price = $2, unit = $3 WHERE id = $4`
	result, err := db.Exec(query, product.Name, product.Price, product.Unit, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "data": product})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM Products WHERE id = $1"
	result, err := db.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm deletion"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "data": "OK"})
}

func createTransaction(c *gin.Context) {
	var transaction TransactionRequest

	// Parsing input dari body request
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Mulai transaksi database
	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to begin transaction"})
		return
	}

	// Insert data ke tabel transactions
	queryTransaction := "INSERT INTO transactions (billDate, entryDate, finishDate, employeeId, customerId) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var transactionID string
	err = tx.QueryRow(queryTransaction, transaction.BillDate, transaction.EntryDate, transaction.FinishDate, transaction.EmployeeId, transaction.CustomerId).Scan(&transactionID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Insert data ke tabel bill_details
	for _, detail := range transaction.BillDetails {
		queryBillDetail := `
			INSERT INTO bill_details (billId, productId, qty)
			VALUES ($1, $2, $3) RETURNING id, productPrice`
		var billDetailID string
		// Dapatkan harga produk dari tabel Products
		var productPrice int
		err = tx.QueryRow("SELECT price FROM Products WHERE id = $1", detail.ProductId).Scan(&productPrice)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product price", "details": err.Error()})
			}
			return
		}

		// Insert ke tabel bill_details
		queryBillDetail = "INSERT INTO bill_details (billId, productId, qty, productPrice) VALUES ($1, $2, $3, $4) RETURNING id"
		err = tx.QueryRow(queryBillDetail, transactionID, detail.ProductId, detail.Qty, productPrice).Scan(&billDetailID)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bill detail", "details": err.Error()})
			return
		}

		detail.ID = billDetailID
		detail.ProductPrice = productPrice
	}

	// Commit transaksi jika semuanya sukses
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// Mengembalikan respons sukses
	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaction created",
		"data": gin.H{
			"id":          transactionID,
			"billDate":    transaction.BillDate,
			"entryDate":   transaction.EntryDate,
			"finishDate":  transaction.FinishDate,
			"employeeId":  transaction.EmployeeId,
			"customerId":  transaction.CustomerId,
			"billDetails": transaction.BillDetails,
		},
	})
}
