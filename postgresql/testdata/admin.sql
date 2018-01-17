CREATE TABLE IF NOT EXISTS public.admins (
  id bigserial NOT NULL,
  name varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  CONSTRAINT pkey_admins PRIMARY KEY (id),
  CONSTRAINT unique_admins_on_name UNIQUE (name)
)
WITH (
    OIDS=FALSE
);

INSERT INTO public.admins(name, password, email) VALUES ('Alata', 'GoseiRed', 'red@goseiger.com');
INSERT INTO public.admins(name, password, email) VALUES ('Eri', 'GoseiPink', 'pink@goseiger.com');
INSERT INTO public.admins(name, password, email) VALUES ('Agri', 'GoseiBlack', 'black@goseiger.com');
INSERT INTO public.admins(name, password, email) VALUES ('Moune', 'GoseiYellow', 'yellow@goseiger.com');
INSERT INTO public.admins(name, password, email) VALUES ('Hyde', 'GoseiBlue', 'blue@goseiger.com');
INSERT INTO public.admins(name, password, email) VALUES ('Duplicate1', 'dp1', 'dups@goseiger.com');
INSERT INTO public.admins(name, password, email) VALUES ('Duplicate2', 'dp2', 'dups@goseiger.com');
