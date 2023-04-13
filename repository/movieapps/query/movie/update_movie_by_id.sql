UPDATE movieapps.movies
SET 
    title = COALESCE(NULLIF($2,''), title),
    description = COALESCE(NULLIF($3,''), description),
    duration = COALESCE(NULLIF($4,''), duration),
    artists = COALESCE(NULLIF($5,''), artists),
    genres = COALESCE(NULLIF($6,''), genres),
    watch_url = COALESCE(NULLIF($7,''), watch_url),
    viewer = COALESCE(NULLIF($8,0), viewer),
    updated_by = $9,
    updated_at = now()
WHERE id = $1;