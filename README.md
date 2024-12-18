# Aplikasi Laundry

### Deskripsi

This repository contains a Laundry Service Management Application developed using Go (Golang) and the Gin framework. It allows managing customers, employees, products, and transactions for a laundry service.
---

### Features
- Manage Customers (CRUD operations)

- Manage Employees (CRUD operations)

- Manage Products (CRUD operations)

- Manage Transactions

  - Create transactions

  - Retrieve transaction details

  - List transactions with optional filters (date range and product name)


### Requirements
- Go(version 1.18 or later)
- PostgreSQL database
- Git

## Installation

1. Clone The Repository:
```cmd 
git clone https://github.com/my-username/laundry-service.git
```
2. Set Up the Database:
- Install PostgreSQL and create a database.
- Update the config.ConnectDB function in the config package with your database connection settings.
3. install Dependencies:
```go
  go mod tidy
```
4. Run the Application:
```go 
go run .
```
5. Access the API: The application will run on http://localhost:8080 by default.

## API Documentation
Use tools like Postman or cURL to test API endpoints

### Customer API

#### Create Customer

Request :

POST /customers

```json
{
  "name": "string",
  "phoneNumber": "string",
  "address": "string"
}
```

Response :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "phoneNumber": "string",
    "address": "string"
  }
}
```

#### Get Customer
GET /customers/:id

Response :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "phoneNumber": "string",
    "address": "string"
  }
}
```

#### Update Customer
PUT /customers/id

Request :
```json
{
  "name": "string",
  "phoneNumber": "string",
  "address": "string"
}
```

Response :
```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "phoneNumber": "string",
    "address": "string"
  }
}
```

#### Delete Customer

Request :

DELETE /customers/:id

Response :

```json
{
  "message": "string",
  "data": "OK"
}
```

### Employee API

#### Create Employee

Request :

POST /employees
```json
{
  "name": "string",
  "phoneNumber": "string",
  "address": "string"
}
```

Response :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "phoneNumber": "string",
    "address": "string"
  }
}
```

#### Get Employee

Request :

GET /employees/:id 

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "phoneNumber": "string",
    "address": "string"
  }
}
```

#### Update Employee

Request :

PUT employees/:id
```json
{
  "name": "string",
  "phoneNumber": "string",
  "address": "string"
}
```

Response :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "phoneNumber": "string",
    "address": "string"
  }
}
```

#### Delete Employee

Request :

DELETE /employees/:id

Response :


```json
{
  "message": "string",
  "data": "OK"
}
```

### Product API

#### Create Product

Request :

POST /products

```json
{
	"name": "string",
  "price": int,
  "unit": "string" (satuan product,cth: Buah atau Kg)
}
```

Response :


```json
{
	"message": "string",
	"data": {
		"id": "string",
		"name": "string",
		"price": int,
		"unit": "string" (satuan product,cth: Buah atau Kg)
	}
}
```

#### List Product

Request :

GET /products

Response :

```json
{
	"message": "string",
	"data": [
		{
			"id": "string",
			"name": "string",
			"price": int,
			"unit": "string" (satuan product,cth: Buah atau Kg)
		},
		{
			"id": "string",
			"name": "string",
			"price": int,
			"unit": "string" (satuan product,cth: Buah atau Kg)
		}
	]
}
```

#### Product By Id

Request :

GET /products/:id

Response :


```json
{
	"message": "string",
	"data": {
		"id": "string",
		"name": "string",
		"price": int,
		"unit": "string" (satuan product,cth: Buah atau Kg)
	}
}
```

#### Update Product

Request :

PUT /products/:id

```json
{
	"name": "string",
	"price": int,
	"unit": "string" (satuan product,cth: Buah atau Kg)
}
```

Response :

```json
{
	"message": "string",
	"data": {
		"id": "string",
		"name": "string",
		"price": int,
		"unit": "string" (satuan product,cth: Buah atau Kg)
	}
}
```

#### Delete Product

Request :

DELETE products/:id

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": "OK"
}
```

### Transaction API

#### Create Transaction

Request :

POST /transactions

```json
{
	"billDate": "string",
	"entryDate": "string",
	"finishDate": "string",
	"employeeId": "string",
	"customerId": "string",
	"billDetails": [
		{
			"productId": "string",
			"qty": int
		}
	]
}
```

Request :

```json
{
	"message": "string",
	"data":  {
		"id":  "string",
		"billDate":  "string",
		"entryDate":  "string",
		"finishDate":  "string",
		"employeeId":  "string",
		"customerId":  "string",
		"billDetails":  [
			{
				"id":	"string",
				"billId":  "string",
				"productId":  "string",
				"productPrice": int,
				"qty": int
			}
		]
	}
}
```

#### Get Transaction

Request :

GET /transactions/:id_bill

Response :

- Status Code: 200 OK
- Body :

```json
{
	"message": "string",
  "data": {
    "id": "string",
    "billDate": "string",
    "entryDate": "string",
    "finishDate": "string",
    "employee": {
      "id": "string",
      "name": "string",
      "phoneNumber": "string",
      "address": "string"
    },
    "customer": {
      "id": "string",
      "name": "string",
      "phoneNumber": "string",
      "address": "string"
    },
    "billDetails": [
      {
        "id": "string",
        "billId": "string",
        "product": {
          "id": "string",
          "name": "string",
          "price": int,
          "unit": "string" (satuan product,cth: Buah atau Kg)
        },
        "productPrice": int,
        "qty": int
      }
    ],
    "totalBill": int
  }
}
```

#### List Transaction

Request :

GET /transactions?startDate=DATE&endDate=DATE&productName=Product_name" for specific data relating to start time, end time, and product name 

GET/transactions for all data


Response :

```json
{
	"message": "string",
  "data": [
    {
      "id": "string",
      "billDate": "string",
      "entryDate": "string",
      "finishDate": "string",
      "employee": {
        "id": "string",
        "name": "string",
        "phoneNumber": "string",
        "address": "string"
      },
      "customer": {
        "id": "string",
        "name": "string",
        "phoneNumber": "string",
        "address": "string"
      },
      "billDetails": [
        {
          "id": "string",
          "billId": "string",
          "product": {
            "id": "string",
            "name": "string",
            "price": int,
            "unit": "string" (satuan product,cth: Buah atau Kg)
          },
          "productPrice": int,
          "qty": int
        }
      ],
      "totalBill": int
    }
  ]
}
```
