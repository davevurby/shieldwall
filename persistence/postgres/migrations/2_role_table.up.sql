create table if not exists "role" (
  id text not null,
  namespaces text[] not null,
  primary key (id)
);