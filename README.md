<p align="center">
    <a href="https://zleague.gg">
        <img alt="zleague" src="https://avatars1.githubusercontent.com/u/70303271?s=200&v=4"> 
    </a>
</p>
<br>
<p align="center">This repository contains all of the code that is related to the API of ZLeague</p>
<br>


## Table of Contents
- [Prerequisites](#req)
  - [Golang](#golang)
  - [Docker](#docker)
- [Installation](#install)
- [Local Development](#localdev)
  - [Enviornment Variables](#env)
  - [Docker Compose](#compose)
- [Resources](#resources)
  - [Backend Wiki](#wiki)
  - [API Documentation](#apidocs)  
  - [Progress Tracker](#progresstracker)

<a name="install"></a>

# Installation
> Instructions to prepare you for development

1. Clone the project to your machine

```bash
# clone the repository
https://github.com/zleague/api.git
# cd into the project
cd api
# install all dependencies
go mod download
```

# Devlopment
> We assume you have the Docker Daemon running.

### Enviornment Variables
Create a .env file in project root directory

```bash
  ├── .env
  ├── .env.example
  ├── docker-compose.yml
  ├── go.mod
  ├── go.sum
  ├── main.go
  └── ...
```
### Example
> Example Contents of the .env file
```bash
DB_URI=mongodb://localhost:27017
DB_USERNAME=root
DB_PASSWORD=AVeryStrongPassword1234
SERVER_USERNAME=backend
SERVER_PASSWORD=hackerman

```
Your .env file should have the same keys as the .env.example. If you create new enviornment variables update the .env.example.

### Docker Compose
```bash
# build the docker container and get it up and running
docker-compose up --build
```

### Starting server with Go

```bash
# go run
go run main.go

```