CREATE DATABASE forum;

CREATE TABLE user (
    id serial,
    name text
);

CREATE TABLE post (
    id serial,
    data text,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE comment (
    id serial,
    data text,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE forum (
    id serial
);