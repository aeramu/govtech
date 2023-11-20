-- +goose Up
CREATE TABLE product_reviews(
    id int not null auto_increment primary key,
    user_id int not null,
    product_id int not null,
    rating int not null,
    comment varchar(256) not null,
    foreign key(product_id) references products(id)
);

-- +goose Down
DROP TABLE product_reviews;
