package main

import (
	"regexp"
	"strings"
	"os"
)

const RESET = "\x1b[0m"

var colors = map[string]string{
	"black": "\x1b[30m",
	"red": "\x1b[31m",
	"green": "\x1b[32m",
	"yellow": "\x1b[33m",
	"blue": "\x1b[34m",
	"magenta": "\x1b[35m",
	"cyan": "\x1b[36m",
	"white": "\x1b[37m",
	"gray": "\x1b[90m",
}

func col(text string) string {
	color := colors[strings.ToLower(os.Getenv("COLORGO_COL"))]
	if color == "" {
		color = colors["red"]
	}
	return color + text + RESET
}

func line(text string) string {
	color := colors[strings.ToLower(os.Getenv("COLORGO_LINE"))]
	if color == "" {
		color = colors["red"]
	}
	return color + text + RESET
}

func file(text string) string {
	color := colors[strings.ToLower(os.Getenv("COLORGO_FILE"))]
	if color == "" {
		color = colors["blue"]
	}
	return color + text + RESET
}

func prefix(text string) string {
	color := colors[strings.ToLower(os.Getenv("COLORGO_PREFIX"))]
	if color == "" {
		color = colors["gray"]
	}
	return color + text + RESET
}

var re = regexp.MustCompile(`^(.*\.go)\:(\d+)\:(?:(\d+):)?(.*)\n?$`)

var trimPrefix = os.Getenv("COLORGO_TRIM_PREFIX")
var addPrefix = os.Getenv("COLORGO_ADD_PREFIX")

func parse(in string) (out string) {
	foundPrefix := false

	if trimPrefix != "" {
		if strings.HasPrefix(in, trimPrefix) {
			foundPrefix = true
			in = strings.TrimPrefix(in, trimPrefix)
		}
	}
	matches := re.FindStringSubmatch(in)
	if len(matches) > 0 {
		if matches[3] == "" {
			out = file(matches[1]) + ":" + line(matches[2]) + matches[4] + "\n"
		} else {
			out = file(matches[1]) + ":" + line(matches[2]) + ":" + col(matches[3]) + matches[4] + "\n"		
		}
		if foundPrefix {
			out = prefix(addPrefix) + out
		}
	} else {
		if strings.HasSuffix(in, "\n") {
			out = in
		} else {
			out = in + "\n"
		}
	}
	return
}
