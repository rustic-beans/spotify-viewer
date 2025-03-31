ALTER TABLE playlists ADD COLUMN IF NOT EXISTS image_url TEXT DEFAULT NULL;

UPDATE playlists a
SET image_url = (
	SELECT url
	FROM playlist_images al
	LEFT JOIN images i ON i.url = al.image_url
	WHERE al.playlist_id = a.id
	ORDER BY (i.width, i.height) DESC
	LIMIT 1
);

DELETE FROM playlists WHERE image_url IS NULL;

ALTER TABLE playlists ALTER COLUMN image_url SET NOT NULL;

DROP TABLE playlist_images;
