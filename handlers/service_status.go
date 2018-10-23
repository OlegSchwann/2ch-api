package handlers

import (
	"2ch_api/types"
	"github.com/valyala/fasthttp"
	"net/http"
)

func (e *Environment) ServiceStatus(ctx *fasthttp.RequestCtx) {
	responseStatus, err := e.ConnPool.ServiceStatus()
	if err != nil {
		response, _ := types.Error{
			Message: "Error on executing 'ServiceStatus': " + err.Error(),
		}.MarshalJSON()
		ctx.Write(response)
		ctx.Response.Header.Set("Content-Type", "application/json; charset=UTF-8")
		ctx.Response.Header.SetStatusCode(http.StatusInternalServerError)
		return
	}
	response, _ := responseStatus.MarshalJSON()
	ctx.Write(response)
	ctx.Response.Header.Set("Content-Type", "application/json; charset=UTF-8")
	ctx.Response.Header.SetStatusCode(http.StatusOK)
	return
}
