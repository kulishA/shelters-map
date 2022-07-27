CREATE TABLE IF NOT EXISTS shelters
(
    id             serial       not null unique,
    city           varchar(255) null,
    address        varchar(255) null,
    address_number varchar(255) null,
    latitude       varchar(255) null,
    longitude      varchar(255) null,
    closer_type    varchar(255) null,
    shelter_type   varchar(255) null,
    building_type  varchar(255) null,
    owner          varchar(255) null,
    phone          varchar(255) null,
    ramp           varchar(255) null
);

CREATE INDEX latitude_longitude_idx ON shelters (latitude, longitude);