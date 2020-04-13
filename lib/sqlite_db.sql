create table if not exists settings (
    key text not null, 
    val text not null, 
    primary key (key)
);

create index IDX_settings_key on settings(key);

insert into settings values('VERSION', '1');

create table if not exists domain (
    id integer generated, 
    domain_name text not null, 
    table_name text not null, 
    primary key(id)
);

create table if not exists default_links (
    short_link text not null,
    long_link text not null,
    is_enabled integer not null default 0 check (is_enabled == 0 or is_enabled == 1), 
    created_on integer not null default (strftime('%s','now'))
);

create unique index UDX_def_short_links on default_links(short_link);

create table admin_tokens (
    token text not null,
    token_description text not null,
    domain_id integer null,
    expire_at integer default 0,
    is_Root integer not null default 0 check (is_Root == 0 or is_Root == 1)
);

create unique index IDX_tokens_uuid on admin_tokens(token);
