FROM alpine:latest

#creates the app directory
RUN mkdir /app

#copies the executable/binary created from the first command in the up_build Makefile command to the /app directory just created in the docker image
COPY loggerApp /app

#provies an entry point for the executable file
CMD [ "/app/loggerApp" ]