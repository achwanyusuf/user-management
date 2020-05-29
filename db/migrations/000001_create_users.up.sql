BEGIN;

CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    userId VARCHAR(36) UNIQUE NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    address VARCHAR(1000) NOT NULL,
    password VARCHAR(100) NOT NULL,
    token VARCHAR(1000) NULL DEFAULT NULL,
    latestLogin TIMESTAMP NULL DEFAULT NULL
);

INSERT INTO users (userId, email, address, password) values ('ab1faa5ac0e342eaafde7073d9aaa0b4', 'achwan.yusuf@gmail.com', 'Jalan', '$2a$04$Zf3eMfnvIU5b61s8DN0miOIGVig0n7VxcuYnBWIqMB1I41JDpzlUG');

COMMIT;