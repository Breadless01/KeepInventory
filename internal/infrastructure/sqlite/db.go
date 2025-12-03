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
            name TEXT NOT NULL,
            sitz TEXT
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle kunden nicht anlegen: %v", err)
	}

	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS projekte (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            kunde TEXT
        );
    `); err != nil {
		log.Fatalf("konnte Tabelle projekte nicht anlegen: %v", err)
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

	return db
}
