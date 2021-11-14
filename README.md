# Goact ![travis-badge](https://app.travis-ci.com/Syssos/goact.svg?branch=main)
Goact is an early stage Go + ReactJS Chat application exoskeleton. What that means is this project is intended for being a way to generate a fast, efficient, and tested chat application, and deploy within minutes after minor styling changes, via its own service or as an integration into another react app as a component.

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

Currently this app will handle multiple users interacting within one chat room. The accounts are not managed by a database so any database required can be implemented. Residing in the [.env](https://github.com/Syssos/goact/blob/main/backend/.env) file are 2 default accounts to be used during testing.

The React front end will prompt the user to sign in before continuing to the application. The authentication is handled by signed JWT tokens, which are saved in the browser. If a user does not have an authentic cookie when interacting with the websocket, or while attempting to make a connection, they will get redirected to the signin page.

### Future Feature Plans

Unit and benchmark tests are currently being worked on.

After testing is added, multiple scripts are going to be written to automate the installation and configuration of a database. Data storage will be vital in implementing multiple chat rooms and will need to be configured before doing so. Databases are a vastly more secure option when dealing with a production environment, this step will be introduced to save time down the line.


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
## More Information

For more information check out the Official Docs page for the project [here](https://docs.codyparal.com/?project=3&category=intro_ga)
