package main

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	err := filepath.WalkDir("gen/go/wb", visitGen)
	if err != nil {
		panic(err)
	}

}

func visitGen(path string, di fs.DirEntry, err error) error {
	if strings.Contains(di.Name(), "oas") {

		f, err := os.OpenFile(path, os.O_RDWR, 0777)
		if err != nil {
			panic("err open file")
		}

		data, err := io.ReadAll(f)
		if err != nil {
			panic("err read file")
		}

		cutted, found := strings.CutPrefix(filepath.Dir(path), "gen/go/wb/")
		if !found {
			panic("prefix not found")
		}

		data = bytes.ReplaceAll(data, []byte("package api"), []byte(fmt.Sprintf("package wb%s", strings.Title(cutted[:len(cutted)-3]))))

		if di.Name() == "oas_schemas_gen.go" {
			r := regexp.MustCompile(`type .* struct {`)

			for {
				i := r.FindIndex(data)
				if i == nil {
					break
				}
				{
					rr := regexp.MustCompile(`}`)
					ii := rr.FindIndex(data[i[1]:])

					data = bytes.ReplaceAll(data, data[i[0]:i[1]+ii[1]], []byte(""))
				}
			}
		}

		f.Truncate(0)
		f.Seek(0, 0)

		_, err = f.Write(data)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
