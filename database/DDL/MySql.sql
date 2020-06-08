CREATE TABLE `userdetails` (
    `username` VARCHAR(64) NOT NULL UNIQUE,
    `password` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`username`)
);
