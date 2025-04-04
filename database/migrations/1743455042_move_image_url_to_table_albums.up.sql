ALTER TABLE albums ADD COLUMN IF NOT EXISTS image_url TEXT;

UPDATE albums a
SET image_url = (
	SELECT url
	FROM album_images al
	LEFT JOIN images i ON i.url = al.image_url
	WHERE al.album_id = a.id
	ORDER BY (i.width, i.height) DESC
	LIMIT 1
);

DELETE FROM albums WHERE image_url IS NULL;

ALTER TABLE albums ALTER COLUMN image_url SET NOT NULL;

DROP TABLE album_images;
