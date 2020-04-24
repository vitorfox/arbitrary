package state

import (
	"github.com/valyala/fasthttp"
	"github.com/vitorfox/arbitrary/pkg/config"
	"time"
)

type StateControl struct {
	CurrentNotThrottled int
	CurrentThrottled int
}

func NewStateControl() *StateControl {
	return &StateControl{}
}

func (st *StateControl) success(ctx *fasthttp.RequestCtx, route *config.Route) {
	ctx.SetStatusCode(route.SuccessfulResponse.StatusCode)
	ctx.SetBody([]byte(route.SuccessfulResponse.Body))
}

func (st *StateControl) toomany(ctx *fasthttp.RequestCtx, route *config.Route) {
	ctx.SetStatusCode(route.ThrottledResponse.StatusCode)
	ctx.SetBody([]byte(route.ThrottledResponse.Body))
}

func (st *StateControl) SetFastHTTPContext(ctx *fasthttp.RequestCtx, route *config.Route, throttling *config.Throttling) {

	if st.CurrentNotThrottled < throttling.MaxSimultaneousRequests {
		st.CurrentNotThrottled++
		time.Sleep(throttling.DelayOnResponse)
		st.success(ctx, route)
		st.CurrentNotThrottled--
		return
	}

	if st.CurrentThrottled < throttling.MaxThrottledRequests {
		st.CurrentThrottled++
		time.Sleep(throttling.DelayOnThrottledResponse)
		st.success(ctx, route)
		st.CurrentThrottled--
		return
	}

	st.toomany(ctx, route)
}