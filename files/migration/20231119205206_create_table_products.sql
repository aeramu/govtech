-- +goose Up
CREATE TABLE products(
    id int not null auto_increment primary key,
    sku varchar(50) not null unique,
    title varchar(50) not null,
    description varchar(256) not null,
    category_id int not null,
    image_url varchar(256) not null,
    weight int not null,
    price int not null,
    rating float not null,
    created_at timestamp not null default now(),
    foreign key(category_id) references categories(id)
);

-- +goose Down
DROP TABLE products;
