# todoapp

### This is a simple todoapp built with Go.

## How to Run it?

** Clone the repo in your local environment:**

~~~~
git clone https://github.com/flakrimjusufi/todoapp.git
~~~~

**In case you have docker-compose installed in your machine, just execute the following:**

~~~~
docker-compose up
~~~~

You should see an output in terminal:

~~~
Starting server in http://localhost:8088
~~~

Open your browser in [http://localhost:8088](http://localhost:8088) or Send a GET request using cURL or Postman/Insomnia:

`curl -X GET -k http://localhost:8088/getTodoList`

...and you should receive a response from the server with a ToDoList.

### In case you don't have docker installed, there are some pre-requisites:

**1. Install go:**

~~~~
[https://golang.org/], any one of the three latest major releases of Go.
For installation instructions, see Go's getting started guide: https://golang.org/doc/install
~~~~

**2. Run the server:**

`go run main.go`

~~~~
You should recieve a response:
Starting server in http://localhost:8088
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