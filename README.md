# How To Run It Locally

1. Install makefile.
2. Install docker
3. run in terminal in the root folder. 'make local-test'
4. export these variables:
export DB_HOST=localhost
export DB_USER=test_user
export DB_PASS=test_password
export DB_PORT=5432
export DB_SSL_MODE=disable
export PORT=8080
export DB_NAME=test_db
5. run in terminal in the root folder. 'go run main.go'
6. run your desired api

# Running using external DB
1. Modify code in main.go openDB to include ssl certs.
2. export these variables:
export DB_HOST={yourhosttoyourdb}
export DB_USER={yourdbuser}
export DB_PASS={yourdbpassword}
export DB_PORT={yourdbport}
export DB_SSL_MODE={verifyca}
export PORT={the port you want to run your application}
export DB_NAME={yourdbname}
3. run in terminal in the root folder. 'go run main.go'
4. run your desired api