CREATE DATABASE IF NOT EXISTS company_hierarchy;

USE company_hierarchy;

DROP PROCEDURE IF EXISTS AddDepartment;
DROP PROCEDURE IF EXISTS UpdateDepartment;
DROP PROCEDURE IF EXISTS ActivateDepartment;
DROP PROCEDURE IF EXISTS DeactivateDepartment;
DROP PROCEDURE IF EXISTS DeleteDepartment;
DROP PROCEDURE IF EXISTS RestoreDepartment;
DROP PROCEDURE IF EXISTS ApproveDepartment;
DROP PROCEDURE IF EXISTS UnapproveDepartment;
DROP PROCEDURE IF EXISTS GetHierarchy;

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
    SET
        name = COALESCE(dept_name, name),
        parent_id = COALESCE(dept_parent_id, parent_id),
        flags = COALESCE(dept_flags, flags)
    WHERE id = dept_id;
END //

CREATE PROCEDURE ActivateDepartment(
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags | 1 -- Mark as active
    WHERE id = dept_id;
END //

CREATE PROCEDURE DeactivateDepartment(
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags & ~1 -- Mark as inactive
    WHERE id = dept_id;
END //

CREATE PROCEDURE DeleteDepartment(
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags | 2 -- Mark as deleted
    WHERE id = dept_id;
END //

CREATE PROCEDURE RestoreDepartment(
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags & ~2 -- Mark as restored (not deleted)
    WHERE id = dept_id;
END //

CREATE PROCEDURE ApproveDepartment(
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags | 4 -- Mark as approved
    WHERE id = dept_id;
END //

CREATE PROCEDURE UnapproveDepartment(
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags & ~4 -- Mark as not approved
    WHERE id = dept_id;
END //

CREATE PROCEDURE GetHierarchy(
    IN dept_parent_id INT
)
BEGIN
    WITH RECURSIVE department_hierarchy AS (
        SELECT id, name, parent_id, flags
        FROM departments
        WHERE id = dept_parent_id AND (flags & 1) = 0
        UNION ALL
        SELECT d.id, d.name, d.parent_id, d.flags
        FROM departments d
        INNER JOIN department_hierarchy dh ON dh.id = d.parent_id
        WHERE (d.flags & 1) = 0
    )
    SELECT * FROM department_hierarchy;
END //

DELIMITER ;
