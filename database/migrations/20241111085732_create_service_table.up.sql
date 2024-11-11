CREATE TABLE service (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id BIGINT(20) UNSIGNED NOT NULL,
  protector_id BIGINT(20) UNSIGNED,
  type ENUM('urgencia', 'acompaname', 'vamos_a_quedar') NOT NULL,
  status ENUM('pendiente', 'en_proceso', 'finalizado') DEFAULT 'pendiente',
  start_time DATETIME(3),
  end_time DATETIME(3),
  location_start_latitude DECIMAL(10, 8),
  location_start_longitude DECIMAL(11, 8),
  location_end_latitude DECIMAL(10, 8),
  location_end_longitude DECIMAL(11, 8),
  created_at DATETIME(3) NOT NULL,
  updated_at DATETIME(3) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (protector_id) REFERENCES protectors(id),
  KEY idx_services_created_at (created_at),
  KEY idx_services_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;