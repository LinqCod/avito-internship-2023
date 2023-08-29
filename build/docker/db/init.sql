CREATE TABLE IF NOT EXISTS users (
                                     id serial PRIMARY KEY,
                                     username varchar
);

CREATE TABLE IF NOT EXISTS segments (
                                        slug varchar PRIMARY KEY,
                                        description text
);

CREATE TABLE IF NOT EXISTS users_segments (
                                              user_id serial,
                                              slug varchar,
                                              CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
                                              CONSTRAINT fk_segment FOREIGN KEY (slug) REFERENCES segments (slug)
);