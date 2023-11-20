-- +goose Up
CREATE TABLE categories(
    id int primary key,
    name varchar(50) not null
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE categories;
-- +goose StatementEnd
