-- +goose Up
CREATE TABLE IF NOT EXISTS `dag` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `desc` varchar(255),
    `cron` varchar(255),
    `vars` text,
    `status` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    INDEX idx_uid (`uid`),
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `task` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` varchar(255) NOT NULL,
    `dag_uid` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `action_name` varchar(255) NOT NULL,
    `timeout_secs` int,
    `params` text,
    `prechecks` text,
    `depend_on` text,
    INDEX idx_uid_daguid (`uid`, `dag_uid`),
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `dag_instance` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` varchar(255) NOT NULL,
    `dag_uid` varchar(255) NOT NULL,
    `trigger` varchar(255) NOT NULL,
    `worker` varchar(255) NOT NULL,
    `vars` text,
    `status` varchar(255) NOT NULL,
    `reason` varchar(255),
    `cmd` varchar(255),
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    INDEX idx_uid_daguid (`uid`, `dag_uid`),
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `task_instance` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` varchar(255) NOT NULL,
    `task_uid` varchar(255) NOT NULL,
    `dag_instance_uid` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `action_name` varchar(255) NOT NULL,
    `timeout_secs` int,
    `params` text,
    `status` varchar(255) NOT NULL,
    `reason` varchar(255),
    `precheck` text,
    `depend_on` text,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    INDEX idx_uid_taskuid_daginsuid (`uid`, `task_uid`, `dag_instance_uid`),
    PRIMARY KEY (`id`)
);

-- +goose Down
DROP TABLE IF EXISTS `dag`;
DROP TABLE IF EXISTS `task`;
DROP TABLE IF EXISTS `dag_instance`;
DROP TABLE IF EXISTS `task_instance`;