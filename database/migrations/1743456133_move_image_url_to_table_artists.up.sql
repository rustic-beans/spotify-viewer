ALTER TABLE artists ADD COLUMN IF NOT EXISTS image_url TEXT DEFAULT NULL;

UPDATE artists a
SET image_url = (
	SELECT url
	FROM artist_images al
	LEFT JOIN images i ON i.url = al.image_url
	WHERE al.artist_id = a.id
	ORDER BY (i.width, i.height) DESC
	LIMIT 1
);

DELETE FROM artists WHERE image_url IS NULL;

ALTER TABLE artists ALTER COLUMN image_url SET NOT NULL;

DROP TABLE artist_images;
