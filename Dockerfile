##########
## This is an example if you would use node.js to solve the challenge.
#FROM node:10-alpine
#WORKDIR /usr/src/app
#COPY package*.json ./
#RUN npm install
#COPY . .

###########
# Keep this port exposed to simplify the acceptance test.
# Make sure your implementation works on this port.
# EXPOSE 3000

#CMD ["node", "src/app.js"]

# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./
COPY pkg/ ./

RUN go build -o /docker-gs-ping

EXPOSE 3000

CMD [ "/docker-gs-ping" ]
