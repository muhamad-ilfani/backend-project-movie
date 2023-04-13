SELECT 
    mv.id "id",
    mv.title "title",
    mv.description "description",
    mv.duration "duration",
    mv.artists "artists",
    mv.genres "genres",
    mv.watch_url "watch_url",
    mv.viewer "viewer"
FROM movieapps.movies mv
where mv.title = COALESCE(NULLIF($3, ''), mv.title) 
    and mv.description = COALESCE(NULLIF($4, ''), mv.description) 
    and mv.artists = COALESCE(NULLIF($5, ''), mv.artists) 
    and mv.genres = COALESCE(NULLIF($6, ''), mv.genres) 
limit $1 offset $2;