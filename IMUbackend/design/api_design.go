package design

import (
	. "goa.design/goa/v3/dsl"
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

var StudentProfile = Type("StudentProfile", func() {
	Attribute("studentID", String)
	Attribute("name", String)
	Attribute("bio", String)
	Attribute("img", File)
})

var ArticleIdName = Type("ArticleIdName", func() {
	Attribute("id", String)
	Attribute("name", String)
	Attribute("updated", Int64)
})

var File = Type("file", func() {
	Attribute("name", String)
	Attribute("content", Bytes)
})

var LoginAttribute = Type("Login", func() {
	Attribute("studentId", String)
	Attribute("password", String)
})

var _ = Service("imubackend", func() {
	Description("markdown file server")
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
		// Payload(func() {
		// 	Attribute("id", String)
		// })
		Payload(String)
		Result(func() {
			Attribute("id", String)
			Attribute("studentID", String)
			Attribute("articleName", String)
			Attribute("content", String)
			Attribute("image", ArrayOf(File))
			Attribute("createdAt", Int64)
			Attribute("updatedAt", Int64)
		})
		HTTP(func() {
			GET("/article/get/{id}")
			Param("id", String)
			Response(StatusOK)
		})
	})

	// User CRUD
	Method("getProfile", func() {
		Security(JWTAuth)
		Description("get student profile")
		Result(StudentProfile)
		Payload(func() {
			TokenField(2, "token", String, func() {
				Description("JWT token")
			})
			Required("token")
		})
		HTTP(func() {
			GET("/student/profile")
			Response(StatusOK)
		})
	})
	Method("getProfiles", func() {
		NoSecurity()
		Description("get students profile")
		Result(ArrayOf(StudentProfile))
		HTTP(func() {
			GET("/students/profile")
			Response(StatusOK)
		})
	})
	Method("login", func() {
		NoSecurity()
		Description("IMU teacher and student login")
		Payload(LoginAttribute)
		Result(String)
		HTTP(func() {
			POST("/student/login")
			Response(StatusOK)
		})
	})

	Method("Signup", func() {
		NoSecurity()
		Description("Uraguchi Nyugaku")
		Payload(func() {
			Attribute("studentID", String)
			Attribute("name", String)
			Attribute("email", String)
			Attribute("password", String)
		})
		Result(String)
		HTTP(func() {
			POST("/student/signup")
			Response(StatusOK)
		})
	})

	Method("refreshToken", func() {
		Security(JWTAuth)
		Description("refresh token (each 5 minutes)")
		Payload(func() {
			Attribute("token", String)
			TokenField(2, "token", String, func() {
				Description("JWT token")
			})
			Required("token")
		})
		Result(func() {
			Attribute("token", String)
		})
		HTTP(func() {
			POST("/refresh")
			Response(StatusOK)
		})
	})

	Method("UpdateBio", func() {
		Security(JWTAuth)
		Description("update student bio")
		Payload(func() {
			Attribute("bio", String)
			TokenField(2, "token", String, func() {
				Description("JWT token")
			})
			Required("bio", "token")
		})
		HTTP(func() {
			PUT("/student/bio")
			Response(StatusOK)
		})
	})

	Method("UpdateImg", func() {
		Security(JWTAuth)
		Description("update student img")
		Payload(func() {
			Attribute("img", File)
			TokenField(2, "token", String, func() {
				Description("JWT token")
			})
			Required("img", "token")
		})
		HTTP(func() {
			MultipartRequest()
			PUT("/student/icon")
			Response(StatusOK)
		})
	})

	Method("UpdateName", func() {
		Security(JWTAuth)
		Description("update student name")
		Payload(func() {
			Attribute("name", String)
			TokenField(2, "token", String, func() {
				Description("JWT token")
			})
			Required("name", "token")
		})
		HTTP(func() {
			PUT("/student/name")
			Response(StatusOK)
		})
	})
})
