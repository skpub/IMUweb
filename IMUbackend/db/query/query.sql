-- USERS CRUD
-- USERS C
-- name: CreateUser :one
INSERT INTO users (id, name, email, password, since) VALUES ($1, $2, $3, $4, NOW()) RETURNING id;

-- USERS R
-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1;
-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1;
-- name: Login :exec
SELECT count(*) FROM users WHERE id = $1 AND password = $2;

-- USERS U
-- name: UpdateUserName :exec
UPDATE users SET name = $2 WHERE id = $1;
-- name: UpdateUserBio :exec
UPDATE users SET bio = $2 WHERE id = $1;

-- USERS D
-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
