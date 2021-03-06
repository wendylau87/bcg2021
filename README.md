# BCG Digital Venture

## Pre-requisites
- mysql database
- go 1.12.x

## Import Database
1. Migrate database from file bcg/database/migrations/schema.sql

## Setting environment
1. Open file .env at folder 
```hashkell
DB_DRIVER=mysql
DB_USER=${dbuser}
DB_PASSWORD=${dbpassword}
DB_HOST=${dbhostname}
DB_PORT=3306
DB_DATABASE=bcg
```
2. Change value for this parameter (for hostname,you can use your IP) :
    - **${dbuser}**
    - **${dbpassword}**
    - **${dbhostname}**

## How to run services
1. ensure you have migrate all databases.
2. ensure you have change the environment file.
3. build binaries using this command
```bash
$ go build -o app
```
4. run binary using this command
```bash
$ ./app
```
5. note : this service default port will be 8090.

## Testing the service
1. open postman
2. import postman file `BCG.postman_collection.json`
3. execute postman 

## Graphql
- folder graphql