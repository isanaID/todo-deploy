-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
  id SERIAL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  created_at varchar(255),
  updated_at varchar(255),
  PRIMARY KEY (id)
);


CREATE TABLE task (
    id SERIAL,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    deadline varchar(255) NOT NULL,
    user_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    status_id BIGINT NOT NULL,
    created_at varchar(255),
    updated_at varchar(255),
    PRIMARY KEY (id)
);


CREATE TABLE status_task (
    id SERIAL,
    status varchar(255) NOT NULL,
    created_at varchar(255),
    updated_at varchar(255),
    PRIMARY KEY (id)
);


CREATE TABLE category (
    id SERIAL,
    name varchar(255) NOT NULL,
    created_at varchar(255),
    updated_at varchar(255),
    PRIMARY KEY (id)
);


alter table task add foreign key (user_id) references users(id);
alter table task add foreign key (status_id) references status_task(id);
alter table task add foreign key (category_id) references category(id);

-- +migrate StatementEnd