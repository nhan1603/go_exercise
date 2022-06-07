# Friend Management System

A simple application that mimic the basic operation of a social network. The application provides a range of APIs that allows users to manage the relationships between users.

The provided operations include simple actions such as create, delete or view on certain relationship between users.

## 1. Summary

- Programming Language: Golang
- Database: Postgresql
- Deployment: Docker, Linux
- Tools: Goland, Git

## 2. List of APIs
```
POST /api/create-user

POST /api/add-friend

POST /api/subscribe

POST /api/block

POST /api/common-friend

POST /api/update-receiver
```

## 3. Deployment

- This project can be deployed by Docker to Linux server at: http://localhost:3000/

### #Deployment process

```
make setup
make boilerplate
make run
```

### #Testing

```
make test
```
