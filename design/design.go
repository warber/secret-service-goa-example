package design

import . "goa.design/goa/v3/dsl"

var _ = API("secret", func() {
	Title("Secret Service")
	Description("Keptn API for creating secrets")
	Server("sercret", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var _ = Service("openapi", func() {
	Files("/openapi.json", "openapi3.json")
})

var _ = Service("secret", func() {
	Description("The secret service can create delete or update secrets")
	Method("getSecret", func() {
		Result(GetSecretsResponseBody)
		HTTP(func() {
			GET("/secret")
			Response(StatusOK)
		})
	})
	Method("createSecret", func() {
		Result(CreateSecretResponseBody)
		Payload(CreateSecretRequestBody)
		Error("secret_already_exists")
		HTTP(func() {
			POST("/secret")
			Response(StatusCreated)
			Response("secret_already_exists", StatusConflict, func() {
				Description("The secret already exists")
			})
		})
	})
})

var GetSecretsResponseBody = Type("GetSecretsResponseBody", func() {
	Description("Get Secrets Result")
	Attribute("secrets", ArrayOf(GetSecretResponseBody))
})

var GetSecretResponseBody = Type("GetSecretResponseBody", func() {
	Description("Secret result")
	Attribute("keys", ArrayOf(String))
	Attribute("name", String)
	Attribute("scope", String)
})

var CreateSecretRequestBody = Type("CreateSecretRequestBody", func() {
	Description("Create Secret Request Body")
	Attribute("data", MapOf(String, String))

	Attribute("name", String)
	Attribute("scope", String)
})
var CreateSecretResponseBody = Type("CreateSecretResponseBody", func() {
	Description("Create Secret Result")
	Attribute("data", MapOf(String, String))
	Attribute("name", String)
	Attribute("scope", String)
})
