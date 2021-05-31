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
- `POST /users/:id`
- `PUT /users/:id`
- `DELETE /users/:id`

### questions

- `GET /questions`
- `GET /questions/:id`
- `POST /questions/ask`
- `PUT /questions/:id/edit`
- `DELETE /questions/:id`

### answers

- `POST /questions/:id`
- `PUT /answers/:id`
- `DELETE /answers/:id`

### categories

- `GET /categories`
- `POST /categories`
- `PUT /categories/:category_name`

## RESTful endpoints users

### GET /users

> Get All users

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

```
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
            "first_name": "marwan",
            "last_name": "juna",
            "email": "marwan@mail.com"
        },
        {
            "id": 2,
            "user_name": "uchihasasuke",
            "first_name": "uchiha",
            "last_name": "saske",
            "email": "uchiha@mail.com"
        }
    ]
}
```

_Response (500 - Internal Server Error)_

```
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

```
not needed
```

_Request Body_

```
{
  "user_name": "<user name to get insert into>",
  "email": "<email to get insert into>",
  "password": "<password to get insert into>"
}
```

_Response (201)_

```
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

```
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

```
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

```
not needed
```

_Request Body_

```
{
  "email": "<email to get compare>",
  "password": "<password to get compare>"
}
```

_Response (200)_

```
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

```
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

```
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

```
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

```
{
    "meta": {
        "message": "get user succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 3,
        "user_name": "boruto",
        "first_name": "uzumaki",
        "last_name": "boruto",
        "email": "boruto@mail.com"
    }
}
```

_Response (500 - Internal Server Error)_

```
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

```
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_

```
{
    "last_name": "boruto"
}
```

_Response (200)_

```
{
    "meta": {
        "message": "update user succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 3,
        "user_name": "boruto",
        "first_name": "boruto",
        "last_name": "boruto",
        "email": "boruto@mail.com"
    }
}
```

_Response (500 - Internal Server Error)_

```
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

```
{
    "meta": {
        "message": "user was deleted successfully",
        "code": 200,
        "status": "success"
    },
    "data": {
        "message": "delete user id 3 succeed",
        "delete_time": "2021-05-27T05:55:25.705047677Z"
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

## RESTful endpoints questions

### GET /questions

> Get all questions

_Request Header_

```
not needed
```

_Request Body_

```
not needed
```

_Response (200)_

```
{
    "meta": {
        "message": "get all questions succeed",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 2,
            "title": "ini judul 2",
            "content": "ini content 2",
            "user_id": 2,
            "category_id": 2,
            "category_name": "Python",
            "answers": []
        }
    ]
}
```

_Response (500 - Internal Server Error)_

```
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
