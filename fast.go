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
// BenchmarkSlow-16              68          18515406 ns/op        20279242 B/op     189843 allocs/op
// BenchmarkFast-16             896           1328198 ns/op          634471 B/op       8405 allocs/op
// PASS
// ok      hw3     2.857s

const (
	atMarker           = "@"
	uniqueBrowsersSize = 200 // naive optimizations accordingly to benchmarks
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
	} // write result of resolved user based on browsers used
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
		err := tempUser.UnmarshalJSON(scanner.Bytes()) // read bytes directly
		if err != nil {
			panic(err)
		}
		handleBrowsers(tempUser, seenBrowsers, &uniqueBrowsers, foundUsersBuilder, idx)
		idx++
	}

	fmt.Fprintln(out, "found users:\n"+foundUsersBuilder.String())
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
