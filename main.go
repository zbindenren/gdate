package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/peterbourgon/ff/v3"
)

const (
	dfltLayout     = "Mon, 2 Jan 2006 15:04:05 MST"
	dfltDateHeader = "Date"
	envVarPrefix   = "GDATE"
)

func main() {
	fs := flag.NewFlagSet("gdate", flag.ContinueOnError)

	var (
		url        = fs.String("url", "", "The url which gdate uses to parse timestamp (also via GDATE_URL env var)")
		layout     = fs.String("layout", dfltLayout, "The layout to use to parse timestamp (also via GDATE_LAYOUT env var)")
		dateHeader = fs.String("header", dfltDateHeader, "The header name that contains the timestamp (also via GDATE_HEADER env var)")
	)

	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix(envVarPrefix),
		ff.WithConfigFileParser(ff.PlainParser),
	)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return
		}

		log.Fatal(err)
	}

	if *url == "" {
		log.Fatal("No URL specified")
	}

	dateCmd, err := exec.LookPath("date")
	if err != nil {
		log.Fatal(err)
	}

	client := resty.New().SetBaseURL(*url)

	r, err := client.R().Get("/")
	if err != nil {
		log.Fatal(err)
	}

	st, err := time.Parse(*layout, r.Header().Get(*dateHeader))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Time Difference:", time.Until(st))

	dateString := st.In(time.Now().Location()).Format("2 Jan 2006 15:04:05")
	fmt.Printf("Setting system date to: %s\n", dateString)
	args := []string{"--set", dateString}

	cmd := exec.Command(dateCmd, args...) // #nosec G204
	if out, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("failed to run '%s': %s", strings.Join(cmd.Args, " "), string(out))
	}
}
