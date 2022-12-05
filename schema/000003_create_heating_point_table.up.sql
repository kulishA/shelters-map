CREATE TABLE IF NOT EXISTS heating_points
(
    id                 serial       not null unique,
    city               varchar(255) null,
    address            varchar(255) null,
    district            varchar(255) null,
    latitude           varchar(255) null,
    longitude          varchar(255) null,
    closer_type        varchar(255) null,
    start_date         varchar(255) null,
    capacity           integer      null,
    mobile_phone       varchar(255) null,
    phone              varchar(255) null,
    responsible_person varchar(255) null,
    kitchen_room       bool         not null,
    mother_room        bool         not null,
    disabilities_room  bool         not null
);

CREATE INDEX latitude_longitude_idx ON heating_points (latitude, longitude);