CREATE TABLE users (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "login" TEXT UNIQUE NOT NULL,
    "name" TEXT NOT NULL,
    "password_hash" TEXT NOT NULL
);

CREATE TABLE notebooks (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "description" TEXT NOT NULL,
    "user_id" BIGINT REFERENCES users("id") ON DELETE CASCADE NOT NULL
);

CREATE TABLE notes (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "title" TEXT NOT NULL,
    "body" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    "user_id" BIGINT REFERENCES users("id") ON DELETE CASCADE NOT NULL,
    "notebook_id" BIGINT REFERENCES notebooks("id") ON DELETE CASCADE NOT NULL
);

CREATE TABLE shared_notes (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "whose_id" BIGINT REFERENCES users("id") ON DELETE CASCADE NOT NULL,
    "whom_id" BIGINT REFERENCES users("id") ON DELETE CASCADE NOT NULL,
    "note_id" BIGINT REFERENCES notes("id") ON DELETE CASCADE NOT NULL,
    "accepted" BOOLEAN NOT NULL,
    UNIQUE("whose_id", "whom_id", "note_id")
);

CREATE TABLE todo_lists (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "title" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,
    "user_id" BIGINT REFERENCES users("id") ON DELETE CASCADE NOT NULL,
    "notebook_id" BIGINT REFERENCES notebooks("id") ON DELETE CASCADE NOT NULL
);

CREATE TABLE todo_items (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "body" TEXT NOT NULL,
    "done" BOOLEAN NOT NULL,
    "todo_list_id" BIGINT REFERENCES todo_lists("id") ON DELETE CASCADE NOT NULL
);

CREATE TABLE shared_todo_lists (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "whose_id" BIGINT REFERENCES users("id") ON DELETE CASCADE NOT NULL,
    "whom_id" BIGINT REFERENCES users("id") ON DELETE CASCADE NOT NULL,
    "todo_list_id" BIGINT REFERENCES todo_lists("id") ON DELETE CASCADE NOT NULL,
    "accepted" BOOLEAN NOT NULL,
    UNIQUE("whose_id", "whom_id", "todo_list_id")
);

CREATE TABLE sessions (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "user_id" BIGINT REFERENCES users("id") ON DELETE CASCADE NOT NULL,
    "refresh_token" UUID NOT NULL,
    "fingerprint" TEXT NOT NULL,
    "expires_in" TIMESTAMPTZ NOT NULL
);
