DROP INDEX IF EXISTS order_item_uidx_order_id_product_id;
DROP TABLE IF EXISTS public.order_item;

DROP INDEX IF EXISTS order_uidx_external_id;
DROP TABLE IF EXISTS public.order;

DROP INDEX IF EXISTS product_uidx_external_id;
DROP TABLE IF EXISTS public.product;

DROP TABLE IF EXISTS public.user;
DROP TABLE IF EXISTS public.relationship;
