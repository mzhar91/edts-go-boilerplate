# golang build
FROM golang:1.18.8-alpine3.16 AS build

LABEL maintainer="Harry Kurniawan <harry.kurniawan@sg-edts.com>"

ENV WORKDIR=/sg-edts.com/klik-scheduler

# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git

WORKDIR ${WORKDIR}

COPY . ${WORKDIR}

RUN cd ${WORKDIR} \
    && go mod tidy

RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/main ./main.go

#=======================
#=======================

FROM nginx:alpine

LABEL maintainer="Harry Kurniawan <harry.kurniawan@sg-edts.com>"

ENV BUILD_DIR=/sg-edts.com/klik-scheduler
ENV WORKDIR=/home/app

WORKDIR ${WORKDIR}

# config

COPY --from=build ${BUILD_DIR}/bin ${WORKDIR}

RUN chmod +x /home/app/main

CMD ["/home/app/main"]

EXPOSE 8080