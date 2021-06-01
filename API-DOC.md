# StuckOverflow

```
https://stuckoverflow.herokuapp.com
```

## List of available endpoints

### users

- `GET /users`
- `POST /users/register`
- `POST /users/login`
- `GET /users/:id`
- `PUT /users/:id`
- `DELETE /users/:id`

### categories

- `GET /categories`
- `POST /categories`

### questions

- `GET /questions`
- `GET /questions/:id`
- `POST /questions/ask`
- `PUT /questions/:id/edit`
- `DELETE /questions/:id`
- `GET /questions/categories/:category_name`

### answers

- `POST /questions/:id`
- `PUT /answers/:id`
- `DELETE /answers/:id`

## RESTful endpoints users

### GET /users

> Get All users

_Request Header_

```go
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "get all users succeed",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "user_name": "marwan",
            "email": "marwan@mail.com",
            "first_name": "marwan",
            "last_name": "marwan",
            "avatar": "",
            "location": "Indonesia"
        },
        {
            "id": 2,
            "user_name": "kezia",
            "email": "kezia@mail.com",
            "first_name": "",
            "last_name": "",
            "avatar": "",
            "location": ""
        }
    ]
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

---

### POST /users/register

> Create new user

_Request Header_

```go
not needed
```

_Request Body_

```go
{
  "user_name": "<user name to get insert into>",
  "email": "<email to get insert into>",
  "password": "<password to get insert into>"
}
```

_Response (201)_

```go
{
    "meta": {
        "message": "insert user data succeed",
        "code": 201,
        "status": "success"
    },
    "data": {
        "id": 1,
        "user_name": "marwan",
        "email": "marwan@mail.com",
    }
}
```

_Response (400 - Bad Request)_

```go
{
  "meta": {
      "message": "input data required",
      "code": 400,
      "status": "bad request"
  },
  "data":
      {
        "errors": []
      }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

---

### POST /users/login

> Compare data login on database with data inputed

_Request Header_

```go
not needed
```

_Request Body_

```go
{
  "email": "<email to get compare>",
  "password": "<password to get compare>"
}
```

_Response (200)_

```go
{
    "meta": {
        "message": "login user succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "token": "<your token>"
    }
}
```

_Response (400 - Bad Request)_

```go
{
  "meta": {
      "message": "input data required",
      "code": 400,
      "status": "bad request"
  },
  "data":
      {
        "errors": []
      }
}
```

_Response (401 - Unauthorized)_

```go
{
    "meta": {
      "message": "input data error",
      "code": 401,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

---

### GET /users/:user_id

> Get user by user ID

_Request Header_

```go
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "get user succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "user_name": "marwan",
        "email": "marwan@mail.com",
        "first_name": "marwan",
        "last_name": "marwan",
        "avatar": "",
        "location": "Indonesia"
    }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  },
  "data" : {
      "error": ""
  }
}
```

---

### PUT /users/:user_id

> Update user by User iD

_Request Header_

```go
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```go
{
    "last_name": "marwan2"
}
```

_Response (200)_

```go
{
    "meta": {
        "message": "update user succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "user_name": "marwan",
        "email": "marwan@mail.com",
        "first_name": "marwan",
        "last_name": "marwan2",
        "avatar": "",
        "location": "Indonesia"
    }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  },
  "data" : {
      "error": ""
  }
}
```

---

### DELETE /users/:user_id

> Delete user by ID

_Request Header_

```go
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "user was deleted successfully",
        "code": 200,
        "status": "success"
    },
    "data": {
        "message": "delete user id 2 succeed",
        "delete_time": "2021-05-31T17:21:59.7646992+07:00"
    }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  },
  "data" : {
      "error": ""
  }
}
```

## RESTful endpoints categories

### GET /categories

> Get all categories

_Request Header_

```go
not needed
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "get all categories succeed",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "category_name": "Golang",
            "description": "description golang"
        },
        {
            "id": 2,
            "category_name": "Python",
            "description": "description python"
        }
    ]
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

---

### POST /categories

> Add new category

_Request Header_

```go
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```go
{
  "category_name": "<category name>",
  "description": "<description here>"
}
```

_Response (200)_

```go
{
    "meta": {
        "message": "insert new category succeed",
        "code": 201,
        "status": "success"
    },
    "data": {
        "id": 4,
        "category_name": "CSS",
        "description": "description"
    }
}
```

_Response (400 - Bad Request)_

```go
{
  "meta": {
      "message": "input data required",
      "code": 400,
      "status": "bad request"
  },
  "data":
      {
        "errors": []
      }
}
```

_Response (401 - Unauthorized)_

```go
{
    "meta": {
      "message": "input data error",
      "code": 401,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

---

### GET /categories/:category_name

> Get questions by category name

_Request Header_

```go
not needed
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "get category succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "category_name": "Golang",
        "description": "",
        "questions": [
            {
                "id": 2,
                "title": "The standard Lorem Ipsum passage, used since the 1500s",
                "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
                "image_file": "",
                "user_id": 1,
                "answers": []
            }
        ]
    }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  },
  "data" : {
      "error": ""
  }
}
```

---

### PUT /categories/:category_name

> update category by category name

_Request Header_

```go
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```go
{
    "category_name" : "<category name to update>",
    "description" : "<description here>"
}
```

_Response (200)_

