-- +goose Up
-- +goose StatementBegin
CREATE TABLE deliveries(
    id          serial             constraint deliveries_pk PRIMARY KEY,
    name        varchar(50),
    phone       varchar(50),
    zip         varchar(50),
    city        varchar(50),
    address     varchar(50),
    region      varchar(50),
    email       varchar(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE deliveries;
-- +goose StatementEnd
