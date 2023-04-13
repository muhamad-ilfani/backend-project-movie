CREATE TABLE IF NOT EXISTS movieapps.users(
    id BIGSERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT false,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255),
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS movieapps.movies(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,
    artists VARCHAR(255) NOT NULL,
    genres VARCHAR(255) NOT NULL,
    watch_url VARCHAR(255) NOT NULL,
    viewer int8 NOT NULL DEFAULT 0::int8,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255),
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS movieapps.vote(
    id BIGSERIAL PRIMARY KEY,
    user_id int8 NOT NULL,
    movie_id int8 NOT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255),
    deleted_at timestamp,
    CONSTRAINT app_fk_users FOREIGN KEY (user_id) REFERENCES movieapps.users (id),
    CONSTRAINT app_fk_movies FOREIGN KEY (movie_id) REFERENCES movieapps.movies (id)
);