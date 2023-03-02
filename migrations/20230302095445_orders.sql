-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders(
    id                  serial          constraint order_pk PRIMARY KEY,
    order_uid           varchar(50),
    track_number        varchar(50),
    entry               varchar(10),
    delivery_id         int             references deliveries(id),
    payment_id          int,
    locale              varchar(5),
    internal_signature  varchar(50),
    customer_id         varchar(50),
    delivery_service    varchar(50),
    shardkey            varchar(50),
    sm_id               int,
    date_created        timestamp,
    oof_shard           varchar(30)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
