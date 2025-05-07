-- +goose Up
-- +goose StatementBegin
CREATE TABLE languages (
    id INTEGER PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL
);

CREATE TABLE translations (
    id INTEGER PRIMARY KEY,
    key TEXT NOT NULL,
    language_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    UNIQUE(key, language_id),
    FOREIGN KEY (language_id) REFERENCES languages(id)
);

CREATE TABLE profile (
    id INTEGER PRIMARY KEY CHECK (id = 1),
    name_key TEXT NOT NULL,
    title_key TEXT NOT NULL,
    about_key TEXT NOT NULL,
    image_url TEXT
);

CREATE TABLE companies (
    id INTEGER PRIMARY KEY,
    name_key TEXT NOT NULL,
    website_url TEXT,
    location_key TEXT
);

CREATE TABLE positions (
    id INTEGER PRIMARY KEY,
    company_id INTEGER NOT NULL,
    title_key TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    description_key TEXT,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE position_items (
    id INTEGER PRIMARY KEY,
    position_id INTEGER NOT NULL,
    item_key TEXT NOT NULL,
    FOREIGN KEY (position_id) REFERENCES positions(id)
);

CREATE TABLE projects (
    id INTEGER PRIMARY KEY,
    title_key TEXT NOT NULL,
    desc_key TEXT NOT NULL
);

CREATE TABLE contacts (
    id INTEGER PRIMARY KEY,
    type TEXT NOT NULL,
    label_key TEXT NOT NULL,
    value TEXT NOT NULL
);

CREATE TABLE open_source (
    id INTEGER PRIMARY KEY,
    title_key TEXT NOT NULL,
    desc_key TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS contacts;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS position_items;
DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS profile;
DROP TABLE IF EXISTS translations;
DROP TABLE IF EXISTS languages;
-- +goose StatementEnd
