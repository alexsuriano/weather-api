FROM golang:1.21.5 as build
WORKDIR /app
COPY . . 
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o weather-api .

FROM scratch
COPY --from=build /app/weather-api .
ENTRYPOINT ["./weather-api"]