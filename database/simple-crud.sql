-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 26 Feb 2023 pada 17.45
-- Versi server: 10.4.24-MariaDB
-- Versi PHP: 8.0.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `simple-crud`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `id` bigint(20) NOT NULL,
  `name` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `price` double NOT NULL,
  `quantity` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`id`, `name`, `description`, `price`, `quantity`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, '', '', 0, 0, '2023-02-23 22:16:13.263', '2023-02-23 22:16:13.263', '2023-02-23 22:17:09.204'),
(3, 'ets', 'tteswttset', 1000, 4, '2023-02-23 22:20:06.204', '2023-02-23 22:20:06.204', '2023-02-25 21:58:01.784'),
(4, '', '', 0, 0, '2023-02-23 22:49:58.659', '2023-02-23 22:49:58.659', NULL),
(5, '', '', 0, 0, '2023-02-23 22:55:50.907', '2023-02-23 22:55:50.907', NULL),
(6, '', '', 0, 0, '2023-02-23 22:55:56.707', '2023-02-23 22:55:56.707', NULL),
(7, '', '', 0, 0, '2023-02-23 23:14:20.132', '2023-02-23 23:14:20.132', NULL),
(8, '', '', 0, 0, '2023-02-23 23:15:41.487', '2023-02-23 23:15:41.487', NULL),
(9, '', '', 0, 0, '2023-02-25 21:27:27.939', '2023-02-25 21:27:27.939', NULL),
(10, '', '', 0, 0, '2023-02-25 21:27:37.453', '2023-02-25 21:27:37.453', NULL),
(11, '', '', 0, 0, '2023-02-25 21:28:46.017', '2023-02-25 21:28:46.017', NULL),
(12, 'Baju baru', 'alhamdulillah', 0, 0, '2023-02-25 21:48:30.031', '2023-02-25 21:48:30.031', NULL),
(13, 'wkwkkwkw', '', 0, 0, '2023-02-25 21:50:59.472', '2023-02-25 21:50:59.472', NULL),
(14, 'hehehehe', 'alhamdulillah', 10000, 20, '2023-02-25 21:54:17.294', '2023-02-25 22:06:41.402', NULL),
(15, 'baju wkwkkwk', 'asdlaksndalskdnkl', 20000, 10, '2023-02-25 21:57:06.450', '2023-02-25 22:00:07.145', '2023-02-25 22:00:27.063');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_products_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
