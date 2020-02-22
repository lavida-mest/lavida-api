DROP TABLE IF EXISTS `trip_category`;
CREATE TABLE `trip_category` (
  `category_id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`category_id`)
);