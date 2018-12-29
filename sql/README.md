# SQL Scripts

## flyway

- local env. setup
- can be automated
- needs to be hardend
- undo functionality requires enterprise version

```sh
$>docker run -d -p 3306:3306 -e MYSQL_USER=<user> -e MYSQL_PASSWORD=<user_password> -e MYSQL_ROOT_PASSWORD=<root_password> -e MYSQL_DATABASE=UserDB mysql:8
$>flyway -user=<root> -password=<root_password> -url=jdbc:mysql://localhost/UserDB -locations=filesystem:<path>/sql migrate
```
