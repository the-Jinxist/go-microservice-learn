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

Day 19:
- We're working on speeding up our work using gRPC. To do this, we used the tools gotten by running these commands: `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27`, `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
- To use gRPC, we first write the protocol file, we compile it, then we get some generated files. After this, we write the client code, we write the server code, then we test things
- Information for installing for protoc can be found here: https://grpc.io/docs/protoc-installation/
- The command we used to generate the files that we would work with corresponding to this, looked something like this: `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative log.proto`
- We wrote up the Logserver struct. It seems it has to always inherit(via composition), the `Unimplemented[Service name specified in the protoc]ServiceServer` interface, and also implement the function that comes with the interface. The signature of this interface can be found in the [proto-file-name]_grpc.pb.go file in an interface that looks like `[Service name specified in the protoc]ServiceServer`

Day 20:
- Listened for connections in the logger-microservice
- Wrote up the client code for communicating with the gPRC server.
- Updated front-end code and tested

Day 21:
- Docker Swarm Overview: `https://docs.docker.com/engine/swarm/`
- Building images for each of our services using a command that looks like: `docker build -f your-service.dockerfile -t your-docker-hub-username/service-name:1.0.0 <path; in our case we use ".">`
- We logged in into our docker hub using `docker login`
- We then pushed our image to docker using `docker push <username>/your-service-name:version`

Day 22:
- Wrote our `swarm.yml`
- We initialized our docker swarm using `docker swarm init` in the project folder. Found some instructions to add a worker node to our swarm in case if there's too much traffic or something
- To regenerate this instruction and token, we can use the command `docker swarm join-token worker` and to add a manager node, `docker swarm join-token manager`
- To deploy docker swarm, we need to execute this command `docker stack deploy -c swarm.yml myapp`

Day 23:
- Faced a big wahala. Mailhog service wasn't starting. Apparently, it was resolving to a different architecture so it couldn't start the service, such nonsense. I had to add a flag to the command to deploy the swarm, it looked like this: `docker stack deploy -c swarm.yml microservice --resolve-image never`
- Learning about scaling services up. i.e increasing the replicas of the services that your deployed. The command looks something like.. `docker service scale microservice_listener-service=3`. To scale down, we use the same command, just reducing the number of replicas: `docker service scale microservice_listener-service=2`
- To update a service, first we're updating our image using the pervious command but with a new version. It'll look like this: `docker build -f listener-service.dockerfile -t neofemo/listener-service:1.0.1 .`. then you push to docker with the new version `docker push neofemo/listener-service:1.0.1`.

After, to update the image that one of our service is using we use this command `docker service update --image <image-name:image-version> <name of service>`

- To stop a service, you just scale it to 0.
- To remove the entire swarm, call `docker stack rm microservice`
- To leave the swarm completely, `docker swarm leave [--force](if it's a manager)`

- We added the frontend service to docker and our swarm. Replicating the steps we did for other services before

Day 24:
- Using Caddy as a proxy to hit our web server: Here: `https://caddyserver.com/` and the image on docker hub: `https://hub.docker.com/_/caddy`
- From tutorial resource, we downloaded two files, `caddy.dockerfile` and `Caddyfile`
- We built an image from  `caddy.dockerfile` and pushed it to docker hub
- We added caddy to our swarm.yml and specified the volumes. Volumes help to store data collected in a docker container, somewhere else because docker containers are ephemeral

Day 25:
- Created folders `caddy_data` and `caddy_config`. They correspond to the folders provided for the volumes for the caddy image in our `swarm.yml`
- We edited /etc/hosts using vim. Used the command: `sudo vi /etc/hosts` and added a backend entry where `localhost` was defined

Day 26:
- Updated front-end with dynamic url gotten from our environemt. Added the code to `main.go` and passed it to `test.page.gohtml`.

# After a very long hiatus, I'm baaaack 
Day 27:
- I forgot to update the binary for broker-service and frontend-service after updating the ports they listen to.
- Updated postgres version on swarm.yml because 14.0 has some malware
- We started working to push our swarm to the internet. I created a vm instance on google cloud, You'll pass Google Compute Engine, create a vm instance. Using the information provided. `Ubuntu 20.4 LTS`, Allow https and http to connect. I ssh-ed into the server (Did this my creating a ssh key pair, uploaded the public key to the ssk key section of the vm instance and ran the command `ssh -i [path-to-private-key](e.g ~/.ssh/goserver) [username](e.g nerosilva522)@[external server ip](e.g 35.188.42.215)`) and ran some commands to setup access
- `sudo adduser neo` (created new user), `sudo usermod -aG sudo neo` (made user have sudo priviledges), `sudo ufw allow ssh`, `sudo ufw allow http`, `sudo ufw allow https` to allow ssh, https and http access respectively
- Enabled access on specific ports for docker swarm. `sudo ufw allow 27017/tcp`, `sudo ufw allow 7946/tcp`, `sudo ufw allow 7946/udp`, `sudo ufw allow 4789/udp`, `sudo ufw allow 8025/tcp`
- We then called `ufw enable` to save our settings. Called `ufw status` to view all our settings
- Installed docker on the vm: https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository
- We updated the hostname of our server `sudo hostnamectl set-hostname node-1`
- We then updates the hosts file with the command: `sudo vi /etc/hosts` Used down arrow to the bottom of the file. Type `o`, copied the external IP address of the server. pasted it, hit tab, wrote `node-1`

Day 28:
- Got a domain, trying to add the IP address of the external VM instance to a subdomain DNS records

Day 29:
- Finally got the subdomains pointed to the external IP address of my google VM
- Added a CNAME subdomain for `broker`, so we can access `broker.tukio.com.ng`. 
- Created a new caddy file to point to the domains we just setup up. it's `Caddyfile.production`, which is used by the `caddy.production.dockerfile` when creating the container. We then created a new swarm configuration file; `swarm.production.yml` file
- The new prod swarm configuration file points to the production caddy container that we deployed to docker hub. We deployed it using the `caddy.production.dockerfile`
- Docker swarm contains didn't work fine because the folders for 
  - ```
    ./db-data/postgres/:/var/lib/postgresql/data/
    ./db-data/mongo/:/data/db
    
    ```
    for mongo and postgres contains in our swarm config file, were not created yet
  
- Found an issue with some of my binaries, they weren't built for amd64, the architecture of the VM instance, added the AARCH command to the golang build commands, rebuilt the binaries, built new tags using the `--no-cache` flag. E,g `docker build --no-cache -f front-end.dockerfile -t neofemo/front-end:1.2.0 .`
  - Pushed the new tags. Pending mail, logger and listener services.
  - Once that is done, will delete the swarm.yml file, use `vi swarm.yml` to create a new one, copy the new production swarm file and try to deploy the swarm again.
  - You can use `docker stack rm myapp` to delete the swarm