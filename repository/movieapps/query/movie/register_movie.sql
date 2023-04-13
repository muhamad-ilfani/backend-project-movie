INSERT INTO movieapps.movies(title, description, duration, artists, genres, watch_url, created_by)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING id;