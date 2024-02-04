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

### 📝How to run:
- Configure .env.example file or use mine configuration and rename it to .env  
- Create a database with the name you specified in the .env file
- Run `docker-compose up --build` to build and run the app
- Install github.com/golang-migrate/migrate for migrations
- Run `make migrate/up` to apply migrations or `make migrate/down` to rollback migrations

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
[_] Finish CRUD handlers
    [_] add author repository and service methods
    [ ] finish author transport methods
    [_] add genre repository and service methods
    [ ] finish genre transport methods
    [_] add book repository and service methods
    [ ] finish book transport methods
    [ ] add json validation  
    [ ] add error responses & logging
[ ] Add registration and authentication    
[ ] Add authorization  
[ ] Add unit tests  
[ ] Add caching with Redis  

[X] Add auto project rebuild
[ ] Add swagger

To Be Continued...
```