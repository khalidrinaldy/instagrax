-- +migrate Up
-- +migrate StatementBegin

create table comments (
    id uuid default uuid_generate_v4(),
    text text not null,
    user_id uuid not null,
    post_id uuid not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    primary key (id)
)

-- +migrate StatementEnd