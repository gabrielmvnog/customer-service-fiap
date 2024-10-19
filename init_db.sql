DROP TABLE IF EXISTS customer;
CREATE TABLE customers(
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    documentNumber TEXT NOT NULL
);
