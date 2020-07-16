# sling_cafe 

A simple REST API made using:
* Golang
* mongo-go MongoDB Driver

## Requirements
- go and dependencies
- mongodb

## Project Structure

Uses the model, repo, handler structure.

```
│   config.json                                   // contains configuration details
│   main.go                                       // entry point
│
├───app                                           
│   ├───handler                                   // all the endpoints defined in this directory
│   │       admin_handler.go
│   │       admin_transaction_handler.go
│   │       caterer_handler.go
│   │       employee_handler.go
│   │       employee_transaction_handler.go
│   │       meal_entry_handler.go
│   │       meal_type_handler.go
│   │       report_handler.go
│   │       search_handler.go
│   │
│   ├───model                                   // data structures
│   │       admin.go
│   │       admin_transaction.go
│   │       caterer.go
│   │       employee.go
│   │       employee_transaction.go
│   │       enums.go
│   │       manager.go
│   │       meal_entry.go
│   │       meal_type.go
│   │
│   └───repo                                    // performs all database related tasks
│           admin_repo.go
│           admin_transaction_repo.go
│           caterer_repo.go
│           employee_repo.go
│           employee_transaction_repo.go
│           events.go
│           manager_repo.go
│           meal_entry_repo.go
│           meal_type_repo.go
│
├───config                                      // helper singleton to load the configuration from config.json
│       config.go
│
├───db                                          // helper singleton to connect to database
│       mongo.go
│
├───dbscripts                                   // scripts for testing db and mongodb queries (not essential for project)
│       db_init.js
│       db_meals_transactions.js
│       db_payments_transactions.js
│       db_spec.js
│       sling_cafe_export.json
│       test.js
│
├───log                                         // logger wrapper
│       logger.go
│
└───util                                        // custom utilities
        custom_datetime.go
        custom_pipeline.go
        cutom_bson.go
        errors.go
        kronika.go
        optional_query.go
        validate.go
```

## Installation

1. Install go, follow the official install instructions
2. Install all go dependencies by running `go get` from the project directory
3. Install mongodb, follow official install instructions

### Building and Running
1. run `go build` from the project directory
2. run the executable


  
