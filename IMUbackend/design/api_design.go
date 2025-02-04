package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("imubackend", func() {
	Title("IMU web backend")
	Server("imubackend", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var Markdown = Type("Markdown", func() {
	Attribute("articleName", String)
	Attribute("content", String)
	Required("articleName", "content")
})

var LoginAttribute = Type("Login", func() {
	Attribute("studentName", String)
	Attribute("password", String)
})

var SignupAttribute = Type("Signup", func() {
	Attribute("student_id", String)
	Attribute("password", String)
	Attribute("studentName", String)
	Attribute("email", String)
})


var _ = Service("imubackend", func() {
	Description("markdown file server.")

	HTTP(func() {
		Path("/api")
	})

	Method("create", func() {
		Description("create markdown file.")
		Payload(Markdown)
		HTTP(func() {
			POST("/article/create")
			Response(StatusOK)
		})
	})

	// User CRUD
	Method("createStudent", func() {
		Description("create student.")
		Payload(SignupAttribute)
		HTTP(func() {
			POST("/student/create")
			Response(StatusOK)
		})
	})
	Method("login", func() {
		Description("IMU teacher and student login.")
		Payload(LoginAttribute)
		Result(String)
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
		})
	})
})
