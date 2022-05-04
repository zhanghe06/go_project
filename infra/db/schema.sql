CREATE DATABASE /*!32312 IF NOT EXISTS */ `sap_cert_mgt` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;


-- SELECT @@SESSION.sql_mode;
-- # 去除 sql_mode中的 NO_ZERO_IN_DATE 和 NO_ZERO_DATE
-- SET SESSION sql_mode = 'modes';

USE `sap_cert_mgt`;

CREATE TABLE IF NOT EXISTS `cert`
(
    `id`                  INT         NOT NULL AUTO_INCREMENT,
    `auth_id`             VARCHAR(32) NOT NULL DEFAULT '' COMMENT '客户端ID',
    `p_version`           VARCHAR(32) NOT NULL DEFAULT '' COMMENT '接口版本',
    `cont_rep`            VARCHAR(32) NOT NULL DEFAULT '' COMMENT '内容存储库',
    `serial_number`       VARCHAR(20) NOT NULL DEFAULT '' COMMENT '证书序列号',
    `version`             TINYINT     NOT NULL DEFAULT 0 COMMENT '证书版本（0:V1,1:V2,2:V3）',
    `issuer_name`         VARCHAR(64) NOT NULL DEFAULT '' COMMENT '颁发机构',
    `signature_algorithm` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '签名算法',
    `not_before`          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '有效期开始时间',
    `not_after`           TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '有效期结束时间',
    `enabled_state`       TINYINT     NOT NULL DEFAULT 0 COMMENT '启用状态（0:已停用,1:已启用）',
    `deleted_state`       TINYINT     NOT NULL DEFAULT 0 COMMENT '删除状态（0:未删除,1:已删除）',
    `created_at`          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`          TIMESTAMP   NULL COMMENT '删除时间',
    `created_by`          VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建人员',
    `updated_by`          VARCHAR(64) NOT NULL DEFAULT '' COMMENT '更新人员',
    `deleted_by`          VARCHAR(64) NOT NULL DEFAULT '' COMMENT '删除人员',
    PRIMARY KEY (`id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='证书管理';

CREATE TABLE IF NOT EXISTS `operation_log`
(
    `id`            INT          NOT NULL AUTO_INCREMENT,
    `op_type`       VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '操作类型（create:创建、update:更新、delete:删除）',
    `rs_type`       VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '资源类型（cert、）',
    `rs_id`         INT          NOT NULL DEFAULT 0 COMMENT '资源ID',
    `op_detail`     VARCHAR(512) NOT NULL DEFAULT '' COMMENT '操作详情',
    `op_error`      VARCHAR(512) NOT NULL DEFAULT '' COMMENT '操作错误',
    `deleted_state` TINYINT      NOT NULL DEFAULT 0 COMMENT '删除状态（0:未删除,1:已删除）',
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`    TIMESTAMP    NULL COMMENT '删除时间',
    `created_by`    VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '创建人员',
    `updated_by`    VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '更新人员',
    `deleted_by`    VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '删除人员',
    PRIMARY KEY (`id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='操作日志';

CREATE TABLE IF NOT EXISTS `notice_strategy`
(
    `id`                INT          NOT NULL AUTO_INCREMENT,
    `notice_type`       TINYINT      NOT NULL DEFAULT 0 COMMENT '证书过期通知类型（0:邮件,1:短信）',
    `trigger_threshold` TINYINT      NOT NULL DEFAULT 0 COMMENT '证书过期触发阈值（单位:天;范围:0-255）',
    `to_emails`         VARCHAR(256) NOT NULL DEFAULT '' COMMENT '接收邮箱（半角逗号分隔）',
    `enabled_state`     TINYINT      NOT NULL DEFAULT 0 COMMENT '启用状态（0:已停用,1:已启用）',
    `deleted_state`     TINYINT      NOT NULL DEFAULT 0 COMMENT '删除状态（0:未删除,1:已删除）',
    `created_at`        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        TIMESTAMP    NULL COMMENT '删除时间',
    `created_by`        VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '创建人员',
    `updated_by`        VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '更新人员',
    `deleted_by`        VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '删除人员',
    PRIMARY KEY (`id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='通知策略';

CREATE TABLE IF NOT EXISTS `notice_conf`
(
    `id`            INT          NOT NULL AUTO_INCREMENT,
    `notice_type`   TINYINT      NOT NULL DEFAULT 0 COMMENT '证书过期通知类型（0:邮件,1:短信）',
    `config_data`   VARCHAR(256) NOT NULL DEFAULT '' COMMENT '配置数据（加密存储:json序列化+对称加密+base64）',
    `deleted_state` TINYINT      NOT NULL DEFAULT 0 COMMENT '删除状态（0:未删除,1:已删除）',
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`    TIMESTAMP    NULL COMMENT '删除时间',
    `created_by`    VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '创建人员',
    `updated_by`    VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '更新人员',
    `deleted_by`    VARCHAR(64)  NOT NULL DEFAULT '' COMMENT '删除人员',
    PRIMARY KEY (`id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='通知配置';

CREATE TABLE IF NOT EXISTS `notice_event`
(
    `id`                 INT         NOT NULL AUTO_INCREMENT,
    `cert_id`            INT         NOT NULL DEFAULT 0 COMMENT '证书ID',
    `notice_strategy_id` INT         NOT NULL DEFAULT 0 COMMENT '策略ID',
    `event_state`        TINYINT     NOT NULL DEFAULT 0 COMMENT '事件状态（0:pending准备,1:waiting等待,2:process进行,3:success成功,4:failure失败）',
    `deleted_state`      TINYINT     NOT NULL DEFAULT 0 COMMENT '删除状态（0:未删除,1:已删除）',
    `created_at`         TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`         TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`         TIMESTAMP   NULL COMMENT '删除时间',
    `created_by`         VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建人员',
    `updated_by`         VARCHAR(64) NOT NULL DEFAULT '' COMMENT '更新人员',
    `deleted_by`         VARCHAR(64) NOT NULL DEFAULT '' COMMENT '删除人员',
    PRIMARY KEY (`id`),
    KEY `idx_cert_id` (`cert_id`),
    KEY `idx_notice_strategy_id` (`notice_strategy_id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='通知事件';
