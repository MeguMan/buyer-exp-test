FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

#RUN psql -c "CREATE DATABASE avito_test_db;"
#RUN psql -c "CREATE DATABASE user avito_test WITH PASSWORD 'secret';"
#RUN psql -c "GRANT ALL PRIVILEGES ON DATABASE abito_test_db TO avito_test;"
RUN make
RUN make migrations

CMD ["apiserver"]