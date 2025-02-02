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
})
