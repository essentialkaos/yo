package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v6/arg"
	"pkg.re/essentialkaos/ek.v6/env"
	"pkg.re/essentialkaos/ek.v6/fmtc"
	"pkg.re/essentialkaos/ek.v6/fsutil"
	"pkg.re/essentialkaos/ek.v6/usage"

	"pkg.re/essentialkaos/go-simpleyaml.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "Yo"
	VER  = "0.0.2"
	DESC = "Command-line YAML processor"
)

const (
	ARG_FROM_FILE = "f:from-file"
	ARG_NO_COLOR  = "nc:no-color"
	ARG_HELP      = "h:help"
	ARG_VER       = "v:version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Query struct {
	Tokens     []Token
	Processors []string
}

type Token struct {
	Key   string
	Index []int
	Range Range
}

type Range struct {
	Start int
	End   int
}

// ////////////////////////////////////////////////////////////////////////////////// //

var argMap = arg.Map{
	ARG_FROM_FILE: {Type: arg.STRING},
	ARG_NO_COLOR:  {Type: arg.BOOL},
	ARG_HELP:      {Type: arg.BOOL},
	ARG_VER:       {Type: arg.BOOL},
}

// ////////////////////////////////////////////////////////////////////////////////// //

func main() {
	args, errs := arg.Parse(argMap)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	configureUI()

	if arg.GetB(ARG_VER) {
		showAbout()
		os.Exit(1)
	}

	if arg.GetB(ARG_HELP) {
		showUsage()
		os.Exit(1)
	}

	if len(args) == 0 && !arg.Has(ARG_FROM_FILE) {
		showUsage()
		os.Exit(1)
	}

	process(strings.Join(args, " "))
}

// configureUI configure user interface
func configureUI() {
	envVars := env.Get()
	term := envVars.GetS("TERM")

	fmtc.DisableColors = true

	if term != "" {
		switch {
		case strings.Contains(term, "xterm"),
			strings.Contains(term, "color"),
			term == "screen":
			fmtc.DisableColors = false
		}
	}

	if arg.GetB(ARG_NO_COLOR) {
		fmtc.DisableColors = true
	}

	if !fsutil.IsCharacterDevice("/dev/stdout") && envVars.GetS("FAKETTY") == "" {
		fmtc.DisableColors = true
	}
}

// readData reads data from standart input or file
func readData() ([]byte, error) {
	if arg.Has(ARG_FROM_FILE) {
		return readFromFile(arg.GetS(ARG_FROM_FILE))
	}

	return readFromStdin()
}

// readFromFile read data from file
func readFromFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

// readFromStdin read data from standart input
func readFromStdin() ([]byte, error) {
	return ioutil.ReadFile("/dev/stdin")
}

// process start data processing
func process(query string) {
	data, err := readData()

	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	yaml, err := simpleyaml.NewYaml(data)

	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}

	execQuery(yaml, query)
}

// execQuery execute query over YAML
func execQuery(yaml *simpleyaml.Yaml, query string) {
	var data []*simpleyaml.Yaml

	for _, q := range parseQuery(query) {
		data = []*simpleyaml.Yaml{yaml}

		for _, t := range q.Tokens {
			if len(data) == 0 {
				break
			}

			if t.IsArrayToken() || data[0].IsArray() || data[0].Get(t.Key).IsArray() {
				data = execArrayTokenSelector(t, data)
			} else {
				data = execBasicTokenSelector(t, data)
			}
		}

		if len(q.Processors) == 0 {
			renderData(data)
		} else {
			processData(q.Processors, data)
		}
	}
}

