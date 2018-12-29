-- Repeated Migration- recreate view

CREATE OR REPLACE VIEW Families AS 
    SELECT Username, FamilyName FROM Users;

