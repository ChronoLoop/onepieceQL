-- +goose Up

CREATE TABLE IF NOT EXISTS characters (
    id BIGSERIAL PRIMARY KEY,
    japanese_name TEXT,
    english_name TEXT,
    debut TEXT,
	affiliations Text,
	origin TEXT,
	age TEXT,
	birthday TEXT,
	blood_type TEXT,
	bounty TEXT,
	description TEXT,
	devil_fruit_name TEXT,
    avatar_src TEXT
) INHERITS (time);

DROP TRIGGER IF EXISTS set_timestamp on "characters";
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON characters
FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose Down
DROP TABLE IF EXISTS characters;
