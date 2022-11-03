create table users (
    user_id int not null unique,
    balance int
);

create table reservedFunds (
    user_id int references users(user_id) not null ,
    service_id int not null,
    order_id int not null,
    price int
);

