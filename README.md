# Goact
Goact is a Go + React based chatting app. This application is still being updated and modified with features, and is intended to be used as a starting point for chat based applications. At its current point, it establishes a connection between the server and client, and elevates it to a web socket. Then multiple users can interact with a chat room. Authentication is required, however to make it easy for others to utilize the code, users are hard coded into the backend. It is recommend that if this application is used for any project, a database, or equally efficient method of data storing is utilized.

<p align="center">
  <img src="https://github.com/Syssos/goact/blob/main/GoactExample.png" alt="goact example img"/>
</p>

## Dependancies
- [Go](https://golang.org/)
	* UUID Package ([google/uuid](https://github.com/google/uuid))v1.2.0
	* JWT Go ([dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go))v3.2.0
	* Websocket Package ([gorilla/websocket](https://github.com/gorilla/websocket))v1.4.2
	* Mux Router ([gorilla/mux](https://github.com/gorilla/mux))v1.8.0
	* CORS Router ([rs/cors](https://github.com/rs/cors))v1.8.0

- [Node.js & npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
	* [axios](https://www.npmjs.com/package/axios)
	* [node-sass](https://www.npmjs.com/package/node-sass)

## Project Plans
From this point forward the app will be modified for more practical means. I would like to use this space to explain some of the features that will be added over the course of the next couple weeks.

1. Chat control system (Means of controlling which poeple are conversating with eachother, Aka chat rooms)
2. Additional User information (A method of storing and modifing additional information like profile pictures or statuses)

## Running Locally
To get this project running local start by cloning this repository to a location on your local machine.

This app will rely on two services running at the same time. To start these services we will need two terminals open, on needs to be in the frontend/ directory. The other in the GoBackend/ directory.

### Terminal 1 (Location: '../goact/GoBackend/')

```bash
go run main.go
```

> **Note:** This should run the 'go get' command for any needed package's that are not installed.
> Alternatively you can start the docker container.
> ```bash
> docker build -t backend .
> ```
> ```bash
> docker run -it -p 8080:8080 backend
> ```

### Terminal 2 (Location: '../goact/frontend/')
The first thing we need to do to get our front end together is installing all of the packages needed for this project. The package.json file is included in this directory. Meaning the command below will get all the packages needed for the front end.

```bash
npm install
```

Once that the packages are installed, and the backend is running in the another terminal, we can start our frontend

```bash
npm start
```

## Known Issues

While this app is buildable and runs, it is not in anymeans ready for a production enviornment. Primarily because the security is very lacking, but also because there is no form of control. These features are currently being worked on, once added the ability to run in a production environment will be more eligable
