-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 24, 2021 at 08:30 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 7.4.24

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `courier`
--

-- --------------------------------------------------------

--
-- Table structure for table `couriers`
--

CREATE TABLE `couriers` (
  `id` int(11) NOT NULL,
  `logistic_name` varchar(100) NOT NULL,
  `amount` int(11) NOT NULL,
  `destination_name` varchar(100) NOT NULL,
  `origin_name` varchar(100) NOT NULL,
  `duration` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `couriers`
--

INSERT INTO `couriers` (`id`, `logistic_name`, `amount`, `destination_name`, `origin_name`, `duration`) VALUES
(3, 'JNE', 10000, 'Jakarta', 'Bandung', '1-2'),
(4, 'JNT', 12000, 'Jakarta', 'Bandung', '1-2'),
(5, 'JNE', 20000, 'Jakarta', 'Surabaya', '3-5'),
(6, 'JNE', 5000, 'Jakarta', 'Bogor', '1-2'),
(7, 'JNE', 12000, 'Surabaya', 'Jogja', '1-2'),
(8, 'JNE', 166000, 'Bogor', 'Malang', '2-4'),
(9, 'SiCepat', 11000, 'Surabaya', 'Jogja', '1-2'),
(10, 'SiCepat', 14000, 'Bogor', 'Malang', '2-3'),
(11, 'SiCepat', 10000, 'Jakarta', 'Bandung', '1-2'),
(12, 'JNT', 12000, 'Jakarta', 'Bandung', '1-2'),
(13, 'JNE', 20000, 'Surabaya', 'Semarang', '2-4'),
(14, 'JNT', 5000, 'Surabaya', 'Bogor', '1-2'),
(15, 'JNE', 12000, 'Surabaya', 'Jogja', '1-2'),
(16, 'JNT', 166000, 'Bogor', 'Malang', '2-4'),
(17, 'JNT', 11000, 'Surabaya', 'Jogja', '1-2'),
(18, 'SiCepat', 14000, 'Malang', 'Bogor', '2-3');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `uuid` varchar(36) NOT NULL,
  `msisdn` varchar(15) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`uuid`, `msisdn`, `username`, `password`) VALUES
('02c4c75d-344e-11ec-8f36-00155d074c9e', '628916235123', 'gumiya', '$2a$14$aaDlSvheFUfpvlIYpPqF3.93VqjawJzOE79qeIm3KJEmI4Nwx3ENm'),
('589e954b-348f-11ec-a51a-00155d074c9e', '628122150122412', 'NewUsername', '$2a$14$Ch6Ec7Qr6Dt5A3E.bEe05umKwHH1KbBA87JgvM0AAihlgkmZ0Sp5C'),
('b94b0606-3439-11ec-8f36-00155d074c9e', '6281221501224', 'nathieqs', '$2a$14$v.aBDMbs7FVROsb1caE2BOwYkwMp6NRRJoW2Ki6jNBHFaK10.Rap.');

--
-- Triggers `users`
--
DELIMITER $$
CREATE TRIGGER `uuid_generator` BEFORE INSERT ON `users` FOR EACH ROW SET new.uuid = uuid()
$$
DELIMITER ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `couriers`
--
ALTER TABLE `couriers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`uuid`),
  ADD UNIQUE KEY `msisdn` (`msisdn`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `couriers`
--
ALTER TABLE `couriers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