// execArrayTokenSelector execute array query token over given data
func execArrayTokenSelector(t Token, data []*simpleyaml.Yaml) []*simpleyaml.Yaml {
	var result []*simpleyaml.Yaml

	if len(t.Index) != 0 {
		for _, item := range data {
			for _, index := range t.Index {
				if t.Key == "" {
					if item.IsIndexExist(index) {
						result = append(result, item.GetIndex(index))
					}
				} else {
					if item.Get(t.Key).IsIndexExist(index) {
						result = append(result, item.Get(t.Key).GetIndex(index))
					}
				}
			}
		}
	} else {
		for _, item := range data {
			if t.Range.Start == -1 && t.Range.End == -1 {
				if item.IsExist(t.Key) {
					result = append(result, item.Get(t.Key))
				}

				continue
			}

		RANGELOOP:
			for index := t.Range.Start; index < t.Range.End; index++ {
				if t.Key == "" {
					if item.IsIndexExist(index) {
						result = append(result, item.GetIndex(index))
					} else {
						break RANGELOOP
					}
				} else {
					if item.Get(t.Key).IsIndexExist(index) {
						result = append(result, item.Get(t.Key).GetIndex(index))
					} else {
						break RANGELOOP
					}
				}
			}
		}
	}

	return result
}

// execBasicTokenSelector execute basic query token over given data
func execBasicTokenSelector(t Token, data []*simpleyaml.Yaml) []*simpleyaml.Yaml {
	var result []*simpleyaml.Yaml

	for _, item := range data {
		if item.IsExist(t.Key) {
			result = append(result, item.Get(t.Key))
		}
	}

	return result
}

// renderData render yaml structs to string
func renderData(data []*simpleyaml.Yaml) {
	for _, item := range data {
		switch {
		case item.IsArray():
			if item.GetIndex(0).IsMap() || item.GetIndex(0).IsArray() {
				encodeYaml(item)
			} else {
				fmt.Println(strings.Join(item.MustStringArray(nil), "\n"))
			}

		case item.IsMap():
			encodeYaml(item)

		default:
			fmt.Printf("%v\n", item.Interface())
		}
	}
}

// processData run processors over given data
func processData(processor []string, data []*simpleyaml.Yaml) {
	var result interface{}

	for _, pf := range processor {
		switch pf {
		case "len", "length":
			result = processorFuncLength(data, result)
		case "keys":
			result = processorFuncKeys(data, result)
		case "sort":
			result = processorFuncSort(result)
		default:
			printError("Unknown function \"%s\"", pf)
			os.Exit(1)
		}
	}

	switch result.(type) {
	case string, int:
		fmt.Println(result)
	case []int:
		for _, v := range result.([]int) {
			fmt.Println(v)
		}
	case []string:
		for _, v := range result.([]string) {
			fmt.Println(v)
		}
	}
}

// processorFuncLength is length processor
func processorFuncLength(data []*simpleyaml.Yaml, k interface{}) []int {
	var result []int

	if k == nil {
		for _, item := range data {
			switch {
			case item.IsArray():
				result = append(result, item.ArraySize())
			case item.IsMap():
				result = append(result, len(item.MustMap(nil)))
			default:
				result = append(result, len(item.MustString("")))
			}
		}
	} else {
		switch k.(type) {
		case string:
			return []int{len(k.(string))}
		case []string:
			return []int{len(k.([]string))}
		}
	}

	return result
}

// processorFuncKeys is keys processor
func processorFuncKeys(data []*simpleyaml.Yaml, k interface{}) []string {
	var result []string

	if k != nil {
		return nil
	}

	for _, item := range data {
		if item.IsMap() {
			keys, _ := item.GetMapKeys()
			result = append(result, keys...)
		}
	}

	return result
}

// processorFuncKeys is sort processor
func processorFuncSort(k interface{}) []string {
	var result []string

	switch k.(type) {
	case string:
		result = []string{k.(string)}
	case []string:
		result = k.([]string)
		sort.Strings(result)
	}

	return result
}

// parseQuery parse query
func parseQuery(query string) []Query {
	var result []Query

	for _, q := range splitQuery(query) {
		result = append(result, parseSubQuery(q))
	}

	return result
}

// parseSubQuery parse sub-query
func parseSubQuery(query string) Query {
	query = strings.TrimSpace(query)

	if !strings.Contains(query, "|") {
		return Query{Tokens: parseTokens(query)}
	}

	qs := strings.Split(query, "|")

	if len(qs) < 2 {
		return Query{Tokens: parseTokens(qs[0])}
	}

	return Query{Tokens: parseTokens(qs[0]), Processors: parseProcessors(qs[1:])}
}

