CREATE TABLE student (
    id          varchar(126) PRIMARY KEY NOT NULL,
    name        varchar(126) NOT NULL,
    bio         varchar(510),
    since       timestamp NOT NULL,
    email       varchar(126) NOT NULL UNIQUE,
    password    TEXT NOT NULL
);