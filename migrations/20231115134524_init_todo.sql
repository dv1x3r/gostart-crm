-- +goose Up
-- +goose StatementBegin
create table todo (
    id integer primary key,
    name text not null,
    description text,
    quantity integer
) strict;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table todo;
-- +goose StatementEnd
