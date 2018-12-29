-- local env. setup; can be automated; needs to be hardend
-- $>docker run -d -p 3306:3306 -e MYSQL_USER=<user> -e MYSQL_PASSWORD=<user_password> -e MYSQL_ROOT_PASSWORD=<root_password> -e MYSQL_DATABASE=UserDB mysql:8
-- $>flyway -user=<root> -password=<root_password> -url=jdbc:mysql://localhost/UserDB -locations=filesystem:<path>/sql migrate

-- V2 - shows an update
-- USE UserDB

-- stored procedures, term `UserXxx` is the domain
-- create
DELIMITER $$

DROP PROCEDURE IF EXISTS `Users.CreateUser.V2` $$
CREATE PROCEDURE `Users.CreateUser.V2` (IN username VARCHAR(24), IN givenname VARCHAR(24), IN familyname VARCHAR(24))
BEGIN
    INSERT INTO Users(Username, Givenname, Familyname) VALUES (CONCAT("u_", username, "_n"), givenname, familyname);
END $$
GRANT EXECUTE ON PROCEDURE `Users.CreateUser.V2` TO UserRole $$

DELIMITER ;
