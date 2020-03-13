-- MySQL dump 10.13  Distrib 5.7.29, for Linux (x86_64)
--
-- Host: localhost    Database: lavida
-- ------------------------------------------------------
-- Server version	5.7.29-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `guide`
--

DROP TABLE IF EXISTS `guide`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `guide` (
  `tour_guide_id` int(11) NOT NULL AUTO_INCREMENT,
  `tour_guide_name` varchar(45) NOT NULL,
  `tour_guide_email` varchar(45) NOT NULL,
  `tour_guide_number` varchar(45) NOT NULL,
  `category_id` int(11) NOT NULL,
  PRIMARY KEY (`tour_guide_id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `guide_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `trip_category` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `guide`
--

LOCK TABLES `guide` WRITE;
/*!40000 ALTER TABLE `guide` DISABLE KEYS */;
INSERT INTO `guide` VALUES (1,'Taste Makers','tmakers@tme.gh','2334243567678',1),(2,'Taste Makers','tmakers@tme.gh','2334243567678',2),(3,'Kofi Adventures','kofi@travelers.com','2334243567678',5),(4,'Kofi Adventures','kofi@travelers.com','2334243567678',6);
/*!40000 ALTER TABLE `guide` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `trip`
--

DROP TABLE IF EXISTS `trip`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trip` (
  `trip_id` int(11) NOT NULL AUTO_INCREMENT,
  `trip_name` varchar(45) NOT NULL,
  `trip_location` varchar(45) NOT NULL,
  `trip_description` varchar(45) NOT NULL,
  `trip_activity` varchar(45) NOT NULL,
  `trip_price` float(20,2) DEFAULT NULL,
  `trip_capacity` int(11) NOT NULL,
  `trip_month` varchar(45) NOT NULL,
  `trip_year` varchar(45) NOT NULL,
  `trip_duration` varchar(48) NOT NULL,
  `trip_type` varchar(45) NOT NULL,
  `traveler_type` varchar(45) NOT NULL,
  `price_visibilty` tinyint(1) NOT NULL DEFAULT '1',
  `trip_availability` tinyint(1) NOT NULL DEFAULT '1',
  `trip_status` varchar(45) NOT NULL,
  `tour_guide` int(11) NOT NULL,
  PRIMARY KEY (`trip_id`),
  KEY `tour_guide` (`tour_guide`),
  CONSTRAINT `trip_ibfk_1` FOREIGN KEY (`tour_guide`) REFERENCES `guide` (`tour_guide_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trip`
--

LOCK TABLES `trip` WRITE;
/*!40000 ALTER TABLE `trip` DISABLE KEYS */;
INSERT INTO `trip` VALUES (1,'Mombasa gateaway','Mombasa','Come enjoy hiking at the foot of mt Kenya','trekking hiking swimming',35.50,12,'January','2023','two week','Beach','group',0,0,'available',1),(2,'Busia Weekend Escape','Ghana','Come enjoy hiking at the foot of mt Kenya','Walking trekking  Nature observation',35.50,12,'January','2023','two week','Beach','As a group',0,0,'available',3),(3,'Bonfire Getaway','Ghana Ada','Come enjoy hiking at the foot of mt Kenya','Seaside Beach Swimming',135.50,12,'May','2020','one week','Beach','As a group',0,0,'available',2),(4,'Hiking Boti Falls','Central Ghana','Come enjoy hiking at Boti falls','Hiking SightSeeing Nature',46.50,12,'September','2020','3 days','Nature Adventure','As a group',0,0,'available',2);
/*!40000 ALTER TABLE `trip` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `trip_category`
--

DROP TABLE IF EXISTS `trip_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trip_category` (
  `category_id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trip_category`
--

LOCK TABLES `trip_category` WRITE;
/*!40000 ALTER TABLE `trip_category` DISABLE KEYS */;
INSERT INTO `trip_category` VALUES (1,'Multi City tours'),(2,'Offbeat tours'),(3,'Boutique travel'),(4,'Off the beaten track'),(5,'Adventure'),(6,'Highlights');
/*!40000 ALTER TABLE `trip_category` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-13  8:38:17
