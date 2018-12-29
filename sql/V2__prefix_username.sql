-- V2 - adds prefix to CreateUser.V2

DELIMITER $$

DROP PROCEDURE IF EXISTS `Users.CreateUser.V2` $$
CREATE PROCEDURE `Users.CreateUser.V2` (IN username VARCHAR(24), IN givenname VARCHAR(24), IN familyname VARCHAR(24))
BEGIN
    INSERT INTO Users(Username, Givenname, Familyname) VALUES (CONCAT("u_", username), givenname, familyname);
END $$
GRANT EXECUTE ON PROCEDURE `Users.CreateUser.V2` TO UserRole $$

DELIMITER ;
