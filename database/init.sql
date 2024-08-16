CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    packs      JSONB     DEFAULT '[]'
);

CREATE TABLE packs
(
    id              SERIAL PRIMARY KEY,
    brand           VARCHAR(255),
    model           VARCHAR(255),
    year_purchased  INT,
    capacity_litres FLOAT,
    weight_empty    FLOAT,
    usage_count     INT
);

CREATE TABLE pack_instances
(
    id           SERIAL PRIMARY KEY,
    pack_id      INT REFERENCES packs (id),
    total_weight INT,
    usage_date   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    sections     JSONB DEFAULT '[]',
    user_id      INT REFERENCES users (id),
    notes        TEXT
);

CREATE TABLE sections
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255),
    description TEXT,
    usage_count INT
);

CREATE TABLE section_instances
    (
    id                        SERIAL PRIMARY KEY,
    section_id                INT REFERENCES sections (id),
    pack_instance_id          INT REFERENCES pack_instances (id),
    total_weight              INT,
    percentage_of_pack_weight FLOAT,
    items                    JSONB DEFAULT '[]',
    notes                    TEXT
);

CREATE TABLE items
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255),
    description TEXT,
    weight      FLOAT,
    usage_count INT
);

CREATE TABLE item_instances
(
    id          SERIAL PRIMARY KEY,
    item_id     INT REFERENCES items (id),
    section_id  INT REFERENCES sections (id),
    notes        TEXT
);
