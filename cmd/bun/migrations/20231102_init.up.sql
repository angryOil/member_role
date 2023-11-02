SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."member_role"
(
    id            SERIAL PRIMARY KEY,
    member_id     int     not null,
    cafe_id       int     not null,
    cafe_role_ids varchar not null,
    created_at    timestamptz
);


create unique index mr_cafe_id_member_id_unique on member_role (member_id, cafe_id);
