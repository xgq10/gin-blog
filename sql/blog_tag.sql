CREATE table `blog_tag`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(100)        DEFAULT '' COMMENT '标签名称',
    `created_on`  int(10) unsigned    DEFAULT 0 COMMENT '创建时间',
    `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
    `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT = '文章标签管理';