package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// http://comptrain.co/individuals/workout/wednesday-%C2%B7-1-31-18/
	// http://comptrain.co/individuals/workout/friday-%C2%B7-11-3-17/
	// first date Saturday · 07.01.17 / July 1. ;-)
	// github.com/gocolly/colly
	// curl -i 'http://comptrain.co/individuals/wp-admin/admin-ajax.php' -H 'Pragma: no-cache' -H 'Origin: http://comptrain.co' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'Accept: text/html, */*; q=0.01' -H 'Cache-Control: no-cache' -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Referer: http://comptrain.co/individuals/workoutarchive/' -H 'DNT: 1' --data 'action=vc_get_vc_grid_data&vc_action=vc_get_vc_grid_data&tag=vc_basic_grid&data%5Bvisible_pages%5D=5&data%5Bpage_id%5D=191&data%5Bstyle%5D=all&data%5Baction%5D=vc_get_vc_grid_data&data%5Bshortcode_id%5D=1498577965240-57c33715-8b38-7&data%5Btag%5D=vc_basic_grid&vc_post_id=191&_vcnonce=3948c1f247' --compressed
	// curl -i 'http://comptrain.co/individuals/wp-admin/admin-ajax.php' -H 'Pragma: no-cache' -H 'Origin: http://comptrain.co' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'Accept: text/html, */*; q=0.01' -H 'Cache-Control: no-cache' -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Referer: http://comptrain.co/individuals/workoutarchive/' -H 'DNT: 1' --data 'action=vc_get_vc_grid_data&vc_action=vc_get_vc_grid_data&tag=vc_basic_grid&data%5Bvisible_pages%5D=5&data%5Bpage_id%5D=191&data%5Bstyle%5D=all&data%5Baction%5D=vc_get_vc_grid_data&data%5Bshortcode_id%5D=1498577965240-57c33715-8b38-7&data%5Btag%5D=vc_basic_grid&vc_post_id=191&_vcnonce=ab03d90854' --compressed

	const Nonce = "3034c0afc0" // changes frequently

	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.CacheDir("./testdata/"),
		//colly.Debugger(&debug.LogDebugger{}),
		colly.Async(false),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36"),
		// 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Referer: http://comptrain.co/individuals/workoutarchive/'
	)
	c.Limit(&colly.LimitRule{
		// seems like this only has an effect on c.Visit, not quite sure.
		DomainGlob:  "http://comptrain.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	//{
	//	fp, err := os.OpenFile("./admin-ajax.html", os.O_RDONLY, 0)
	//	tr := &cstesting.HTTPTrip{
	//		// use buffer pool
	//		GenerateResponse: func(_ *http.Request) *http.Response {
	//			return &http.Response{
	//				StatusCode: 200,
	//				Body:       fp,
	//				Header: http.Header{
	//					"Content-Type": []string{"html"},
	//				},
	//			}
	//		},
	//		Err:          err,
	//		RequestCache: make(map[*http.Request]struct{}),
	//	}
	//	c.WithTransport(tr)
	//}

	// Find and visit next page links
	c.OnHTML(`a.vc_gitem-link[href]`, func(e *colly.HTMLElement) {
		//log.Println("Next", e.Attr("href"))
		e.Request.Visit(e.Attr("href"))
	})

	f, err := os.Create("./z_comptrain.html")
	fatalCheckErr(err)

	f.WriteString(`<html><body>`)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		_, err := fmt.Fprintf(f, "\n<h1>%s</h1>\n", e.Text)
		fatalCheckErr(err)
	})

	c.OnHTML("div.wpb_text_column div.wpb_wrapper", func(e *colly.HTMLElement) {
		html, err := e.DOM.Html()
		fatalCheckErr(err)
		_, err = f.WriteString(strings.TrimSpace(html))
		fatalCheckErr(err)
	})

	fatalCheckErr(c.Post("http://comptrain.co/individuals/wp-admin/admin-ajax.php", map[string]string{
		"action":              "vc_get_vc_grid_data",
		"vc_action":           "vc_get_vc_grid_data",
		"tag":                 "vc_basic_grid",
		"data[visible_pages]": "5",
		"data[page_id]":       "191",
		"data[style]":         "all",
		"data[action]":        "vc_get_vc_grid_data",
		"data[shortcode_id]":  "1498577965240-57c33715-8b38-7",
		"data[tag]":           "vc_basic_grid",
		"vc_post_id":          "191",
		"_vcnonce":            Nonce,
	}))

	c.Wait()
	f.WriteString(` </body></html>`)
	fatalCheckErr(f.Close())

	// Display collector's statistics
	log.Println(c)

}

func fatalCheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
