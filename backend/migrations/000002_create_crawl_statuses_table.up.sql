CREATE TABLE crawl_statuses (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    url_id      INT NOT NULL UNIQUE,                          -- one status per URL
    status      ENUM('queued', 'running', 'done', 'error') NOT NULL DEFAULT 'queued',
    error_msg   TEXT NULL,                                    -- populated only if status = 'error'
    started_at  TIMESTAMP NULL,                               -- set when status moves to 'running'
    finished_at TIMESTAMP NULL,                               -- set when status moves to 'done' or 'error'
    FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE
);
