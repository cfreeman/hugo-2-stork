/*
 * Copyright (c) Clinton Freeman 2021
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"flag"
	"fmt"
	"github.com/gernest/front"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.Println("hugo-2-stork v0.0.1")

	var srcDir string
	var urlStem string
	flag.StringVar(&srcDir, "src", ".", "The hugo directory containing posts to index")
	flag.StringVar(&urlStem, "url", "https://tomsachsarchive.org/posts/", "The URL prefix for search results")
	flag.Parse()

	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)

	buf := []byte("[input]\n")
	buf = append(buf, []byte("frontmatter_handling=\"Parse\"\n")...)
	buf = append(buf, []byte("base_directory = \""+srcDir+"\"")...)
	buf = append(buf, []byte("\n\nfiles = [\n")...)

	var files []string

	filepath.Walk(srcDir, func(src string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, src)
		}

		return nil
	})

	i := 0
	for _, src := range files {
		// It's a regular file, open it and look for YAML.
		file, err := os.Open(src)
		if err != nil {
			log.Fatal("Unable to open file %v", err)
		}
		defer file.Close()

		meta, _, err := m.Parse(file)
		if err == nil {
			title := fmt.Sprintf("%v", meta["title"])
			base := filepath.Base(src)
			ext := filepath.Ext(base)
			url := urlStem + base[:strings.LastIndex(base, ext)]

			// Add a comma to seperate entries for each post if we have more than
			// one.
			if i > 0 {
				buf = append(buf, []byte(",\n")...)
			}
			buf = append(buf, []byte("\t{path=\""+base+
				"\", url=\"" + url + "\", title=\""+title+"\"}")...)
			i = i + 1
		}
	}

	buf = append(buf, []byte("\n]\n\n")...)


	err := ioutil.WriteFile("stork.toml", buf, 0644)
	if err != nil {
		log.Fatal("Unable to write config file %v", err)
	}
}
