-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_items(
    id          serial      constraint order_items_pk PRIMARY KEY,
    order_id    int         references orders(id),
    item_id     int         references items(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_items;
-- +goose StatementEnd
