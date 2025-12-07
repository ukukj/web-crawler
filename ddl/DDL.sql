CREATE TABLE jobs
(
    id           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    source       VARCHAR(50)  NOT NULL,
    ref          TEXT,
    company_name VARCHAR(255),
    work_place   VARCHAR(255),
    career       VARCHAR(255),
    education    VARCHAR(255),
    crawled_at   DATETIME   NOT NULL,
    created_at   DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active    TINYINT(1) NOT NULL DEFAULT 1,

    UNIQUE KEY uq_source_ref (source, ref (255))
);
