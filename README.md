# About

This project implements a backend server written in Go for a bookstore. It persists data in a mysql database (MariaDB) hosted locally on the machine. The server supports the following CRUD APIs:-
1. Create book -> POST on /book
2. Update book by ID -> PUT on /book/{Id}
3. Read all books -> READ on /book
4. Read book by ID -> READ on /book/{Id}
5. Delete book by ID -> DELETE on /book/{Id}

These APIs have been tested and verified using postman

# Steps to run the project

1. Install MariaDB on your machine
2. Create a Database 'bookstore'
3. Run the following command in MariaDB to grant access to the database for a new user (which will be used by the Go project to access the database).

grant all on bookstore.* to '{yourUserName}'@localhost identified by '{yourPasswd}'

4. Open the file /pkg/config/app.go
5. Edit line 11 with {yourUserName} and {yourPasswd} set in step 3 as follows:-

d, err := gorm.Open("mysql", "{yourUserName}:{yourPasswd}!!!@/bookstore?charset=utf8mb4&parseTime=True&loc=Local")

6. Run "go build" in the /cmd/main directory in your terminal
7. Run "go run main.go" command in the terminal
8. Use postman to test the APIs defined above