package handlers

import (
	"github.com/valyala/fasthttp"
	"net/http"

	"github.com/OlegSchwann/2ch_api/types"
)

// truncate table для всех таблиц, быстро уничтожает данные.
func (e *Environment) ServiceClear(ctx *fasthttp.RequestCtx) {
	err := e.ConnPool.ServiceClear()
	if err != nil {
		response, _ := types.Error{
			Message: "unable to truncate table: " + err.Error(),
		}.MarshalJSON()
		ctx.Write(response)
		ctx.Response.Header.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.Response.Header.SetStatusCode(http.StatusOK)
	return
}
