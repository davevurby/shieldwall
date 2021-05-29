create table if not exists "policy" (
    subject text not null,
    namespace text not null,
    policy text not null,
    effect text not null,
    primary key (subject, namespace, policy, effect)
);