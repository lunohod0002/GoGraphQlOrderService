-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       fio TEXT NOT NULL,
                       balance INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          price INTEGER NOT NULL CHECK (price >= 0)
);

CREATE TABLE carts (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       total_sum INTEGER DEFAULT 0,
                       discount INTEGER  DEFAULT 0,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE items (
                       id SERIAL PRIMARY KEY,
                       cart_id INTEGER NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
                       product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                       quantity INTEGER NOT NULL CHECK (quantity > 0),
                       UNIQUE(cart_id, product_id)
);

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        cart_id INTEGER NOT NULL REFERENCES carts(id) ON DELETE RESTRICT,
                        status TEXT NOT NULL DEFAULT 'pending',
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;