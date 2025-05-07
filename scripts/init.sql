DROP DATABASE IF EXISTS `OnlineMeeting`;
CREATE DATABASE `OnlineMeeting` DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

USE `OnlineMeeting`;

SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `t_meeting`;
CREATE TABLE `t_meeting` (
    `id` BIGINT(20) AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `room_number` CHAR(9) NOT NULL COMMENT '会议ID', -- 9位随机数
    `topic` VARCHAR(128) NOT NULL COMMENT '会议主题',
    `mode` TINYINT(1) NOT NULL COMMENT '会议模式', -- 1:漫游模式, 2:会议室模式, 3:虚拟人模式
    `distance` INT NOT NULL COMMENT '电子围栏半径', -- 漫游模式下，多少米以内的用户自动加入会议
    `type` TINYINT(1) NOT NULL COMMENT '会议类型', -- 1:即时会议, 2:预约会议
    `status` TINYINT(1) NOT NULL COMMENT '会议状态', -- 1:待开始, 2:进行中, 3:已结束, 4:已取消
    `location` VARCHAR(128) NOT NULL COMMENT '会议地点',
    `creator_id` VARCHAR(128) NOT NULL COMMENT '会议创建人ID',
    `description`VARCHAR(128) NOT NULL COMMENT '会议描述信息',
    `create_time` DATETIME NOT NULL COMMENT '会议创建时间',
    `start_time` DATETIME NOT NULL COMMENT '会议开始时间', -- 预约会议最长3个月
    `end_time` DATETIME COMMENT '会议结束时间', -- 根据实际结束时间设置
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_room_number` (`room_number`),
    KEY `idx_topic` (`topic`)
) COMMENT='会议表';

DROP TABLE IF EXISTS `t_meeting_participant`;
CREATE TABLE `t_meeting_participant` (
    `id` BIGINT(20) AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `m_room_number` CHAR(9) NOT NULL COMMENT '会议ID, 关联t_meeting表room_number字段',
    `user_id` VARCHAR(128) NOT NULL COMMENT '被邀请用户的ID',
    `user_name` VARCHAR(128) NOT NULL COMMENT '被邀请的用户名称',
    `role` INT UNSIGNED NOT NULL COMMENT '角色, 1:管理员, 2:主持人, 3:普通成员',
    `status` TINYINT(1) NOT NULL COMMENT '邀请状态, 1:待处理, 2:已接受, 3:已拒绝',
    `update_time` DATETIME COMMENT '接受/拒绝时间',
    `join_time` DATETIME COMMENT '加入时间',
    `exit_time` DATETIME COMMENT '退出时间',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`m_room_number`) REFERENCES `t_meeting`(`room_number`),
    UNIQUE KEY `idx_room_number_user_id` (`m_room_number`, `user_id`),
    KEY `idx_user_id` (`user_id`)
)COMMENT='会议参与者表';

DROP TABLE IF EXISTS `t_file`;
CREATE TABLE `t_file` (
    `id` BIGINT(20) NOT NULL COMMENT '主键ID',
    `m_room_number` CHAR(9) NOT NULL COMMENT '关联会议ID，对应t_meeting.room_number',
    `file_name` VARCHAR(128) NOT NULL COMMENT '文件原始名称',
    `save_name` VARCHAR(128) NOT NULL COMMENT '文件保存名称',
    `file_size` INT UNSIGNED NOT NULL COMMENT '文件大小 (单位: 字节)',  -- 限制文件最大为100MB
    `file_type` INT NOT NULL COMMENT '文件类型',
    `storage_path` VARCHAR(1024) NOT NULL COMMENT '文件存储路径或标识符 (如 UUID 或云存储 Key)',
    `uploader_id` VARCHAR(128) NOT NULL COMMENT '上传者ID',
    `uploader_name` VARCHAR(128) NOT NULL COMMENT '上传者名字',
    `upload_time` DATETIME NOT NULL COMMENT '上传时间',
    `deletor_id` VARCHAR(128) COMMENT '删除者ID',
    `deletor_name` VARCHAR(128) COMMENT '删除者名字',
    `delete_time` DATETIME COMMENT '删除时间',
    `status` TINYINT(1) NOT NULL COMMENT '状态: 1-有效, 2-已删除',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`m_room_number`) REFERENCES `t_meeting`(`room_number`),
    KEY `idx_m_room_number` (`m_room_number`)
) COMMENT='会议文件表';

SET FOREIGN_KEY_CHECKS = 1;

select 
    t1.room_number, t1.topic, t1.mode, t1.distance, t1.status, t1.location, 
    t2.user_name, t2.status, t2.update_time, t2.join_time, t2.exit_time 
from t_meeting t1 left join t_meeting_participant t2 on t1.room_number = t2.m_room_number
where t2.user_id = '003';

EXPLAIN UPDATE `t_meeting_participant` SET `status`=2,`update_time`='2025-05-12 12:49:00' WHERE `m_room_number`='889308385' AND `user_id`='004';
