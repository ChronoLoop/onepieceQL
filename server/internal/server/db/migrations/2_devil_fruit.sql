-- +goose Up

CREATE TABLE IF NOT EXISTS devil_fruits (
    id BIGSERIAL PRIMARY KEY,
    japanese_name TEXT,
    english_name TEXT,
	meaning Text,
	usage_debut TEXT,
	type TEXT,
	description TEXT,
	current_owner TEXT,
	previous_owner TEXT,
	avatar_src TEXT
) INHERITS (time);

DROP TRIGGER IF EXISTS set_timestamp on "devil_fruits";
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON devil_fruits
FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose Down
DROP TABLE IF EXISTS devil_fruits;
