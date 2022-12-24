CREATE TABLE IF NOT EXISTS public.users
(
    id          uuid            not null unique,
    name        varchar(127),
    email       varchar(127)    not null unique,
    phone       varchar(127)
);
