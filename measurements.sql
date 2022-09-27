SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";



CREATE TABLE `measurements` (
  `id` int(11) NOT NULL,
  `temperature` double NOT NULL,
  `heatIndex` double NOT NULL,
  `humidity` double NOT NULL,
  `timestamp` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `averages` (
  `id` int(11) NOT NULL,
  `temperature_av` double NOT NULL,
  `heatIndex_av` double NOT NULL,
  `humidity_av` double NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




ALTER TABLE `measurements`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `averages`
  ADD PRIMARY KEY (`id`);



ALTER TABLE `averages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=0;
ALTER TABLE `measurements`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=0;

COMMIT;
