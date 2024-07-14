## Real Estate Management System

This project constitutes a comprehensive full-stack real estate application featuring CRUD (Create, Read, Update, Delete) functionalities for property listings and purchases.
It also includes an array of security measures such as rate limiting and unit testing for property endpoints.

## Getting Started

Clone the Repository: <br>
`git clone <repository-url>`

### Install Dependencies:

To run unit tests for the server-side code, navigate to the server directory and execute:<br>
`go mod tidy`

### Set Up Environment Variables:

Create a .env file in the server directory using .env.example as the template and update the necessary configurations.

### Run the Server:

`go run main.go`

### Unit Testing: Running Tests

To run unit tests for the server-side code, navigate to the server directory and execute:<br>

`cd unit_test`<br>
`go test`

## Features

- Robust user authentication with secure session management.
- CRUD operations for efficient management of property listings.
- Implementation of rate limiting to control the number of requests to the API.
- Unit testing for property-related endpoints to ensure reliability.
- Authorization checks for controlling access to specific routes or actions based on user roles.
