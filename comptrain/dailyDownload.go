package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := mainDownload(ctx, "https://comptrain.co/wod/", "wod")
		if err != nil {
			println(err.Error())
		}
		return nil
	})
	g.Go(func() error {
		err := mainDownload(ctx, "https://comptrain.co/home-gym/", "hg")
		if err != nil {
			println(err.Error())
		}
		return nil
	})
	g.Wait()
}

func mainDownload(ctx context.Context, urlPrefix, filePrefix string) error {
	usaToday := strings.ToLower(time.Now().Format("Monday-1-2-2006"))
	euToday := strings.ToLower(time.Now().Format("2006-01-02")) // 2006 January 2nd
	// urlToday := fmt.Sprintf("%s%s/", urlPrefix, usaToday)
	println("Downloading ", urlPrefix)

	req, err := http.NewRequest("GET", urlPrefix, nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("/home/pi/comptrain/%s_%s.html", filePrefix, euToday))
	if err != nil {
		return err
	}
	defer f.Close()
	if n, err := io.Copy(f, resp.Body); err != nil {
		return err
	} else {
		fmt.Printf("Downloaded HTML for %s with size %2.fkb\n", usaToday, float64(n)/1024)
	}
	return nil
}
