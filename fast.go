package hw3

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// goos: windows
// goarch: amd64
// pkg: hw3
// cpu: AMD Ryzen 9 5900HX with Radeon Graphics
// BenchmarkSlow-8               69          17937375 ns/op        20191368 B/op     189831 allocs/op
// BenchmarkFast-8              735           1602824 ns/op         1858347 B/op      10407 allocs/op
// PASS
// ok      hw3     2.859s

const (
	atMarker           = "@"
	uniqueBrowsersSize = 1000
)

func handeUniqueBrowser(seenBrowsers map[string]struct{}, browser string, uniqueBrowsers *int) {
	if _, seen := seenBrowsers[browser]; !seen {
		seenBrowsers[browser] = struct{}{}
		*uniqueBrowsers++
	}
}

func handleBrowsers(user *User, seenBrowsers map[string]struct{}, uniqueBrowsers *int, writer *strings.Builder, idx int) {
	var AndroidPresented, MSIEPResented bool
	for i := 0; i < len(user.Browsers); i++ {
		browser := user.Browsers[i]
		switch {
		case strings.Contains(browser, "Android"):
			AndroidPresented = true
			handeUniqueBrowser(seenBrowsers, browser, uniqueBrowsers)
		case strings.Contains(browser, "MSIE"):
			MSIEPResented = true
			handeUniqueBrowser(seenBrowsers, browser, uniqueBrowsers)
		}
	}
	if !(AndroidPresented && MSIEPResented) {
		return
	}
	writer.WriteString(fmt.Sprintf("[%d] %s <%s>\n", idx, user.Name, strings.ReplaceAll(user.Email, atMarker, " [at] ")))
}

// This function avoids collecting intermediate results during parsing target text file.
func FastSearch(out io.Writer) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var seenBrowsers = make(map[string]struct{}, uniqueBrowsersSize)
	var uniqueBrowsers, idx int

	var tempUser *User
	var foundUsersBuilder *strings.Builder = &strings.Builder{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tempUser = &User{}
		err := tempUser.UnmarshalJSON([]byte(scanner.Text()))
		if err != nil {
			panic(err)
		}
		handleBrowsers(tempUser, seenBrowsers, &uniqueBrowsers, foundUsersBuilder, idx)
		idx++
	}

	fmt.Fprintln(out, "found users:\n"+foundUsersBuilder.String())
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
