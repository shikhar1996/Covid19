# Covid19
Get COVID-19 cases information in Indian Geography


## Hosted app on Heroku
```
https://sleepy-wave-66147.herokuapp.com/swagger/index.html
```

## Follow the instructions below to get a local app running

## Install golang

with [homebrew](http://mxcl.github.io/homebrew/):

```Shell
sudo brew install go
```

with [apt](http://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

## Project Setup

Clone the repository

```
git clone https://github.com/shikhar1996/Covid19.git
```

Enter the directory

```
cd Covid19
```

Clean up dependencies
```
go mod tidy
```

Update the database location and credentials in the following files
* src/database/connect.go
* src/database/redis.go
* src/geoencoding/reverse_encoding.go

Generate Swagger docs
```
swag init -g src/server/server.go
```

Start server
```
go run main.go
```

Access the API and Swagger UI through the connection on the connected port.
```
localhost:port/
localhost:port/swagger/index.html
```

## Limitations

* There is monthly limit of 25,000 requests for reverse encoding API.
* Redis Cache 30MB
* Heroku 550 free dyno hours each month. There the service might be idle after 30 minutes of inactivity. In this case the first request might be slow

## Postman Collection

* https://www.getpostman.com/collections/403a8f620e9d0473ee3e

## Issues

* If the server is running on secured domain use HTTPS for sending request.
* If the server is running on localhost HTTP can be used.
