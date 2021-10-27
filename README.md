# todoapp

This is a simple todoapp build with Go. 

## Pre-requisites

### 1. Go

~~~~
[https://golang.org/], any one of the three latest major releases of Go.
For installation instructions, see Go's getting started guide: https://golang.org/doc/install
~~~~

## How to Run it?

**1. Clone the repo in your local environment:**

~~~~
git clone https://github.com/flakrimjusufi/todoapp.git
~~~~

**2. Run the server:**

`go run main.go`

~~~~
You should recieve a response:
Starting server in localhost:8088
~~~~

**3. Send a GET request using cURL or Postman/Insomnia:**

`curl -X GET -k http://localhost:8088/getTodoList`

~~~~
You should have a response from server: 

[
  {
    "id": 1,
    "labels": "Clean dishes",
    "comments": "Don't forget the garbage too..",
    "due_date": "2021-10-28T07:40:00Z"
  },
  {
    "id": 2,
    "labels": "Send kids to school",
    "comments": "Don't forget the Umbrella..",
    "due_date": "2021-10-28T08:10:00Z"
  },
  {
    "id": 3,
    "labels": "Go to work",
    "comments": "Do some small-talks with colleagues before starting the work..",
    "due_date": "2021-10-28T08:30:00Z"
  }
]
~~~~