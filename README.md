# devops
Web-HW1

## TEAM MEMBERS

|Name       |
|:----------:|
|*Sara Zahedi*|
|*Sina Elahimanesh*|
|*Darioush Amiri*|

## Project Structure
This project is the first practice on DevOps of Team. We developed this project with these thecnologies:
- Go programming language --> Gin Framework
- Node.js Programming language --> Express Framework
- Frontend of the project is developed by HTML,CSS and BootStrap Components
- Locust Tool for load testing
- Docker Images and Containers and finally Docker Composers
- Nginx Reverse Proxy
- Redis Tool for caching data

## Project Endpoints
Nginx is listening on port 80. 
There are some endpoints at the nginx file configuration (nginx/conf.d/default.conf):
- / -> get index page and redirect user to a page which has 2 buttons to choose which page they want to go to.
- /node/ -> get the node page and redirect user to Node page form
- /go/ -> get the go page and redirect user to Go page form
- /go/sha256/ -> post an input
- /go/sha256 -> get the value of hashed string which has saved in Redis
- /node/sha256/ -> post an input
- /node/sha256 -> get the value of hashed string which has saved in Redis

## Porst and Images
Each section of our project **has been dockerized** with Dockerfiles as shown bellow:
- web-nginx Image -> this image has nginx and its configurations on an alpine based system. Its port is **80**.
- go-backend Image -> this image has go backend files and modules and its port is **8080**.
- node-backend Image -> this image has node backend files and its port is **3000**.
- redis Image -> redis is run on port **6379**.

