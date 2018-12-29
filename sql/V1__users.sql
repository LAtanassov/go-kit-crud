-- V1 - users database initialization script, term `UserXxx` is the domain

-- assume mysql service user with name `UserService` exists
CREATE ROLE IF NOT EXISTS UserRole;
GRANT UserRole TO UserService;

-- tables
CREATE TABLE IF NOT EXISTS Users (
    ID MEDIUMINT NOT NULL AUTO_INCREMENT,
    Username VARCHAR(24) NOT NULL,
    Givenname VARCHAR(24),
    Familyname VARCHAR(24),
    PRIMARY KEY (ID)
);

CREATE INDEX IDX_Users_Username ON Users(Username);


-- stored procedures, term `UserXxx` is the domain
-- create
DELIMITER $$

DROP PROCEDURE IF EXISTS `Users.CreateUser.V1` $$
CREATE PROCEDURE `Users.CreateUser.V1` (IN username VARCHAR(24), IN givenname VARCHAR(24), IN familyname VARCHAR(24))
BEGIN
    INSERT INTO Users(Username, Givenname, Familyname) VALUES (username, givenname, familyname);
END $$
GRANT EXECUTE ON PROCEDURE `Users.CreateUser.V1` TO UserRole $$

-- read
DROP PROCEDURE IF EXISTS `Users.ReadUser.V1` $$
CREATE PROCEDURE `Users.ReadUser.V1` (IN username VARCHAR(24))
BEGIN
    SELECT u.Username, u.Givenname, u.Familyname 
    FROM Users as u
    WHERE u.Username = username;
END $$
GRANT EXECUTE ON PROCEDURE `Users.ReadUser.V1` TO UserRole $$

-- update
DROP PROCEDURE IF EXISTS `Users.UpdateUser.V1` $$
CREATE PROCEDURE `Users.UpdateUser.V1` (IN username VARCHAR(24), IN givenname VARCHAR(24), IN familyname VARCHAR(24))
BEGIN
    UPDATE Users as u
    SET u.Givenname = givenname, u.Familyname = familyname
    WHERE u.Username = username;
END $$
GRANT EXECUTE ON PROCEDURE `Users.UpdateUser.V1` TO UserRole $$

-- delete
DROP PROCEDURE IF EXISTS `Users.DeleteUser.V1` $$
CREATE PROCEDURE `Users.DeleteUser.V1` (IN username VARCHAR(24))
BEGIN
    DELETE
    FROM Users
    WHERE Username = username;
END $$
GRANT EXECUTE ON PROCEDURE `Users.DeleteUser.V1` TO UserRole $$

DELIMITER ;
