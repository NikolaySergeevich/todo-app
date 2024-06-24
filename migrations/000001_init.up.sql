CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists -- таблица списка
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255) --описание
);

CREATE TABLE users_lists -- таблица для соотношения пользователя со списками задач
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items -- таблица для описания задачи
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255), --описание
    done boolean not null default false
);

CREATE TABLE lists_items -- таблица соотношения задачи со списком
(
    id serial not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);