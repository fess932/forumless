CREATE DATABASE forum;

CREATE TABLE forum (
    id SERIAL UNIQUE
);

CREATE TABLE "user" (
    id SERIAL UNIQUE,
    name text
);

CREATE TABLE post (
    id SERIAL UNIQUE,
    data text,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES "user"(id)
);

CREATE TABLE comment (
    id SERIAL UNIQUE,
    data text,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES "user"(id)
);