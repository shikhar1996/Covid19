# Covid19
Get COVID-19 cases information in Indian Geography


## Install go(lang)

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

## API Limit

* There is monthly limit of 25,000 requests for reverse encoding API.


## Issues
* If the server is running on secured domain use HTTPS for sending request.
* If the server is running on localhost HTTP can be used.
