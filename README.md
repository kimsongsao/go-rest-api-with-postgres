[![Github](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/kimsongsao)
[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/kimsongsao/)

![Golang RESTful API](https://raw.githubusercontent.com/kimsongsao/go-rest-api-with-postgres/main/banner.png)


# Go RESTful API using GIN & GORM with PostgreSQL
An example RESTful API using GIN, GORM, PostgreSQL, and Swagger UI.

## Prerequisites
1. A working Go installation (version 1.21.1)
2. PostgreSQL
3. VS Code

## Getting started
1. Clone project from repository
```
git clone https://github.com/kimsongsao/go-rest-api-with-postgres.git
```
2. Create .env file by cloning from .env.example
3. Config DB_URL based on your database environment.
4. Run Migration
```
go run migrations/migration.go
```
4. Run
```
go run main.go
```
5. Open broswer with link below
```
http://localhost:3000/docs/index.html
```

## Project Structures
![Project Structures](https://raw.githubusercontent.com/kimsongsao/go-rest-api-with-postgres/main/projectstrutures.png)