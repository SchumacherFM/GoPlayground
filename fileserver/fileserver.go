package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"

	"github.com/minio/highwayhash"
)

var highwayHashKey [highwayhash.Size]byte

func main() {
	n, err := rand.Read(highwayHashKey[:])
	fmt.Printf("%d\n%s\n%#v\n", n, err, highwayHashKey)
	return
	log.Println("Version:", runtime.Version())
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":9000", nil)
}

/*
$ wrk -t12 -c400 -d10m --latency http://127.0.0.1:9000

Go 1.5beta3
$ wrk -t12 -c400 -d10m --latency http://127.0.0.1:9000
Running 10m test @ http://127.0.0.1:9000
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    16.45ms   11.42ms 491.89ms   78.34%
    Req/Sec     1.67k     0.87k    6.23k    66.66%
  Latency Distribution
     50%   15.86ms
     75%   20.10ms
     90%   27.36ms
     99%   45.05ms
  8944168 requests in 10.00m, 2.93GB read
  Socket errors: connect 155, read 0, write 0, timeout 0
  Non-2xx or 3xx responses: 3205216
Requests/sec:  14903.52
Transfer/sec:      5.01MB

go 1.4.2
$ wrk -t12 -c400 -d10m --latency http://127.0.0.1:9000
Running 10m test @ http://127.0.0.1:9000
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    14.40ms    7.15ms 363.19ms   77.42%
    Req/Sec     1.88k     0.88k    7.26k    76.34%
  Latency Distribution
     50%   14.07ms
     75%   16.61ms
     90%   22.43ms
     99%   34.63ms
  10080244 requests in 10.00m, 3.40GB read
  Socket errors: connect 155, read 39, write 0, timeout 0
  Non-2xx or 3xx responses: 2635072
Requests/sec:  16796.94
Transfer/sec:      5.80MB

/////////////////////////////////////////////////

Go 1.4.2
$ wrk -c 20 -d 1m -t 4 --latency http://127.0.0.1:9000
Running 1m test @ http://127.0.0.1:9000
  4 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.98ms    3.61ms  63.69ms   96.07%
    Req/Sec     3.50k   511.11     4.20k    83.08%
  Latency Distribution
     50%    1.35ms
     75%    1.94ms
     90%    2.78ms
     99%   20.05ms
  836677 requests in 1.00m, 351.08MB read
Requests/sec:  13924.23
Transfer/sec:      5.84MB

Go 1.5beta3
$ wrk -c 20 -d 1m -t 4 --latency http://127.0.0.1:9000
Running 1m test @ http://127.0.0.1:9000
  4 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.44ms    1.97ms  37.83ms   96.26%
    Req/Sec     4.27k   395.62     4.96k    79.08%
  Latency Distribution
     50%    1.12ms
     75%    1.53ms
     90%    2.11ms
     99%   11.33ms
  1019235 requests in 1.00m, 427.69MB read
Requests/sec:  16983.66
Transfer/sec:      7.13MB


@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

go 1.4.2

lib code imported
$ wrk -t12 -c200 -d1m --latency http://127.0.0.1:9000
Running 1m test @ http://127.0.0.1:9000
  12 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.79ms    5.49ms  79.13ms   77.43%
    Req/Sec     1.26k    97.74     2.22k    71.57%
  Latency Distribution
     50%   13.58ms
     75%   15.44ms
     90%   17.81ms
     99%   27.43ms
  905121 requests in 1.00m, 379.80MB read
  Non-2xx or 3xx responses: 13
Requests/sec:  15072.45
Transfer/sec:      6.32MB

with 3 lines of code and use of library
$ wrk -t12 -c200 -d1m --latency http://127.0.0.1:9000
Running 1m test @ http://127.0.0.1:9000
  12 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    14.97ms    9.39ms 224.99ms   79.51%
    Req/Sec     1.12k   207.72     1.90k    78.59%
  Latency Distribution
     50%   14.55ms
     75%   17.44ms
     90%   22.44ms
     99%   48.97ms
  803311 requests in 1.00m, 337.01MB read
  Socket errors: connect 0, read 5, write 0, timeout 0
  Non-2xx or 3xx responses: 244
Requests/sec:  13366.92
Transfer/sec:      5.61MB


go 1.5beta3
$ wrk -t12 -c200 -d1m --latency http://127.0.0.1:9000
Running 1m test @ http://127.0.0.1:9000
  12 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    14.77ms    9.47ms 243.05ms   79.78%
    Req/Sec     1.13k   146.82     1.55k    79.35%
  Latency Distribution
     50%   14.95ms
     75%   17.53ms
     90%   21.44ms
     99%   47.20ms
  811442 requests in 1.00m, 340.16MB read
  Non-2xx or 3xx responses: 1430
Requests/sec:  13511.53
Transfer/sec:      5.66MB
*/
