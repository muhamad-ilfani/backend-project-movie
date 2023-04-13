# backend-project-movie
Movie Festival App

## Technologies
1. Go Language
2. Echo Framework
3. Postgresql DB

## DB Setting
This service use postgreSQL database.
Please see .env file. Default value as follow :
>DB_HOST=localhost  
DB_DRIVER=postgres  
DB_USER=xxxx    
DB_PASSWORD=xxxx    
DB_NAME=postgres    
DB_PORT=5432

Change above configuration based on your local connection for Postgresql.

Make sure user has access to create schema and table. This service will automathically create schema and table.

 Or you can directly create schema with name "movieapps"

## API Secret for Authentication
Please see .env file. Default value as follow :
>API_SECRET=verysecret

## How to Run
1. Clone this repository with
    > git clone https://github.com/muhamad-ilfani/backend-project-movie.git
2. Setting your DB connection and match it with .env file

3. Run service with command
    > go run main.go

4. Make sure service is running in port 8000 with command
    > curl -X GET localhost:3030/

    Response :
    > {"message":"welcome"}

## List API
1. POST /register-admin

    This API is used to create user with admin role

    command :   
    curl -X POST 'localhost:3030/register-admin' -H "Content-type:application/json" -d "{\"user_name\": \"admin\", \"email\": \"admin@gmail.com\", \"password\": \"admin\"}"
2. POST /register-user

    This API is used to create common user

    command:    
    curl -X POST 'localhost:3030/register-user' -H "Content-type:application/json" -d \
"{\"user_name\": \"user1\", \"email\": \"user1@gmail.com\", \"password\": \"user1\"}"

3. POST /login
    This API is used for login process and create token

    command:    
    curl -X POST 'localhost:3030/login' -H "Content-type:application/json" -d \
"{\"email\": \"user1@gmail.com\", \"password\": \"user1\"}"

4. POST /user

    This API is used to get all movies with authentication user.
    You can filter the movie by title, desctiption, artists, genres and with pagination (limit, offset). Default value for limit=10, offset=0.

    command:    
    curl -X POST localhost:3030/user/ -H "Content-type:application/json" -H "Authorization: Bearer {token}"} -d "{\"limit\": 10, \"offset\": 0, \"title\":\"title\"}"

5. GET /user/:id

    This API is used to get user by ID with authentication.
    If user hit this API, it will add viewers value in DB.

    command:    
    curl -X GET localhost:3030/user/:id -H "Content-type:application/json" -H "Authorization: Bearer {token}"}

6. POST /admin/register-movie
    
    This API is used to register movie with authentication and admin authorization.

    command:    
    curl -X POST 'localhost:3030/admin/register-movie' -H "Content-type:application/json" -H "Authorization: Bearer {token}"} -d \\\
    "{
    \"title\": \"title\", 
    \"description\": \"description\",
    \"duration\": \"duration\",
    \"artists\": \"artists\",
    \"genres\": \"genres\",
    \"watch_url\": \"watch_url\", 
    }"

7. PATCH /update-movie/:id

    This API is used to update movie with authentication and admin authorization. You can fill several data which you want to update based on movie id.

    command:    
    curl -X PATCH 'localhost:3030/admin/update-movie/1' -H "Content-type:application/json" -H "Authorization: Bearer {token}"} -d \\\
    "{
    \"title\": \"new title\", 
    \"artists\": \"new artist\"
    }"

    nb : you can add desctiption, duration, genres,watch_url if you want to change