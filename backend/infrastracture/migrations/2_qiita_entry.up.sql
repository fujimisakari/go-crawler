CREATE TABLE `qiita_entry` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `crawl_entry_id` INT(11) NOT NULL,
  `user` VARCHAR(255) NOT NULL DEFAULT "",
  `title` VARCHAR(255) NOT NULL DEFAULT "",
  `link` VARCHAR(255) NOT NULL DEFAULT "",
  `posted_at` datetime DEFAULT NULL,
  INDEX idx_crawl_entry_id (`crawl_entry_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
