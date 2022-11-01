# Goact ![travis-badge](https://app.travis-ci.com/Syssos/goact.svg?branch=main)
Goact is a concurrency favoring Go/ReactJS project exoskeleton. This essentually means it is intended for being a way to generate a fast, efficient, and tested chat component, and be able to fully intergrate it into a current front end within minutes.

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

## Project Details
### Current Features:

Currently this app will handle multiple users interacting within one chat room. The accounts are not managed by a database so any database desired to be used can be implemented fairly easily. Residing in the [.env](https://github.com/Syssos/goact/blob/main/backend/.env) file are 2 default accounts used for testing out the application and ensureing functionality is working.

The React front end will prompt the user to sign in before continuing to the component outlining the chat feature. The authentication is handled by signed JWT tokens, which are saved in the browser as a cookie. If a user does not have an authentically signed cookie when attempting to access to the websocket connection, they will get redirected to the signin page.


## Running Locally
To get this project running local start by cloning this repository to a location on your local machine.

The Go Back end will need to be running at all times in order to have chat functionality. It is responsable for validating users, creating TCP/IP websocket connections, and tracking each chat room in a Go routine while any user is connected.

The Front end will optain a cookie string for the current user, this string will be valid for a predetermand amount of time and the browser uses it as a sort of key to tell the server the user is the person they claim to be. 

In order for the project to run both the front and back end will need to be ran, these services will run on different ports and will both need to be started in their own termian session, or have one run in the background.


### Terminal 1 (../goact/backend/)

```bash
$ go test ./...

ok      github.com/Syssos/goact
ok      github.com/Syssos/goact/models/chatroom
ok      github.com/Syssos/goact/routes

$ go build .
$ ./goact
Starting server on localhost:8080

```

> **Note:** This should run the 'go get' command for any needed package's that are not installed.
> Alternatively you can start the docker container.
> ```bash
> docker build -t backend .
> ```
> ```bash
> docker run -it -p 8080:8080 backend
> ```

### Terminal 2 (../goact/frontend/)
The first thing we need to do to get our front end together is install any packages needed by this project. The package.json file will in this directory includeds all of the required dependancies meaning the npm install command can handle grabbing them for us.

```bash
npm install
```

With dependancies gathered, and the backend running, we can start our frontend and begin interacting with the app.

```bash
npm start
```
