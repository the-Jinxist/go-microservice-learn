# Microservices in Golang
Following the Trevor Sawler Udemy course

Day 1:
- First real thing I did was download the front-end file resources, moved it into the project.

Day 2:
 - Launched the static front end file on the localhost:80
 - Ran `go mod init` in `broker-service` module. I guess this makes it it's own application
 - We're using `go-chi` for our routing needs, not fiber, not gin. Found here: https://github.com/go-chi/chi
    ```
        go get -u github.com/go-chi/chi/v5
        go get -u github.com/go-chi/cor
        go get -u github.com/go-chi/chi/v5/middleware
   ```
- We created a route on go-chi, uses basic `http.HandleFunc` interface.

Day 3:
- We're working on creating a docker image of our broker service
   - We tried the first way you could do this; using the docker compose file, We made our dockerfile for the broker-service first, file called `broker-service.dockerfile`. Then we created the actual compose file in the project folder called `docker-compose.yml`. Added comments to each command to explain them more. We ran the compose up code using the following command: `docker compose up -d`

Day 4:
- Tested our our broker service by accessing the service from our `test.page.gohtml`. Had to write HTML which wasn't something I liked smh
- Created helper functions to write and read json, and also to send an error message
- We created a MakeFile to run some commands easily, we also edited our dockerfile to take advantage of the binary built with the first command of the 
   up_build Makefile command

Day 5:
- Wrote up stubs for authentication microservice

   