CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    firstname TEXT,
    lastname TEXT,
    email TEXT UNIQUE NOT NULL,
    pass TEXT NOT NULL
);

CREATE TABLE users_telegram (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    tuser_id TEXT NOT NULL,
    chat_id TEXT NOT NULL,
    PRIMARY KEY (user_id, tuser_id)
);

CREATE TABLE words (
    id BIGSERIAL PRIMARY KEY,
    word TEXT NOT NULL,
    lang TEXT NOT NULL,
    part_of_speech TEXT,
    transcription TEXT
);

CREATE TABLE translation (
    meaning_id BIGINT REFERENCES words(id) ON DELETE CASCADE,
    translation_id BIGINT REFERENCES words(id) ON DELETE CASCADE,
    PRIMARY KEY (meaning_id, translation_id)
);

CREATE TABLE word_groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    lang_from TEXT NOT NULL,
    lang_to TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT 'active'
);

CREATE TABLE word_groups_words (
    group_id INTEGER REFERENCES word_groups(id) ON DELETE CASCADE,
    word_meaning_id BIGINT REFERENCES words(id) ON DELETE CASCADE,
    level TEXT,
    status TEXT NOT NULL DEFAULT 'active',
    last_right TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (group_id, word_meaning_id)
);