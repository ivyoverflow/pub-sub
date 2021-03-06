CREATE TABLE books (
    id VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE,
    date_of_issue VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    rating DECIMAL(4, 2) NOT NULL,
    price DECIMAL(6, 2) NOT NULL,
    in_stock BOOLEAN NOT NULL
);
