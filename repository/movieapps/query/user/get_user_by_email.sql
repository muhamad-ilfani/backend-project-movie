SELECT
    us.id "id",
    us.password "password",
    us.is_admin "is_admin"
FROM movieapps.users us
WHERE us.email = $1
LIMIT 1;