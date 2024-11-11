CREATE TABLE protector (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id BIGINT(20) UNSIGNED NOT NULL,
  rating DECIMAL(3,2) DEFAULT 0.0,
  status ENUM('disponible', 'ocupado', 'fuera_de_servicio') DEFAULT 'disponible',
  last_latitude DECIMAL(10, 8),      -- Coordenada de latitud
  last_longitude DECIMAL(11, 8),     -- Coordenada de longitud
  created_at DATETIME(3) NOT NULL,
  updated_at DATETIME(3) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  KEY idx_protectors_created_at (created_at),
  KEY idx_protectors_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
