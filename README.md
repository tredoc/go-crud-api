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
💡Check docs dir for more info about db schema. 
Feel free to insert db schema from txt file to dbdiagram.io to see it in a more convenient way.

## ✅Flow:
```
[X] Add gitignore
[X] Init module  
[X] Design a schema for a relational database with dbdiagram.io  
[X] Add basic server  
[X] Add Slog logger  
[X] Mock route handlers for Books CRUD  
[X] Add repository  
[ ] Dockerize the go app and postgresql  
[ ] Add configuration  
[ ] Add migration  
[ ] Finish CRUD handlers
    [ ] validate input json  
    [ ] unify error responses  
[ ] Add registration and authentication    
[ ] Add authorization  
[ ] Add unit tests  
[ ] Add cashing with Redis  

To Be Continued...
```