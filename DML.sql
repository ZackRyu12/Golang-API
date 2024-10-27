CREATE TABLE Employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phonenumber VARCHAR(15) NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(20, 2) NOT NULL,
    address VARCHAR(255) NOT NULL
);