-- +goose Up
INSERT INTO categories(id, name)
    VALUES (1, 'Food'), (2, 'Pet'), (3, 'Furniture'), (4, 'Art');

-- +goose Down
DELETE FROM categories;
