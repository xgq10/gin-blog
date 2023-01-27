CREATE TABLE `blog_article`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id`      int(10) unsigned    DEFAULT '0' COMMENT '标签ID',
    `title`       varchar(100)        DEFAULT '' COMMENT '文章标题',
    `desc`        varchar(255)        DEFAULT '' COMMENT '简述',
    `content`     text,
    `created_on`  int(11)             DEFAULT NULL,
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(255)        DEFAULT '' COMMENT '修改人',
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用1为启用',
    PRIMARY KEY (`id`)
) ENGINE = innoDB
  DEFAULT CHARSET = utf8 COMMENT = '文章管理';