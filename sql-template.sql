PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE passwords(key varchar(50), type varchar(8), value text);
CREATE TABLE version(version DECIMAL(2,2))
INSERT INTO "version" VALUES (1.0)
COMMIT;
