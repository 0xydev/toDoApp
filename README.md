# ToDoApp with GO

This project is a REST API that enables users to manage their to-do list. The API supports GET, POST, DELETE, and PUT HTTP requests and offers multiple endpoints for interacting with the data.

## Prerequisites

Before you begin, ensure you have installed:

- [Go](https://go.dev/) (at least version 1.16, Go Modules are used)

## Cloning the Repository

To get a local copy of the code, clone it using git:

```bash
git clone https://github.com/0xydev/toDoApp.git
cd ToDoApp
```
## Managing dependencies

This project uses Go Modules for managing dependencies. To download all necessary dependencies, use the following command in your terminal:
```bash
go mod tidy
```

This command aligns the go.mod file with the actual source code dependencies.
Starting the Application

Run the following command to start the application:

```bash
go run main.go
```

This command starts the server, and it's now listening to HTTP requests.
Using the Application

After running the server, you can interact with the API using the following endpoints:

    GET /todos: List all to-do items.
    
    POST /todos: Add a new to-do item.
    
    GET /todos/{id}: Fetch details of a single to-do item by ID.
    
    PUT /todos/{id}: Update details of a particular to-do item by ID.
    
    DELETE /todos/{id}: Delete a particular to-do item by ID. 

## Testing

You can test the application endpoints using tools such as cURL, Postman, or any other HTTP client.
Contributing

For any changes, please open an issue first to discuss what you would like to contribute. Pull requests are welcome.
## License

This project is licensed under the MIT License.
