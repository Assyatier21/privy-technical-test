-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 10, 2022 at 05:52 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `technical_privy`
--
CREATE DATABASE IF NOT EXISTS `technical_privy` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `technical_privy`;

-- --------------------------------------------------------

--
-- Table structure for table `privy_cakes`
--

DROP TABLE IF EXISTS `privy_cakes`;
CREATE TABLE `privy_cakes` (
  `id` int(11) NOT NULL,
  `title` text NOT NULL,
  `description` text NOT NULL,
  `rating` float NOT NULL,
  `image` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `privy_cakes`
--

INSERT INTO `privy_cakes` (`id`, `title`, `description`, `rating`, `image`, `created_at`, `updated_at`) VALUES
(1, 'First Cake', 'This is very first cake in this store', 9, 'https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg', '2022-12-08 04:39:09', '2022-12-08 10:40:02'),
(4, 'New Cakes', 'This is red velvet cakes', 8.2, 'https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg', '2022-12-09 20:47:40', '2022-12-09 20:47:40');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `privy_cakes`
--
ALTER TABLE `privy_cakes`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `privy_cakes`
--
ALTER TABLE `privy_cakes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
