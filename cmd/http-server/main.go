package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/vitorfox/arbitrary/pkg/config"
	"github.com/vitorfox/arbitrary/pkg/state"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	conf = &config.Base{}
	routes = make(map[string]*config.Route)
	throtteling *config.Throttling
	st = state.NewStateControl()
)

func handler (ctx *fasthttp.RequestCtx) {

	key := string(ctx.Path())+"-"+strings.ToLower(string(ctx.Method()))
	route, ok := routes[key]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	st.SetFastHTTPContext(ctx, route, throtteling)
}

func statHandler (ctx *fasthttp.RequestCtx) {
	ctx.SetBody([]byte(fmt.Sprintf("--- st:\n%v\n\n", st)))
	ctx.SetStatusCode(200)
}

func main() {
	file, err := os.Open(os.Getenv("CONFIG_FILE"))
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	err = yaml.Unmarshal(data, conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- config:\n%v\n\n", conf)

	for _, v := range conf.Routes {
		routes[v.Path+"-"+strings.ToLower(v.Method)] = &v
	}
	throtteling = &conf.Throttling

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		err = fasthttp.ListenAndServe("0.0.0.0:8088", handler)
		log.Fatal("Failed to listen", err)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		err = fasthttp.ListenAndServe("0.0.0.0:9088", statHandler)
		log.Fatal("Failed to listen", err)
		wg.Done()
	}()

	wg.Wait()

}

