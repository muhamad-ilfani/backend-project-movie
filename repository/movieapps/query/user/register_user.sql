INSERT INTO movieapps.users(user_name, email, password, is_admin)
VALUES($1, $2, $3, $4)
RETURNING id;