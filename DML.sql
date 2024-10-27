CREATE TABLE Employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phonenumber VARCHAR(15) NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    unit VARCHAR(255) NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    billDate DATE NOT NULL,
    entryDate DATE NOT NULL,
    finishDate DATE,
    employeeId INT REFERENCES employees(id) ON DELETE SET NULL,
    customerId INT REFERENCES customers(id) ON DELETE SET NULL
);

CREATE TABLE bill_details (
    id SERIAL PRIMARY KEY,
    billId INT REFERENCES transactions(id) ON DELETE CASCADE,
    productId INT REFERENCES products(id) ON DELETE SET NULL,
    productPrice INT NOT NULL,
    qty INT NOT NULL
);
