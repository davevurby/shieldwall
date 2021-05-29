create table if not exists "identity" (
    id text not null,
    namespace text not null,
    roles text[] not null default '{}',
    primary key (id, namespace)
);