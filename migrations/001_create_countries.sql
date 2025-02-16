create table country (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  country_name varchar not null
);

---- create above / drop below ----

drop table country;
