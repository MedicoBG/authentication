CREATE TYPE moderator_type AS ENUM (
    'doctor',
    'pharmacy',
    'citizen',
    'medicament'
    );


CREATE TABLE moderator_auth
(
    id       uuid           NOT NULL DEFAULT gen_random_uuid(),
    email    text           NOT NULL,
    password bytea          NOT NULL,
    type     moderator_type NOT NULL
);