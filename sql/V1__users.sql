-- V1 - users database initialization script, term `UserXxx` is the domain

-- assume mysql service user with name `UserService` exists
CREATE ROLE IF NOT EXISTS UserRole;
GRANT UserRole TO UserService;

-- tables
CREATE TABLE IF NOT EXISTS Users (
    ID MEDIUMINT NOT NULL AUTO_INCREMENT,
    Username VARCHAR(24) UNIQUE NOT NULL,
    Givenname VARCHAR(24),
    Familyname VARCHAR(24),
    PRIMARY KEY (ID)
);

