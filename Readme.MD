# DIKIRUIM TEST SOLUTION

This is the solution to a test in respect to an application for a golang position

## Task
Implement an authenticated GraphQL API that allows users to track todos. The schema should include the mutations and queries as shown in the attached picture.

Authenticated access should be handled with the Authorization header and Bearer tokens. You can use the golang-jwt/jwt (https://pkg.go.dev/github.com/golang-jwt/jwt) package to generate tokens. If the user isn’t authenticated, the API should return a 401 using middleware/directives.

Your GraphQL schema should implement nested types and field resolvers or preloads to fetch them from the DB. This should work recursively
![Text Image](Backend%20Callenge.JPG)

## Table of Contents

- [Project Title](#project-title)
    - [Table of Contents](#table-of-contents)
    - [Description](#description)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Contributing](#contributing)
    - [License](#license)

## Description
This project is a high-performance, scalable and secure web application developed using the latest technologies and industry best practices. The application is built with a package oriented design, utilizing GORM as the database layer, JWT for secure authentication, and Data Loader to handle N+1 queries for optimal performance.

The application is fully dockerized and ready for deployment in production environments. It has been thoroughly tested to ensure it can handle high traffic loads and maintain optimal performance under stress.

With the use of Viper config, the application is highly configurable, making it easy to deploy in a variety of environments with different configurations.

In addition, Zap has been used as the logging library, ensuring that detailed logs are available for debugging and monitoring.

Overall, this project is a robust, reliable and efficient web application that is well suited for production use. Its use of cutting-edge technologies and best practices ensures that it will continue to perform optimally and maintain high levels of security and reliability.

## Installation
To install and run the application, you need to have Docker installed on your machine.
If you don't have it installed, please follow the instructions on how to install Docker [here](https://docs.docker.com/get-docker/).

After Clone the repository and navigate to the project directory and run `docker compose up -d`

## Usage
### Sign Up
To sign up a new user, make a POST request to `http://localhost:7500/query` with the following query:
```mutation {
  signUp(username: "john_doe", email: "john.doe@example.com", password: "password123") {
    success
    message
    user {
      userid
      username
      email
    }
    authToken
  }
}
```

A sample response would look like
```
{
    "data": {
        "signUp": {
            "success": true,
            "message": "Successfully created user ",
            "user": {
                "username": "john_doe",
                "email": "john.doe@example.com",
                "userid": "b63fe0dc-d990-4b24-8421-f0bc4ddac77e"
            },
            "authToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiIiwidXNlcm5hbWUiOiJhbGkiLCJlbWFpbCI6ImFiY0BleGFsZWUxMzQuY29tIiwiaXNzIjoiZG9tYWluX3VybCIsInN1YiI6Ik9yZW9sdXdhIiwiZXhwIjoxNjgzMjcxNjY5LCJpYXQiOjE2ODMyNzA0Njl9.K1waM_aC38OTZFQajtVAmdIIYlDrZ3baqK8rOv_gqrs"
        }
    }
}
```

[Learn more about GraphQL mutations and queries](https://graphql.org/learn/queries/)


## Contact
If you have any suggestions or requests, please feel free to reach out to me. I'll do my best to respond promptly and accommodate your needs as best I can. Thank you for your interest in this project!
