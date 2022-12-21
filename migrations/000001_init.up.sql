CREATE TABLE users
(
    id          serial          not null unique,
    name        varchar(127)    not null,
    login       varchar(127)    not null unique,
    password    varchar(127)    not null
);

CREATE TABLE products
(
    id      serial          not null unique,
    name    varchar(127)    not null unique,
    price   real            not null
);

CREATE TABLE orders
(
    id      serial                                      not null unique,
    user_id int references users(id) on delete cascade  not null
);

CREATE TABLE store
(
    id          serial                                          not null unique,
    order_id    int references orders(id) on delete cascade     not null,
    product_id  int references products(id) on delete cascade   not null,
    count       int                                             not null
);