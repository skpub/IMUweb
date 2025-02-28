// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	db "IMUbackend/db"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Querier is an autogenerated mock type for the Querier type
type Querier struct {
	mock.Mock
}

// CreateImg provides a mock function with given fields: ctx, name
func (_m *Querier) CreateImg(ctx context.Context, name string) (uuid.UUID, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for CreateImg")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uuid.UUID, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uuid.UUID); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMarkdown provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateMarkdown(ctx context.Context, arg db.CreateMarkdownParams) (uuid.UUID, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateMarkdown")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateMarkdownParams) (uuid.UUID, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateMarkdownParams) uuid.UUID); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateMarkdownParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMarkdownImgRel provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateMarkdownImgRel(ctx context.Context, arg db.CreateMarkdownImgRelParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateMarkdownImgRel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateMarkdownImgRelParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateStudent provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateStudent(ctx context.Context, arg db.CreateStudentParams) (string, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateStudent")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateStudentParams) (string, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateStudentParams) string); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateStudentParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteImg provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteImg(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteImg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMarkdown provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteMarkdown(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMarkdown")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMarkdownImgRel provides a mock function with given fields: ctx, arg
func (_m *Querier) DeleteMarkdownImgRel(ctx context.Context, arg db.DeleteMarkdownImgRelParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMarkdownImgRel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.DeleteMarkdownImgRelParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMarkdownImgRelByMarkdownID provides a mock function with given fields: ctx, markdownID
func (_m *Querier) DeleteMarkdownImgRelByMarkdownID(ctx context.Context, markdownID uuid.UUID) error {
	ret := _m.Called(ctx, markdownID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMarkdownImgRelByMarkdownID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, markdownID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStudent provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteStudent(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteStudent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindImages provides a mock function with given fields: ctx, markdownID
func (_m *Querier) FindImages(ctx context.Context, markdownID uuid.UUID) ([]uuid.UUID, error) {
	ret := _m.Called(ctx, markdownID)

	if len(ret) == 0 {
		panic("no return value specified for FindImages")
	}

	var r0 []uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]uuid.UUID, error)); ok {
		return rf(ctx, markdownID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []uuid.UUID); ok {
		r0 = rf(ctx, markdownID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, markdownID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindImgByID provides a mock function with given fields: ctx, id
func (_m *Querier) FindImgByID(ctx context.Context, id uuid.UUID) (db.Img, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindImgByID")
	}

	var r0 db.Img
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (db.Img, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) db.Img); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Img)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMarkdownByID provides a mock function with given fields: ctx, id
func (_m *Querier) FindMarkdownByID(ctx context.Context, id uuid.UUID) (db.Markdown, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindMarkdownByID")
	}

	var r0 db.Markdown
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (db.Markdown, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) db.Markdown); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Markdown)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMarkdownImgRelByMarkdownID provides a mock function with given fields: ctx, markdownID
func (_m *Querier) FindMarkdownImgRelByMarkdownID(ctx context.Context, markdownID uuid.UUID) ([]db.MarkdownImgRel, error) {
	ret := _m.Called(ctx, markdownID)

	if len(ret) == 0 {
		panic("no return value specified for FindMarkdownImgRelByMarkdownID")
	}

	var r0 []db.MarkdownImgRel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]db.MarkdownImgRel, error)); ok {
		return rf(ctx, markdownID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []db.MarkdownImgRel); ok {
		r0 = rf(ctx, markdownID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.MarkdownImgRel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, markdownID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindStudentByEmail provides a mock function with given fields: ctx, email
func (_m *Querier) FindStudentByEmail(ctx context.Context, email string) (db.Student, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindStudentByEmail")
	}

	var r0 db.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.Student, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.Student); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(db.Student)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindStudentByID provides a mock function with given fields: ctx, id
func (_m *Querier) FindStudentByID(ctx context.Context, id string) (db.Student, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindStudentByID")
	}

	var r0 db.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.Student, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.Student); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Student)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindStudents provides a mock function with given fields: ctx
func (_m *Querier) FindStudents(ctx context.Context) ([]db.Student, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindStudents")
	}

	var r0 []db.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]db.Student, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []db.Student); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Student)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticle provides a mock function with given fields: ctx, id
func (_m *Querier) GetArticle(ctx context.Context, id uuid.UUID) ([]db.GetArticleRow, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetArticle")
	}

	var r0 []db.GetArticleRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]db.GetArticleRow, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []db.GetArticleRow); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.GetArticleRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMarkdown provides a mock function with given fields: ctx
func (_m *Querier) ListMarkdown(ctx context.Context) ([]db.ListMarkdownRow, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListMarkdown")
	}

	var r0 []db.ListMarkdownRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]db.ListMarkdownRow, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []db.ListMarkdownRow); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.ListMarkdownRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMarkdownID provides a mock function with given fields: ctx
func (_m *Querier) ListMarkdownID(ctx context.Context) ([]uuid.UUID, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListMarkdownID")
	}

	var r0 []uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]uuid.UUID, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []uuid.UUID); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, arg
func (_m *Querier) Login(ctx context.Context, arg db.LoginParams) (int64, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.LoginParams) (int64, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.LoginParams) int64); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.LoginParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateImg provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateImg(ctx context.Context, arg db.UpdateImgParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateImg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateImgParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMarkdown provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateMarkdown(ctx context.Context, arg db.UpdateMarkdownParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMarkdown")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateMarkdownParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStudentBio provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateStudentBio(ctx context.Context, arg db.UpdateStudentBioParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStudentBio")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateStudentBioParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStudentImg provides a mock function with given fields: ctx, id
func (_m *Querier) UpdateStudentImg(ctx context.Context, id string) (uuid.NullUUID, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStudentImg")
	}

	var r0 uuid.NullUUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uuid.NullUUID, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uuid.NullUUID); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(uuid.NullUUID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStudentName provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateStudentName(ctx context.Context, arg db.UpdateStudentNameParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStudentName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateStudentNameParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQuerier creates a new instance of Querier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *Querier {
	mock := &Querier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
