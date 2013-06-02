package datagen

import (
	"sort"
	"testing"
)

func Test_TextGen_1(t *testing.T) {
	// Gen("Text | MinSize:10 | MaxSize:30 | Count:5")
	td := TextData{
		MinSize: 10,
		MaxSize: 15,
		Count:   5,
	}
	a := TextGen(td)
	if len(a) != td.Count {
		t.Errorf("Err: Expected 5 random strings.  Received %d.", len(a))
	} else {
		t.Logf("OK: Expected 5 random strings.  Received %v.", a)
	}

	for i := 0; i < len(a); i++ {
		if len(a[i]) < td.MinSize {
			t.Errorf("Expected string of minimum size %d. Received size was %d.", td.MinSize, len(a[i]))
		} else {
			t.Logf("Expected string of minimum size %d. Received %v of size %d.", td.MinSize, a[i], len(a[i]))
		}

		if len(a[i]) > td.MaxSize {
			t.Errorf("Expected string of maximum size %d. Received size was %d.", td.MaxSize, len(a[i]))
		} else {
			t.Logf("Expected string of maximum size %d. Received %v of size %d.", td.MaxSize, a[i], len(a[i]))
		}
	}
}

func Test_GetFileData_Seq(t *testing.T) {
	s, err := GetFileData([]string{"test_data_file.txt"}, "", false, 5)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 5 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 5, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 5, len(s))
	}

	exp := []string{
		"Afghanistan",
		"Albania",
		"Algeria",
		"Andorra",
		"Angola",
	}

	for i := 0; i < 5; i++ {
		if exp[i] != s[i] {
			t.Errorf("Expected %s at %d, but found %s", exp[i], i, s[i])
		} else {
			t.Logf("Expected %s at %d, and found %s", exp[i], i, s[i])
		}
	}

}

func Test_GetFileData_Regex(t *testing.T) {
	s, err := GetFileData([]string{"test_data_file.txt"}, "^B.*", false, 5)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 5 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 5, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 5, len(s))
	}

	exp := []string{
		"Bahamas",
		"Bahrain",
		"Bangladesh",
		"Barbados",
		"Belarus",
	}

	for i := 0; i < 5; i++ {
		if exp[i] != s[i] {
			t.Errorf("Expected %s at %d, but found %s", exp[i], i, s[i])
		} else {
			t.Logf("Expected %s at %d, and found %s", exp[i], i, s[i])
		}
	}

}

func Test_GetFileData_Random(t *testing.T) {
	s, err := GetFileData([]string{"test_data_file.txt"}, "", true, 5)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 5 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 5, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 5, len(s))
	}

	//the chances that it will be in sorted order are very very low. But not impossible.
	if sort.StringsAreSorted(s) {
		t.Errorf("Chances that random array is sorted is very very low.  But they are here: %v", s)
	}
	t.Logf("Received values: %v", s)
}

func Test_Country(t *testing.T) {

	s, err := GenElement("country", 1)

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 1 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 1, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 1, len(s))
	}

	exp := []string{
		"Afghanistan",
	}

	for i := 0; i < 1; i++ {
		if exp[i] != s[i] {
			t.Errorf("Expected %s at %d, but found %s", exp[i], i, s[i])
		} else {
			t.Logf("Expected %s at %d, and found %s", exp[i], i, s[i])
		}
	}
}

func Test_Country_Regex(t *testing.T) {

	s, err := GenElement("country | reGEX: ^C.* ", 6) //case shouldn't matter

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 6 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 6, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 6, len(s))
	}

	exp := []string{
		"Cambodia",
		"Cameroon",
		"Canada",
		"Cape Verde",
		"Central African Rep",
		"Chad",
		"Chile",
	}

	for i := 0; i < 6; i++ {
		if exp[i] != s[i] {
			t.Errorf("Expected %s at %d, but found %s", exp[i], i, s[i])
		} else {
			t.Logf("Expected %s at %d, and found %s", exp[i], i, s[i])
		}
	}
}

func Test_Country_Random(t *testing.T) {
	s, err := GenElement("Country | regex: ^C.* | RanDOM", 5) //case shouldn't matter

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 5 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 5, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 5, len(s))
	}

	//the chances that it will be in sorted order are very very low. But not impossible.
	if sort.StringsAreSorted(s) {
		t.Errorf("Chances that random array is sorted is very very low.  But they are here: %v", s)
	}
	t.Logf("Received values: %v", s)
}

func Test_FirstName(t *testing.T) {

	s, err := GenElement("firstname", 1)

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 1 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 1, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 1, len(s))
	}

	exp := []string{
		"AARON",
	}

	for i := 0; i < 1; i++ {
		if exp[i] != s[i] {
			t.Errorf("Expected %s at %d, but found %s", exp[i], i, s[i])
		} else {
			t.Logf("Expected %s at %d, and found %s", exp[i], i, s[i])
		}
	}
}

