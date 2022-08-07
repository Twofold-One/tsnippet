CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created TIMESTAMP NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);