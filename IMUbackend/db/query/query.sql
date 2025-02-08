-- student CRUD
-- student C
-- name: CreateStudent :one
INSERT INTO student (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING id;

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


-- Defines queries for a data unit that must maintain transactional consistency.
-- RESOURCES: markdown, img, markdown_img_rel 
-- markdown CRUD
-- markdown C
-- name: CreateMarkdown :one
INSERT INTO markdown (student_id, title, content_path) VALUES ($1, $2, $3) RETURNING id;

-- markdown R
-- name: FindMarkdownByID :one
SELECT * FROM markdown WHERE id = $1;
-- name: ListMarkdownID :many
SELECT id FROM markdown;

-- markdown U
-- name: UpdateMarkdown :exec
UPDATE markdown SET title = $2 AND updated = NOW() WHERE id = $1;

-- markdown D
-- name: DeleteMarkdown :exec
DELETE FROM markdown WHERE id = $1;


-- img CRUD
-- img C
-- name: CreateImg :one
INSERT INTO img (name) VALUES ($1) RETURNING id;

-- img R
-- name: FindImgByID :one
SELECT * FROM img WHERE id = $1;

-- img U
-- name: UpdateImg :exec
UPDATE img SET name = $2 WHERE id = $1;

-- img D
-- name: DeleteImg :exec
DELETE FROM img WHERE id = $1;


-- md_img_rel CRUD
-- md_img_rel C
-- name: CreateMarkdownImgRel :exec
INSERT INTO markdown_img_rel (markdown_id, img_id) VALUES ($1, $2);

-- md_img_rel R
-- name: FindMarkdownImgRelByMarkdownID :many
SELECT * FROM markdown_img_rel WHERE markdown_id = $1;

-- md_img_rel D
-- name: DeleteMarkdownImgRelByMarkdownID :exec
DELETE FROM markdown_img_rel WHERE markdown_id = $1;

-- md_img_rel D
-- name: DeleteMarkdownImgRel :exec
DELETE FROM markdown_img_rel WHERE markdown_id = $1 AND img_id = $2;


--- md, img, md_img_rel R
-- name: FindImages :many
SELECT i.id AS img_path
FROM img i 
JOIN markdown_img_rel mir ON i.id = mir.img_id
WHERE mir.markdown_id = $1;

-- name: GetArticle :many
SELECT *
FROM markdown m
JOIN markdown_img_rel mir ON m.id = mir.markdown_id
JOIN img i ON i.id = mir.img_id 
WHERE m.id = $1;
