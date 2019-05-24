
-- +migrate Up
CREATE TABLE public.users (
	id serial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	name text NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


-- +migrate Down
DROP TABLE users;