```go
{
    "meta": {
        "message": "success update category by name",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 2,
        "category_name": "Python",
        "description": "description python"
    }
}
```

_Response (400 - Bad Request)_

```go
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  },
  "data" : {
      "errors" : []
  }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  },
  "data" : {
      "error" : ""
  }
}
```

## RESTful endpoints questions

### GET /questions

> Get all questions

_Request Header_

```go
not needed
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "get all questions succeed",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 2,
            "title": "The standard Lorem Ipsum passage, used since the 1500s",
            "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
            "user_id": 1,
            "user_name": "marwan",
            "category_id": 1,
            "category_name": "Golang",
            "answers": []
        },
        {
            "id": 3,
            "title": "The standard Lorem Ipsum passage, used since the 1500s",
            "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
            "user_id": 1,
            "user_name": "marwan",
            "category_id": 2,
            "category_name": "Python",
            "answers": []
        }
    ]
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

---

### GET /questions/:id

> Get a question by question id

_Request Header_

```go
not needed
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "get question succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 2,
        "title": "The standard Lorem Ipsum passage, used since the 1500s",
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        "user_id": 1,
        "user_name": "marwan",
        "category_id": 1,
        "category_name": "Golang",
        "answers": []
    }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

### POST /questions/ask

> Create a new question

_Request Header_

```go
{
    "Authorization" : "<your Authorization>"
}
```

_Request Body_

```go
{
  "title" : "<title to get insert into>",
  "content" : "<content to get insert into>",
  "category_id": "<category id to get insert into>",
}
```

_Response (201)_

```go
{
    "meta": {
        "message": "insert new question succeed",
        "code": 201,
        "status": "success"
    },
    "data": {
        "id": 3,
        "title": "The standard Lorem Ipsum passage, used since the 1500s",
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        "user_id": 1,
        "category_id": 2
    }
}
```

_Response (400 - Bad Request)_

```go
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  },
  "data" :
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  },
  "data" :
      {
        "error" : ""
      }
}
```

---

### PUT /questions/:id/edit

> Update question by id

_Request Header_

```go
{
    "Authorization" : "<your Authorization>"
}
```

_Request Body_

```go
{
  "title" : "<title to update>",
  "content" : "<content to update>",
  "category_id": "<category id to update>",
}
```

_Response (201)_

```go
{
    "meta": {
        "message": "update question succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 2,
        "title": "Edit The standard Lorem Ipsum passage, used since the 1500s",
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        "user_id": 1,
        "user_name": "marwan",
        "category_id": 1,
        "category_name": "Golang",
        "answers": []
    }
}
```

_Response (400 - Bad Request)_

```go
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  },
  "data" :
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  },
  "data" :
      {
        "error" : ""
      }
}
```

---

### DELETE /questions/:id

> Delete question by ID

_Request Header_

```
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "question was deleted successfully",
        "code": 200,
        "status": "success"
    },
    "data": {
        "message": "delete question id 3 succeed",
        "delete_time": "2021-05-31T22:15:59.8456443+07:00"
    }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  },
  "data" : {
      "error": ""
  }
}
```

### GET /questions/categories/:category_name

> Get all questions by category name

_Request Header_

```go
not needed
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "get all questions by category succeed",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 2,
            "title": "Edit The standard Lorem Ipsum passage, used since the 1500s",
            "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
            "user_id": 1,
            "user_name": "marwan",
            "category_id": 1,
            "category_name": "Golang",
            "answers": []
        },
        {
            "id": 5,
            "title": "The standard Lorem Ipsum passage, used since the 1500s",
            "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
            "user_id": 1,
            "user_name": "marwan",
            "category_id": 1,
            "category_name": "Golang",
            "answers": []
        }
    ]
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  },
  "data":
      {
        "error": ""
      }
}
```

## RESTful endpoints answers

### POST /questions/:id

> Create a new answer by question id

_Request Header_

```go
{
    "Authorization" : "<your Authorization>"
}
```

_Request Body_

```go
{
  "content" : "<content to get insert into>"
}
```

_Response (201)_

```go
{
    "meta": {
        "message": "insert new answer succeed",
        "code": 201,
        "status": "success"
    },
    "data": {
        "id": 1,
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        "image_file": "",
        "user_id": 1,
        "question_id": 3
    }
}
```

_Response (400 - Bad Request)_

```go
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  },
  "data" :
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  },
  "data" :
      {
        "error" : ""
      }
}
```

---

### PUT /answers/:id

> Update answer by id

_Request Header_

```go
{
    "Authorization" : "<your Authorization>"
}
```

_Request Body_

```go
{
  "content" : "<content to update>"
}
```

_Response (201)_

```go
{
    "meta": {
        "message": "update answer succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "content": "Edit Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        "user_id": 1,
        "question_id": 3
    }
}
```

_Response (400 - Bad Request)_

```go
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  },
  "data" :
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  },
  "data" :
      {
        "error" : ""
      }
}
```

---

### DELETE /questions/:id

> Delete question by ID

_Request Header_

```go
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```go
not needed
```

_Response (200)_

```go
{
    "meta": {
        "message": "question was deleted successfully",
        "code": 200,
        "status": "success"
    },
    "data": {
        "message": "delete answer id 1 succeed",
        "delete_time": "2021-05-31T22:51:09.0561205+07:00"
    }
}
```

_Response (500 - Internal Server Error)_

```go
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  },
  "data" : {
      "error": ""
  }
}
```

##
