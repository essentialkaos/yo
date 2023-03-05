package cli

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/options"
	"github.com/essentialkaos/ek/v12/usage"
	"github.com/essentialkaos/ek/v12/usage/completion/bash"
	"github.com/essentialkaos/ek/v12/usage/completion/fish"
	"github.com/essentialkaos/ek/v12/usage/completion/zsh"
	"github.com/essentialkaos/ek/v12/usage/man"
	"github.com/essentialkaos/ek/v12/usage/update"

	"github.com/essentialkaos/go-simpleyaml/v2"

	"github.com/essentialkaos/yo/cli/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "Yo"
	VER  = "0.5.6"
	DESC = "Command-line YAML processor"
)

const (
	OPT_FROM_FILE = "f:from-file"
	OPT_NO_COLOR  = "nc:no-color"
	OPT_HELP      = "h:help"
	OPT_VER       = "v:version"

	OPT_VERB_VER     = "vv:verbose-version"
	OPT_COMPLETION   = "completion"
	OPT_GENERATE_MAN = "generate-man"
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

var optMap = options.Map{
	OPT_FROM_FILE: {Type: options.STRING},
	OPT_NO_COLOR:  {Type: options.BOOL},
	OPT_HELP:      {Type: options.BOOL},
	OPT_VER:       {Type: options.BOOL},

	OPT_VERB_VER:     {Type: options.BOOL},
	OPT_COMPLETION:   {},
	OPT_GENERATE_MAN: {Type: options.BOOL},
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Init is main function
func Init(gitRev string, gomod []byte) {
	preConfigureUI()

	args, errs := options.Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	configureUI()

	switch {
	case options.Has(OPT_COMPLETION):
		os.Exit(genCompletion())
	case options.Has(OPT_GENERATE_MAN):
		os.Exit(genMan())
	case options.GetB(OPT_VER):
		showAbout(gitRev)
		os.Exit(0)
	case options.GetB(OPT_VERB_VER):
		support.ShowSupportInfo(APP, VER, gitRev, gomod)
		os.Exit(0)
	case options.GetB(OPT_HELP),
		len(args) == 0 && !options.Has(OPT_FROM_FILE):
		showUsage()
		os.Exit(0)
	}

	process(strings.Join(args.Strings(), " "))
}

// preConfigureUI preconfigures UI based on information about user terminal
func preConfigureUI() {
	term := os.Getenv("TERM")

	fmtc.DisableColors = true

	if term != "" {
		switch {
		case strings.Contains(term, "xterm"),
			strings.Contains(term, "color"),
			term == "screen":
			fmtc.DisableColors = false
		}
	}

	if !fsutil.IsCharacterDevice("/dev/stdout") && os.Getenv("FAKETTY") == "" {
		fmtc.DisableColors = true
	}

	if os.Getenv("NO_COLOR") != "" {
		fmtc.DisableColors = true
	}
}

// configureUI configures user interface
func configureUI() {
	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}
}

// readData reads data from standart input or file
func readData() ([]byte, error) {
	if options.Has(OPT_FROM_FILE) {
		return readFromFile(options.GetS(OPT_FROM_FILE))
	}

	return readFromStdin()
}

// readFromFile reads data from file
func readFromFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

// readFromStdin reads data from standart input
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

// execQuery executes query over YAML
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

// execArrayTokenSelector executes array query token over given data
func execArrayTokenSelector(t Token, data []*simpleyaml.Yaml) []*simpleyaml.Yaml {
	var result []*simpleyaml.Yaml

	if len(t.Index) != 0 {
		for _, item := range data {
			for _, index := range t.Index {
				if t.Key == "" {
					if item.IsIndexExist(index) {
						result = append(result, item.GetByIndex(index))
					}
				} else {
					if item.Get(t.Key).IsIndexExist(index) {
						result = append(result, item.Get(t.Key).GetByIndex(index))
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
						result = append(result, item.GetByIndex(index))
					} else {
						break RANGELOOP
					}
				} else {
					if item.Get(t.Key).IsIndexExist(index) {
						result = append(result, item.Get(t.Key).GetByIndex(index))
					} else {
						break RANGELOOP
					}
				}
			}
		}
	}

	return result
}

// execBasicTokenSelector executes basic query token over given data
func execBasicTokenSelector(t Token, data []*simpleyaml.Yaml) []*simpleyaml.Yaml {
	var result []*simpleyaml.Yaml

	for _, item := range data {
		if item.IsExist(t.Key) {
			result = append(result, item.Get(t.Key))
		}
	}

	return result
}

// renderData renders yaml structs to string
func renderData(data []*simpleyaml.Yaml) {
	for _, item := range data {
		switch {
		case item.IsArray():
			if item.GetByIndex(0).IsMap() || item.GetByIndex(0).IsArray() {
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

// processData runs processors over given data
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

// processorFuncLength is a length processor
func processorFuncLength(data []*simpleyaml.Yaml, k interface{}) []int {
	var result []int

	if k == nil {
		for _, item := range data {
			switch {
			case item.IsArray():
				result = append(result, len(item.MustArray(nil)))
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

// processorFuncKeys is a keys processor
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

// processorFuncKeys is a sort processor
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

// parseQuery parses query
func parseQuery(query string) []Query {
	var result []Query

	for _, q := range splitQuery(query) {
		result = append(result, parseSubQuery(q))
	}

	return result
}

// parseSubQuery parses sub-query
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

// parseTokens splits query to tokens
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

// parseToken parses token
func parseToken(token string) Token {
	if strings.Contains(token, "[") && strings.Contains(token, "]") {
		is := strings.Index(token, "[")
		return parseArrayToken(token[:is], token[is:])
	}

	return Token{Key: token, Range: Range{-1, -1}}
}

// parseArrayToken parses array token
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

// parseProcessors parses processors
func parseProcessors(processors []string) []string {
	var result []string

	for _, p := range processors {
		result = append(result, strings.TrimSpace(p))
	}

	return result
}

// splitQuery splits query
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

// converEnum converts string slice to int slice
func converEnum(s []string) []int {
	var result []int

	for _, i := range s {
		result = append(result, str2int(i, 0))
	}

	return result
}

// str2int converts string to int
func str2int(s string, def int) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return def
	}

	i, _ := strconv.Atoi(s)

	return i
}

// encodeYaml encodes yaml struct to string
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

// IsArrayToken returns true if it array selector token
func (t Token) IsArrayToken() bool {
	if len(t.Index) != 0 || t.Range.Start != -1 || t.Range.End != -1 {
		return true
	}

	return false
}

// ////////////////////////////////////////////////////////////////////////////////// //

// showUsage prints usage info
func showUsage() {
	genUsage().Render()
}

// showAbout prints info about version
func showAbout(gitRev string) {
	genAbout(gitRev).Render()
}

// genCompletion generates completion for different shells
func genCompletion() int {
	info := genUsage()

	switch options.GetS(OPT_COMPLETION) {
	case "bash":
		fmt.Printf(bash.Generate(info, "yo"))
	case "fish":
		fmt.Printf(fish.Generate(info, "yo"))
	case "zsh":
		fmt.Printf(zsh.Generate(info, optMap, "yo"))
	default:
		return 1
	}

	return 0
}

// genMan generates man page
func genMan() int {
	fmt.Println(
		man.Generate(
			genUsage(),
			genAbout(""),
		),
	)

	return 0
}

// genUsage generates usage info
func genUsage() *usage.Info {
	info := usage.NewInfo("", "query")

	info.AddOption(OPT_FROM_FILE, "Read data from file", "filename")
	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show version")

	info.AddRawExample("cat file.yml | yo '.foo'", "Return value for key foo")
	info.AddExample("-f file.yml '.foo'", "Return value for key foo")
	info.AddExample("-f file.yml '.foo | length'", "Print value length")
	info.AddExample("-f file.yml '.foo[]'", "Return all items from array")
	info.AddExample("-f file.yml '.bar[2:]'", "Return subarray started from item with index 2")
	info.AddExample("-f file.yml '.bar[1,2,5]'", "Return items with index 1, 2 and 5 from array")
	info.AddExample("-f file.yml '.bar[] | length'", "Print array size")
	info.AddExample("-f file.yml '.xyz | keys'", "Print hash map keys")
	info.AddExample("-f file.yml '.xyz | keys | length'", "Print number of hash map keys")
	info.AddExample("-f file.yml '.xyz | keys | sort'", "Print sorted list of keys")

	return info
}

// genAbout generates info about version
func genAbout(gitRev string) *usage.About {
	about := &usage.About{
		App:           APP,
		Version:       VER,
		Desc:          DESC,
		Year:          2006,
		Owner:         "ESSENTIAL KAOS",
		License:       "Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>",
		UpdateChecker: usage.UpdateChecker{"essentialkaos/yo", update.GitHubChecker},
	}

	if gitRev != "" {
		about.Build = "git:" + gitRev
	}

	return about
}