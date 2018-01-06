CREATE TABLE `crawl_entry` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `date` VARCHAR(128) NOT NULL,
  UNIQUE KEY `date` (`date`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `hatena_hotentry` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `crawl_entry_id` INT(11) NOT NULL,
  `title` VARCHAR(255) NOT NULL DEFAULT "",
  `link` VARCHAR(255) NOT NULL DEFAULT "",
  `description` TEXT,
  INDEX `idx_crawl_entry_id` (`crawl_entry_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
