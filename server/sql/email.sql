CREATE TABLE `emails` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `email_subject` varchar(64) NOT NULL COMMENT '邮件标题',
    `email_body` text(65535) NOT NULL COMMENT '邮件正文',
    `email_creator` varchar(64) NOT NULL COMMENT '邮件创建者ID',
    `email_receiver` varchar(256) NOT NULL COMMENT '邮件接收者',
    `email_public` varchar(16) NOT NULL COMMENT '邮件是否公开',
    `content_type` varchar(16) NOT NULL COMMENT '邮件正文类型，http还是plain',
    `triggle_mode` varchar(16) NOT NULL COMMENT '触发模式，定向/随机',
    `email_state` int(10) NOT NULL COMMENT '邮件状态。0待发送，1已发送，2发送失败',
    `triggle_time` int(10) unsigned NOT NULL COMMENT '邮件发送触发时间',
    `create_time` int(10) unsigned NOT NULL COMMENT '创建时间',
    `update_time` int(10) unsigned NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `key_info` (`email_subject`, `email_state`, `triggle_time`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

