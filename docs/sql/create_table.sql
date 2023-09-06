CREATE TABLE `user_infos` (
    `id` bigint unsigned AUTO_INCREMENT,
    `created_at` datetime(3) NULL,
    `updated_at` datetime(3) NULL,
    `deleted_at` datetime(3) NULL,
    `uid` bigint unsigned,
    `username` longtext,
    `nickname` longtext,
    `avatar_url` longtext,
    PRIMARY KEY (`id`),
    INDEX `idx_user_infos_deleted_at` (`deleted_at`)
);

CREATE TABLE `user_auths` (
    `id` bigint unsigned AUTO_INCREMENT,
    `created_at` datetime(3) NULL,
    `updated_at` datetime(3) NULL,
    `deleted_at` datetime(3) NULL,
    `uid` bigint unsigned,
    `username` longtext,
    `password_hash` longtext,
    PRIMARY KEY (`id`),
    INDEX `idx_user_auths_deleted_at` (`deleted_at`)
)
