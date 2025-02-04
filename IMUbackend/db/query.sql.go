// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createImg = `-- name: CreateImg :one
INSERT INTO img (name, img_path) VALUES ($1, $2) RETURNING id
`

type CreateImgParams struct {
	Name    string `json:"name"`
	ImgPath string `json:"img_path"`
}

// img CRUD
// img C
func (q *Queries) CreateImg(ctx context.Context, arg CreateImgParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createImg, arg.Name, arg.ImgPath)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createMarkdown = `-- name: CreateMarkdown :one
INSERT INTO markdown (student_id, title, content_path) VALUES ($1, $2, $3) RETURNING id
`

type CreateMarkdownParams struct {
	StudentID   string `json:"student_id"`
	Title       string `json:"title"`
	ContentPath string `json:"content_path"`
}

// Defines queries for a data unit that must maintain transactional consistency.
// RESOURCES: markdown, img, markdown_img_rel
// markdown CRUD
// markdown C
func (q *Queries) CreateMarkdown(ctx context.Context, arg CreateMarkdownParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createMarkdown, arg.StudentID, arg.Title, arg.ContentPath)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createMarkdownImgRel = `-- name: CreateMarkdownImgRel :exec
INSERT INTO markdown_img_rel (markdown_id, img_id) VALUES ($1, $2)
`

type CreateMarkdownImgRelParams struct {
	MarkdownID int64 `json:"markdown_id"`
	ImgID      int64 `json:"img_id"`
}

// md_img_rel CRUD
// md_img_rel C
func (q *Queries) CreateMarkdownImgRel(ctx context.Context, arg CreateMarkdownImgRelParams) error {
	_, err := q.db.ExecContext(ctx, createMarkdownImgRel, arg.MarkdownID, arg.ImgID)
	return err
}

const createStudent = `-- name: CreateStudent :one
INSERT INTO student (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING id
`

type CreateStudentParams struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// student CRUD
// student C
func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createStudent,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteImg = `-- name: DeleteImg :exec
DELETE FROM img WHERE id = $1
`

// img D
func (q *Queries) DeleteImg(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteImg, id)
	return err
}

const deleteMarkdown = `-- name: DeleteMarkdown :exec
DELETE FROM markdown WHERE id = $1
`

// markdown D
func (q *Queries) DeleteMarkdown(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMarkdown, id)
	return err
}

const deleteMarkdownImgRel = `-- name: DeleteMarkdownImgRel :exec
DELETE FROM markdown_img_rel WHERE markdown_id = $1 AND img_id = $2
`

type DeleteMarkdownImgRelParams struct {
	MarkdownID int64 `json:"markdown_id"`
	ImgID      int64 `json:"img_id"`
}

// md_img_rel D
func (q *Queries) DeleteMarkdownImgRel(ctx context.Context, arg DeleteMarkdownImgRelParams) error {
	_, err := q.db.ExecContext(ctx, deleteMarkdownImgRel, arg.MarkdownID, arg.ImgID)
	return err
}

const deleteMarkdownImgRelByMarkdownID = `-- name: DeleteMarkdownImgRelByMarkdownID :exec
DELETE FROM markdown_img_rel WHERE markdown_id = $1
`

// md_img_rel D
func (q *Queries) DeleteMarkdownImgRelByMarkdownID(ctx context.Context, markdownID int64) error {
	_, err := q.db.ExecContext(ctx, deleteMarkdownImgRelByMarkdownID, markdownID)
	return err
}

const deleteStudent = `-- name: DeleteStudent :exec
DELETE FROM student WHERE id = $1
`

// student D
func (q *Queries) DeleteStudent(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteStudent, id)
	return err
}

const findImages = `-- name: FindImages :many
SELECT i.img_path AS img_path
FROM img i 
JOIN markdown_img_rel mir ON i.id = mir.img_id
WHERE mir.markdown_id = $1
`

// - md, img, md_img_rel R
func (q *Queries) FindImages(ctx context.Context, markdownID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, findImages, markdownID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var img_path string
		if err := rows.Scan(&img_path); err != nil {
			return nil, err
		}
		items = append(items, img_path)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findImgByID = `-- name: FindImgByID :one
SELECT id, name, img_path FROM img WHERE id = $1
`

// img R
func (q *Queries) FindImgByID(ctx context.Context, id int64) (Img, error) {
	row := q.db.QueryRowContext(ctx, findImgByID, id)
	var i Img
	err := row.Scan(&i.ID, &i.Name, &i.ImgPath)
	return i, err
}

const findMarkdownByID = `-- name: FindMarkdownByID :one
SELECT id, student_id, title, content_path, since, updated FROM markdown WHERE id = $1
`

// markdown R
func (q *Queries) FindMarkdownByID(ctx context.Context, id int64) (Markdown, error) {
	row := q.db.QueryRowContext(ctx, findMarkdownByID, id)
	var i Markdown
	err := row.Scan(
		&i.ID,
		&i.StudentID,
		&i.Title,
		&i.ContentPath,
		&i.Since,
		&i.Updated,
	)
	return i, err
}

const findMarkdownImgRelByMarkdownID = `-- name: FindMarkdownImgRelByMarkdownID :many
SELECT markdown_id, img_id FROM markdown_img_rel WHERE markdown_id = $1
`

// md_img_rel R
func (q *Queries) FindMarkdownImgRelByMarkdownID(ctx context.Context, markdownID int64) ([]MarkdownImgRel, error) {
	rows, err := q.db.QueryContext(ctx, findMarkdownImgRelByMarkdownID, markdownID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MarkdownImgRel
	for rows.Next() {
		var i MarkdownImgRel
		if err := rows.Scan(&i.MarkdownID, &i.ImgID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findStudentByEmail = `-- name: FindStudentByEmail :one
SELECT id, name, bio, since, email, password FROM student WHERE email = $1
`

func (q *Queries) FindStudentByEmail(ctx context.Context, email string) (Student, error) {
	row := q.db.QueryRowContext(ctx, findStudentByEmail, email)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.Since,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const findStudentByID = `-- name: FindStudentByID :one
SELECT id, name, bio, since, email, password FROM student WHERE id = $1
`

// student R
func (q *Queries) FindStudentByID(ctx context.Context, id string) (Student, error) {
	row := q.db.QueryRowContext(ctx, findStudentByID, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.Since,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const login = `-- name: Login :exec
SELECT count(*) FROM student WHERE id = $1 AND password = $2
`

type LoginParams struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) Login(ctx context.Context, arg LoginParams) error {
	_, err := q.db.ExecContext(ctx, login, arg.ID, arg.Password)
	return err
}

const updateImg = `-- name: UpdateImg :exec
UPDATE img SET name = $2 AND img_path = $3 WHERE id = $1
`

type UpdateImgParams struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	ImgPath string `json:"img_path"`
}

// img U
func (q *Queries) UpdateImg(ctx context.Context, arg UpdateImgParams) error {
	_, err := q.db.ExecContext(ctx, updateImg, arg.ID, arg.Name, arg.ImgPath)
	return err
}

const updateMarkdown = `-- name: UpdateMarkdown :exec
UPDATE markdown SET title = $2 AND updated = NOW() WHERE id = $1
`

type UpdateMarkdownParams struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// markdown U
func (q *Queries) UpdateMarkdown(ctx context.Context, arg UpdateMarkdownParams) error {
	_, err := q.db.ExecContext(ctx, updateMarkdown, arg.ID, arg.Title)
	return err
}

const updateStudentBio = `-- name: UpdateStudentBio :exec
UPDATE student SET bio = $2 WHERE id = $1
`

type UpdateStudentBioParams struct {
	ID  string         `json:"id"`
	Bio sql.NullString `json:"bio"`
}

func (q *Queries) UpdateStudentBio(ctx context.Context, arg UpdateStudentBioParams) error {
	_, err := q.db.ExecContext(ctx, updateStudentBio, arg.ID, arg.Bio)
	return err
}

const updateStudentName = `-- name: UpdateStudentName :exec
UPDATE student SET name = $2 WHERE id = $1
`

type UpdateStudentNameParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// student U
func (q *Queries) UpdateStudentName(ctx context.Context, arg UpdateStudentNameParams) error {
	_, err := q.db.ExecContext(ctx, updateStudentName, arg.ID, arg.Name)
	return err
}
