package datagen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type TextData struct {
	MinSize int
	MaxSize int
	Count   int
}

type MarkerOptions struct {
	BlockBegin, BlockEnd     string
	OptionsBegin, OptionsEnd string
	ElementBegin, ElementEnd string
	Separator                string
	LastSeparator            string
}

var DEFAULT = MarkerOptions{
	"{{{", "}}}",
	"[[[", "]]]",
	"{{", "}}",
	"\n", //Separator
	"\n", //Last Separator
}

var CSV = MarkerOptions{
	"{{", "}}",
	"[", "]",
	"{", "}",
	",",  //Separator
	"\n", //Last Separator
}

var XML = MarkerOptions{
	"{{", "}}",
	"[[", "]]",
	"{", "}",
	"\n", //Separator
	"\n", //Last Separator
}

var DOLLAR = MarkerOptions{
	"$(", ")$",
	"$[", "]$",
	"${", "}$",
	" ",  //Separator
	"\n", //Last Separator
}

var smallLetters = "abcdefghijklmnopqrstuvwxyz"
var capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var letters = smallLetters + capitalLetters
var numbers = "0123456789"
var specialChars = "~!@#$%^&*()_-+=<>,./;:'\"[]{}\\|"

func TextGen(ds TextData) []string {
	var a []string
	for c := 0; c < ds.Count; c++ {
		str := ""
		size := ds.MinSize + rand.Intn(ds.MaxSize-ds.MinSize+1)
		for i := 0; i < size; i++ {
			pos := rand.Intn(len(letters))
			str = str + string(letters[pos])
		}
		a = append(a, str)
	}
	return a
}

func GetFileData(fnames []string, regex string, random bool, count int) ([]string, error) {

	sequential := !random

	r := new(regexp.Regexp)
	if regex != "" {
		var err error
		r, err = regexp.Compile(regex)
		if err != nil {
			return nil, err
		}
	}

	var bs [][]byte

	//open file and read contents as an array
	//combine multiple files
	for _, fname := range fnames {
		b, err := ioutil.ReadFile(fname)
		if err != nil {
			return nil, err
		}
		tmpBs := bytes.Split(b, []byte{'\n'})
		if len(tmpBs) > 0 {
			bs = append(bs, tmpBs...)
		}
	}

	//see if regex is non empty
	if regex != "" {
		//else filter array 
		for i := len(bs) - 1; i >= 0; i-- {
			if !r.Match(bs[i]) {
				bs = append(bs[:i], bs[i+1:]...)
			}
		}
	}

	// if len(allLines) == 0 {
	if len(bs) == 0 {
		return nil, nil
	}

	//if sequential, then take first "count" items
	var retLines []string
	if sequential {
		if count > len(bs) {
			count = len(bs)
		}

		for i := 0; i < count; i++ {
			retLines = append(retLines, string(bs[i]))
		}
		return retLines, nil
	}

	//if random, then loop for "count" items randomly
	if random {
		for i := 0; i < count; i++ {
			pos := rand.Intn(len(bs))
			retLines = append(retLines, string(bs[pos]))
		}
		return retLines, nil
	}

	return nil, nil
}

func getOptionsMap(s string) map[string]string {
	// func getOptionsMap(s string, in interface{}) error {
	parts := strings.Split(s, "|")
	mParts := make(map[string]string)
	for _, onePart := range parts {
		onePart = strings.TrimSpace(onePart)
		subParts := strings.Split(onePart, ":")
		key := strings.ToLower(strings.TrimSpace(subParts[0]))
		val := ""
		if len(subParts) > 1 { // part following : exists
			val = strings.TrimSpace(subParts[1])
		}
		mParts[key] = val
	}

	return mParts
}

