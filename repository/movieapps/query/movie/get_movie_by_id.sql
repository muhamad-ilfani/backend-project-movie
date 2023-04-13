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
WHERE mv.id = $1
limit 1;