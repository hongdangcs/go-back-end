## Architecture Layers of the project

- Router
- Controller
- Usecase
- Repository
- Domain

![Go Backend Clean Architecture Diagram](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-backend-arch-diagram.png?raw=true)

## Major Packages used in this project

- **gin**: Gin is an HTTP web framework written in Go (Golang). It features a Martini-like API with much better
  performance -- up to 40 times faster. If you need a smashing performance, get yourself some Gin.
- **mongo go driver**: The Official Golang driver for MongoDB.
- **jwt**: JSON Web Tokens are an open, industry-standard RFC 7519 method for representing claims securely between two
  parties. Used for Access Token and Refresh Token.
- **viper**: For loading configuration from the `.env` file. Go configuration with fangs. Find, load, and unmarshal a
  configuration file in JSON, TOML, YAML, HCL, INI, envfile, or Java properties formats.
- **bcrypt**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **testify**: A toolkit with common assertions and mocks that plays nicely with the standard library.
- **mockery**: A mock code autogenerator for Golang used in testing.
- Check more packages in `go.mod`.

### Public API Request Flow without JWT Authentication Middleware

![Public API Request Flow](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-public-api-request-flow.png?raw=true)

### Private API Request Flow with JWT Authentication Middleware

> JWT Authentication Middleware for Access Token Validation.

![Private API Request Flow](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-private-api-request-flow.png?raw=true)

### How to run this project?

We can run this Go Backend Clean Architecture project with or without Docker. Here, I am providing both ways to run this
project.

- Clone this project

```bash
# Move to your workspace
cd your-workspace

# Clone this project into your workspace
git clone https://github.com/hongdangcseiu/go-back-end.git

# Move to the project root directory
cd go-backend-clean-architecture
```

#### Run without Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install `go` if not installed on your machine.
- Install `MongoDB` if not installed on your machine.
- Important: Change the `DB_HOST` to `localhost` (`DB_HOST=localhost`) in `.env` configuration file. `DB_HOST=mongodb`
  is needed only when you run with Docker.
- Run `go run cmd/main.go`.
- Access API using `http://localhost:8080`

#### Run with Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run `docker-compose up -d`.
- Access API using `http://localhost:8080`

### How to run the test?

```bash
# Run all tests
go test ./...
```

### How to generate the mock code?

In this project, to test, we need to generate mock code for the use-case, repository, and database.

```bash
# Generate mock code for the usecase and repository
mockery --dir=domain --output=domain/mocks --outpkg=mocks --all

# Generate mock code for the database
mockery --dir=mongo --output=mongo/mocks --outpkg=mocks --all
```

Whenever you make changes in the interfaces of these use-cases, repositories, or databases, you need to run the
corresponding command to regenerate the mock code for testing.

### The Complete Project Folder Structure

```
.
├── Dockerfile
├── api
│   ├── controller
│   │   ├── login_controller.go
│   │   ├── profile_controller.go
│   │   ├── profile_controller_test.go
│   │   ├── refresh_token_controller.go
│   │   ├── signup_controller.go
│   │   ├── user_controller.go
│   │   ├── comment_controller.go
│   │   └── post_controller.go
│   ├── middleware
│   │   └── jwt_auth_middleware.go
│   └── route
│       ├── login_route.go
│       ├── profile_route.go
│       ├── refresh_token_route.go
│       ├── route.go
│       ├── signup_route.go
│       ├── user_route.go
│       ├── comment_route.go
│       └── post_route.go
├── bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd
│   └── main.go
├── docker-compose.yaml
├── domain
│   ├── error_response.go
│   ├── jwt_custom.go
│   ├── login.go
│   ├── profile.go
│   ├── refresh_token.go
│   ├── signup.go
│   ├── success_response.go
│   ├── comment.go
│   ├── post.go
│   └── user.go
├── go.mod
├── go.sum
├── internal
│   └── tokenutil
│       └── tokenutil.go
├── mongo
│   └── mongo.go
├── repository
│   ├── comment_repository.go
│   ├── post_repository.go
│   ├── user_repository.go
│   └── user_repository_test.go
└── usecase
    ├── login_usecase.go
    ├── comment_usecase.go
    ├── post_usecase.go
    ├── profile_usecase.go
    ├── refresh_token_usecase.go
    ├── signup_usecase.go
    └── user_usecase.go
```

### Example API Request and Response

- Signup

    - Request

  ```
  fetch('http://localhost:8080/api/signup', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded'
  },
  body: 'email=test@gmail.com&password=test&name=Test%20Name'
  }
  )
  ```

    - Response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

- Login

    - Request

  ```
  fetch('http://localhost:8080/api/login', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded'
  },
  body: 'email=test@gmail.com&password=test'
  }
  )
  ```

    - Response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

- Profile

    - Request

  ```
  fetch('http://localhost:8080/api/profile', {
  method: 'GET',
  headers: {
    'Authorization': 'Bearer access_token'
  }
  }
  )
  ```

    - Response

  ```json
  {
    "name": "Test Name",
    "email": "test@gmail.com"
  }
  ```

- Blog post
    - Create a blog post

        - Request

      ```
      fetch('http://localhost:8080/api/post/new', {method: 'POST',  headers: {
        'Authorization': 'Bearer access_token',
        'Content-Type': 'application/x-www-form-urlencoded'
        },  
      body: 'title=Blog title&content=Blog content'  
      }
      )
      ```

        - Response

      ```json
      {  
      "message": "Post created successfully"  
      }
      ``` 
    - Get all blog post

        - Request

      ```
      fetch('http://localhost:8080/api/post', 
      {
      method: 'GET'
      }
      )
      ```

    - Get a blog post by ID

        - Request

      ```
      fetch('http://localhost:8080/api/post/:id', 
      {
      method: 'GET',
      }
      )
      ```

        - Response

      ```json
      {
        "title": "Post title",
        "userID": "user id",
        "content": "Post content",
        "date_create": "",
        "date_update": "",
        "categories": [],
        "approved": ""
        }
      ``` 
- Comment
    - Create a comment

        - Request

      ```
      fetch('http://localhost:8080/api/comment/:postID', {method: 'POST',  headers: {
        'Authorization': 'Bearer access_token',
        'Content-Type': 'application/x-www-form-urlencoded'
        },  
      body: 'content=Comment content'  
      }
      )
      ```

        - Response

      ```json
      {  
      "message": "Comment created successfully"  
      }
      ``` 
    - Get all comment of a blog post

        - Request

      ```
      fetch('http://localhost:8080/api/comment/post/:postId', 
      {
      method: 'GET'
      }
      )
      ```
        - Response

      ```json
      [
      {
          "userID": "user id",
          "postID": "post id",
          "content": "comment content",
          "date_create": "date"
      }
      ]
      ``` 

### License

```
   Copyright (C) 2023 Amit Shekhar

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```