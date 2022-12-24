package main

import "github.com/snowmerak/lux/swagger"

var appSwagger = &swagger.Info{
	Title:       "configer",
	Description: "A config manager",
	Version:     "0.0.1",
	Contact: struct {
		Email string "json:\"email,omitempty\""
	}{Email: "snowmerak@outlook.com"},
}

var rootGetSwagger = &swagger.Router{
	Summary:     "Get config value",
	Description: "Get config value",
}

var rootPostSwagger = &swagger.Router{
	Summary:     "Set config value",
	Description: "Set config value to the database with the given body",
}

var rootPutSwagger = &swagger.Router{
	Summary:     "Set config value",
	Description: "Set config value to the database with the given URL query (:name?value=...)",
}

var rootDeleteSwagger = &swagger.Router{
	Summary:     "Delete config value",
	Description: "Delete config value from the database",
}
