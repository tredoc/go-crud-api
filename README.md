#  🚀Book CRUD api
## 💡Consider it as a junior developer testing task

### ❗Requirements
#### 💻API:
- The application should be able to, via HTTP requests:
- Retrieve by id (GET /book)
- Retrieve a list of all entities (GET /books)
- Create (POST/PUT)
- Update an entity by id (POST/PUT)
- Delete an entity by id (DELETE)

#### 💻Book entity:
- Log errors with any logger of your choice 
- Use JSON for data transmission 
- Use any library of your choice for routing
- Data storage of your choice: PostgreSQL, MSSQL, MYSQL, MongoDB, etc.
- Use SQL/ORM for working with relational databases

#### 💻Additional tasks:
- Caching requests
- Dockerize the application
- Authorization
- Automatic testing through Postman
- Unit tests

---
💡Check db/docs for more info about db schema. 
Feel free to insert db schema from txt file to dbdiagram.io to see it in a more convenient way.

To generate migration use `migrate create -ext sql -dir db/migrations -seq migration_name`

### 📝How to run:
- Configure .env.example file or use mine configuration and rename it to .env
- Run `docker-compose up --build` to build and run the app
- Create a database with the name you specified in the .env file
- Install `github.com/golang-migrate/migrate` for migrations
- Run `make migrate/up` to apply migrations or `make migrate/down` to rollback migrations
- Install `go install github.com/vektra/mockery/v2@v2.42.0` for mock generation
- Run `mock/service` to generate mock files for services
- Run `make test` to run tests
- Run `make run/dev` to run the app in dev mode

## ✅Flow:
```
[X] Add gitignore
[X] Init module  
[X] Design a schema for a relational database with dbdiagram.io  
[X] Add basic server  
[X] Add Slog logger  
[X] Mock route handlers for Books CRUD  
[X] Add repository  
[X] Dockerize the go app and postgresql  
[X] Add configuration  
[X] Add migration  
[X] Finish CRUD handlers
    [X] add author repository and service methods
    [X] finish author transport methods
    [X] add genre repository and service methods
    [X] finish genre transport methods
    [X] add book repository and service methods
    [X] finish book transport methods
    [X] add json validation  
    [X] add error responses & logging
    [X] add tests and refactor CRUD handlers
[X] Add registration and authentication    
[X] Add authorization  
[_] Add caching with Redis  

[X] Add project auto rebuild  
[ ] Move to 1.22 native router  
[ ] Add graceful shutdown
[ ] Add pagination, sorting and filtration  
[ ] Add db constraints for isbn, authors, genres and etc  
[ ] Add swagger  

To Be Continued...
```