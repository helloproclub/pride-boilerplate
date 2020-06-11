#BUILD STAGE
FROM golang:alpine as build
RUN apk add git
ADD . /src
RUN cd /src && go build -o main

#RUN
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /src/main /app/
ENTRYPOINT ./main
