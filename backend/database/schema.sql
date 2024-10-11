CREATE TABLE users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    username VARCHAR(20) NOT NULL,
    password TEXT NOT NULL UNIQUE,
    terms date NOT NULL,
    bio VARCHAR(300)
);

CREATE TABLE avatar (
    avatar_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    avatar TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE topics(
    topic_id INTEGER PRIMARY KEY AUTOINCREMENT,
    topic_name VARCHAR(30) NOT NULL
);

CREATE TABLE journals(
    journal_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    topic_id INTEGER NOT NULL,
    journal_name VARCHAR(20) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (topic_id) REFERENCES topics(topic_id) ON DELETE CASCADE
);

CREATE TABLE pages(
    page_id INTEGER PRIMARY KEY AUTOINCREMENT,
    journal_id INTEGER NOT NULL,
    page_entry VARCHAR(500),
    page_date date NOT NULL,
    FOREIGN KEY (journal_id) REFERENCES journals(journal_id) ON DELETE CASCADE
);