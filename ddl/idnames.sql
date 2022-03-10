DROP TABLE IF EXISTS idnames;
CREATE TABLE IF NOT EXISTS idnames (
	 id       INT
	,name     VARCHAR(16)
	,created  TIMESTAMP DEFAULT NOW()
	,modified TIMESTAMP DEFAULT NOW()
);

COMMENT ON COLUMN idnames.id       IS 'This is a id';
COMMENT ON COLUMN idnames.name     IS 'This is a name';
COMMENT ON COLUMN idnames.created  IS 'Timestamp when the row was created';
COMMENT ON COLUMN idnames.modified IS 'Timestamp when the row was modified';
