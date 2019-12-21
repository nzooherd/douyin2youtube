CREATE TABLE `tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table upload
# ------------------------------------------------------------

CREATE TABLE `upload` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `upload_id` varchar(64) NOT NULL DEFAULT '' COMMENT '上传ID',
  `upload_result` smallint(6) NOT NULL DEFAULT '1' COMMENT '0:上传成功 1:上传失败',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table upload_info
# ------------------------------------------------------------

CREATE TABLE `upload_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `upload_id` varchar(64) NOT NULL DEFAULT '' COMMENT '上传ID',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '标题',
  `description` text NOT NULL COMMENT '描述',
  `category` int(11) NOT NULL DEFAULT '0' COMMENT '类别',
  `keywords` varchar(128) NOT NULL DEFAULT '' COMMENT '关键词(json)',
  `privacy` varchar(32) NOT NULL DEFAULT '' COMMENT '是否公开',
  `video_path` varchar(128) NOT NULL DEFAULT '' COMMENT '文件路径',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table upload_video_ref
# ------------------------------------------------------------

CREATE TABLE `upload_video_ref` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `upload_id` varchar(64) NOT NULL DEFAULT '' COMMENT '上传id',
  `video_id` varchar(64) NOT NULL DEFAULT '' COMMENT '视频id',
  `order` smallint(6) NOT NULL DEFAULT '0' COMMENT '视频位次',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table video
# ------------------------------------------------------------

CREATE TABLE `video` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `video_id` varchar(64) NOT NULL DEFAULT '' COMMENT '视频ID',
  `author_id` varchar(32) NOT NULL DEFAULT '' COMMENT '用户id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table video_info
# ------------------------------------------------------------

CREATE TABLE `video_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `video_id` varchar(64) NOT NULL DEFAULT '' COMMENT '视频ID',
  `url_list` text NOT NULL COMMENT '视频地址(json)',
  `duration` int(11) NOT NULL DEFAULT '0' COMMENT '持续时间',
  `comment_count` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
  `like_count` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
  `description` text NOT NULL COMMENT '视频描述',
  `has_download` smallint(6) NOT NULL DEFAULT '0' COMMENT '0:下载失败, 1:下载成功',
  `video_path` varchar(128) NOT NULL DEFAULT '' COMMENT '视频路径',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_video` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table video_tag_ref
# ------------------------------------------------------------

CREATE TABLE `video_tag_ref` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `video_id` varchar(64) NOT NULL DEFAULT '0' COMMENT '视频ID',
  `tag_id` int(11) NOT NULL DEFAULT '0' COMMENT '标签ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
