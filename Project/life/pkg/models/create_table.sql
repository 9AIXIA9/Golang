CREATE TABLE `user`
(
    `id`          bigint(20)                             NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(20)                             NOT NULL,
    `username`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password`    varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `email`       varchar(64) COLLATE utf8mb4_general_ci,
    `gender`      tinyint(4)                             NOT NULL DEFAULT 0,
    `create_time` timestamp                              NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp                              NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_general_ci;

CREATE TABLE `capability`
(
    `id`               bigint(20)                             NOT NULL AUTO_INCREMENT,
    `user_id`          bigint(20)                             NOT NULL,
    `capability_name`  varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `capability_score` int                                    NOT NULL DEFAULT 0,
    `create_time`      timestamp                              NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time`      timestamp                              NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_capability` (`user_id`, `capability_name`) USING BTREE,
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;