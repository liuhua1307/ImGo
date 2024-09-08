create table users(
    id bigint unsigned primary key  auto_increment not null ,
    phone_number varchar(20) not null,
    password varchar(100)

);

create unique index idx_id_phn ON users (phone_number,id);

create table user_info(
    id bigint unsigned primary key auto_increment not null,
    name text,
    sex tinyint unsigned,
    comment text
);

create table friendship(
    id  bigint unsigned primary key auto_increment not null,
    user_id bigint not null ,
    join_id bigint not null ,
    status  tinyint not null ,
    created_at timestamp,
    deleted_at timestamp
)