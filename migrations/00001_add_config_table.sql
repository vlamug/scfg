-- +goose Up
CREATE TABLE config (
  cfg_id SERIAL NOT NULL,
  ckey CHARACTER VARYING(100),
  cset TEXT
);

-- +goose Down
DROP TABLE config;
