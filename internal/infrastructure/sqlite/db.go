package sqlite

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func OpenDB(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("konnte SQLite-DB nicht öffnen: %v", err)
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		log.Fatalf("konnte foreign_keys aktivieren: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS bauteile (
            id                        INTEGER PRIMARY KEY AUTOINCREMENT,
            teil_name                 TEXT    NOT NULL,
            kunde_id                  INTEGER REFERENCES kunden (id) ON DELETE SET NULL,
            projekt_id                INTEGER REFERENCES projekte (id) ON DELETE SET NULL,
            erstelldatum              TEXT    NOT NULL,
            typ_id                    INTEGER NOT NULL,
            herstellungsart_id        INTEGER NOT NULL,
            verschleissteil_id        INTEGER NOT NULL,
            funktion_id               INTEGER NOT NULL,
            material_id               INTEGER NOT NULL,
            oberflaechenbehandlung_id INTEGER NOT NULL,
            farbe_id                  INTEGER NOT NULL,
            reserve_id                INTEGER NOT NULL,
            sachnummer                TEXT    NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle bauteile nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		CREATE VIRTUAL TABLE IF NOT EXISTS bauteile_fts USING fts5(
		  teil_name,
		  sachnummer,
		  content='bauteile',
		  content_rowid='id'
		);
	`); err != nil {
		log.Fatalf("konnte virtuelle Tabelle bauteile nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		INSERT INTO bauteile_fts (rowid, teil_name, sachnummer) SELECT id, teil_name, sachnummer FROM bauteile
	`); err != nil {
		log.Fatalf("Insert Fehler: %v", err)
	}

	if _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS typ (
            id      INTEGER PRIMARY KEY AUTOINCREMENT,
            name    TEXT NOT NULL,
            symbol  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle typ nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS herstellungsart (
            id      INTEGER PRIMARY KEY AUTOINCREMENT,
            name    TEXT NOT NULL,
            symbol  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle herstellungsart nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS verschleissteil (
            id      INTEGER PRIMARY KEY AUTOINCREMENT,
            name    TEXT NOT NULL,
            symbol  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle verschleissteil nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS funktion (
            id      INTEGER PRIMARY KEY AUTOINCREMENT,
            name    TEXT NOT NULL,
            symbol  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle funktion nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS material (
            id      INTEGER PRIMARY KEY AUTOINCREMENT,
            name    TEXT NOT NULL,
            symbol  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle material nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS oberflaechenbehandlung (
            id      INTEGER PRIMARY KEY AUTOINCREMENT,
            name    TEXT NOT NULL,
            symbol  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle oberflaechenbehandlung nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS farbe (
            id      INTEGER PRIMARY KEY AUTOINCREMENT,
            name    TEXT NOT NULL,
            symbol  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle farbe nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS reserve (
            id      INTEGER PRIMARY KEY AUTOINCREMENT,
            name    TEXT NOT NULL,
            symbol  INTEGER NOT NULL
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle reserve nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS kunden (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            sitz TEXT
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle kunden nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		CREATE VIRTUAL TABLE IF NOT EXISTS kunden_fts USING fts5(
		  name,
		  sitz,
		  content='kunden',
		  content_rowid='id'
		);
	`); err != nil {
		log.Fatalf("konnte virtuelle Tabelle kunden nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		INSERT INTO kunden_fts (rowid, name, sitz) SELECT id, name, sitz FROM kunden
	`); err != nil {
		log.Fatalf("Insert Fehler: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS projekte (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            kunde TEXT
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle projekte nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		CREATE VIRTUAL TABLE IF NOT EXISTS projekte_fts USING fts5(
		  name,
		  kunde,
		  content='projekte',
		  content_rowid='id'
		);
	`); err != nil {
		log.Fatalf("konnte virtuelle Tabelle projekte nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		INSERT INTO projekte_fts (rowid, name, kunde) SELECT id, name, kunde FROM projekte
	`); err != nil {
		log.Fatalf("Insert Fehler: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS lieferanten (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            sitz TEXT
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle lieferanten nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		CREATE VIRTUAL TABLE IF NOT EXISTS lieferanten_fts USING fts5(
		  name,
		  sitz,
		  content='lieferanten',
		  content_rowid='id'
		);
	`); err != nil {
		log.Fatalf("konnte virtuelle Tabelle lieferanten nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		INSERT INTO lieferanten_fts (rowid, name, sitz) SELECT id, name, sitz FROM lieferanten
	`); err != nil {
		log.Fatalf("Insert Fehler: %v", err)
	}

	if _, err := db.Exec(`
      	CREATE TRIGGER IF NOT EXISTS bauteile_ai AFTER INSERT ON bauteile
		BEGIN
		  INSERT INTO bauteile_fts(rowid, teil_name, sachnummer)
		  VALUES (new.id, new.teil_name, new.sachnummer);
		END;        
    `); err != nil {
		log.Fatalf("konnte Tigger für bauteile After Insert nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
  		CREATE TRIGGER IF NOT EXISTS bauteile_au AFTER UPDATE ON bauteile
		BEGIN
		  INSERT INTO bauteile_fts(bauteile_fts, rowid, teil_name, sachnummer)
		  VALUES('delete', old.id, old.teil_name, old.sachnummer);
		
		  INSERT INTO bauteile_fts(rowid, teil_name, sachnummer)
		  VALUES (new.id, new.teil_name, new.sachnummer);
		END;
    `); err != nil {
		log.Fatalf("konnte Tigger für bauteile After Update nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TRIGGER IF NOT EXISTS bauteile_ad AFTER DELETE ON bauteile
		BEGIN
		  INSERT INTO bauteile_fts(bauteile_fts, rowid, teil_name, sachnummer)
		  VALUES('delete', old.id, old.teil_name, old.sachnummer);
		END;
    `); err != nil {
		log.Fatalf("konnte Tigger für bauteile After Delete nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
      	CREATE TRIGGER IF NOT EXISTS kunden_ai AFTER INSERT ON kunden
		BEGIN
		  INSERT INTO kunden_fts(rowid, name, sitz)
		  VALUES (new.id, new.name, new.sitz);
		END;        
    `); err != nil {
		log.Fatalf("konnte Tigger für kunden After Insert nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
  		CREATE TRIGGER IF NOT EXISTS kunden_au AFTER UPDATE ON kunden
		BEGIN
		  INSERT INTO kunden_fts(kunden_fts, rowid, name, sitz)
		  VALUES('delete', old.id, old.name, old.sitz);
		
		  INSERT INTO kunden_fts(rowid, name, sitz)
		  VALUES (new.id, new.new, new.sitz);
		END;
    `); err != nil {
		log.Fatalf("konnte Tigger für kunden After Update nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TRIGGER IF NOT EXISTS kunden_ad AFTER DELETE ON kunden
		BEGIN
		  INSERT INTO kunden_fts(kunden_fts, rowid, namw, sitz)
		  VALUES('delete', old.id, old.name, old.sitz);
		END;
    `); err != nil {
		log.Fatalf("konnte Tigger für kunden After Delete nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
      	CREATE TRIGGER IF NOT EXISTS projekte_ai AFTER INSERT ON projekte
		BEGIN
		  INSERT INTO projekte_fts(rowid, name, kunde)
		  VALUES (new.id, new.name, new.kunde);
		END;        
    `); err != nil {
		log.Fatalf("konnte Tigger für projekte After Insert nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
  		CREATE TRIGGER IF NOT EXISTS projekte_au AFTER UPDATE ON projekte
		BEGIN
		  INSERT INTO projekte_fts(projekte_fts, rowid, name, kunde)
		  VALUES('delete', old.id, old.name, old.kunde);
		
		  INSERT INTO projekte_fts(rowid, name, kunde)
		  VALUES (new.id, new.new, new.kunde);
		END;
    `); err != nil {
		log.Fatalf("konnte Tigger für projekte After Update nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TRIGGER IF NOT EXISTS projekte_ad AFTER DELETE ON projekte
		BEGIN
		  INSERT INTO projekte_fts(projekte_fts, rowid, name, kunde)
		  VALUES('delete', old.id, old.name, old.kunde);
		END;
    `); err != nil {
		log.Fatalf("konnte Tigger für projekte After Delete nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
      	CREATE TRIGGER IF NOT EXISTS lieferanten_ai AFTER INSERT ON lieferanten
		BEGIN
		  INSERT INTO lieferanten_fts(rowid, name, sitz)
		  VALUES (new.id, new.name, new.sitz);
		END;        
    `); err != nil {
		log.Fatalf("konnte Tigger für lieferanten After Insert nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
  		CREATE TRIGGER IF NOT EXISTS lieferanten_au AFTER UPDATE ON lieferanten
		BEGIN
		  INSERT INTO lieferanten_fts(lieferanten_fts, rowid, name, sitz)
		  VALUES('delete', old.id, old.name, old.sitz);
		
		  INSERT INTO lieferanten_fts(rowid, name, sitz)
		  VALUES (new.id, new.new, new.sitz);
		END;
    `); err != nil {
		log.Fatalf("konnte Tigger für lieferanten After Update nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TRIGGER IF NOT EXISTS lieferanten_ad AFTER DELETE ON lieferanten
		BEGIN
		  INSERT INTO lieferanten_fts(lieferanten_fts, rowid, name, sitz)
		  VALUES('delete', old.id, old.name, old.sitz);
		END;
    `); err != nil {
		log.Fatalf("konnte Tigger für projekte After Delete nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS lieferant_bauteil (
			lieferant_id INTEGER NOT NULL,
			bauteil_id INTEGER NOT NULL,
			PRIMARY KEY (lieferant_id, bauteil_id),
			FOREIGN KEY (lieferant_id) REFERENCES lieferanten(id),
			FOREIGN KEY (bauteil_id) REFERENCES bauteile(id))
	`); err != nil {
		log.Fatalf("konnte Junction Table für lieferant_bauteil nicht anlegen: %v", err)
	}

	return db
}
