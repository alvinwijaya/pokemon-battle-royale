CREATE TABLE matches
(
    `id`   INT(11) UNSIGNED AUTO_INCREMENT NOT NULL,

    `created_at`  DATETIME NOT NULL DEFAULT NOW(),
    `updated_at`  DATETIME NOT NULL DEFAULT NOW(),

    PRIMARY KEY (`id`),
    INDEX (`created_at`)
);