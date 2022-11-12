CREATE TABLE match_details
(
    `id`       INT(11) UNSIGNED AUTO_INCREMENT NOT NULL,
    `match_id` INT(11) UNSIGNED NOT NULL,

    `pokemon_id` INT(11) UNSIGNED NOT NULL,
    `score`      INT(11) UNSIGNED NOT NULL,

    `created_at`  DATETIME NOT NULL DEFAULT NOW(),
    `updated_at`  DATETIME NOT NULL DEFAULT NOW(),
    `deleted_at`  DATETIME NULL,

    PRIMARY KEY (`id`),
    INDEX (`match_id`)
);