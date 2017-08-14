package main

import "strings"

func prefixLines(prefix, text string) string {
	ret := ""
	for _, s := range strings.Split(text, "\n") {
		ret += prefix + s + "\n"
	}
	// discard latest newline
	return ret[:len(ret)-1]
}

func suffixLines(suffix, text string) string {
	ret := ""
	for _, s := range strings.Split(text, "\n") {
		ret += s + suffix + "\n"
	}
	// discard latest newline
	return ret[:len(ret)-1]
}
