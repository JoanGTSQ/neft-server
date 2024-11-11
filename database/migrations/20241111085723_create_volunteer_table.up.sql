CREATE TABLE volunteer (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id BIGINT(20) UNSIGNED NOT NULL,
  training VARCHAR(100),
  points INT DEFAULT 0,
  created_at DATETIME(3) NOT NULL,
  updated_at DATETIME(3) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  KEY idx_volunteers_created_at (created_at),
  KEY idx_volunteers_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
