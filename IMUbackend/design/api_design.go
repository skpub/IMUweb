package design

import (
	. "goa.design/goa/v3/dsl"
)

// var JWTAuth = JWTServerInterceptor("jwt", func() {
// 	Description("Secures student only endpoints")
// })

var JWTAuth = Interceptor("JWTAuth", func() {
	Description("Serverside validation")

	ReadPayload(func() {
		Attribute("token", String)
	})
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
	Attribute("studentID", String)
	Attribute("password", String)
})

var _ = Service("imubackend", func() {
	Description("markdown file server")
	HTTP(func() {
		Path("/api")
	})

	Method("createArticle", func() {
		ServerInterceptor(JWTAuth)
		Description("create markdown file")
		Payload(func() {
			Attribute("articleName", String)
			Attribute("content", String)
			Attribute("image", ArrayOf(File))
			Attribute("token", String)
			Required("articleName", "content", "token")
		})
		HTTP(func() {
			Cookie("token")
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
		ServerInterceptor(JWTAuth)
		Description("get student profile")
		Result(StudentProfile)
		Payload(func() {
			Attribute("token", String)
			Required("token")
		})
		HTTP(func() {
			GET("/student/profile")
			Cookie("token")
			Response(StatusOK)
		})
	})
	Method("getProfiles", func() {
		
		Description("get students profile")
		Result(ArrayOf(StudentProfile))
		HTTP(func() {
			GET("/students/profile")
			Response(StatusOK)
		})
	})
	Method("login", func() {
		Description("IMU teacher and student login")
		Payload(LoginAttribute)
		Result(String)
		HTTP(func() {
			POST("/student/login")
			Response(StatusOK, func() {
				Cookie("token")
				CookieMaxAge(60 * 60 * 24)
				CookiePath("/")
			})
		})
	})

	Method("Signup", func() {
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
		ServerInterceptor(JWTAuth)
		Description("refresh token (each 5 minutes)")
		Payload(func() {
			Attribute("token", String)
			Required("token")
		})
		Result(func() {
			Attribute("token", String)
		})
		HTTP(func() {
			POST("/refresh")
			Cookie("token")
			Response(StatusOK, func() {
				Cookie("token")
				CookieMaxAge(60 * 60 * 24)
				CookiePath("/")
			})
		})
	})

	Method("UpdateBio", func() {
		ServerInterceptor(JWTAuth)
		Description("update student bio")
		Payload(func() {
			Attribute("bio", String)
			Attribute("token", String)
			Required("bio", "token")
		})
		HTTP(func() {
			PUT("/student/bio")
			Cookie("token")
			Response(StatusOK)
		})
	})

	Method("UpdateImg", func() {
		ServerInterceptor(JWTAuth)
		Description("update student img")
		Payload(func() {
			Attribute("img", File)
			Attribute("token", String)
			Required("img", "token")
		})
		HTTP(func() {
			MultipartRequest()
			PUT("/student/icon")
			Cookie("token")
			Response(StatusOK)
		})
	})

	Method("UpdateName", func() {
		ServerInterceptor(JWTAuth)
		Description("update student name")
		Payload(func() {
			Attribute("name", String)
			Attribute("token", String)
			Required("name", "token")
		})
		HTTP(func() {
			PUT("/student/name")
			Cookie("token")
			Response(StatusOK)
		})
	})
})
