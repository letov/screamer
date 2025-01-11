-- +goose Up
CREATE TYPE "public"."metric_type_enum" AS ENUM('counter', 'gauge');

-- +goose Down
DROP TYPE "public"."metric_type_enum";
