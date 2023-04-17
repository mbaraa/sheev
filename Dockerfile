FROM alpine:latest as build

RUN apk add cairo cairo-dev go

WORKDIR /app
COPY . .

RUN go get
RUN go build -ldflags="-w -s"

FROM node:16-alpine as build-front

WORKDIR /app
COPY ./client/ ./
RUN npm i
RUN npm run build

FROM alpine:latest as run

RUN apk add cairo

WORKDIR /app

COPY --from=build /app/sheev ./run
COPY --from=build /app/res/fonts/*.ttf /usr/share/fonts/
COPY --from=build /app/res/ ./res/
COPY --from=build-front /app/ ./client/

EXPOSE 4200

CMD ["./run"]
