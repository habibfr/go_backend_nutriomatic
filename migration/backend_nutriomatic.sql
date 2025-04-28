-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 28 Apr 2025 pada 05.39
-- Versi server: 10.4.25-MariaDB
-- Versi PHP: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `backend_nutriomatic`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `activity_levels`
--

CREATE TABLE `activity_levels` (
  `al_id` varchar(36) NOT NULL,
  `al_type` bigint(20) DEFAULT NULL,
  `al_desc` longtext DEFAULT NULL,
  `al_value` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `activity_levels`
--

INSERT INTO `activity_levels` (`al_id`, `al_type`, `al_desc`, `al_value`, `created_at`, `updated_at`) VALUES
('2ca761d7-3ef6-4892-8a0b-2a81e7c425f6', 2, 'Lightly active', 1.375, '2025-04-24 18:42:17.616', '2025-04-24 18:42:17.616'),
('3c12a580-f13d-42cd-a323-b2715cea47d1', 1, 'Sedentary', 1.2, '2025-04-24 18:42:17.603', '2025-04-24 18:42:17.603'),
('6610e79b-c06b-4c3f-986a-e5f75a2d444d', 5, 'Super active', 1.9, '2025-04-24 18:42:17.626', '2025-04-24 18:42:17.626'),
('7c9f3b04-6aec-43f4-88b0-ef335b2a1c5e', 3, 'Moderately active', 1.55, '2025-04-24 18:42:17.619', '2025-04-24 18:42:17.619'),
('c2ab5f2e-a836-4f43-bcf4-3d2b1e137d54', 4, 'Very active', 1.725, '2025-04-24 18:42:17.623', '2025-04-24 18:42:17.623');

-- --------------------------------------------------------

--
-- Struktur dari tabel `health_goals`
--

CREATE TABLE `health_goals` (
  `hg_id` varchar(36) NOT NULL,
  `hg_type` bigint(20) DEFAULT NULL,
  `hg_desc` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `health_goals`
--

INSERT INTO `health_goals` (`hg_id`, `hg_type`, `hg_desc`, `created_at`, `updated_at`) VALUES
('2bf9677c-eadf-481b-be19-9396a26bfb18', 2, 'Maintain Weight', '2025-04-24 18:42:17.630', '2025-04-24 18:42:17.630'),
('6c4418bd-1bac-40bf-a30d-01ea0681071a', 1, 'Lose Weight', '2025-04-24 18:42:17.628', '2025-04-24 18:42:17.628'),
('7002d438-dc87-4f43-b3f9-13bb47856922', 3, 'Gain Weight', '2025-04-24 18:42:17.632', '2025-04-24 18:42:17.632');

-- --------------------------------------------------------

--
-- Struktur dari tabel `nutrition_info`
--

CREATE TABLE `nutrition_info` (
  `ni_id` varchar(36) NOT NULL,
  `ni_type` varchar(255) DEFAULT NULL,
  `ni_text` varchar(500) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `payments`
--

CREATE TABLE `payments` (
  `payment_id` varchar(255) NOT NULL,
  `payment_method` varchar(255) NOT NULL,
  `payment_link` varchar(512) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `payments`
--

INSERT INTO `payments` (`payment_id`, `payment_method`, `payment_link`, `created_at`, `updated_at`) VALUES
('bde8e5bf-0dab-40ab-821d-153500fe5283', 'default', 'www.qris.com', '2025-04-24 17:15:38', '2025-04-24 17:15:38');

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `product_id` varchar(36) NOT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `product_price` decimal(10,2) DEFAULT NULL,
  `product_desc` text DEFAULT NULL,
  `product_isshow` tinyint(1) DEFAULT NULL,
  `product_lemaktotal` decimal(10,2) DEFAULT NULL,
  `product_protein` decimal(10,2) DEFAULT NULL,
  `product_karbohidrat` decimal(10,2) DEFAULT NULL,
  `product_garam` decimal(10,2) DEFAULT NULL,
  `product_servingsize` decimal(10,2) DEFAULT NULL,
  `product_picture` varchar(255) DEFAULT NULL,
  `product_expshow` timestamp NULL DEFAULT NULL,
  `createdat` timestamp NULL DEFAULT current_timestamp(),
  `updatedat` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `store_id` varchar(36) DEFAULT NULL,
  `pt_id` varchar(36) DEFAULT NULL,
  `product_grade` varchar(10) DEFAULT NULL,
  `product_energi` double DEFAULT 0,
  `product_gula` double DEFAULT 0,
  `product_saturatedfat` double DEFAULT 0,
  `product_fiber` double DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`product_id`, `product_name`, `product_price`, `product_desc`, `product_isshow`, `product_lemaktotal`, `product_protein`, `product_karbohidrat`, `product_garam`, `product_servingsize`, `product_picture`, `product_expshow`, `createdat`, `updatedat`, `store_id`, `pt_id`, `product_grade`, `product_energi`, `product_gula`, `product_saturatedfat`, `product_fiber`) VALUES
('021590d5-18b4-495e-b988-6c31facc3516', 'test', '54.00', 'test', 1, '4.00', '8.00', '6.00', '0.01', '5.00', 'http://localhost:8080/uploads/88fa0587-4bff-4640-a2c3-c13b8c3eafdf.jpg', '2025-05-25 07:03:37', '2025-04-25 06:51:23', '2025-04-25 07:03:37', '439f9a2b-736d-40a5-ba35-d2aa84273d9f', '6bb279ce-df46-40d5-925c-5554dc64e047', 'B', 5, 9, 7, 8),
('7aff29c2-129e-437e-be47-782e082f5f5e', 'test', '64.00', 'test', 1, '9.00', '4.00', '3.00', '0.00', '6.00', 'http://localhost:8080/uploads/66a6e222-7b48-4a9d-bf11-14b191db2818.jpg', '2025-05-25 09:14:28', '2025-04-25 09:13:36', '2025-04-25 09:14:28', 'a5ba4d5a-8df3-4bf9-88c8-fde38484064e', '3eae74c2-883a-4ded-8ea7-af1e2103d272', 'B', 6, 3, 8, 2),
('940d1ed4-8355-4018-9090-8596c343360d', 'test', '5.00', 'e\nd\n', 3, '5.00', '3.00', '6.00', '0.01', '4.00', 'http://localhost:8080/uploads/8f8079b7-4ef3-43b6-867e-c9ced16a23b9.jpg', '2025-04-25 09:15:19', '2025-04-25 09:15:19', '2025-04-25 09:18:27', 'a5ba4d5a-8df3-4bf9-88c8-fde38484064e', '772f48cb-ba60-42e2-9d26-9a7f54e4065c', 'B', 8, 7, 4, 8);

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_types`
--

CREATE TABLE `product_types` (
  `pt_id` varchar(36) NOT NULL,
  `pt_name` longtext DEFAULT NULL,
  `pt_type` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `product_types`
--

INSERT INTO `product_types` (`pt_id`, `pt_name`, `pt_type`, `created_at`, `updated_at`) VALUES
('2b1a3d86-446e-439d-873b-e6619f9c9ca8', 'Juice', 302, '2025-04-24 17:08:15.087', '2025-04-24 17:08:15.087'),
('3eae74c2-883a-4ded-8ea7-af1e2103d272', 'Bread', 204, '2025-04-24 17:08:15.078', '2025-04-24 17:08:15.078'),
('4121c7f0-ed47-4cd5-ad94-cae50e719d0e', 'Ice Cream', 203, '2025-04-24 17:08:15.073', '2025-04-24 17:08:15.073'),
('6bb279ce-df46-40d5-925c-5554dc64e047', 'Cheese', 205, '2025-04-24 17:08:15.081', '2025-04-24 17:08:15.081'),
('772f48cb-ba60-42e2-9d26-9a7f54e4065c', 'Snacks', 201, '2025-04-24 17:08:15.067', '2025-04-24 17:08:15.067'),
('7d9d219d-c334-40bc-be0a-bf9be8727da5', 'Water', 301, '2025-04-24 17:08:15.085', '2025-04-24 17:08:15.085'),
('8980770e-822b-46a8-91fa-12e635c6ed11', 'Tea', 305, '2025-04-24 17:08:15.095', '2025-04-24 17:08:15.095'),
('a1f1d096-7794-4859-aabc-fc058d00a5a2', 'Milk', 303, '2025-04-24 17:08:15.090', '2025-04-24 17:08:15.090'),
('bc1ef188-722d-4083-af50-b157c5c9c233', 'Chocolate Bar', 202, '2025-04-24 17:08:15.071', '2025-04-24 17:08:15.071'),
('dd8791e0-8fd2-444c-a847-cf2327b7aefb', 'Frozen Food', 206, '2025-04-24 17:08:15.083', '2025-04-24 17:08:15.083'),
('ff98f0ec-731c-42e3-ae00-ea1c7c068c7a', 'Coffee', 304, '2025-04-24 17:08:15.092', '2025-04-24 17:08:15.092');

-- --------------------------------------------------------

--
-- Struktur dari tabel `scanned_nutritions`
--

CREATE TABLE `scanned_nutritions` (
  `sn_id` varchar(100) NOT NULL,
  `sn_productname` varchar(255) DEFAULT NULL,
  `sn_producttype` varchar(255) DEFAULT NULL,
  `sn_info` text DEFAULT NULL,
  `sn_picture` text DEFAULT NULL,
  `sn_energy` double DEFAULT NULL,
  `sn_protein` double DEFAULT NULL,
  `sn_fat` double DEFAULT NULL,
  `sn_carbohydrate` double DEFAULT NULL,
  `sn_sugar` double DEFAULT NULL,
  `sn_salt` double DEFAULT NULL,
  `sn_grade` varchar(10) DEFAULT NULL,
  `sn_saturatedfat` double DEFAULT NULL,
  `sn_fiber` double DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `user_id` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `scanned_nutritions`
--

INSERT INTO `scanned_nutritions` (`sn_id`, `sn_productname`, `sn_producttype`, `sn_info`, `sn_picture`, `sn_energy`, `sn_protein`, `sn_fat`, `sn_carbohydrate`, `sn_sugar`, `sn_salt`, `sn_grade`, `sn_saturatedfat`, `sn_fiber`, `created_at`, `updated_at`, `user_id`) VALUES
('659b6705-e57d-4fcc-a49a-03b2a3674c61', 'scan_1745572347662', '', '', 'http://localhost:8080/uploads/787611eb-b88f-4779-8d7e-49bb3efea5fb.jpg', 1155.71, 8, 9.4, 3, 2.1, 0.041, 'C', 0, 10, '2025-04-25 16:12:38', '2025-04-25 16:12:38', 'e198a10b-eab8-42e8-ad70-96c8a4d8e26b'),
('b71dfd88-0006-4763-b9e3-fe7590ce4be2', 'scan_1745570110868', '', '', 'http://localhost:8080/uploads/b6c5ba3b-3935-46ee-b1c9-953d4a7614f5.jpg', 1155.71, 6.1, 9.4, 3, 2.1, 0.24, 'C', 0, 10, '2025-04-25 15:35:18', '2025-04-25 15:35:18', 'bd43ad57-5b8c-45a1-b938-a0094a0303e3'),
('c2c8b14f-24eb-41c0-bfd3-c1f5292d8d4d', 'scan_1745569899961', '', '', 'http://localhost:8080/uploads/1555a0d0-e2d6-4c41-b412-577dca13da5b.jpg', 1155.71, 6.1, 7.8, 17.86, 4.2, 0.24, 'C', 0, 0, '2025-04-25 15:31:41', '2025-04-25 15:31:41', 'bd43ad57-5b8c-45a1-b938-a0094a0303e3');

-- --------------------------------------------------------

--
-- Struktur dari tabel `stores`
--

CREATE TABLE `stores` (
  `store_id` varchar(36) NOT NULL,
  `store_name` longtext DEFAULT NULL,
  `store_username` longtext DEFAULT NULL,
  `store_address` longtext DEFAULT NULL,
  `store_contact` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` varchar(36) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `stores`
--

INSERT INTO `stores` (`store_id`, `store_name`, `store_username`, `store_address`, `store_contact`, `created_at`, `updated_at`, `user_id`) VALUES
('439f9a2b-736d-40a5-ba35-d2aa84273d9f', 'habib1 store', 'habib1', 'surabaya', '1234567890', '2025-04-25 13:49:57.332', '2025-04-25 13:49:57.332', 'bd43ad57-5b8c-45a1-b938-a0094a0303e3'),
('a5ba4d5a-8df3-4bf9-88c8-fde38484064e', 'habib2', 'habib2', 'sby', '123458976', '2025-04-25 16:12:55.896', '2025-04-25 16:12:55.896', 'e198a10b-eab8-42e8-ad70-96c8a4d8e26b');

-- --------------------------------------------------------

--
-- Struktur dari tabel `tokens`
--

CREATE TABLE `tokens` (
  `user_id` varchar(36) NOT NULL,
  `token` varchar(255) NOT NULL,
  `expires_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `tokens`
--

INSERT INTO `tokens` (`user_id`, `token`, `expires_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
('8b750fd4-766a-42ba-abe0-91946daf8dfd', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTA5NTQsIm5hbWUiOiJhZG1pbiIsInN1YiI6MX0.uu0_OtYyhZu4Nbb5YFMIDQaRFXvuEeqeBB3U_nw8cEM', '2025-04-26 14:02:34', '2025-04-25 14:02:34', '2025-04-25 14:02:34', '2025-04-25 16:14:21'),
('bd43ad57-5b8c-45a1-b938-a0094a0303e3', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTAxNzYsIm5hbWUiOiJoYWJpYjEiLCJzdWIiOjF9.rSeHoNQwcDXSd8bKf8w2LO27AuUJKEBIpuwoxbt6lm4', '2025-04-26 13:49:36', '2025-04-25 13:49:36', '2025-04-25 13:49:36', '2025-04-25 14:04:00'),
('bd43ad57-5b8c-45a1-b938-a0094a0303e3', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTEwNDAsIm5hbWUiOiJoYWJpYjEiLCJzdWIiOjF9._IWKyMJdBytBlXviHfnU4dLbMg9-_HuBO6GqCqcVJC8', '2025-04-26 14:04:00', '2025-04-25 14:04:00', '2025-04-25 14:04:00', NULL),
('8b750fd4-766a-42ba-abe0-91946daf8dfd', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTg4NjEsIm5hbWUiOiJhZG1pbiIsInN1YiI6MX0.llkGK5hwsj8GBMxQnOtNJgPBGkUrDPaqLSSKy9d4Fog', '2025-04-26 16:14:21', '2025-04-25 16:14:21', '2025-04-25 16:14:21', '2025-04-25 16:16:11'),
('e198a10b-eab8-42e8-ad70-96c8a4d8e26b', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTg4ODgsIm5hbWUiOiJoYWJpYjIiLCJzdWIiOjF9.6ixWWPWoAoWDG6qmOFhBpxzvWFErQTrnIGQ4qxmFfrE', '2025-04-26 16:14:48', '2025-04-25 16:14:48', '2025-04-25 16:14:48', '2025-04-25 16:18:45'),
('8b750fd4-766a-42ba-abe0-91946daf8dfd', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTg5NzEsIm5hbWUiOiJhZG1pbiIsInN1YiI6MX0.YyAdDB8YsIaL3gnMFDnhsduT-68PuJqbsM7bYXdY9FY', '2025-04-26 16:16:11', '2025-04-25 16:16:11', '2025-04-25 16:16:11', NULL),
('c8c2b21e-fc46-42a3-84aa-3f22e61cc80a', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTgwODgsIm5hbWUiOiJoYWJpYjMiLCJzdWIiOjF9.BU5ck-rqwq2QIMhLRQ61wBFuSWuTxjE3pxpzQfm3zhg', '2025-04-26 16:01:28', '2025-04-25 16:01:28', '2025-04-25 16:01:28', NULL),
('e198a10b-eab8-42e8-ad70-96c8a4d8e26b', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTkxMjUsIm5hbWUiOiJoYWJpYjIiLCJzdWIiOjF9.wctqcMdcoW2ulG-Dx-_OLdvMFLuBjzDcB6SGKNy7PS8', '2025-04-26 16:18:45', '2025-04-25 16:18:45', '2025-04-25 16:18:45', NULL),
('e198a10b-eab8-42e8-ad70-96c8a4d8e26b', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU2NTY2NjMsIm5hbWUiOiJoYWJpYjIiLCJzdWIiOjF9.7blwWIdJHgfQQ0LU4YBEi4D6F-03FaJrXiMUKLNfBaw', '2025-04-26 15:37:43', '2025-04-25 15:37:43', '2025-04-25 15:37:43', '2025-04-25 16:14:48');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `tsc_id` varchar(255) NOT NULL,
  `tsc_price` double NOT NULL,
  `tsc_virtualaccount` varchar(255) DEFAULT NULL,
  `tsc_start` datetime NOT NULL,
  `tsc_end` datetime NOT NULL,
  `tsc_status` varchar(100) DEFAULT NULL,
  `tsc_bukti` text DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `product_id` varchar(255) DEFAULT NULL,
  `store_id` varchar(255) DEFAULT NULL,
  `payment_id` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`tsc_id`, `tsc_price`, `tsc_virtualaccount`, `tsc_start`, `tsc_end`, `tsc_status`, `tsc_bukti`, `created_at`, `updated_at`, `product_id`, `store_id`, `payment_id`) VALUES
('144a6b7f-afc1-461b-be91-e93248eea9f4', 5000, '', '2025-04-25 16:15:24', '2025-04-26 16:15:24', 'declined', 'http://localhost:8080/uploads/1b59e5ed-f94d-4881-97f9-b11af686909b.jpg', '2025-04-25 16:15:24', '2025-04-25 16:18:27', '940d1ed4-8355-4018-9090-8596c343360d', 'a5ba4d5a-8df3-4bf9-88c8-fde38484064e', 'bde8e5bf-0dab-40ab-821d-153500fe5283'),
('20480d0c-fcdf-474a-90fa-452424a6390b', 5000, '', '2025-04-25 16:13:40', '2025-04-26 16:13:40', 'accepted', '', '2025-04-25 16:13:40', '2025-04-25 16:14:28', '7aff29c2-129e-437e-be47-782e082f5f5e', 'a5ba4d5a-8df3-4bf9-88c8-fde38484064e', 'bde8e5bf-0dab-40ab-821d-153500fe5283'),
('2455da2c-52f6-4f68-9d2d-69efc75e68b6', 5000, '', '2025-04-25 14:00:50', '2025-04-26 14:00:50', 'accepted', '', '2025-04-25 14:00:50', '2025-04-25 14:03:37', '021590d5-18b4-495e-b988-6c31facc3516', '439f9a2b-736d-40a5-ba35-d2aa84273d9f', 'bde8e5bf-0dab-40ab-821d-153500fe5283');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `username` varchar(191) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `role` longtext DEFAULT NULL,
  `gender` bigint(20) DEFAULT NULL,
  `telp` longtext DEFAULT NULL,
  `profpic` longtext DEFAULT NULL,
  `birthdate` longtext DEFAULT NULL,
  `place` longtext DEFAULT NULL,
  `height` double DEFAULT NULL,
  `weight` double DEFAULT NULL,
  `weight_goal` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `hg_id` varchar(36) DEFAULT NULL,
  `al_id` varchar(36) DEFAULT NULL,
  `calories` double DEFAULT NULL,
  `classification` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `username`, `name`, `email`, `password`, `role`, `gender`, `telp`, `profpic`, `birthdate`, `place`, `height`, `weight`, `weight_goal`, `created_at`, `updated_at`, `hg_id`, `al_id`, `calories`, `classification`) VALUES
('23a32ea1-1ea4-417a-90de-aaf603c7272a', '', 'budi1', 'budi1@example.com', '$2a$10$p4GIC0y/OgoB9y0HzZOBP.QJw7OUXdn8dZHDki48gQxN4613xVKOC', '', 0, '', '', '', '', 0, 0, 0, '2025-04-24 18:46:51.543', '2025-04-24 18:46:51.543', '6c4418bd-1bac-40bf-a30d-01ea0681071a', '3c12a580-f13d-42cd-a323-b2715cea47d1', 0, ''),
('8b750fd4-766a-42ba-abe0-91946daf8dfd', '', 'admin', 'admin@admin.com', '$2a$10$BG.BtDcV7phkZDF/MQuAQOG/Ia2wyXYz47sY0irmdJYdUb08LphTS', 'admin', 0, '', '', '', '', 0, 0, 0, '2025-04-25 09:24:23.718', '2025-04-25 09:24:23.718', '6c4418bd-1bac-40bf-a30d-01ea0681071a', '3c12a580-f13d-42cd-a323-b2715cea47d1', 0, ''),
('bd43ad57-5b8c-45a1-b938-a0094a0303e3', '', 'habib1', 'habib1@gmail.com', '$2a$10$MUCgHwtaxTK4lHfh8.cGn.mGhSKYgrzwThknq/eZEkDso3NqfI/ei', '', 0, '', '', '', '', 0, 0, 0, '2025-04-25 13:49:23.014', '2025-04-25 13:49:23.014', '6c4418bd-1bac-40bf-a30d-01ea0681071a', '3c12a580-f13d-42cd-a323-b2715cea47d1', 0, ''),
('c8c2b21e-fc46-42a3-84aa-3f22e61cc80a', '', 'Adam', 'habib3@gmail.com', '$2a$10$.otA.ahQd1ISv2ZewMz5JOAWekM/DKIJgcWxHVwedNlhhYBjYJY8q', '', 2, '082182138121', '', '2000-01-01', '', 0, 0, 0, '2025-04-25 16:01:04.350', '2025-04-25 16:10:14.050', '2bf9677c-eadf-481b-be19-9396a26bfb18', '7c9f3b04-6aec-43f4-88b0-ef335b2a1c5e', 0, ''),
('e198a10b-eab8-42e8-ad70-96c8a4d8e26b', '', 'habib2', 'habib2@gmail.com', '$2a$10$JbApzhkzoDr1FCOZN.TDUepGy4ktKa1vxSDCVmQJQ8umbUj8ylrTm', '', 0, '1234585679', 'http://localhost:8080/uploads/3b6ffd5b-0592-4b0a-8aa9-a4e01c7092f9.jpg', '2000-01-01', '', 180, 50, 70, '2025-04-25 15:37:33.372', '2025-04-25 16:19:06.860', '7002d438-dc87-4f43-b3f9-13bb47856922', '7c9f3b04-6aec-43f4-88b0-ef335b2a1c5e', 2464, 'Overweight_Level_I');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `activity_levels`
--
ALTER TABLE `activity_levels`
  ADD PRIMARY KEY (`al_id`);

--
-- Indeks untuk tabel `health_goals`
--
ALTER TABLE `health_goals`
  ADD PRIMARY KEY (`hg_id`);

--
-- Indeks untuk tabel `nutrition_info`
--
ALTER TABLE `nutrition_info`
  ADD PRIMARY KEY (`ni_id`);

--
-- Indeks untuk tabel `payments`
--
ALTER TABLE `payments`
  ADD PRIMARY KEY (`payment_id`);

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`product_id`),
  ADD KEY `store_id` (`store_id`),
  ADD KEY `pt_id` (`pt_id`);

--
-- Indeks untuk tabel `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`pt_id`);

--
-- Indeks untuk tabel `scanned_nutritions`
--
ALTER TABLE `scanned_nutritions`
  ADD PRIMARY KEY (`sn_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `stores`
--
ALTER TABLE `stores`
  ADD PRIMARY KEY (`store_id`),
  ADD KEY `fk_stores_user` (`user_id`);

--
-- Indeks untuk tabel `tokens`
--
ALTER TABLE `tokens`
  ADD UNIQUE KEY `token` (`token`),
  ADD UNIQUE KEY `idx_tokens_deleted_at` (`deleted_at`) USING BTREE,
  ADD KEY `unique_user_id` (`user_id`) USING BTREE;

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`tsc_id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `store_id` (`store_id`),
  ADD KEY `payment_id` (`payment_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_email` (`email`),
  ADD KEY `FK_users_health_goals` (`hg_id`),
  ADD KEY `FK_users_activity_levels` (`al_id`);

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`store_id`) REFERENCES `stores` (`store_id`),
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`pt_id`) REFERENCES `product_types` (`pt_id`);

--
-- Ketidakleluasaan untuk tabel `scanned_nutritions`
--
ALTER TABLE `scanned_nutritions`
  ADD CONSTRAINT `scanned_nutritions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Ketidakleluasaan untuk tabel `stores`
--
ALTER TABLE `stores`
  ADD CONSTRAINT `fk_stores_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Ketidakleluasaan untuk tabel `tokens`
--
ALTER TABLE `tokens`
  ADD CONSTRAINT `fk_tokens_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Ketidakleluasaan untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`store_id`) REFERENCES `stores` (`store_id`),
  ADD CONSTRAINT `transactions_ibfk_3` FOREIGN KEY (`payment_id`) REFERENCES `payments` (`payment_id`);

--
-- Ketidakleluasaan untuk tabel `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `FK_users_activity_levels` FOREIGN KEY (`al_id`) REFERENCES `activity_levels` (`al_id`),
  ADD CONSTRAINT `FK_users_health_goals` FOREIGN KEY (`hg_id`) REFERENCES `health_goals` (`hg_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
