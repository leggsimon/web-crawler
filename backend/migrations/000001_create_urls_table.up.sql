-- Initial schema: migrations versioning is handled by golang-migrate (schema_migrations table).

CREATE TABLE urls (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    url         VARCHAR(2048) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY unique_url (url(768))
);
