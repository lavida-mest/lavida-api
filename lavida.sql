DROP TABLE IF EXISTS `trip_category`;
CREATE TABLE `trip_category` (
  `category_id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`category_id`)
);

DROP TABLE IF EXISTS `guide`;
CREATE TABLE `guide` (
  `tour_guide_id` int(11) NOT NULL AUTO_INCREMENT,
  `tour_guide_name` varchar(45) NOT NULL,
  `tour_guide_email` varchar(45) NOT NULL,
  `tour_guide_number` varchar(45) NOT NULL,
  `category` int(11) NOT NULL,
  FOREIGN KEY(`category`) REFERENCES trip_category(`category_id`),
  PRIMARY KEY(`tour_guide_id`)
);