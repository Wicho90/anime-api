
CREATE TABLE seasons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    number SMALLINT NOT NULL,
    slug VARCHAR(100) NOT NULL,
    image_url VARCHAR(255) NOT NULL
);


CREATE TABLE episodes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    number SMALLINT NOT NULL,
    duration INTERVAL NOT NULL,
    url VARCHAR(255) NOT NULL,
    slug VARCHAR(100) NOT NULL,
    season_id INTEGER REFERENCES seasons(id)
);


