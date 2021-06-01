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
- `GET /categories/:category_name`

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

### jobs

- `GET /jobs/all`
- `GET/jobs/all/?page={}&limit={}`
- `GET /jobs/:id`

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

## RESTful endpoints jobs

### GET /jobs/all

> Get all job list

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
        "message": "get all jobs succeed",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "title": "(Senior) PHP developer",
            "url": "https://remotive.io/remote-jobs/software-dev/senior-php-developer-690639",
            "company_name": "IvyPanda",
            "job_type": "full_time",
            "publication_date": "2021-06-01T01:39:14",
            "description": "<div class=\"h3\">Company Description</div>\n<p>IvyPanda is a student success hub designed to improve the educational outcomes and learning capabilities of students around the world by connecting them with academic experts and by providing highly-efficient self-study services and online tools.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Job Description</div>\n<ul>\n<li>Write well-designed, clean code according to detailed specifications</li>\n<li>Ensure strong optimization and functionality of the core product software through testing, maintaining, and troubleshooting</li>\n<li>Make contributions at all stages of the development process</li>\n<li>Develop and deploy new features</li>\n<li>Communicate and collaborate with the rest of the development team</li>\n<li>Follow and learn the new best practices of the industry</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Qualifications</div>\n<ul>\n<li>PHP software development experience</li>\n<li>Understanding and knowledge of open-source projects and web technologies</li>\n<li>Experience in common third-party APIs</li>\n<li>Attention to detail and good communication skills</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Additional Information</div>\n<p>All your information will be kept confidential according to EEO guidelines.</p>\n",
            "candidate_required_location": "USA Only"
        },
        {
            "id": 2,
            "title": "Backend Developer",
            "url": "https://remotive.io/remote-jobs/software-dev/backend-developer-690718",
            "company_name": "iBizpeople Ltd",
            "job_type": "full_time",
            "publication_date": "2021-06-01T01:39:12",
            "description": "<div class=\"h3\">Company Description</div>\n<p>iBizpeople is an international Talent Acquisition consultancy that helps top performers in the iGaming, Digital, eCommerce and Tech industries, as well as up and coming talent, find permanent positions with innovative start ups and established forward thinking organisations. Not only do we find you the role, but we also guide you through every step of the interview process so that you will be fully prepared for what happens next. Our aim is to match you with the perfect role for you in companies that will help develop your talent and further your career.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Job Description</div>\n<p>My client, a fast growing iGaming Affiliate Marketing company, is looking to recruit a talented Backend Developer to join the team on a remote basis. If you have excellent programming skills and a passion for developing applications or improving existing ones, my client would like to meet you. As a Back-End Developer, you’ll work closely with their engineers to ensure system consistency and improve the user experience.</p>\n<p>You should be able to develop and maintain functional and stable web applications and API's to meet our needs.</p>\n<p> </p>\n<p><strong>Responsibilities</strong></p>\n<ul>\n<li>Participate in the entire application life-cycle, focusing on coding and debugging</li>\n<li>Write clean code to develop functional web applications and API's</li>\n<li>Troubleshoot and debug applications</li>\n<li>Perform tests to optimize performance</li>\n<li>Manage cutting-edge technologies to improve legacy sensitive and large-scale applications</li>\n<li>Collaborate with Front-End Developers to integrate user-facing elements with server-side logic</li>\n<li>Gather and address technical and design requirements</li>\n<li>Provide training and support to internal teams</li>\n<li>Build reusable code and libraries for future use</li>\n<li>Liaise with developers, designers, and system administrators to identify new features</li>\n<li>Follow emerging technologies</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Qualifications</div>\n<ul>\n<li>4+ years of commercial experience within a similar role e.g. Backend Developer, PHP Developer, Backend Engineer.</li>\n<li>Experience working with PHP frameworks such as Laravel.</li>\n<li>Experience with Wordpress development, REST / GraphQL API and VueJS experience is a plus.</li>\n<li>Strong problem solving skills and ability to troubleshoot operational issues.</li>\n<li>Ability to write and optimize MySQL queries and display an overall understanding of relational databases.</li>\n<li>Excellent communicator with strong verbal and writing skills in English.</li>\n<li>Degree in Computer Science, Computer Engineering, Applied Math or in any related field is a plus but not a must.</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Additional Information</div>\n<ul>\n<li>An agile and multicultural company</li>\n<li>Competitive salary</li>\n<li>Relocation Assistance</li>\n<li>Health Insurance</li>\n<li>Daily refreshments</li>\n<li>Social events</li>\n<li>Focus on your growth and development</li>\n</ul>\n<p>**Applications in English only please</p>\n",
            "candidate_required_location": "Malta"
        },
        ...
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

