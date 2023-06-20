create table mawbs (
                    id serial primary key,
                    number text,
                    origin text,
                    destination text,
                    created_at timestamp not null default current_timestamp,
                    updated_at timestamp not null default current_timestamp
);