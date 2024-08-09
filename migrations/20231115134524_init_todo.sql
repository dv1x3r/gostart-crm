-- +goose Up
-- +goose StatementBegin
create table todo (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [name] text not null,
    [description] text,
    [quantity] integer
) strict;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table todo;
-- +goose StatementEnd
