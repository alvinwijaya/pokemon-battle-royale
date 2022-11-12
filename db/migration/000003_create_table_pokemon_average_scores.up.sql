CREATE TABLE pokemon_average_scores
(
    `id`       INT(11) UNSIGNED AUTO_INCREMENT NOT NULL,

    `pokemon_id`    INT(11) UNSIGNED NOT NULL,
    `average_score` DECIMAL(6,5) NOT NULL,
    `match_count`    INT(11) UNSIGNED NOT NULL,

    `created_at`  DATETIME NOT NULL DEFAULT NOW(),
    `updated_at`  DATETIME NOT NULL DEFAULT NOW(),

    PRIMARY KEY (`id`),
    INDEX (`pokemon_id`),
    INDEX (`average_score`)
);