-- +goose Up
-- +goose StatementBegin
create table attribute_group (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [name] text not null,
    unique ([name] collate nocase)
) strict;

create table attribute_set (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [attribute_group_id] integer not null references attribute_group(id) on delete restrict,
    [name] text not null,
    [in_box] integer not null,
    [in_filter] integer not null,
    [position] integer not null default 0,
    unique ([attribute_group_id], [name] collate nocase)
) strict;

create table attribute_value (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [attribute_set_id] integer not null references attribute_set(id) on delete restrict,
    [name] text not null,
    [position] integer not null default 0,
    unique ([attribute_set_id], [name] collate nocase)
) strict;

create table brand (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [name] text not null,
    unique ([name] collate nocase)
) strict;

create table category (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [slug] text not null,
    [name] text not null,
    [icon] text,
    [no_tax] integer not null,
    [is_published] integer not null,
    [attribute_group_id] integer references attribute_group(id) on delete restrict,
    [parent_id] integer references category(id) on delete restrict,
    [mp_path] text,
    [mp_level] integer not null default 0,
    [mp_position] integer not null default 0
) strict;

create unique index category_slug_ix on category(coalesce([parent_id], 0), [slug] collate nocase);
create index category_mp_level_ix on category(mp_level); -- Product.getQuerySelectBase
create index category_parent_id_ix on category(parent_id); -- Product.getQuerySelectBase

create table supplier (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [slug] text not null,
    [code] text not null,
    [name] text not null,
    [description] text,
    [is_published] integer not null,
    [position] integer not null default 0,
    unique ([slug] collate nocase),
    unique ([code] collate nocase),
    unique ([name] collate nocase)
) strict;

create table product_status (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [name] text not null,
    [color] text not null,
    [position] integer not null default 0,
    unique ([name] collate nocase)
) strict;

create table product (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [slug] text not null,
    [code] text not null,
    [name] text not null,
    [description] text,
    [quantity] real not null,
    [price] real,
    [brand_id] integer not null references brand(id) on delete restrict,
    [category_id] integer not null references category(id) on delete restrict,
    [supplier_id] integer not null references supplier(id) on delete restrict,
    [status_id] integer references product_status(id) on delete restrict,
    [is_published] integer not null,
    [hash] text not null,
    unique ([supplier_id], [slug] collate nocase),
    unique ([supplier_id], [code] collate nocase)
) strict;

create index product_supplier_code_ix on product(supplier_id, code); -- ExtCatalog.MergeProducts.mergeSQL
create index product_category_supplier_ix on product(category_id, supplier_id) where quantity > 0 and is_published = 1; -- Category.getQueryFindByParentID

create table product_attribute (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [product_id] integer not null references product(id) on delete cascade,
    [attribute_set_id] integer not null references attribute_set(id) on delete restrict,
    [attribute_value_id] integer not null references attribute_value(id) on delete restrict,
    unique ([product_id], [attribute_set_id])
) strict;

create index product_attribute_ix on product_attribute(attribute_set_id, product_id, attribute_value_id); -- Filter.getQueryFindAttributeFacets

create table product_media (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [product_id] integer not null references product(id) on delete cascade,
    [name] text not null,
    [position] integer not null default 0,
    [file] blob not null,
    [thumbnail] blob,
    unique ([product_id], [name] collate nocase)
) strict;

create table payment_method (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [name] text not null,
    [position] integer not null default 0,
    unique ([name] collate nocase)
) strict;

create table order_status (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [name] text not null,
    [color] text not null,
    [position] integer not null default 0,
    unique ([name] collate nocase)
) strict;


create table order_header (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [order_status_id] integer not null references order_status(id) on delete restrict,
    [payment_method_id] integer not null references payment_method(id) on delete restrict,
    [email] text not null,
    [language] text not null,
    [first_name] text not null,
    [last_name] text not null,
    [phone_number] text not null,
    [delivery_address] text not null,
    [comment] text,
    [notes] text
) strict;

create table order_line (
    [id] integer primary key autoincrement,
    [created_at] integer not null default (unixepoch()),
    [updated_at] integer not null default (unixepoch()),
    [deleted_at] integer,
    [order_id] integer not null references order_header(id) on delete cascade,
    [product_id] integer references product(id) on delete set null,
    [product_code] text not null,
    [product_snapshot] text not null,
    [quantity] integer not null,
    [price] real not null
) strict;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table attribute_group;
drop table attribute_set;
drop table attribute_value;
drop table brand;
drop table category;
drop table supplier;
drop table product_status;
drop table product;
drop table product_attribute;
drop table product_media;
drop table payment_method;
drop table order_status;
drop table order_header;
drop table order_line;
-- +goose StatementEnd