func Test_FirstName_Regex(t *testing.T) {

	s, err := GenElement("firstname | reGEX: ^D[aA].* ", 6) //case shouldn't matter

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 6 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 6, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 6, len(s))
	}

	exp := []string{
		"DALE",
		"DALLAS",
		"DALTON",
		"DAMIAN",
		"DAMIEN",
		"DAMION",
	}

	for i := 0; i < 6; i++ {
		if exp[i] != s[i] {
			t.Errorf("Expected %s at %d, but found %s", exp[i], i, s[i])
		} else {
			t.Logf("Expected %s at %d, and found %s", exp[i], i, s[i])
		}
	}
}

func Test_FirstName_Random(t *testing.T) {
	s, err := GenElement("FirstName | regex: ^C.* | RanDOM", 5) //case shouldn't matter

	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if len(s) != 5 {
		t.Errorf("Expected string array of size %d. Received size was %d.", 5, len(s))
	} else {
		t.Logf("Expected string array of size %d. Received size was %d.", 5, len(s))
	}

	//the chances that it will be in sorted order are very very low. But not impossible.
	if sort.StringsAreSorted(s) {
		t.Errorf("Chances that random array is sorted is very very low.  But they are here: %v", s)
	}
	t.Logf("Received values: %v", s)
}

func Test_GetBlockOptions(t *testing.T) {
	bo, err := getBlockOptions("", DEFAULT)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	s := `
			
	`
	expBO := blockOptions{
		Count:         1,
		Separator:     "\n",
		LastSeparator: "\n",
		ElementBegin:  "{{",
		ElementEnd:    "}}",
	}

	if bo, err = getBlockOptions(s, DEFAULT); err != nil || *bo != expBO {
		t.Errorf("Expected %+v. Received %+v.", expBO, *bo)
	} else {
		t.Logf("Expected %+v. Received %+v.", expBO, *bo)
	}

	s = ` [[[ count: 5 | elementBegin: {[{ | elementEnd: []} ]]]`
	expBO = blockOptions{
		Count:         5,
		Separator:     "\n",
		LastSeparator: "\n",
		ElementBegin:  "{[{",
		ElementEnd:    "[]}",
	}

	if bo, err = getBlockOptions(s, DEFAULT); err != nil || *bo != expBO {
		t.Errorf("Expected %+v. Received %+v.", expBO, *bo)
	} else {
		t.Logf("Expected %+v. Received %+v.", expBO, *bo)
	}

}

func Test_GenBlock(t *testing.T) {
	block := `{{{ [[[ count: 3 ]]] {{ firstname }} }}}`

	exp := `AARON
ABDUL
ABE
`

	// s, err := GenBlock(block, "{{{", "}}}", 
	s, err := GenBlock(block)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if s != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v", exp, s)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v", exp, s)
	}

	block =
		`{{{[[[ count:3| Separator: ",
" | LastSeparator: "
" ]]]{
"name": "{{firstname | regex:^C.* }}",
"location": "{{ country | regex:^C.*  }}"
} }}}`

	exp = `{
"name": "CALEB",
"location": "Cambodia"
},
{
"name": "CALVIN",
"location": "Cameroon"
},
{
"name": "CAMERON",
"location": "Canada"
}
`

	// s, err := GenBlock(block, "{{{", "}}}", 
	s, err = GenBlock(block)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		// return
	}

	if s != exp {
		t.Errorf("FAIL. Expected %+v. \nReceived %+v.", exp, s)
	} else {
		t.Logf("PASS. Expected %+v. \nReceived %+v.", exp, s)
	}
}

