-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Waktu pembuatan: 03 Jul 2022 pada 21.18
-- Versi server: 8.0.29-0ubuntu0.20.04.3
-- Versi PHP: 7.4.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `wisata`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `categories`
--

CREATE TABLE `categories` (
  `id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `avatar` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `categories`
--

INSERT INTO `categories` (`id`, `name`, `avatar`) VALUES
(1, 'Pantai', 'images/beach.png'),
(2, 'Taman', 'images/garden.png'),
(3, 'Pendidikan', 'images/reading.png'),
(11, 'Perternakan', 'images/cattle.png'),
(15, 'Pergunungan', 'images/mountain.png');

-- --------------------------------------------------------

--
-- Struktur dari tabel `galleries`
--

CREATE TABLE `galleries` (
  `id` int NOT NULL,
  `tourist_id` int NOT NULL,
  `avatar` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `galleries`
--

INSERT INTO `galleries` (`id`, `tourist_id`, `avatar`) VALUES
(1, 2, 'images/download.jpeg'),
(2, 2, 'images/download(3).jpeg'),
(5, 4, 'images/download (4).jpeg'),
(6, 4, 'images/download (5).jpeg'),
(7, 4, 'images/download (6).jpeg'),
(8, 4, 'images/download (7).jpeg'),
(9, 4, 'images/images.jpeg'),
(10, 4, 'images/images (1).jpeg'),
(11, 2, 'images/142581-0_665_374.jpg'),
(12, 10, 'images/download (8).jpeg'),
(13, 10, 'images/download (9).jpeg'),
(14, 10, 'images/download (10).jpeg'),
(15, 14, 'images/download (8).jpeg'),
(16, 14, 'images/Puncak-Harfat.jpg'),
(17, 14, 'images/download (10).jpeg');

-- --------------------------------------------------------

--
-- Struktur dari tabel `itineraries`
--

CREATE TABLE `itineraries` (
  `id` int NOT NULL,
  `user_id` int NOT NULL,
  `initial_local` varchar(255) NOT NULL,
  `start_day` varchar(255) NOT NULL,
  `end_day` varchar(255) NOT NULL,
  `start_time` varchar(255) NOT NULL,
  `end_time` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `itineraries`
--

INSERT INTO `itineraries` (`id`, `user_id`, `initial_local`, `start_day`, `end_day`, `start_time`, `end_time`) VALUES
(47, 1, 'Rumah Makan Siang Malam', '3/6/2022', '3/6/2022', '15:19', '15: 19'),
(48, 1, 'Rumah Makan Siang Malam 2', '3/6/2022', '3/6/2022', '15:19', '15: 19');

-- --------------------------------------------------------

--
-- Struktur dari tabel `timelines`
--

CREATE TABLE `timelines` (
  `id` int NOT NULL,
  `itinerary_id` int NOT NULL,
  `time` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `latitude` varchar(255) NOT NULL,
  `longitude` varchar(255) NOT NULL,
  `jarak` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `timelines`
--

INSERT INTO `timelines` (`id`, `itinerary_id`, `time`, `title`, `description`, `latitude`, `longitude`, `jarak`) VALUES
(3, 27, 'Day ', '23:45', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(4, 31, 'Day 0', '23:56', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(5, 32, 'Day 0', '23:57', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(6, 33, 'Day 1', '23:59', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(7, 33, 'Day 2', '23:59', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(8, 34, 'Day 1', '0:2', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(9, 34, 'Day 2', '0:2', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(10, 34, 'Day 3', '0:2', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(11, 35, 'Day 1', '0:3', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(12, 35, 'Day 2', '0:3', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(13, 35, 'Day 3', '0:3', 'Pantai Samas Yogykarta', '7.818346', '110.343805', 0),
(14, 36, 'Day 0', '9:1', 'Pantai Hamadi', '2.5264303', '140.6120972', 656149),
(15, 36, 'Day 0', '9:1', 'Puncak Harfat', '-1.97346', '130.465494', 1112361),
(16, 36, 'Day 0', '9:1', 'Pantai Harlem', '-2.4620136', '140.3677778', 103988),
(17, 38, 'Day 0', '9:54', 'Pantai Hamadi', '2.5264303', '140.6120972', 1305523),
(18, 38, 'Day 0', '9:54', 'Pantai Harlem', '-2.4620136', '140.3677778', 1105768),
(19, 38, 'Day 0', '9:54', 'Puncak Harfat', '-1.97346', '130.465494', 158010),
(20, 40, 'Day 0', '9:59', 'Pantai Harlem', '-2.4620136', '140.3677778', 1105768),
(21, 40, 'Day 0', '9:59', 'Pantai Hamadi', '2.5264303', '140.6120972', 1305523),
(22, 40, 'Day 0', '9:59', 'Puncak Harfat', '-1.97346', '130.465494', 158010),
(23, 42, '21758.7 - 21758.75 jam', '10:26', 'Pantai Hamadi', '2.5264303', '140.6120972', 1305523),
(24, 42, '18429.5 - 18429.55 jam', '10:26', 'Pantai Harlem', '-2.4620136', '140.3677778', 1105768),
(25, 42, '2633.5 - 2633.55 jam', '10:26', 'Puncak Harfat', '-1.97346', '130.465494', 158010),
(26, 44, '21758.7 - 21758.75 jam', '10:30', 'Pantai Hamadi', '2.5264303', '140.6120972', 1305523),
(27, 44, '18429.5 - 18429.55 jam', '10:30', 'Pantai Harlem', '-2.4620136', '140.3677778', 1105768),
(28, 44, '2633.5 - 2633.55 jam', '10:30', 'Puncak Harfat', '-1.97346', '130.465494', 158010),
(29, 45, 'Day 0', '2.2 hours', 'Pantai Hamadi', '2.5264303', '140.6120972', 1305528),
(30, 45, 'Day 0', '1.8 hours', 'Pantai Harlem', '-2.4620136', '140.3677778', 1105769),
(31, 45, 'Day 0', '2.6 hours', 'Puncak Harfat', '-1.97346', '130.465494', 158020),
(32, 46, 'Day 0', '184.3 hours', 'Pantai Harlem', '-2.4620136', '140.3677778', 1105769),
(33, 46, 'Day 0', '263.4 hours', 'Puncak Harfat', '-1.97346', '130.465494', 158020),
(34, 47, 'Day 0', '184.3 hours', 'Pantai Harlem', '-2.4620136', '140.3677778', 1105769),
(35, 47, 'Day 0', '217.6 hours', 'Pantai Hamadi', '2.5264303', '140.6120972', 1305528),
(36, 47, 'Day 0', '263.4 hours', 'Puncak Harfat', '-1.97346', '130.465494', 158020),
(37, 48, 'Day 0', '217.6 hours', 'Pantai Hamadi', '2.5264303', '140.6120972', 1305528),
(38, 48, 'Day 0', '184.3 hours', 'Pantai Harlem', '-2.4620136', '140.3677778', 1105769),
(39, 48, 'Day 0', '263.4 hours', 'Puncak Harfat', '-1.97346', '130.465494', 158020);

-- --------------------------------------------------------

--
-- Struktur dari tabel `tourist_sites`
--

CREATE TABLE `tourist_sites` (
  `id` int NOT NULL,
  `category_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` longtext NOT NULL,
  `address` text NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `website` varchar(255) NOT NULL,
  `latitude` varchar(255) NOT NULL,
  `longitude` varchar(255) NOT NULL,
  `link_video` text NOT NULL,
  `image_primary` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `tourist_sites`
--

INSERT INTO `tourist_sites` (`id`, `category_id`, `name`, `description`, `address`, `email`, `phone`, `website`, `latitude`, `longitude`, `link_video`, `image_primary`) VALUES
(2, 1, 'Pantai Hamadi', 'Pantai Hamadi terletak kurang lebih 5 kilometer sebelah selatan pusat kota Jayapura, tepatnya di Distrik Jayapura Selatan Kotamadya Jayapura, Provinsi Papua. Pantai ini juga merupakan tempat pendaratan pertama pasukan amphibi sekutu pada tahun 1944. Selain dapat menikmati indahnya suasana pantai yang teduh ini, anda juga dapat mengetahui sejarah tentang perang dunia kedua pada masa lampau yang tersimpan di Pantai Hamadi.\r\n\r\nPantai yang kurang lebih memiliki luas sekitar 5 hektar persegi ini memiliki dari tarik tersendiri. Beberapa meter sebelum memasuki kawasan Pantai Hamadi, angin sepoi-sepoi khas suasana pantai akan terasa di sepanjang jalan, karena memang letaknya bersebelahan dengan jalan raya. Anda akan melawati pintu gerbang kawasan Pantai Hamadi dan dikenakan tarif Rp. 10.000,- untuk kendaraan roda dua dan Rp. 20.000,- untuk kendaraan roda empat.\r\n\r\nKawasan pantai ini di kelilingi pepohonan hijau, yang memberikan suasana kesejukan saat menelusuri jalan menuju pantai. Suasana erotis dan harmonis pun terasa saat anda berada diantara hutan-hutan bakau yang terdapat disekitar Pantai Hamadi ini, memberikan warna dan citra rasa tersendiri untuk anda nikmati. Pasir putih yang terbentang luas dari ujung pantai, seakan-akan memanjakan kita untuk sejenak berjemur dipinggiran pantai. Selain anda dapat memanjakan diri dengan pasir putih pada pantai ini, anda juga dapat melihat secara dekat dan langsung beberapa makam peninggalan pasukan sekutu yang terdapat pada bagian selatan Pantai Hamadi.', 'Tobati, Distrik Jayapura Selatan, Kota Jayapura, Papua', 'samasjogja@gmail.com', '0834939934', 'https://www.nativeindonesia.com/pantai-samas/', '2.5264303', '140.6120972', 'https://www.youtube.com/embed/pTGo9zje9s8', 'images/WhatsApp Image 2022-06-13 at 13.05.30.jpeg'),
(4, 1, 'Pantai Harlem', 'Pantai Harlem merupakan salah satu destinasi wisata yang selalu menjadi andalan di Jayapura. Pasalnya, destinasi wisata yang satu ini mempunyai alam indah yang akan membuat decak kagum siapapun yang melihatnya. Tak heran apabila seluruh masyarakat Jayapura tentu sudah tidak asing lagi dengan destinasi wisata bahari tersebut.\r\n\r\nBagi Anda yang suka berkunjung ke pantai, tak lengkap rasanya apabila belum berkunjung ke sana. banyak sekali turis yang menghabiskan waktu santainya ke pantai tersebut. Berbagai macam pemandangan indah disana dijamin tidak akan membuat para turis merasa bosan.\r\n\r\nTidak hanya turis lokal saja yang tertarik untuk berlibur ke destinasi wisata yang satu ini. Namun, ada juga turis asing yang rela datang jauh-jauh untuk bisa berwisata ke pantai ini. Hal tersebut pastinya merupakan suatu kebanggaan tersendiri. Oleh karena itu, artikel kali ini akan membahas berbagai macam informasi yang berkaitan dengan destinasi wisata Pantai Harlem.', 'Desa Tablasupa, Kecamatan Depapre, Kota Jayapura, Papua 99353, Indonesia.', 'pantaiharlem@gmail.com', '083149395058', 'https://www.pantainesia.com/pantai-harlem', '-2.4620136', '140.3677778', 'https://www.youtube.com/embed/Cm0elLpx__k', 'images/Pantai-Harlem.jpg'),
(14, 1, 'Puncak Harfat', 'Raja Ampat Papua Barat adalah juaranya. Ya, pulau terindah yang ada di ujung Indonesia ini memiliki pesona menakjubkan layaknya surga dunia. Berada di Raja Ampat, kurang seru tanpa mengunjungi spot terbaik, yaitu Puncak Harfat Jaya di Misool. Puncak Harfat Jaya biasa disebut Dapunlol oleh masyarakat sekitar. Lokasi ini kerap menjadi incaran turis dunia ketika travelling ke Raja Ampat. Di sini, traveler dapat melakukan berbagai kegiatan seperti hiking, wisata, fotografi, dan lainnya.\r\n\r\nArtikel ini telah tayang di www.inews.id dengan judul \" Spot Tercantik Raja Ampat, Puncak Harfat Jaya Hits di Kalangan Turis \", Klik untuk baca: https://www.inews.id/travel/destinasi/spot-tercantik-raja-ampat-puncak-harfat-jaya-hits-di-kalangan-turis.\r\nPuncak Harfat Jaya menawarkan berbagai jenis objek wisata alam seperti karst, flora, laguna, dan lainnya. Puncak ini memang terletak di Misool yang terkenal memiliki pemandangan alam indah, apalagi jika Anda melihatnya dari sisi atas. Untuk menjangkau bagian atas gunung membutuhkan stamina, karena rute yang akan ditempuh agak sedikit sulit. Tetapi, rasa lelah itu tidak akan terasa karena Anda akan disuguhkan dengan pemandangan menakjubkan. Dari atas puncak, Anda akan melihat Pulau Misool secara keseluruhan, termasuk laguna, karst, pohon rimbun, terumbu karang, dan lainnya. Meski lokasinya terpencil, Harfat Jaya dikenal oleh semua wisatawan domestik dan asing. Hiking adalah kegiatan paling populer di sini dan jangan lupan membawa kamera untuk berburu spot cantik lainnya di Pulau Misool, Papua Barat.\r\n\r\nArtikel ini telah tayang di www.inews.id dengan judul \" Spot Tercantik Raja Ampat, Puncak Harfat Jaya Hits di Kalangan Turis \", Klik untuk baca: https://www.inews.id/travel/destinasi/spot-tercantik-raja-ampat-puncak-harfat-jaya-hits-di-kalangan-turis.\r\n', 'Desa Tablasupa, Kecamatan Depapre, Kota Jayapura, Papua 99353, Indonesia.', 'puncakharfat@gmail.com', '083149395058', 'https://www.libur.co/tempat-wisata-raja-ampat', '-1.97346', '130.465494', 'https://www.youtube.com/embed/MK-eMrQwWxg', 'images/download (9).jpeg');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `avatar` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `phone`, `avatar`, `role`) VALUES
(1, 'wandi', 'wandipratama@gmail.com', '$2a$04$cisZAauu9CYNmunX2ZAZge5RWkrZs78MBHYdGP4Nj2.dwafzh5p0q', '082345566778', 'images/a5d4c32d-c433-4d57-aae7-1b2272a04f9f.jpg', 'user'),
(6, 'Admin', 'admin@gmail.com', '$2a$04$L0YO7TQGl0n5MTLhPL.APek7WjEQNdh8MK1CTU4fmd.dygaAfxmsu', '028738734', 'images/avatar_default.png', 'admin'),
(14, 'Alisawan', 'alisawan@gmail.com', '$2a$04$Vqt.T2GadXL11SXUMPVwG.16gMc6GSBDqCwlZm/AGJYb3kqjx2hEe', '083149395058', 'images/avatar_default.png', 'user'),
(15, 'Khairunnisa', 'khairunnisa@gmail.com', '$2a$04$cZfrjDsz12SwhpNYDRO/9uY/CKuY76UdANP7jTU8RUfs0paS37IXK', '083149395058', 'images/avatar_default.png', 'user');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `galleries`
--
ALTER TABLE `galleries`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `itineraries`
--
ALTER TABLE `itineraries`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `timelines`
--
ALTER TABLE `timelines`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `tourist_sites`
--
ALTER TABLE `tourist_sites`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT untuk tabel `galleries`
--
ALTER TABLE `galleries`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT untuk tabel `itineraries`
--
ALTER TABLE `itineraries`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=49;

--
-- AUTO_INCREMENT untuk tabel `timelines`
--
ALTER TABLE `timelines`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=40;

--
-- AUTO_INCREMENT untuk tabel `tourist_sites`
--
ALTER TABLE `tourist_sites`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
