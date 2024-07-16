CREATE TABLE IF NOT EXISTS clients
(
    id              serial       not null,
    client_name     varchar(255) not null,
    password_hash   varchar(255) not null,
    created_at      timestamp    not null default now(),
    update_at       timestamp    not null default now()
);

CREATE UNIQUE INDEX clients_name_on_password_idx ON clients (client_name, password_hash);

-- CREATE TABLE IF NOT EXISTS algorithm_status(
--     id          serial  NOT NULL,
--     client_id   integer NOT NULL references clients(client_id) on delete cascade,
--     vwap        boolean NOT NULL default false,
--     twap        boolean NOT NULL default false,
--     hft         boolean NOT NULL default false
-- );