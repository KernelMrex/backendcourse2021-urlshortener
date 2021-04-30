CREATE TABLE IF NOT EXISTS `redirects`
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    key         TEXT,
    destination TEXT
);