func Test_getSubBlocks(t *testing.T) {
	s := `
abcd
{{{ 1234 }}}
efgh
`

	sub := getSubBlock(s, DEFAULT)

	if sub.block != "{{{ 1234 }}}" {
		t.Errorf("FAIL. Expected %+v. \nReceived %+v.", "{{{ 1234 }}}", sub.block)
	} else {
		t.Logf("PASS. Expected %+v. \nReceived %+v.", "{{{ 1234 }}}", sub.block)
	}

	s = `
abcd
{{{ 1234 }}}
{{{ 5678 }}}
efgh
`

	sub = getSubBlock(s, DEFAULT)
	if sub.block != "{{{ 1234 }}}" {
		t.Errorf("FAIL. Expected %+v. \nReceived %+v.", "{{{ 1234 }}}", sub.block)
	} else {
		t.Logf("PASS. Expected %+v. \nReceived %+v.", "{{{ 1234 }}}", sub.block)
	}

	s = `
abcd
{{{ 1234 {{{ 5678 }}} }}}

efgh
`

	sub = getSubBlock(s, DEFAULT)
	if sub.block != "{{{ 5678 }}}" {
		t.Errorf("FAIL. Expected %+v. \nReceived %+v.", "{{{ 5678 }}}", sub.block)
	} else {
		t.Logf("PASS. Expected %+v. \nReceived %+v.", "{{{ 5678 }}}", sub.block)
	}

	s = `
abcd {{{ 1234 {{{5678 {{{ 9012 }}} }}} }}}

{{{ a1234 {{{b5678 {{{ c9012 }}} }}} }}} efgh
`

	sub = getSubBlock(s, DEFAULT)
	if sub.block != "{{{ 9012 }}}" {
		t.Errorf("FAIL. Expected %+v. \nReceived %+v.", "{{{ 9012 }}}", sub.block)
	} else {
		t.Logf("PASS. Expected %+v. \nReceived %+v.", "{{{ 9012 }}}", sub.block)
	}

	s = `
abcd
{{{ 1234 {{{ 
	5678

	{{{ 90
12 }}}

	 }}} 
	}}}

efgh
`

	sub = getSubBlock(s, DEFAULT)
	if sub.block != "{{{ 90\n12 }}}" {
		t.Errorf("FAIL. Expected %+v. \nReceived %+v.", "{{{ 90\n12 }}}", sub.block)
	} else {
		t.Logf("PASS. Expected %+v. \nReceived %+v.", "{{{ 90\n12 }}}", sub.block)
	}
}

func Test_GenBlock_2(t *testing.T) {

	block := `{{{ [[[ count: 2 ]]] {{{ [[[ count: 3 ]]] {{ country }} }}} {{ firstname }} }}}`

	exp := `Afghanistan
Albania
Algeria
AARON
Afghanistan
Albania
Algeria
ABDUL
`

	// s, err := GenBlock(block, "{{{", "}}}", 
	s, err := GenBlock(block)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if s != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v.", exp, s)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v.", exp, s)
	}
}

func Test_DOLLAR(t *testing.T) {
	block := "$( $[count:3 | separator: ',' | lastseparator: '' ]$ ${ country }$ )$"

	exp := `Afghanistan,Albania,Algeria`

	// s, err := GenBlock(block, "{{{", "}}}", 
	s, err := GenBlockX(block, DOLLAR)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if s != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v.", exp, s)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v.", exp, s)
	}

	block = "$( $[count:3]$ ${ country }$ )$"

	exp = `Afghanistan Albania Algeria
`

	// s, err := GenBlock(block, "{{{", "}}}", 
	s, err = GenBlockX(block, DOLLAR)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if s != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v.", exp, s)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v.", exp, s)
	}
}

func Test_CSV(t *testing.T) {
	block := "{{ [count:3 | separator: ',' | lastseparator: ';' ] { country } }}"

	exp := `Afghanistan,Albania,Algeria;`

	// s, err := GenBlock(block, "{{{", "}}}", 
	s, err := GenBlockX(block, CSV)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if s != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v.", exp, s)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v.", exp, s)
	}

	block = " {{ [count:3] { country } }}"

	exp = `Afghanistan,Albania,Algeria
`

	// s, err := GenBlock(block, "{{{", "}}}", 
	s, err = GenBlockX(block, CSV)
	if err != nil {
		t.Errorf("Unexpected error. %v", err)
		return
	}

	if s != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v.", exp, s)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v.", exp, s)
	}
}

func Test_getSubBlockOuter(t *testing.T) {

	block := " 12 {{ efgh {{ abcd}} }} 34 "

	exp := `{{ efgh {{ abcd}} }}`

	// s, err := GenBlock(block, "{{{", "}}}", 
	sub := getSubBlockOuter(block, "{{", "}}")

	if sub.block != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v.", exp, sub.block)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v.", exp, sub.block)
	}

	block = " 12 {{ efgh {{ {{ijkl}} {{ mnop }} abcd}} }} 3412 {{ 1efgh {{ 1abcd}} }} 3456 "

	exp = `{{ efgh {{ {{ijkl}} {{ mnop }} abcd}} }}`

	// s, err := GenBlock(block, "{{{", "}}}", 
	sub = getSubBlockOuter(block, "{{", "}}")

	if sub.block != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v.", exp, sub.block)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v.", exp, sub.block)
	}
}

func Test_Gen(t *testing.T) {

	block := " 12 $( ${country}$ ${firstname}$ )$ 34 "

	exp := ` 12 Afghanistan AARON
 34 `

	// s, err := GenBlock(block, "{{{", "}}}", 
	gen, err := Gen(block, DOLLAR)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if gen != exp {
		t.Errorf("FAIL. Expected %+v. Received %+v.", exp, gen)
	} else {
		t.Logf("PASS. Expected %+v. Received %+v.", exp, gen)
	}
}
