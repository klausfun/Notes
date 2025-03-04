CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    password_hash varchar(255) not null,
    email         varchar(255) not null unique
);

CREATE TABLE notes
(
    id          serial                                      not null unique,
    user_id     int references users (id) on delete cascade not null,
    title       varchar(255)                                not null,
    description varchar(8191)                               not null
);