func setOptions(mParts map[string]string, in interface{}) error {

	v := reflect.ValueOf(in).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		// fmt.Println(v.Field(i).String(), " ; ", v.Field(i).Kind())
		// fmt.Println(t.Field(i).Name)
		fieldName := t.Field(i).Name
		// fieldVal := ""
		for key, val := range mParts {
			if strings.ToLower(fieldName) == key {
				// fieldVal = val
				switch v.Field(i).Kind() {
				case reflect.Int:
					//convert fieldVal to int and then assign it
					tmpInt, err := strconv.ParseInt(val, 0, 32)
					if err != nil {
						return err
					}
					v.Field(i).SetInt(tmpInt)
				case reflect.Bool:
					v.Field(i).SetBool(true)
				case reflect.String:
					//if val contains quotes, remove it
					val = strings.TrimLeft(val, "\"'")
					val = strings.TrimRight(val, "\"'")
					//TODO: have to find a proper way to do \n type chars
					v.Field(i).SetString(val)
				}
			}
		}
	}

	return nil
}

var mFiles = map[string][]string{
	"country":   []string{"country.txt"},
	"firstname": []string{"firstname.txt"},
}

func GenFileElement(fnames []string, mParts map[string]string, count int) ([]string, error) {
	// func City(s string) ([]string, error) {

	opts := struct {
		Regex  string
		Random bool
		// Count  int
	}{
		"",
		false,
		// 1,
	}

	// err := getOptionsMap(s, &opts)
	err := setOptions(mParts, &opts)
	if err != nil {
		return nil, err
	}

	return GetFileData(fnames, opts.Regex, opts.Random, count)
}

// Generate string data for a single element
// city/firstname | regex: | random 
func GenElement(eb string, count int) ([]string, error) {
	mOpts := getOptionsMap(eb)
	for key, val := range mOpts {
		if val == "" { //these won't have a val part.  TODO: Anyways, not foolproof.
			switch key {
			case "country":
				// return City(mOpts, count)
				return GenFileElement(mFiles["country"], mOpts, count)
			case "firstname":
				// return FirstName(mOpts, count)
				return GenFileElement(mFiles["firstname"], mOpts, count)
			default:
				return nil, fmt.Errorf("Unknown element type: %s", key)
			}
		}
	}
	return nil, nil
}

type blockOptions struct {
	Count         int
	Separator     string
	LastSeparator string
	ElementBegin  string
	ElementEnd    string
}

// Parses options string to give back options.  Input string includes the enclosing begin and end separators for the options block 
func getBlockOptions(s string, mo MarkerOptions) (*blockOptions, error) {
	//set defaults
	// bo := blockOptions{
	// 	Count:         1,
	// 	Separator:     "\n",
	// 	LastSeparator: "\n",
	// 	ElementBegin:  "{{",
	// 	ElementEnd:    "}}",
	// }

	// the default values are first set
	bo := blockOptions{
		Count:         1,
		Separator:     mo.Separator,
		LastSeparator: mo.LastSeparator,
		ElementBegin:  mo.ElementBegin,
		ElementEnd:    mo.ElementEnd,
	}

	// if any other options are specified, then the defaults are overridden
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return &bo, nil
	}

	if len(mo.OptionsBegin) > 0 {
		s = strings.TrimPrefix(s, mo.OptionsBegin)
	}
	if len(mo.OptionsEnd) > 0 {
		s = strings.TrimSuffix(s, mo.OptionsEnd)
	}
	s = strings.TrimSpace(s)

	mOpts := getOptionsMap(s)

	//TODO: should add a report extra params as error option in setOptions
	if err := setOptions(mOpts, &bo); err != nil {
		return nil, err
	}

	return &bo, nil
}

type subBlock struct {
	block string
	start int
	end   int
}

/*
func getSubBlocks(s string, mo.BlockBegin, mo.BlockEnd string) []subBlock {

	var blocks []subBlock
	// see if data block has further config blocks
	subBegin := strings.Index(s, mo.BlockBegin)
	// subEnd := -1
	if subBegin >= 0 {
		nxtmo.BlockEnd := strings.Index(s, mo.BlockEnd)
		nxtmo.BlockBegin := strings.IndexAny(s, mo.BlockBegin)

		if nxtmo.BlockEnd == -1 {
			panic("No matching block end within: " + s)
		}

		// if yes, extract block part and call GenBlock again
		if nxtmo.BlockBegin == -1 || nxtmo.BlockBegin > nxtmo.BlockEnd {
			subS := s[subBegin : nxtmo.BlockEnd+len(mo.BlockEnd)]
			blocks = append(blocks, subBlock{block: subS, start: subBegin, end: nxtmo.BlockEnd + len(mo.BlockEnd)})
		}
	}

	return blocks
}
*/

