CREATE TABLE public.users
(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(100) NOT NULL,
    email varchar(100) unique ,
    mobile varchar(100) unique,
    password_hash  varchar(100)
);