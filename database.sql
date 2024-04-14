-- Тут просто лежит контент по базе

CREATE DATABASE banner_service_db;

CREATE TABLE IF NOT EXISTS features (
	id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS tags (
	id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS banners(
	id SERIAL PRIMARY KEY,
	feature_id INT NOT NULL,
	tag_ids INT[] NOT NULL,
	title VARCHAR(255) NOT NULL,
  	text TEXT NOT NULL,
	url TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	is_active BOOLEAN NOT NULL
);

CREATE INDEX IF NOT EXISTS ix_banners_id ON banners (id);
CREATE INDEX IF NOT EXISTS ix_banners_feature_id ON banners (feature_id);
CREATE INDEX IF NOT EXISTS ix_banners_tag_ids ON banners USING gin(tag_ids);

INSERT INTO features DEFAULT VALUES;
INSERT INTO features DEFAULT VALUES;
INSERT INTO features DEFAULT VALUES;
INSERT INTO features DEFAULT VALUES;
INSERT INTO features DEFAULT VALUES;

INSERT INTO tags DEFAULT VALUES;
INSERT INTO tags DEFAULT VALUES;
INSERT INTO tags DEFAULT VALUES;
INSERT INTO tags DEFAULT VALUES;
INSERT INTO tags DEFAULT VALUES;
INSERT INTO tags DEFAULT VALUES;
INSERT INTO tags DEFAULT VALUES;

INSERT INTO banners(feature_id, tag_ids, title, text, url, created_at, updated_at, is_active) VALUES (
	1, '{2, 3, 4}', 'AMOGUS', 'amogus', 'www.amogus.com', NOW(), NOW(), true);

INSERT INTO banners(feature_id, tag_ids, title, text, url, created_at, updated_at, is_active) VALUES (
	4, '{1, 7, 6}', 'Carpicorn', 'Lorem ipsum', 'www.astrology.com', NOW(), NOW(), true);

INSERT INTO banners(feature_id, tag_ids, title, text, url, created_at, updated_at, is_active) VALUES (
	3, '{5, 1, 4}', 'Big truck', 'big truck for sale', 'www.trucks.com', NOW(), NOW(), true);



SELECT * FROM banners WHERE feature_id = 1 AND
(tag_ids <@ '{4, 2, 3}'  and tag_ids @> '{4, 2, 3}');

SELECT * FROM banners WHERE '{4}' <@ tag_ids;

SELECT * FROM banners WHERE feature_id = ANY('{4,3}');

SELECT * FROM banners 
WHERE (feature_id = ANY('{2,1}') OR ARRAY[feature_id] @> '{2,1}') 
AND '{3,4}' <@ tag_ids
ORDER BY id
LIMIT 10 OFFSET 0;

SELECT NOT EXISTS (
    SELECT unnest('{1,2,3,8}'::int[]) EXCEPT
    SELECT id FROM tags
) AS ok;

 SINSERT INTO banners (feature_id, tag_ids, title, text, url, created_at, updated_at, is_active)
SELECT feature_id, tag_ids FROM banners WHERE (feature_id, tag_ids) NOT IN (SELECT feature_id, tag_ids FROM banners)
 VALUES (1, '{7}', "t", "t", "u", NOW(), NOW(), true)

UPDATE banners SET ... WHERE id = 1;

