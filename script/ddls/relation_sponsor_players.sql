CREATE TABLE IF NOT EXISTS relation_sponsor_players (
    id_players INT NOT NULL,
    id_sponsor INT NOT NULL,
    status_code TINYINT NOT NULL DEFAULT 0,
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

ALTER TABLE relation_sponsor_players ADD FOREIGN KEY (id_players) REFERENCES players(id);
ALTER TABLE relation_sponsor_players ADD FOREIGN KEY (id_sponsor) REFERENCES sponsor(id);