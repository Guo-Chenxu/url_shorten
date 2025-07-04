create table url_shorten.short_url
(
    id          bigint auto_increment
        primary key,
    origin_url  text                               not null,
    code        varchar(255)                       not null,
    expire_time datetime                           not null,
    create_time datetime default CURRENT_TIMESTAMP not null,
    constraint short_url_code_expire_time_uindex
        unique (code, expire_time),
    constraint short_url_code_uindex
        unique (code)
);

