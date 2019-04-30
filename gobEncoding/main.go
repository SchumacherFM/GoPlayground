package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/corestoreio/pkg/util/pseudo"
)

type ProductCategory struct {
	AssetID int64
	Key     string
}

type Brand struct {
	AssetID    int64
	Sizetables []struct {
		SizeTableID int
		Category    interface{}
	}
	Features        Features
	RelatedBrandIds []int64
}

type ProductSAP struct {
	Class                  string
	Subclass               string
	SubclassCode           string
	ClassCode              string
	Warengruppe            string
	Sammelartikel          string
	SteuerKlassifikation   string
	WarengruppeCode        string
	Saison                 string
	AktivsterArtikelStatus string
	LieferantBezeichnung   string
	LieferantID            string
}
type Features []*Feature

type Feature struct {
	Key     string
	Cluster string
}
type ProductItemGroup struct {
	AssetID  int64
	ID       string
	Features Features
	Material Features
}
type Product struct {
	AssetID           int64
	Categories        []*ProductCategory
	ID                string
	Brand             *Brand
	SAP               *ProductSAP
	Features          Features
	ProductItemGroups []*ProductItemGroup
	Material          Features
}

const productCount = 1000

func main() {
	gob.Register(([]*Product)(nil))
	// generateGOB()
	generateGOBParallel()

	// the parallel version is with 1000 products approx 5s faster.
	// readGOB()
}

func generateGOB() {
	fw, err := os.Create("products.gob") // approx file size: 1GB
	mustErr(err)
	defer fw.Close()

	enc := gob.NewEncoder(fw)
	ps, err := pseudo.NewService(32168, &pseudo.Options{})
	mustErr(err)

	start := time.Now()
	products := make([]*Product, 0, productCount)
	for i := 0; i < productCount; i++ {
		p := new(Product)
		mustErr(ps.FakeData(p))
		products = append(products, p)
	}
	mustErr(enc.Encode(products))
	mustErr(fw.Sync())
	fmt.Printf("\nDone Encoding! Took %s\n", time.Since(start))
}

func generateGOBParallel() {
	done := make(chan struct{})
	workerDone := make(chan int)
	ps, err := pseudo.NewService(32168, &pseudo.Options{})
	mustErr(err)

	start := time.Now()
	prodChanIN := make(chan *Product)
	prodChanOUT := make(chan *Product)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(i int, ps *pseudo.Service, prodChanIN <-chan *Product, prodChanOUT chan<- *Product) {
			println("Worker Started", i)
			for p := range prodChanIN {
				mustErr(ps.FakeData(p))
				prodChanOUT <- p
			}
			workerDone <- i
			println("Worker Ended", i)
		}(i, ps, prodChanIN, prodChanOUT)
	}

	go func() {
		for i := 0; i < runtime.NumCPU(); i++ {
			<-workerDone
		}
		close(prodChanOUT)
	}()

	go func() {
		fw, err := os.Create("products2.gob") // approx file size: 1GB
		mustErr(err)
		defer fw.Close()
		enc := gob.NewEncoder(fw)

		products := make([]*Product, 0, productCount)
		for p := range prodChanOUT {
			products = append(products, p)
		}

		mustErr(enc.Encode(products))
		mustErr(fw.Sync())
		done <- struct{}{}
		close(done)
		println("All channels closed")
	}()

	go func() {
		for i := 0; i < productCount; i++ {
			prodChanIN <- new(Product)
		}
		close(prodChanIN)
	}()
	<-done
	fmt.Printf("\nDone Encoding Parallel! Took %s\n", time.Since(start))
}

func readGOB() {
	fr, err := os.Open("products.gob")
	mustErr(err)
	defer fr.Close()

	var products []*Product
	err = gob.NewDecoder(bufio.NewReaderSize(fr, 1024*64)).Decode(&products)
	mustErr(err) // when does it return io.EOF?
	fmt.Printf("Total Entries: %d\n", len(products))
}

func mustErr(err error) {
	if err != nil {
		panic(err)
	}
}
