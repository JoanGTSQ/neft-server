CREATE TABLE warning (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  reported_by BIGINT(20) UNSIGNED NOT NULL,
  description TEXT NOT NULL,
  location_latitude DECIMAL(10, 8),
  location_longitude DECIMAL(11, 8),
  created_at DATETIME(3) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (reported_by) REFERENCES users(id) ON DELETE SET NULL,
  KEY idx_warnings_created_at (created_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;