create table hawbs (
                    id serial primary key,
                    origin text not null,
                    destination text not null,
                    consignee text not null,
                    consignor text not null,
                    content text not null,
                    weight integer not null,
                    number text,
                    pieces integer,
                    mawb_id integer references mawbs(id),
                    created_at timestamp not null default current_timestamp,
                    updated_at timestamp not null default current_timestamp
);

