-- +goose Up
-- +goose StatementBegin
CREATE TABLE items(
    id              serial             constraint items_pk PRIMARY KEY,
    chrt_id         int,
    track_number    varchar(50),
    price           real,
    rid             varchar(50),
    name            varchar(50),
    sale            real,
    size            varchar(50),
    total_price     real,
    nm_id           int,
    brand           varchar(50),
    status          int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
