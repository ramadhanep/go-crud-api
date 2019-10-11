package helpers

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

func JSON(ctx *fasthttp.RequestCtx, data map[string]interface{}) {

	// Set Header
	ctx.Response.Header.Set("Content-Type", "application/json")

	res, err := json.Marshal(data)

	if err != nil {
		log.Println("Error Convert to JSON")
		data["error"] = err
	}

	// Write Output
	ctx.Write(res)

	// Set Status HTTP CODE
	ctx.SetStatusCode(fasthttp.StatusOK)
}
