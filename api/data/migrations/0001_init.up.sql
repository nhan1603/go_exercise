CREATE TABLE IF NOT EXISTS public.user
(
    id          SERIAL PRIMARY KEY,
    email       TEXT                     NOT NULL CHECK (email <> ''::text)
    );

CREATE TABLE IF NOT EXISTS public.relationship
(
    id              SERIAL PRIMARY KEY,
    first_email_id  INT NOT NULL CHECK (first_email_id <> second_email_id),
    second_email_id INT NOT NULL,
    status          TEXT NOT NULL DEFAULT ''
    );