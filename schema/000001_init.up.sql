CREATE TABLE IF NOT EXISTS clients
(
    id              serial       not null unique,
    client_name     varchar(255) not null,
    password_hash   varchar(255) not null,
    created_at      timestamp    not null default now(),
    update_at       timestamp    not null default now()
);

CREATE UNIQUE INDEX clients_name_on_password_idx ON clients (client_name, password_hash);

CREATE TABLE IF NOT EXISTS accounts(
    id          serial        not null,
    client_id   integer       not null references clients(id) on delete cascade,
    balance     numeric(10,2) not null default 0,
    update_at   timestamp     not null default now()
);