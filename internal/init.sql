DROP DATABASE IF EXISTS `OnlineMeeting`;
CREATE DATABASE `OnlineMeeting` DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

USE `OnlineMeeting`;

SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `t_meeting`;
CREATE TABLE `t_meeting` (
    `id` bigint(20) AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `room_number` char(9) NOT NULL COMMENT '会议ID',
    `topic` varchar(30) NOT NULL COMMENT '会议主题',
    `mode` tinyint(1) NOT NULL COMMENT '会议模式', -- 1:漫游模式, 2:会议室模式, 3:虚拟人模式
    `distance` int COMMENT '电子围栏半径', -- 漫游模式下，多少米以内的用户自动加入会议
    `type` tinyint(1) NOT NULL COMMENT '会议类型', -- 1:即时会议, 2:预约会议
    `status` tinyint(1) NOT NULL COMMENT '会议状态', -- 1:待开始, 2:进行中, 3:已结束, 4:已取消
    `location` varchar(100) NOT NULL COMMENT '会议地点',
    `creator_id` varchar(36) NOT NULL COMMENT '会议创建人ID',
    `description` varchar(100) NOT NULL COMMENT '会议描述信息',
    `create_time` datetime NOT NULL DEFAULT NOW() COMMENT '会议创建时间',
    `start_time` datetime NOT NULL COMMENT '会议开始时间', -- 预约会议最长3个月
    `end_time` datetime COMMENT '会议结束时间', -- 根据实际结束时间设置
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_room_number` (`room_number`),
    KEY `idx_topic` (`topic`)
) COMMENT='会议表';

DROP TABLE IF EXISTS `t_meeting_participant`;
CREATE TABLE `t_meeting_participant` (
    `id` bigint(20) AUTO_INCREMENT NOT NULL COMMENT '主键ID',
    `m_room_number` char(9) NOT NULL COMMENT '会议ID, 关联t_meeting表room_number字段',
    `user_id` varchar(255) NOT NULL COMMENT '被邀请用户的ID',
    `user_name` varchar(255) NOT NULL COMMENT '被邀请的用户名称',
    `role` int unsigned NOT NULL COMMENT '角色, 1:主持人, 2:管理员, 3:观众',
    `status` tinyint(1) NOT NULL COMMENT '邀请状态, 1:待处理, 2:已接受, 3:已拒绝',
    `update_time` datetime COMMENT '更新时间',
    `join_time` datetime COMMENT '加入时间',
    `exit_time` datetime COMMENT '退出时间',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`m_room_number`) REFERENCES `t_meeting`(`room_number`) ON DELETE CASCADE,
    UNIQUE KEY `idx_room_number_user_id` (`m_room_number`, `user_id`),
    KEY `idx_user_id` (`user_id`)
)COMMENT='会议参与者表';

SET FOREIGN_KEY_CHECKS = 1;