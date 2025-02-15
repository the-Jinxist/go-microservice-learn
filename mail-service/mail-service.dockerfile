FROM --platform=linux/amd64 alpine:latest

RUN mkdir /app

COPY mailApp /app

#this copies our email templates over to the docker container/image
COPY templates /templates

CMD [ "/app/mailApp" ]