// parseTokens split query to tokens
func parseTokens(query string) []Token {
	query = strings.TrimSpace(query)

	var result []Token

	for i, t := range strings.Split(query, ".") {
		if i == 0 || t == "" {
			continue
		}

		result = append(result, parseToken(t))
	}

	return result
}

// parseToken parse token
func parseToken(token string) Token {
	if strings.Contains(token, "[") && strings.Contains(token, "]") {
		is := strings.Index(token, "[")
		return parseArrayToken(token[:is], token[is:])
	}

	return Token{Key: token, Range: Range{-1, -1}}
}

// parseArrayToken parse array token
func parseArrayToken(key, index string) Token {
	if index == "[]" {
		return Token{Key: key, Range: Range{0, 999999999}}
	}

	index = strings.TrimLeft(index, "[")
	index = strings.TrimRight(index, "]")

	if strings.Contains(index, ":") {
		is := strings.Split(index, ":")

		return Token{
			Key: key,
			Range: Range{
				str2int(is[0], 0),
				str2int(is[1], 999999999),
			},
		}
	} else if strings.Contains(index, ",") {
		return Token{Key: key, Range: Range{-1, -1}, Index: converEnum(strings.Split(index, ","))}
	} else {
		return Token{Key: key, Range: Range{-1, -1}, Index: []int{str2int(index, 0)}}
	}
}

// parseProcessors parse processors
func parseProcessors(processors []string) []string {
	var result []string

	for _, p := range processors {
		result = append(result, strings.TrimSpace(p))
	}

	return result
}

// splitQuery split query
func splitQuery(query string) []string {
	var result []string
	var buffer string
	var isArray bool

	for _, r := range query {
		switch r {
		case '[':
			isArray = true
		case ']':
			isArray = false
		}

		if r == ',' && !isArray {
			result = append(result, buffer)
			buffer = ""
			continue
		}

		buffer += string(r)
	}

	if buffer != "" {
		result = append(result, buffer)
	}

	return result
}

// converEnum convert string slice to int slice
func converEnum(s []string) []int {
	var result []int

	for _, i := range s {
		result = append(result, str2int(i, 0))
	}

	return result
}

// str2int convert string to int
func str2int(s string, def int) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return def
	}

	i, _ := strconv.Atoi(s)

	return i
}

// encodeYaml encode yaml struct to string
func encodeYaml(yaml *simpleyaml.Yaml) {
	data, _ := yaml.MarshalYAML()

	// Print encoded YAML without new line symbol
	fmt.Println(string(data[:len(data)-1]))
}

// printError prints error message to console
func printError(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{r}"+f+"{!}\n", a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsArrayToken return true if it array selector token
func (t Token) IsArrayToken() bool {
	if len(t.Index) != 0 || t.Range.Start != -1 || t.Range.End != -1 {
		return true
	}

	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //

func showUsage() {
	usage.Breadcrumbs = true

	info := usage.NewInfo("", "query")

	info.AddOption(ARG_FROM_FILE, "Read data from file", "filename")
	info.AddOption(ARG_NO_COLOR, "Disable colors in output")
	info.AddOption(ARG_HELP, "Show this help message")
	info.AddOption(ARG_VER, "Show version")

	info.AddExample("'.foo'", "Return value for key foo")
	info.AddExample("'.foo | length'", "Print value length")
	info.AddExample("'.foo[]'", "Return all items from array")
	info.AddExample("'.bar[2:]'", "Return subarray started from item with index 2")
	info.AddExample("'.bar[1,2,5]'", "Return items with index 1, 2 and 5 from array")
	info.AddExample("'.bar[] | length'", "Print array size")
	info.AddExample("'.xyz | keys'", "Print hash map keys")
	info.AddExample("'.xyz | keys | length'", "Print number of hash map keys")

	info.Render()
}

func showAbout() {
	about := &usage.About{
		App:     APP,
		Version: VER,
		Desc:    DESC,
		Year:    2006,
		Owner:   "ESSENTIAL KAOS",
		License: "Essential Kaos Open Source License <https://essentialkaos.com/ekol>",
	}

	about.Render()
}
