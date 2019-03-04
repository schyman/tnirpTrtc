# Application details

## External Libraries used

 -  `github.com/lib/pq` - Driver for communicating with Postgres server


## Application Structure


    --
      src/
        - config.json  - Contains application configuration details (Rest port, database connection details)
        - app.go  - Reads config file, connects to database and starts rest server
        - models.go - Contains object models used in application
        - database.go - Module responsible for communicating with Postgres server
        - queries.go - Contains various database queries used in the application
        - rest.go - Module responsible for REST server and routing

        