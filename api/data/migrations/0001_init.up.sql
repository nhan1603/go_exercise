CREATE TABLE IF NOT EXISTS public.product
(
    id          SERIAL PRIMARY KEY,
    external_id TEXT                     NOT NULL CHECK (external_id <> ''::text),
    name        TEXT                     NOT NULL CHECK (name <> ''::text),
    description TEXT                     NOT NULL CHECK (description <> ''::text),
    status      TEXT                     NOT NULL CHECK (status <> ''::text),
    price       BIGINT                   NOT NULL CHECK (price > 0),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );
CREATE UNIQUE INDEX IF NOT EXISTS product_uidx_external_id ON public.product (external_id);

CREATE TABLE IF NOT EXISTS public.order
(
    id          BIGSERIAL PRIMARY KEY,
    external_id TEXT                     NOT NULL CHECK (external_id <> ''::text),
    user_id     TEXT                     NOT NULL CHECK (user_id <> ''::text),
    status      TEXT                     NOT NULL CHECK (status <> ''::text),
    total_cost  BIGINT                   NOT NULL CHECK (total_cost >= 0),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );
CREATE UNIQUE INDEX IF NOT EXISTS order_uidx_external_id ON public.order (external_id);

CREATE TABLE IF NOT EXISTS public.order_item
(
    id         BIGSERIAL PRIMARY KEY,
    order_id   BIGINT                   NOT NULL REFERENCES public.order (id),
    product_id INT                      NOT NULL REFERENCES public.product (id),
    quantity   INT                      NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );
CREATE UNIQUE INDEX IF NOT EXISTS order_item_uidx_order_id_product_id ON public.order_item (order_id, product_id);

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
    status          INT NOT NULL DEFAULT 0
    );