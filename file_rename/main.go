package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alecthomas/repr"
	"github.com/dsoprea/go-exif/v3"
)

func main() {
	if err := filepath.Walk(os.Args[1], func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(strings.ToLower(path), ".jpg") {
			return nil
		}
		if strings.Contains(path, "noExif") {
			return nil
		}

		rawExif, err := exif.SearchFileAndExtractExif(path)
		if err != nil {
			fmt.Println("L27:", err, " => ", path)
			return nil
		}

		tags, _, err := exif.GetFlatExifData(rawExif, &exif.ScanOptions{})
		if err != nil {
			fmt.Println("L33:", err, " => ", path)
			return nil
		}
		var dateTimeOriginal time.Time
		for _, tag := range tags {
			if tag.TagName == "DateTimeOriginal" && tag.Formatted != "" {
				dateTimeOriginal, err = time.Parse("2006:01:02 15:04:05", tag.Formatted)
				if err != nil {
					fmt.Println("L41:", err, " => ", path)
					return nil
				}
			}
		}

		if dateTimeOriginal.IsZero() {
			dir, file := filepath.Split(path)

			np := filepath.Join(dir, "noExif")
			if _, err := os.Stat(np); os.IsNotExist(err) {
				if err := os.Mkdir(np, 0o755); err != nil {
					return err
				}
			}
			fmt.Println("date zero => ", path, "np moved")
			if err := os.Rename(path, filepath.Join(dir, "noExif", file)); err != nil {
				fmt.Println(err, "rename failed")
			}
			return nil
		}

		if info.ModTime().Before(dateTimeOriginal) || info.ModTime().Truncate(time.Hour).Equal(dateTimeOriginal.Truncate(time.Hour)) {
			fmt.Println("date equal => ", path, "exif time:", dateTimeOriginal.String(), "file time:", info.ModTime().String())
			return nil
		}

		println("=====================================================")
		if err := os.Chtimes(path, dateTimeOriginal, dateTimeOriginal); err != nil {
			panic(err)
		}
		repr.Println(dateTimeOriginal.String(), path)

		return nil
	}); err != nil {
		panic(err)
	}

	//p1 := strings.IndexRune(dir.Name(), '(')
	//p2 := strings.IndexRune(dir.Name(), ')')
	//if p1 < 1 || p2 < 1 {
	//	continue
	//}
	//folderName := dir.Name()[p1+1 : p2]
	////_ = os.Mkdir(filepath.Join(os.Args[1], folderName), 0o755)
	//
	//fileNameSplit := strings.Split(dir.Name(), "-")
	//fileNameSplit2 := append([]string{}, fileNameSplit[0])
	//fileNameSplit2 = append(fileNameSplit2, " "+folderName+" ")
	//fileNameSplit2 = append(fileNameSplit2, fileNameSplit[1:]...)
	//
	//old := filepath.Join(os.Args[1], dir.Name())
	//new := filepath.Join(os.Args[1], strings.Join(fileNameSplit2, "-"))
	//repr.Println(old, "=>", new)
	//
	//if err := os.Rename(old, new); err != nil {
	//	repr.Println(err)
	//}
}