/*
func getSubBlock(s string, mo.BlockBegin, mo.BlockEnd string) subBlock {
	// var blocks []subBlock

	regStr := "(?s)" + mo.BlockBegin + ".*?" + mo.BlockEnd
	re := regexp.MustCompile(regStr)

	subS := ""
	matchBegin := -1
	matchEnd := -1
	newS := s
	for {
		indexes := re.FindStringIndex(newS)
		if len(indexes) == 0 {
			break //reached inner most
		}
		if len(indexes) > 0 {
			subS = newS[indexes[0]:indexes[1]]
			matchBegin = matchBegin + indexes[0]
			matchEnd = matchBegin + indexes[1]
			newS = newS[indexes[0]+len(mo.BlockBegin):]
		}
	}

	return subBlock{subS, matchBegin, matchEnd}
}
*/

//return innermost block with BlockBegin and BlockEnd markers
func getSubBlock(s string, mo MarkerOptions) subBlock {

	//find the first occurrence of ending marker
	end := strings.Index(s, mo.BlockEnd)
	if end < 0 {
		return subBlock{}
	}

	//from the first end marker, search backwards to find the first beginning marker
	partS := s[:end+len(mo.BlockEnd)]
	begin := strings.LastIndex(partS, mo.BlockBegin)
	if begin < 0 {
		panic("getSubBlock: end marker found but no matching beginning marker.")
	}

	return subBlock{s[begin : end+len(mo.BlockEnd)], begin, end}
}

func getSubBlockOuter(s string, beginMark, endMark string) subBlock {

	beginCtr := 0
	begin := strings.Index(s, beginMark)
	end := begin + len(beginMark)
	if begin < 0 {
		return subBlock{}
	} else {
		beginCtr = beginCtr + 1
		newS := s[begin+len(beginMark):]
		for beginCtr != 0 {
			// fmt.Println("newS: ", newS)
			// fmt.Println("beginCtr: ", beginCtr)
			// fmt.Println("begin: ", begin)
			// fmt.Println("end: ", end)
			aBegin := strings.Index(newS, beginMark)
			aEnd := strings.Index(newS, endMark)
			// fmt.Println("aBegin: ", aBegin)
			// fmt.Println("aEnd: ", aEnd)
			if aBegin == -1 && aEnd == -1 {
				panic("No matching ending marker: " + endMark)
			}
			if (aBegin == -1 && aEnd > 0) || (aEnd < aBegin) { //this is an innermost match
				beginCtr = beginCtr - 1
				newS = newS[aEnd+len(endMark):]
				end = end + aEnd + len(endMark)
			}
			if aEnd > aBegin && aBegin != -1 { //have to go a level deeper
				beginCtr = beginCtr + 1
				newS = newS[aBegin+len(beginMark):]
				end = end + aBegin + len(beginMark)
			}
		}
	}

	return subBlock{s[begin:end], begin, end}
}

