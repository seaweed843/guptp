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
	"fmt"
	"strings"
	"strconv"
	"time"
	"regexp"
	"reflect"
	"net/url"
)

func ParseUriPathToFields(path *string, template *string, structPtr interface{}) (errRet error) {
	if strings.Contains(*template, "{") {
		pathSplit := strings.Split(*template,"/")
		for index, value := range pathSplit {
			if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
				valueLen := len(pathSplit[index])
				varName := regexp.QuoteMeta(pathSplit[index][1:valueLen-1])
				pathSplit[index] = "(?P<" + varName +">[^/]+)"
			}else{
				pathSplit[index] = regexp.QuoteMeta(pathSplit[index])
			}
		}

		var regExp string
		for _, value := range pathSplit {
			if value != ""{
				regExp = regExp +  "/" + value 
			}
		}

		reg := regexp.MustCompile(regExp)
		matches := reg.FindStringSubmatch(*path)
		if len(matches) > 0{
			rv := reflect.ValueOf(structPtr)
			for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
				rv = rv.Elem()
			}
			subExpNames := reg.SubexpNames()
			for i:=1; i<len(subExpNames);i++{
				matchUnescaped, _ := url.PathUnescape(matches[i])
				f := rv.FieldByName(subExpNames[i])
					if f.IsValid() {
						if f.CanSet() {
							switch fKind := f.Kind(); fKind {
							case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
								if intConverted, err :=  strconv.ParseInt(matchUnescaped, 10, 0); err == nil {
									switch fKind {
									case reflect.Int:
										f.Set(reflect.ValueOf(int(intConverted)))
									case reflect.Int8:
										f.Set(reflect.ValueOf(int8(intConverted)))
									case reflect.Int16:
										f.Set(reflect.ValueOf(int16(intConverted)))
									case reflect.Int32:
										f.Set(reflect.ValueOf(int32(intConverted)))
									case reflect.Int64:
										f.Set(reflect.ValueOf(int64(intConverted)))
									}	
								}else{
									errRet = err
									return
								}
							case reflect.Uint,reflect.Uint8, reflect.Uint16, reflect.Uint32,reflect.Uint64:
								if uintConverted, err :=  strconv.ParseUint(matchUnescaped, 10, 0); err == nil {
									switch fKind {
									case reflect.Uint:
										f.Set(reflect.ValueOf(uint(uintConverted)))
									case reflect.Uint8:
										f.Set(reflect.ValueOf(uint8(uintConverted)))
									case reflect.Uint16:
										f.Set(reflect.ValueOf(uint16(uintConverted)))
									case reflect.Uint32:
										f.Set(reflect.ValueOf(uint32(uintConverted)))
									case reflect.Uint64:
										f.Set(reflect.ValueOf(uint64(uintConverted)))
									}
								}else{
									errRet = err
									return
								}
							case reflect.Bool:
								if boolConverted, err :=  strconv.ParseBool(matchUnescaped); err == nil {
									f.Set(reflect.ValueOf(boolConverted))
								}else{
									errRet = err
									return
								}
							case reflect.Float32:
								if float32Converted, err :=  strconv.ParseFloat(matchUnescaped, 32); err == nil {
									f.Set(reflect.ValueOf(float32Converted))
								}else{
									errRet = err
									return
								}
							case reflect.Float64:
								if float64Converted, err :=  strconv.ParseFloat(matchUnescaped, 64); err == nil {
									f.Set(reflect.ValueOf(float64Converted))
								}else{
									errRet = err
									return
								}
							case reflect.TypeOf(time.Time{}).Kind():
								if sf, ok := reflect.TypeOf(rv).FieldByName(subExpNames[i]); ok{
									timeFormat := sf.Tag.Get("guptp")
									if timeFormat == "" {
										timeFormat = time.RFC3339
									}
									if timeParsed, err := time.Parse(timeFormat, matchUnescaped); err == nil {
										f.Set(reflect.ValueOf(timeParsed))
									}else{
										errRet = err
										return
									}
								}
							case reflect.String:
								f.Set(reflect.ValueOf(matchUnescaped))
							default:
								errRet = fmt.Errorf("/*unexpected field type*/ %T", structPtr)
							}
							
						}
					}
			}
		}else{
			errRet = fmt.Errorf("nothing matches:/*path*/ %s /*temple*/ %s",*path,*template)
		}

	}

	return
}

func ParseUriPathToMapStr(path, template *string) (result map[string]string){
	result = make(map[string]string)
	if strings.Contains(*template, "{") {
		pathSplit := strings.Split(*template,"/")
		for index, value := range pathSplit {
			if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
				valueLen := len(pathSplit[index])
				varName := regexp.QuoteMeta(pathSplit[index][1:valueLen-1])
				result[varName] = ""
				pathSplit[index] = "(?P<" + varName +">[^/]+)"
			}else{
				pathSplit[index] = regexp.QuoteMeta(pathSplit[index])
			}
		}

		var regExp string
		for _, value := range pathSplit {
			if value != ""{
				regExp = regExp +  "/" + value 
			}
		}

		reg := regexp.MustCompile(regExp)
		matches := reg.FindStringSubmatch(*path)
		if len(matches) > 0 {
			for k, _ := range result{
				matchUnescaped, _ := url.PathUnescape(matches[reg.SubexpIndex(k)])
				result[k] = matchUnescaped
			 }
		}
	}

	return
}