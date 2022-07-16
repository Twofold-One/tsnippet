# PostgreSQL Notes

## General

Connect to psql:

`psql -U <postgres_username> -d <dbname>`

## General psql

To list all users:

`\du`

To list all db's:

`\l`

To list all tables in the current db:

`\dt`

Run .sql file:

`\i <file_path>`

## Create DB

`CREATE DATABASE <dbname> OWNER <postgres_username>`
