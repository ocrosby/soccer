CREATE TABLE IF NOT EXISTS countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(3) UNIQUE NOT NULL
);

-- Index for the 'code' column
CREATE INDEX IF NOT EXISTS idx_countries_code ON countries (code);

-- Index for the 'name' column
CREATE INDEX IF NOT EXISTS idx_countries_name ON countries (name);
