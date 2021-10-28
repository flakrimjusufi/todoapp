# todoapp

### This is a simple todoapp built with Go.

## How to Run it?

**Clone the repo in your local environment:**

~~~~
git clone https://github.com/flakrimjusufi/todoapp.git
~~~~

*Remove ".example" from both .env.example and /seeds/init.sql-example and populate them with your environment variables.* 


**In case you have docker-compose installed in your machine, just execute the following:**

~~~~
docker-compose up
~~~~

Docker-compose will build all the dependencies and will add a PostgreSQL image in your container so that we can interact 
with data. 

*Once the docker-compose is finished, you should see an output in terminal:*

~~~
Starting server in http://localhost:8088
~~~

Open your browser in [http://localhost:8088](http://localhost:8088) or Send a GET request using cURL or Postman/Insomnia:

`curl -X GET -k http://localhost:8088/getSampleTodoList`

...and you should receive a response from the server with a ToDoList.

### In case you don't have docker installed, you need to do the following:

**1. Install go:**

~~~~
[https://golang.org/], any one of the three latest major releases of Go.
For installation instructions, see Go's getting started guide: https://golang.org/doc/install
~~~~

**2. Install PostgreSQL (or any SQL database)**

~~~~
[Version 9+]
For installation instructions, please refer to this link: https://www.postgresql.org/download/
~~~~

**2. Remove ".example" from both .env.example and /seeds/init.sql-example and populate them with your environment variables.**

**3. Execute the script for creating the database in /seeds/init.sql**

**4. Run the server:**

`go run main.go`

~~~~
You should recieve a response:
Starting server in http://localhost:8088
~~~~

**3. Send a GET request using cURL or Postman/Insomnia:**

`curl -X GET -k http://localhost:8088/getSampleTodoList`

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

## Interacting with API 

Once the server is up and running, you can hit the endpoints listed in this table to interact with data: 

| HTTP call        | Endpoint           | Description  |
| :-------------: |:-------------:| :-----:|
| GET     | http://localhost:8088/getSampleTodoList | Will list a sample of ToDoList |
| POST      | http://localhost:8088/toDoList     |  Will create a ToDoList in database |
| GET | http://localhost:8088/toDoList      |   Will list all toDos created in database |
| PUT | http://localhost:8088/toDoList/1    |   Will update a toDo in database and send back the response |
| DELETE | http://localhost:8088/toDoList/1    |   Will delete a toDo by Id in database and send a successful message back |
| GET     | http://localhost:8088/ | Will render a UI in browser with some sample ToDoList |
| GET     | http://localhost:8088/viewToDoList | Will render a UI in browser with ToDos created in database|