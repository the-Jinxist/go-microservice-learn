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

Day 6:
- Working on connecting authentication service to postgres database, used the following drivers `go get -u github.com/jackc/pgconn`, `go get -u github.com/jackc/pgx/v4`, `go get -u github.com/jackc/pgx/v4/stdlib`
- We'll be adding a postgres container to our docker-compose.yml and we'll be accessing that
- We also added the environmental variable in our `authentication-service.dockerfile`

Day 7:
- We connected the broker service to the authenticated service, by using the name of the service in the `docker-compose.yml` file
   Created another handle/route/endpoint that will process an action from the body and call the auth service

Day 8:
- Acces the authentication service from the front end via the broker service `/handle` endpoint

Day 9: 
- Started work on the logger microservice. Added stubs for interacting with the mongo database
- Also added data models for interacting with mongodb

Day 10:
- Finished up data models, added Update() and Drop().
- Added server code and handles for routes

Day 11:
- Added `mongo` to our `docker-compose.yml`, added a service called `mongo`, the setup was very similar to postgres's so it should be
   straightforward to understand again
- We then added the corresponding commands to our Makefile. 
- Added connection from broker service to logger service 

Day 12:
- Finished up logger microservice

Day 13:
- Started work on Mail microservice. We're using Mailhog. It's best practice to simulate a mail server in development instead of sending a real email
- We added mailhog to our `docker-compose.yml`
- Started writing boilerplate code for Mailer service
- We imported some third party packages to make our mailing life easier 
   `github.com/vanng822/go-premailer/premailer` -> helps us to use css with our email without fuss
   `github.com/xhit/go-simple-mail/v2` our mailer package
- We wrote the huge logic to send mail

Day 14:
- Added routes to mail service to accept requests and send mails
- Added mail service to `docker-compose.yml` and added the command to build the binary `build_mail` in the Makefile
- Also added the `mail-service.dockerfile` to collect the binary and put it in the docker container
- `sendMail()` function for broker service to access the mail-service

Day 15:
- Started work with learning rabbitmq. Used the library here `https://github.com/rabbitmq/amqp091-go`
- Established a connection with rabbitmq
- Found this article useful: https://dev.to/koddr/working-with-rabbitmq-in-golang-by-examples-2dcn

Day 16:
- Utlized the rabbitmq logging more

Day 17:
- Added emitter to push events
- Tested things

Day 18:
- Set up an RPC server in the Logger microservice and we also started listening for RPC calls in the same microservice
- Sent event to RPC from broker. Tested everything out

   