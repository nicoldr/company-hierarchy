CREATE DATABASE IF NOT EXISTS company_hierarchy;

USE company_hierarchy;

CREATE TABLE IF NOT EXISTS departments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    parent_id INT NULL,
    flags INT DEFAULT 0
);

DELIMITER //

CREATE PROCEDURE AddDepartment(
    IN dept_name VARCHAR(50),
    IN dept_parent_id INT,
    IN dept_flags INT
)
BEGIN
    INSERT INTO departments (name, parent_id, flags) 
    VALUES (dept_name, dept_parent_id, dept_flags);
END //

CREATE PROCEDURE UpdateDepartment(
    IN dept_id INT,
    IN dept_name VARCHAR(50),
    IN dept_parent_id INT,
    IN dept_flags INT
)
BEGIN
    UPDATE departments
    SET name = dept_name, parent_id = dept_parent_id, flags = dept_flags
    WHERE id = dept_id;
END //

CREATE PROCEDURE DeleteDepartment(
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags | 1 -- Set bit 1 to mark as deleted
    WHERE id = dept_id;
END //

CREATE PROCEDURE GetHierarchy(
    IN dept_parent_id INT
)
BEGIN
    SELECT * FROM departments
    WHERE parent_id = dept_parent_id AND (flags & 1) = 0; -- Exclude deleted departments
END //

DELIMITER ;