### GET /jobs/all/?page={page}&limit={how many jobs to show}

> Get jobs with query Ex: GET /jobs/all/?page=1&limit=5

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
        "message": "get all jobs succeed",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "title": "(Senior) PHP developer",
            "url": "https://remotive.io/remote-jobs/software-dev/senior-php-developer-690639",
            "company_name": "IvyPanda",
            "job_type": "full_time",
            "publication_date": "2021-06-01T01:39:14",
            "description": "<div class=\"h3\">Company Description</div>\n<p>IvyPanda is a student success hub designed to improve the educational outcomes and learning capabilities of students around the world by connecting them with academic experts and by providing highly-efficient self-study services and online tools.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Job Description</div>\n<ul>\n<li>Write well-designed, clean code according to detailed specifications</li>\n<li>Ensure strong optimization and functionality of the core product software through testing, maintaining, and troubleshooting</li>\n<li>Make contributions at all stages of the development process</li>\n<li>Develop and deploy new features</li>\n<li>Communicate and collaborate with the rest of the development team</li>\n<li>Follow and learn the new best practices of the industry</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Qualifications</div>\n<ul>\n<li>PHP software development experience</li>\n<li>Understanding and knowledge of open-source projects and web technologies</li>\n<li>Experience in common third-party APIs</li>\n<li>Attention to detail and good communication skills</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Additional Information</div>\n<p>All your information will be kept confidential according to EEO guidelines.</p>\n",
            "candidate_required_location": "USA Only"
        },
        {
            "id": 2,
            "title": "Backend Developer",
            "url": "https://remotive.io/remote-jobs/software-dev/backend-developer-690718",
            "company_name": "iBizpeople Ltd",
            "job_type": "full_time",
            "publication_date": "2021-06-01T01:39:12",
            "description": "<div class=\"h3\">Company Description</div>\n<p>iBizpeople is an international Talent Acquisition consultancy that helps top performers in the iGaming, Digital, eCommerce and Tech industries, as well as up and coming talent, find permanent positions with innovative start ups and established forward thinking organisations. Not only do we find you the role, but we also guide you through every step of the interview process so that you will be fully prepared for what happens next. Our aim is to match you with the perfect role for you in companies that will help develop your talent and further your career.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Job Description</div>\n<p>My client, a fast growing iGaming Affiliate Marketing company, is looking to recruit a talented Backend Developer to join the team on a remote basis. If you have excellent programming skills and a passion for developing applications or improving existing ones, my client would like to meet you. As a Back-End Developer, you’ll work closely with their engineers to ensure system consistency and improve the user experience.</p>\n<p>You should be able to develop and maintain functional and stable web applications and API's to meet our needs.</p>\n<p> </p>\n<p><strong>Responsibilities</strong></p>\n<ul>\n<li>Participate in the entire application life-cycle, focusing on coding and debugging</li>\n<li>Write clean code to develop functional web applications and API's</li>\n<li>Troubleshoot and debug applications</li>\n<li>Perform tests to optimize performance</li>\n<li>Manage cutting-edge technologies to improve legacy sensitive and large-scale applications</li>\n<li>Collaborate with Front-End Developers to integrate user-facing elements with server-side logic</li>\n<li>Gather and address technical and design requirements</li>\n<li>Provide training and support to internal teams</li>\n<li>Build reusable code and libraries for future use</li>\n<li>Liaise with developers, designers, and system administrators to identify new features</li>\n<li>Follow emerging technologies</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Qualifications</div>\n<ul>\n<li>4+ years of commercial experience within a similar role e.g. Backend Developer, PHP Developer, Backend Engineer.</li>\n<li>Experience working with PHP frameworks such as Laravel.</li>\n<li>Experience with Wordpress development, REST / GraphQL API and VueJS experience is a plus.</li>\n<li>Strong problem solving skills and ability to troubleshoot operational issues.</li>\n<li>Ability to write and optimize MySQL queries and display an overall understanding of relational databases.</li>\n<li>Excellent communicator with strong verbal and writing skills in English.</li>\n<li>Degree in Computer Science, Computer Engineering, Applied Math or in any related field is a plus but not a must.</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Additional Information</div>\n<ul>\n<li>An agile and multicultural company</li>\n<li>Competitive salary</li>\n<li>Relocation Assistance</li>\n<li>Health Insurance</li>\n<li>Daily refreshments</li>\n<li>Social events</li>\n<li>Focus on your growth and development</li>\n</ul>\n<p>**Applications in English only please</p>\n",
            "candidate_required_location": "Malta"
        },
        {
            "id": 3,
            "title": "Sr. Java Developer",
            "url": "https://remotive.io/remote-jobs/software-dev/sr-java-developer-690846",
            "company_name": "TMS LLC",
            "job_type": "contract",
            "publication_date": "2021-06-01T01:39:11",
            "description": "<div class=\"h3\">Company Description</div>\n<p>Role: Sr Java Developer</p>\n<p>Location: Remote Opportunity</p>\n<p>Duration: 6+ Months</p>\n<p>Sr. JAVA Full Stack Application Developer (Microservices, AWS Cloud, JAVA, J2EE, needed with AWS Cloud, DevOps, DevSecOps, Microservices Development / Arch / Implementation / Scaling of Microservices, JAVA / J2EE, Spring framework, Spring Boot, Apache KAFKA, JSON Formats / Strings, Kubernetes Clusters, JAVA Closures, XAPI's, Terraforms, RabbitMQ, Kubernetes Clusters, Docker Containers, Jenkins, Automated deployment tools and Docker Orchestration.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Job Description</div>\n<ul>\n<li>Hands-on development experience in AWS cloud environment.</li>\n</ul>\n<ul>\n<li>Experience in Microservices Development, Architecture and hands-on implementation experience.</li>\n<li>Strong skills in JAVA / J2EE, Spring framework and Spring Boot - we have some options with language.</li>\n<li>Strong skills in Apache KAFKA.</li>\n<li>Experience with JAVA Closures, XAPI's, Terraforms and RabbitMQ.</li>\n<li>Strong hands-on skills with JSON Formats / Strings.</li>\n<li>Good experience with DevOps &amp; DevSecOps.</li>\n<li>Scaling of Microservices – Kubernetes Clusters / Docker Containers.</li>\n<li>Experience with Automated deployment tools and Docker Orchestration.</li>\n<li>Experience with Jenkins.</li>\n</ul>\n<div class=\"h3\"> </div>\n<p> </p>\n",
            "candidate_required_location": "USA Only"
        },
        {
            "id": 4,
            "title": "JavaScript Reverse Engineer",
            "url": "https://remotive.io/remote-jobs/software-dev/javascript-reverse-engineer-690470",
            "company_name": "Eyeo",
            "job_type": "full_time",
            "publication_date": "2021-06-01T01:39:09",
            "description": "<div class=\"h3\">Get to know us</div>\n<p>eyeo is an open-source software company that builds products like Adblock Plus, Adblock Browser and Flattr. By leveraging distribution partnerships, we bring ad-blocking technology everywhere, giving users control over their online experience while offering creators, publishers and advertisers more ways to earn money for the free content they provide.</p>\n<p>In combining our reach based on distribution partnerships and our own products, our technology runs on over 230 million devices.</p>\n<p>At eyeo, we’re passionate about user agency, personal privacy, sustainability and keeping the web an open, fair resource for everyone.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">How we work</div>\n<p>eyeo colleagues are based all over the world. We practice agile and work in distributed, cross-functional teams that span nearly every timezone. Many of our tech teams prefer to work asynchronously.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">What you'll do</div>\n<p>You will be responsible for developing small Javascript functions to be used by adblocking filterlist authors, identify/reverse engineer adblocking circumvention technologies used and propose ways to counter them. To do so, skills in software development and engineering, as well reverse engineering skills are essential.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">After your morning coffee, you'll be expected to...</div>\n<ul>\n<li>Collaborate closely with other Developers, QAs and Product Owners</li>\n<li>Debug compiled JS code, deepdive into HTML/CSS page sources in order to write snippets which prevent circumvention of ad blocking technologies</li>\n<li>Write unit tests, review merge requests, pair program with other engineers in order to monitor and improve code quality </li>\n<li>Drive innovation by suggesting new ideas related to ongoing initiatives or existing challenges to the team and following up on these ideas</li>\n<li>Provide technical expertise wherever needed</li>\n<li>Foster open communication and close collaboration with other team members to strengthen the team and increase overall productivity</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">What you bring to the table...</div>\n<ul>\n<li>3+ years of software development experience with good JavaScript skills</li>\n<li>Strong knowledge of HTML/CSS</li>\n<li>Experience in reverse engineering </li>\n<li>Knowledge and experience writing Unit Tests</li>\n<li>A solid base knowledge in Agile ways of working</li>\n<li>High desire to collaborate and improve team culture</li>\n<li>Fluent in both verbal and written English</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">It's awesome, but not required, if you know about...</div>\n<ul>\n<li>Knowledge of browser internals</li>\n<li>Knowledge around Video Streaming Technologies and Protocols</li>\n<li>Experience with intercepting network requests</li>\n<li>Understanding of the ad tech landscape</li>\n<li>Experience with browser extension development</li>\n<li>Experience working on open source projects</li>\n<li>Experience around collaborating with or even managing tech communities</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">What we offer</div>\n<ul>\n<li>Work from home or one of our offices —we trust you to find what works best for you</li>\n<li>Stipend for one of the following: home office, co-working space, or relocation</li>\n<li>Flexible working hours</li>\n<li>28 days paid vacation days </li>\n<li>Your choice of hardware and setup</li>\n<li>Personal and professional development budget</li>\n<li>Monthly childcare stipend for children under 6</li>\n<li>Offsite team days and annual summer company retreat in Cologne</li>\n<li>Company-sponsored hackathons</li>\n</ul>\n<div class=\"h4\"> </div>\n<div class=\"h4\">Privacy Notice</div>\n<p><em>When you apply, you’ll be automatically forwarded to our recruitment platform operated by an external service provider called Greenhouse (seated in the US). Greenhouse collects some information on its website, such as anonymous usage statistics, by using cookies, server logs, and other similar technology. For more information, please refer to Greenhouse’s Privacy Policy. All documents and information provided by you are stored with Greenhouse. In order to ensure an adequate level of data protection, eyeo and Greenhouse have entered into the EU Standard Contractual Clauses (“processors”) - Commission Decision C(2010)593. You can request a copy of this by contacting us at privacy[at]eyeo.com. If you don’t want your data forwarded to Greenhouse, please do not apply. For detailed and further information, please refer to our Privacy Policy at https://eyeo.com/en/privacy.</em></p>\n",
            "candidate_required_location": "Germany Only"
        },
        {
            "id": 5,
            "title": "Principal Software Engineer, Marketplace",
            "url": "https://remotive.io/remote-jobs/software-dev/principal-software-engineer-marketplace-690676",
            "company_name": "TCGplayer",
            "job_type": "full_time",
            "publication_date": "2021-06-01T01:39:08",
            "description": "<div class=\"h1\"><strong>Principal Software Engineer</strong></div>\n<p><strong>Location: Remote</strong></p>\n<div class=\"h2\"><strong>Who We Are</strong></div>\n<p>The Engineering Team at TCGplayer creates and maintains a suite of products in support of our role as the leading online marketplace for trading card games and collectibles. We build applications and technologies that connect thousands of businesses with customers across the collectible gaming industry, powering sales through physical stores, websites, mobile apps and the TCGplayer Marketplace. We are looking for a highly experienced developer for our eCommerce platform level architectures who is results-driven, inventive, an analytical problem solver, and a good communicator.</p>\n<div class=\"h2\"> </div>\n<div class=\"h2\"><strong>Who You Are</strong></div>\n<p>If you are a seasoned developer with full stack experience in numerous technologies that’s ready to get our Marketplace set to support massive growth in an incredibly fun industry, then TCGplayer is the right place for you.</p>\n<div class=\"h2\"> </div>\n<div class=\"h2\"><strong>The Impact You Will Make Here</strong></div>\n<ul>\n<li>Set development trends and lead industry-level growth and improvement across the company. </li>\n<li>Be a driver for many of our key projects here at TCGplayer.</li>\n<li>Lead large technical projects and establish architectural direction with technical teams, and mentor more junior engineers. </li>\n<li>Architect, design and develop platform-level features of our eCommerce website working with some of the best engineers in development today. </li>\n<li>Work on complex systems for managing product searches, APIs for internal and 3rd party consumption, background system jobs, stream-based processing and caching. </li>\n<li>Design best-in-class solutions for our buyers and create platforms for monitoring, managing and easing throughput of our services to ensure service response times are fast, orders are processed appropriately, and data integrity is maintained.</li>\n</ul>\n<div class=\"h2\"> </div>\n<div class=\"h2\"><strong>What You Bring To The Team</strong></div>\n<ul>\n<li>Bachelor’s degree in computer science, information technology or a related field, or a combination of education and experience</li>\n<li>12+ years of  C#, .NET MVC Razor, Entity Framework, .NET Core and ASP.NET experience</li>\n<li>Demonstrated success leading large, strategic, or complex projects that have clear impact to the success of your company.</li>\n<li>Experience setting architectural direction at the organization level and evolving systems toward it for long term success.</li>\n<li>A track record of influencing and delivering projects with organizational impact and priority.</li>\n<li>The ability to quickly analyze and resolve tactical issues across a wide variety of areas. </li>\n<li>Mastery in C# and experience with at least one additional programming language, such as Javascript, Go or Java.</li>\n<li>Expert SQL programming skills: Stored Procedures, Views, and Functions</li>\n<li>Understanding of a NoSQL datastore such as Redis, DynamoDB and MongoDB</li>\n<li>Experience designing and developing microservices using cloud technology with considerations for concurrency and parallelism</li>\n<li>Experience with event processing mechanisms such as Kafka, AWS SNS/SQS, RabbitMQ</li>\n<li>Experience working with Kubernetes, Docker, and public cloud deployment platforms</li>\n</ul>\n<div class=\"h2\"> </div>\n<div class=\"h2\"><strong>What We Provide</strong></div>\n<p>Our benefits program is one of the most flexible and progressive in the country. Plus, benefits start on day one, so you have everything you need to make a stress-free transition to life at TCGplayer.</p>\n<ul>\n<li>Comprehensive medical insurance with a variety of plan options to suit your needs</li>\n<li>Dental and vision insurance</li>\n<li>Unlimited Paid Time Off (PTO) </li>\n<li>100% company paid Family Leave</li>\n<li>401k plan with 4% match </li>\n<li>Stock options for all employees</li>\n<li>Free coaching, counseling and mental health services</li>\n<li>100% company paid life insurance </li>\n<li>Paid trips to work with remote teammates</li>\n</ul>\n<div class='\"content-conclusion\"'><hr>\n<p>TCGplayer is the largest marketplace for trading card games in the world, with nearly 1 billion dollars in sales. Founded by Chedy Hampson &amp; Ray Moore as a digital media platform within the collectible hobby space, we have grown from our initial roots working inside Syracuse-based hobby stores selling comic books, sports cards, CD’s, collectible cards, action figures, and tabletop games, into an expansive eCommerce marketplace that connects a global community of millions of buyers with tens of thousands of retailers. TCGplayer maintains the largest authoritative database of historical and current prices for the most popular trading card games, and this data powers and connects an entire industry of publishers, buyers, sellers, influencers &amp;amp; developers.</p>\n<p><strong>We have been ranked amongst New York State’s 50 best employers and Fortune.com’s top 100 companies for women in the U.S and certified a Great Place to Work by our employees 5 years in a row.</strong> Our entire team prides itself on creating a culture that fosters camaraderie, embraces diversity, and exudes passion. We believe every team member contributes to our success and as such, we provide stock options to all 400 employees, growing to 500 strong by the end of 2021.</p>\n<p>With the largest marketplace for collectible card games, in-store tools for local gaming stores and industry leading fulfillment center to deliver products from sellers to buyers, we bring our innovative technologies and customer focused approach to the $25B global collectibles hobby market. The company’s mission is to establish TCGplayer as the most valued and admired company in the hobby collectibles universe by 2025.</p>\n</div>\n",
            "candidate_required_location": "USA Only"
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

### GET /jobs/:id

> Get a job by id

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
        "id": 1,
        "title": "(Senior) PHP developer",
        "url": "https://remotive.io/remote-jobs/software-dev/senior-php-developer-690639",
        "company_name": "IvyPanda",
        "job_type": "full_time",
        "publication_date": "2021-06-01T01:39:14",
        "description": "<div class=\"h3\">Company Description</div>\n<p>IvyPanda is a student success hub designed to improve the educational outcomes and learning capabilities of students around the world by connecting them with academic experts and by providing highly-efficient self-study services and online tools.</p>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Job Description</div>\n<ul>\n<li>Write well-designed, clean code according to detailed specifications</li>\n<li>Ensure strong optimization and functionality of the core product software through testing, maintaining, and troubleshooting</li>\n<li>Make contributions at all stages of the development process</li>\n<li>Develop and deploy new features</li>\n<li>Communicate and collaborate with the rest of the development team</li>\n<li>Follow and learn the new best practices of the industry</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Qualifications</div>\n<ul>\n<li>PHP software development experience</li>\n<li>Understanding and knowledge of open-source projects and web technologies</li>\n<li>Experience in common third-party APIs</li>\n<li>Attention to detail and good communication skills</li>\n</ul>\n<div class=\"h3\"> </div>\n<div class=\"h3\">Additional Information</div>\n<p>All your information will be kept confidential according to EEO guidelines.</p>\n",
        "candidate_required_location": "USA Only"
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
