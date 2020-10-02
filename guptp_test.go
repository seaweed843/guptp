/*MIT License Copyright (c) 2020 seaweed843

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
==============================================================================*/

package guptp

import (
	"testing"
	"reflect"
)

func TestParseUriPathToMapStr(t *testing.T) {
	var want = map[string]string{
		"Op": "create",
		"Id":"843",
		"Name":"Sea Weed",
	}
	uri :=`/api/v1/create/843/Sea%20Weed`
	template := `/api/v1/{Op}/{Id}/{Name}`
	got := ParseUriPathToMapStr(&uri, &template)
	eq := reflect.DeepEqual(want, got)
	if !eq {
    	t.Errorf("ParseUriPathToMapStr() = %q, want %q", got, want)
	}
}

func TestParseUriPathToFields(t *testing.T) {
	want := struct{Op string; Id int; Name string}{"create", 843, "Sea Weed"}
	got := struct{Op string; Id int; Name string}{}
	uri :=`/api/v1/create/843/Sea%20Weed`
	template := `/api/v1/{Op}/{Id}/{Name}`
	err := ParseUriPathToFields(&uri, &template, &got)

    eq := reflect.DeepEqual(want, got)
	if !eq || err!=nil {
    	t.Errorf("ParseUriPathToFields() = %q, want %q", got, want)
	}
}