CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR
);

CREATE TABLE IF NOT EXISTS segments (
                                        slug VARCHAR PRIMARY KEY,
                                        description TEXT
);

CREATE TABLE IF NOT EXISTS users_segments (
                                              user_id INTEGER REFERENCES users(id),
                                              slug VARCHAR REFERENCES segments(slug),
                                              CONSTRAINT users_segments_pk PRIMARY KEY (user_id, slug)
);