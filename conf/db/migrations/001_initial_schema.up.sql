CREATE TABLE accounts (
    id              serial PRIMARY KEY,
    email           varchar(256) UNIQUE NOT NULL,
    created         timestamp without time zone NOT NULL
);

CREATE INDEX idx_account_email ON accounts ((lower(email)));


CREATE TABLE devices (
    id              serial PRIMARY KEY,
    account_id      integer REFERENCES accounts(id) NOT NULL,
    name            varchar(100) NOT NULL,
    password_hash   bytea NOT NULL,
    created         timestamp without time zone NOT NULL,
    last_login      timestamp without time zone
);

CREATE INDEX idx_device_account_id ON devices (account_id);


CREATE TABLE rooms (
    id              serial PRIMARY KEY,
    name            varchar(25) UNIQUE NOT NULL,
    password_hash   bytea NOT NULL,
    created         timestamp without time zone NOT NULL
);

CREATE INDEX idx_room_name ON rooms (name);


CREATE TABLE channels (
    id              serial PRIMARY KEY,
    room_id         integer REFERENCES rooms(id) NOT NULL,
    name            varchar(25) NOT NULL,
    created         timestamp without time zone NOT NULL,
    UNIQUE (room_id, name)
);

CREATE INDEX idx_channel_room_id ON channels (room_id);


CREATE TABLE members (
    id              serial PRIMARY KEY,
    room_id         integer REFERENCES rooms(id) NOT NULL,
    account_id      integer REFERENCES accounts(id) NOT NULL,
    channel_id      integer REFERENCES channels(id),
    name            varchar(25) NOT NULL,
    admin           boolean NOT NULL,
    banned          boolean NOT NULL,
    created         timestamp without time zone NOT NULL,
    last_login      timestamp without time zone,
    UNIQUE (room_id, account_id),
    UNIQUE (room_id, name)
);

CREATE INDEX idx_member_room_id ON members (room_id);
CREATE INDEX idx_member_account_id ON members (account_id);
