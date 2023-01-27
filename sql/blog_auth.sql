CREATE TABLE `blog_auth`
(
    `id`       int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) DEFAULT '' COMMENT '账号',
    `password` varchar(50) DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES (null, 'xgq', '123456');