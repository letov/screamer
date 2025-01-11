-- +goose Up
CREATE TABLE metrics (
		id SERIAL PRIMARY KEY,
		type metric_type_enum NOT NULL,
		name VARCHAR(50) NOT NULL,
		value FLOAT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE metrics;