//Generate data for a datagen block
// func GenBlockX(s, mo.BlockBegin, mo.BlockEnd, mo.OptionsBegin, mo.OptionsEnd string) (string, error) {
func GenBlockX(s string, mo MarkerOptions) (string, error) {

	fmt.Println("Given string: " + s)
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "", nil
	}

	//remove enclosing {{{ and }}}, or whatever is prefix and suffix of the entire block
	if len(mo.BlockBegin) > 0 {
		s = strings.TrimPrefix(s, mo.BlockBegin)
	}
	if len(mo.BlockEnd) > 0 {
		s = strings.TrimSuffix(s, mo.BlockEnd)
	}
	s = strings.TrimSpace(s)
	fmt.Println("Given string after trimming: " + s)

	//check if there are further sub blocks
	fmt.Println("begin s:" + s)
	sub := getSubBlock(s, mo)
	fmt.Println("subBlock:" + sub.block)
	if sub.block != "" {
		genSub, err := GenBlock(sub.block)
		if err != nil {
			return "", err
		}
		fmt.Println("genSub:" + genSub)

		// dataS = dataS[:sub.start] + genSub + dataS[sub.end:]
		// s = s[:sub.start] + genSub + s[sub.end+len(mo.BlockEnd):]
		s = strings.TrimSpace(s[:sub.start]) + genSub + strings.TrimSpace(s[sub.end+len(mo.BlockEnd):])
		fmt.Println("next s: " + s)
	}

	fmt.Println("After recursive calls s:" + s)

	optBeginPos := strings.Index(s, mo.OptionsBegin)
	optEndPos := strings.Index(s, mo.OptionsEnd)
	if optBeginPos >= 0 && optEndPos < 0 {
		return "", fmt.Errorf("Beginning of options marker found (" + mo.OptionsBegin + ") but did not find end marker (" + mo.OptionsEnd + ").")
	} else if optBeginPos < 0 && optEndPos >= 0 {
		return "", fmt.Errorf("End of options marker found (" + mo.OptionsEnd + ") but did not find beginning marker (" + mo.OptionsBegin + ").")
	} else if optBeginPos > optEndPos {
		return "", fmt.Errorf("Position of beginning of options marker (" + mo.OptionsBegin + ") is after position of end of options marker (" + mo.OptionsEnd + ").")
	}

	optionsS := ""
	if optEndPos > optBeginPos {
		optionsS = strings.TrimSpace(s[optBeginPos+len(mo.OptionsBegin) : optEndPos])
	}

	bo, err := getBlockOptions(optionsS, mo)
	if err != nil {
		return "", err
	}

	dataS := s
	if optEndPos > 0 {

		dataS = strings.TrimSpace(s[:optBeginPos] + s[optEndPos+len(mo.OptionsEnd):])
	}
	fmt.Println("dataS begin:" + dataS)

	// else read individual element parts
	mElements := make(map[string]string)
	markerCnt := 0
	for {
		elBeginPos := strings.Index(dataS, bo.ElementBegin)
		elEndPos := strings.Index(dataS, bo.ElementEnd)
		if elBeginPos < 0 || elEndPos < 0 {
			break
		}

		//assuming here that it is syntatically ok.  TODO: fix this.

		//replace element definition with a marker
		nxtMarker := "<$" + strconv.FormatInt(int64(markerCnt), 10) + "$>"
		elementS := dataS[elBeginPos+len(bo.ElementBegin) : elEndPos]
		mElements[nxtMarker] = elementS
		dataS = dataS[:elBeginPos] + nxtMarker + dataS[elEndPos+len(bo.ElementEnd):]
		markerCnt = markerCnt + 1
	}

	mGenElements := make(map[string][]string)
	// for each element, call GenElement with count 
	for marker, elDef := range mElements {
		data, err := GenElement(elDef, bo.Count)
		if err != nil {
			return "", err
		}
		mGenElements[marker] = data
	}
	// fmt.Printf("mGenElements: %+v \n", mGenElements)

	//substitue data block with strings from GenElements
	fullS := ""
	for i := 0; i < bo.Count; i++ {
		tmpS := dataS
		// fmt.Printf("tmpS %+v \n", tmpS)
		for marker, data := range mGenElements {
			tmpS = strings.Replace(tmpS, marker, data[i], 1)
		}

		if i == bo.Count-1 {
			tmpS += bo.LastSeparator
		} else {
			tmpS += bo.Separator
		}
		fullS = fullS + tmpS
	}

	fmt.Println("fullS: " + fullS)
	return fullS, nil
}

func GenBlock(s string) (string, error) {
	// return GenBlockX(s, "{{{", "}}}", "[[[", "]]]")
	return GenBlockX(s, DEFAULT)
}

//Generate data for an entire input.  To be called recursively.
func Gen(s string, mo MarkerOptions) (string, error) {
	for {
		sub := getSubBlockOuter(s, mo.BlockBegin, mo.BlockEnd)
		if sub.block == "" {
			break
		}
		
		gen, err := GenBlockX(sub.block, mo)
		if err != nil {
			return "", err
		}
		s = s[:sub.start] + gen + s[sub.end:]
	}
	return s, nil
}
