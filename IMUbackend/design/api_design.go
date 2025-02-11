package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var JWTAuth = JWTSecurity("jwt", func() {
	Description("Secures student only endpoints")
})

var _ = API("imubackend", func() {
	Title("IMU web backend")
	Server("imubackend", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var ArticleIdName = Type("ArticleIdName", func() {
	Attribute("id", String)
	Attribute("name", String)
	Attribute("updated", String)
})

var File = Type("file", func() {
	Attribute("name", String)
	Attribute("content", Bytes)
})

var LoginAttribute = Type("Login", func() {
	Attribute("studentId", String)
	Attribute("password", String)
})

var SignupAttribute = Type("Signup", func() {
	Attribute("student_id", String)
	Attribute("password", String)
	Attribute("studentName", String)
	Attribute("email", String)
})

var _ = Service("imubackend", func() {
	Description("markdown file server")
	cors.Origin("*")

	HTTP(func() {
		Path("/api")
	})

	Method("createArticle", func() {
		Security(JWTAuth)
		Description("create markdown file")
		Payload(func() {
			Attribute("articleName", String)
			Attribute("content", String)
			Attribute("image", ArrayOf(File))
			TokenField(2, "token", String, func() {
				Description("JWT token")
			})
			Required("articleName", "content", "token")
		})
		HTTP(func() {
			MultipartRequest()
			POST("/article/create")
			Response(StatusOK)
		})
	})

	Method("listArticle", func() {
		Description("list article")
		Result(func() {
			Attribute("list", ArrayOf(ArticleIdName))
		})
		HTTP(func() {
			GET("/article/list")
			Response(StatusOK)
		})
	})

	Method("getArticle", func() {
		Description("get article")
		Payload(func() {
			Attribute("id", String)
		})
		Result(func() {
			Attribute("id", String)
			Attribute("studentID", String)
			Attribute("articleName", String)
			Attribute("content", String)
			Attribute("image", ArrayOf(File))
			Attribute("createdAt", String)
			Attribute("updatedAt", String)
		})
		HTTP(func() {
			GET("/article/get")
			Response(StatusOK)
		})
	})

	// User CRUD
	Method("createStudent", func() {
		Description("create student")
		Payload(SignupAttribute)
		HTTP(func() {
			POST("/student/create")
			Response(StatusOK)
		})
	})
	Method("login", func() {
		Description("IMU teacher and student login")
		Payload(LoginAttribute)
		Result(String)
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
		})
	})
})
