CREATE TABLE IF NOT EXISTS public.admins (
  id bigserial NOT NULL,
  title varchar(255) NOT NULL,
  published_at timestamp with time zone NOT NULL,
  pages smallint NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  CONSTRAINT pkey_books PRIMARY KEY (id)
)
WITH (
    OIDS=FALSE
);

TRUNCATE TABLE public.books RESTART IDENTITY CASCADE;