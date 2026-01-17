Has basics:
Members is many-to-many (many conversation - many users)
Messages is one-to-many (one conversatiion - many messages)
Each message belongs to exactly 1 conversation
Each message has exactly one sender

Conversations table

create table conversations (
    id text primary key,
    title text,
    created_at timestamp default now()
)

Member table

create table users (
    id text primary key,
    name text, 
    email text unique,
    password_hash text,
    created_at timestamp default now() -- if away for ..., will be deleted for freeing up spaces
)

create table conversation_members (
    conversation_id text references conversations(id),
    user_id text references users(id),
    primary key (conversation_id, user_id)
)

Message table (because it could be huge)
create table messages (
    id text primary key,
    conversation_id text references conversations(id),
    sender_id text references users(id),
    content text not null,
    created_at timestamp default now()
)