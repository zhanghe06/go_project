CREATE DATABASE /*!32312 IF NOT EXISTS */ `go_project` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci */;

USE `go_project`;

CREATE TABLE IF NOT EXISTS `user`
(
    `id`            INT         NOT NULL AUTO_INCREMENT,
    `name`          VARCHAR(20) NOT NULL DEFAULT '' COMMENT '姓名',
    `gender`        TINYINT     NOT NULL DEFAULT 0 COMMENT '性别（1:男,2:女）',
    `enabled_state` TINYINT     NOT NULL DEFAULT 0 COMMENT '启用状态（0:已停用,1:已启用）',
    `deleted_state` TINYINT     NOT NULL DEFAULT 0 COMMENT '删除状态（0:未删除,1:已删除）',
    `created_at`    TIMESTAMP   NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
    `updated_at`    TIMESTAMP   NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
    `deleted_at`    TIMESTAMP   NULL COMMENT '删除时间',
    `created_by`    VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建人员',
    `updated_by`    VARCHAR(64) NOT NULL DEFAULT '' COMMENT '更新人员',
    `deleted_by`    VARCHAR(64) NOT NULL DEFAULT '' COMMENT '删除人员',
    PRIMARY KEY (`id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户';
