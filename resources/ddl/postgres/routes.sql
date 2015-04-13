CREATE TABLE %s.routes (
  route_id integer NOT NULL,
  agency_id integer NOT NULL,
  route_short_name varchar(45) DEFAULT NULL,
  route_long_name varchar(128) DEFAULT NULL,
  route_desc varchar(64) DEFAULT NULL,
  route_type integer DEFAULT NULL,
  route_url varchar(45) DEFAULT NULL,
  route_color char(6) DEFAULT NULL,
  route_text_color char(6) DEFAULT NULL,
  PRIMARY KEY (route_id)
);