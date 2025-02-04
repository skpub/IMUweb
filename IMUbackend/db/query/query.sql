-- student CRUD
-- student C
-- name: CreateStudent :one
INSERT INTO student (id, name, email, password, since) VALUES ($1, $2, $3, $4, NOW()) RETURNING id;

-- student R
-- name: FindStudentByID :one
SELECT * FROM student WHERE id = $1;
-- name: FindStudentByEmail :one
SELECT * FROM student WHERE email = $1;
-- name: Login :exec
SELECT count(*) FROM student WHERE id = $1 AND password = $2;

-- student U
-- name: UpdateStudentName :exec
UPDATE student SET name = $2 WHERE id = $1;
-- name: UpdateStudentBio :exec
UPDATE student SET bio = $2 WHERE id = $1;

-- student D
-- name: DeleteStudent :exec
DELETE FROM student WHERE id = $1;


-- markdown CRUD
-- markdown C
