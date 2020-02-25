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

DROP TABLE IF EXISTS `trip`;
CREATE TABLE `trip` (
  `trip_id` int(11) NOT NULL AUTO_INCREMENT,
  `trip_name` varchar(45) NOT NULL,
  `trip_location` varchar(45) NOT NULL,
  `trip_description` varchar(45) NOT NULL,
  `trip_activity` varchar(45) NOT NULL,
  `trip_price` float(20,2),
  `trip_capacity` int(11) NOT NULL,
  `trip_month` varchar(45) NOT NULL,
  `trip_year` varchar(45) NOT NULL,
  `trip_duration` varchar(48) NOT NULL,
  `trip_type` varchar(45) NOT NULL,
  `traveler_type` varchar(45) NOT NULL,
  `price_visibilty` boolean NOT NULL DEFAULT true,
  `trip_availability` boolean NOT NULL DEFAULT true,
  `trip_status` varchar(45) NOT NULL,
  `tour_guide` int(11) NOT NULL,
  FOREIGN KEY(`tour_guide`) REFERENCES guide(`tour_guide_id`),
  PRIMARY KEY(`trip_id`)
);
