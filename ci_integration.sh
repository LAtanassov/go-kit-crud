#!/bin/sh

CUR_DIR=$(pwd)

DB_DIB=$(docker run -d -p 3306:3306 -e MYSQL_USER=UserService -e MYSQL_PASSWORD=password -e MYSQL_ROOT_PASSWORD=toor -e MYSQL_DATABASE=UserDB mysql:8)
SERVICE_DID=$(docker run -d -p 8080:8080 latanassov/usersvc:0.1.0)

while ! docker exec $DB_DIB mysqladmin ping --silent ; do
    echo "Waiting for database..."
    sleep 2
done
sleep 5

flyway -user=root -password=toor -url=jdbc:mysql://localhost:3306/UserDB -locations=filesystem://$CUR_DIR/sql migrate

go test -race -tags=integration ./...

docker stop $SERVICE_DID >/dev/null
docker stop $DB_DIB >/dev/null