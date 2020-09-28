-- MySQL dump 10.13  Distrib 5.7.30, for Linux (x86_64)
--
-- Host: localhost    Database: blog
-- ------------------------------------------------------
-- Server version	5.7.30

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
-- Table structure for table `blog`
--

DROP TABLE IF EXISTS `blog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog` (
  `blog_id` varchar(50) NOT NULL,
  `title` varchar(50) NOT NULL,
  `content` varchar(20000) NOT NULL,
  `image_url` varchar(1000) NOT NULL,
  `published` tinyint(1) NOT NULL,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `type_id` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`blog_id`),
  KEY `blog_type` (`type_id`),
  CONSTRAINT `blog_type` FOREIGN KEY (`type_id`) REFERENCES `type` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog`
--

LOCK TABLES `blog` WRITE;
/*!40000 ALTER TABLE `blog` DISABLE KEYS */;
INSERT INTO `blog` VALUES ('01aa27b8-197c-423b-b18e-5bdf78a963ea','test-blog-asdf','# this is title\n - an item\n - another item','',1,'2020-08-15 14:34:55','2020-08-15 14:34:55','bf8dfd2f-3564-4e69-a86b-79b191498607'),('0cdc499b-acf6-4cce-a810-daceb75d39e8','test db blog','test db blog content','',1,'2020-08-09 19:00:27','2020-08-09 19:00:27','f9eb5c3d-66f2-4168-8b79-f5976d3082f4'),('12ef150c-f770-4576-84aa-a44872680245','test-blog-tag-association','test blog association with tag','',1,'2020-08-18 23:26:33','2020-08-18 23:26:33','bf8dfd2f-3564-4e69-a86b-79b191498607'),('758a6934-1615-40e5-9a12-bb460350b191','test db blog','test db blog content','',1,'2020-08-12 00:18:14','2020-08-12 00:18:14','f9eb5c3d-66f2-4168-8b79-f5976d3082f4'),('84cb2663-f5c2-4ac5-bc2c-319dae1d3c49','test-ui-title','### test content from frontend ui','',1,'2020-09-21 20:54:56','2020-09-24 22:08:32','3270d506-4eca-4c84-b49b-f0eb8afb6b18'),('bf38c537-9063-4d1e-88d2-8b3b8485c351','test db blog','test db blog content','',1,'2020-08-15 14:07:30','2020-08-15 14:07:30','f9eb5c3d-66f2-4168-8b79-f5976d3082f4');
/*!40000 ALTER TABLE `blog` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_tag`
--

DROP TABLE IF EXISTS `blog_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `blog_tag` (
  `id` varchar(50) NOT NULL,
  `blog_id` varchar(50) NOT NULL,
  `tag_id` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `blog_id` (`blog_id`),
  KEY `tag_id` (`tag_id`),
  CONSTRAINT `blog_tag_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `blog` (`blog_id`),
  CONSTRAINT `blog_tag_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_tag`
--

LOCK TABLES `blog_tag` WRITE;
/*!40000 ALTER TABLE `blog_tag` DISABLE KEYS */;
INSERT INTO `blog_tag` VALUES ('5d60450d-3c76-4ace-a0e8-fa44fddcffdf','12ef150c-f770-4576-84aa-a44872680245','23a3efcc-f7ae-4fc8-a882-01f89ce254c1'),('a07ce715-d106-4333-bfa7-32f892874398','12ef150c-f770-4576-84aa-a44872680245','9c472c07-6505-46c0-a213-e04a8925776a'),('a18ebd31-41f5-41c5-84ad-8fab17644d59','84cb2663-f5c2-4ac5-bc2c-319dae1d3c49','9c472c07-6505-46c0-a213-e04a8925776a'),('c7081ff8-a63e-42fd-8648-492638d15f60','84cb2663-f5c2-4ac5-bc2c-319dae1d3c49','23a3efcc-f7ae-4fc8-a882-01f89ce254c1'),('c74207c4-45bd-46a4-a234-1dcdef62321c','01aa27b8-197c-423b-b18e-5bdf78a963ea','23a3efcc-f7ae-4fc8-a882-01f89ce254c1'),('f7789b12-23b8-43e3-9c35-e2c51e1e8226','0cdc499b-acf6-4cce-a810-daceb75d39e8','23a3efcc-f7ae-4fc8-a882-01f89ce254c1');
/*!40000 ALTER TABLE `blog_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `comment_id` varchar(50) NOT NULL,
  `nickname` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `blog_id` varchar(50) NOT NULL,
  `content` varchar(2000) NOT NULL,
  `avatar_url` varchar(1000) DEFAULT NULL,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`comment_id`),
  KEY `blog_id` (`blog_id`),
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`blog_id`) REFERENCES `blog` (`blog_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES ('3a8f0460-938e-41da-b304-da1f2abc0be8','asdf','qwer@asdf.com','01aa27b8-197c-423b-b18e-5bdf78a963ea','test test','','2020-09-12 11:08:26'),('5dd8b371-1c5e-49d1-9abe-a0d87ac40854','nic2','nic@123.com','01aa27b8-197c-423b-b18e-5bdf78a963ea','my thoughts','','2020-09-09 23:02:24'),('c976e1b4-5d52-4037-968b-524473b0588f','123','asd@qew.com','01aa27b8-197c-423b-b18e-5bdf78a963ea','asdfqwer','','2020-09-09 23:02:59'),('e2863b75-aa2d-4cb8-ba02-ea957cead142','test-comment-nickname','123@4567.com','01aa27b8-197c-423b-b18e-5bdf78a963ea','test comment content','','2020-08-19 22:54:49');
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag`
--

DROP TABLE IF EXISTS `tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tag` (
  `tag_id` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag`
--

LOCK TABLES `tag` WRITE;
/*!40000 ALTER TABLE `tag` DISABLE KEYS */;
INSERT INTO `tag` VALUES ('23a3efcc-f7ae-4fc8-a882-01f89ce254c1','test-update-tag'),('5e29ed79-0c2e-4b11-94fa-68fdc9cc79ef','test-tag-add'),('9c472c07-6505-46c0-a213-e04a8925776a','123123');
/*!40000 ALTER TABLE `tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `type`
--

DROP TABLE IF EXISTS `type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `type` (
  `type_id` varchar(50) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `type`
--

LOCK TABLES `type` WRITE;
/*!40000 ALTER TABLE `type` DISABLE KEYS */;
INSERT INTO `type` VALUES ('3270d506-4eca-4c84-b49b-f0eb8afb6b18','test-ui'),('bf8dfd2f-3564-4e69-a86b-79b191498607','test-update-234'),('f9eb5c3d-66f2-4168-8b79-f5976d3082f4','test-type');
/*!40000 ALTER TABLE `type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `user_id` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `avatar_url` varchar(1000) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('98aa7fc9-a5de-4242-9b02-be1e172abc52','test-1','test','test@1234.com','',0,'2020-08-13 23:59:20','2020-08-13 23:59:20'),('acbb1545-a266-4c1b-a961-9e013fe4fcc5','test-1','test','test@1234.com','',2,'2020-08-13 23:41:48','2020-08-13 23:41:48'),('d54af854-9e92-41e9-a7d8-a71db49a1cbd','admin','admin','test@admin.com','',1,'2020-08-12 15:06:40','2020-08-12 15:06:40');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;