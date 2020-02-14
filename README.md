# Simple Messaging API

API for sending and retrieve message. 

## Getting Started

### Prerequisites

Clone this project

```
git clone git@github.com:chandrafortuna/simple-messaging-api.git
```

### Running the application

Navigate to `simple-messaging-api` folder and build and run it:

```
cd simple-messaging-api/
go build -o messaging-api
./messaging-api
```

The following table shows the HTTP methods and URLs that represent the action supported in the API.

Request | Description | Use case 
--- | --- | --- | --- 
GET /message | Read all messages | Show a list of messages
POST /message?text={message} | Send a message | Create new message, takes the text params 

Make a POST request to the /message endpoint by navigating http://localhost:8000/message?text=test message
The result will be a message that has been sent.

To get all message, navigate to http://localhost:8000/message. the API will return all messages

### Run the client

An example client to retrieve message after send at realtime

```
cd client/
npm install http-server -g
http-server
```

You should see the client is running as following:
```
Starting up http-server, serving ./
Available on:
  http://127.0.0.1:8080
  http://192.168.8.102:8080
```

Open http://127.0.0.1:8080 in your browser

## Running the tests

```
go test ./test/
```