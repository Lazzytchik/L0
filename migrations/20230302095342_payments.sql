-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments(
    id              serial             constraint payments_pk PRIMARY KEY,
    transaction     varchar(50),
    request_id      varchar(50),
    currency        varchar(5),
    provider        varchar(50),
    amount          real,
    payment_dt      bigint,
    bank            varchar(50),
    delivery_cost   real,
    goods_total     int,
    custom_fee      real
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd
