CREATE DATABASE IF NOT EXISTS student;


CREATE TABLE IF NOT EXISTS student(
    user_name varCHAR UNIQUE,
    id INT PRIMARY KEY,
    point INT
);