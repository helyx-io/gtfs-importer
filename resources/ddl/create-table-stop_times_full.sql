CREATE TABLE `gtfs_%s`.`stop_times_full` (
  `stop_id` varchar(45) NOT NULL,
  `stop_code` varchar(45) DEFAULT NULL,
  `stop_name` varchar(64) DEFAULT NULL,
  `stop_desc` varchar(128) DEFAULT NULL,
  `stop_lat` FLOAT DEFAULT NULL,
  `stop_lon` FLOAT DEFAULT NULL,
  `stop_geo` GEOMETRY DEFAULT NULL,
  `zone_id` varchar(45) DEFAULT NULL,
  `stop_url` varchar(45) DEFAULT NULL,
  `location_type` int(11) DEFAULT NULL,
  `parent_station` varchar(45) DEFAULT NULL,
  `arrival_time` time DEFAULT NULL,
  `departure_time` time DEFAULT NULL,
  `stop_sequence` int(11) DEFAULT NULL,
  `stop_head_sign` varchar(8) DEFAULT NULL,
  `pickup_type` int(11) DEFAULT NULL,
  `drop_off_type` int(11) DEFAULT NULL,
  `service_id` varchar(45) DEFAULT NULL,
  `trip_id` varchar(45) NOT NULL,
  `trip_headsign` varchar(45) DEFAULT NULL,
  `direction_id` int(11) DEFAULT NULL,
  `block_id` varchar(45) DEFAULT NULL,
  `shape_id` varchar(45) DEFAULT NULL,
  `agency_id` varchar(45) DEFAULT NULL,
  `route_id` varchar(45) DEFAULT NULL,
  `route_short_name` varchar(45) DEFAULT NULL,
  `route_long_name` varchar(128) DEFAULT NULL,
  `route_desc` varchar(64) DEFAULT NULL,
  `route_type` int(11) DEFAULT NULL,
  `route_url` varchar(45) DEFAULT NULL,
  `route_color` varchar(45) DEFAULT NULL,
  `route_text_color` varchar(45) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;