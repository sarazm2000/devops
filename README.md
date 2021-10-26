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

## Docker Compose
The whole project and docker images has been dockerized with docker compose.
The yml file is in github and project has been dockerzied with this file.

## Scripts
Because we did not have public ip, so to test this project using Docker Compose we just made a creativness method.
We wrote a bash file to set a domain (**web.hw1**) in the directory **/etc/** and file **hosts**. This file add the whole docker compose IP to OS hosts as a new domain. and finally with this method our project is ran on domain **web.hw1** in the system and with the power of docker composer project has been OS Independant (Except Windows :)). 
So, There is a **setip.sh** bash file that is for setting ip and domain.
And There is no need to use this file. You just need to run **run.sh** bash file to docker compse and also set ip and project after running this bash file will be fully up.

## Run Project
To run the project just we should do these steps:
1. Docker Pull Images (web-nginx, go-back, node-back, redis) `docker pull $image_name`
2. just run **run.sh** bash file by running this line of code in terminal:
`bash run.sh`
