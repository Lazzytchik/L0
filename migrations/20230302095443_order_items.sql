-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_items(
    id          int,
    order_id    int,
    item_id     int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_items;
-- +goose StatementEnd
