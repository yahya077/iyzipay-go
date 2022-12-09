package builder

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"unsafe"
)

type RequestStringBuilder struct {
}

func NewRequestStringBuilder() RequestStringBuilder {
	return RequestStringBuilder{}
}

func (b RequestStringBuilder) GetRequestString(v interface{}) string {
	rs := new(RequestString)
	rv := reflect.ValueOf(v)
	num := rv.NumField()
	for i := 0; i < num; i++ {
		fv := rv.Field(i)
		st := rv.Type().Field(i)
		jsonName, err := st.Tag.Lookup("json")
		if err != true {
			rs.appendString(st.Name + "=")
		} else {
			rs.appendString(fmt.Sprintf("%s=", jsonName))
		}
		switch fv.Kind() {
		case reflect.String:
			rs.appendString(fv.String())
		case reflect.Int:
			rs.appendInt(fv.Int())
		case reflect.Struct:
			rs.appendString(b.GetRequestString(fv.Interface()))
		case reflect.Slice:
			rs.appendString("[")
			for k := 0; k < fv.Len(); k++ {
				e := fv.Index(k)
				rs.appendString(b.GetRequestString(e.Interface()))
				if k != fv.Len()-1 {
					rs.appendString(", ")
				}
			}
			rs.appendString("]")
		}
		if i != num-1 {
			rs.appendComma()
		}
	}
	if rs.isLastCharComma() {
		rs.s = rs.s[0 : rs.len()-1]
	}
	rs.wrap("[", "]")
	return rs.String()
}

func (b RequestStringBuilder) GetRandomString(size int) string {
	var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]byte, size)
	rand.Read(s)
	for i := 0; i < size; i++ {
		s[i] = alphabet[s[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&s))
}

func (b RequestStringBuilder) HashPkiString(pki string) string {
	encoder := sha1.New()

	encoder.Write([]byte(pki))
	hash := base64.URLEncoding.EncodeToString(encoder.Sum(nil))

	hash = strings.Replace(hash, "_", "/", -1)
	hash = strings.Replace(hash, "-", "+", -1)

	return hash
}

func (b RequestStringBuilder) BuildPki(apiKey, apiSecret string, request interface{}) (randomString string, pki string) {
	randomString = b.GetRandomString(8)
	requestString := b.GetRequestString(request)

	hashedPki := b.HashPkiString(apiKey + randomString + apiSecret + requestString)

	return randomString, hashedPki
}

type RequestString struct {
	s string
}

func (s *RequestString) String() string {
	return s.s
}

func (s *RequestString) appendString(str string) {
	s.s += fmt.Sprintf("%s", str)
}

func (s *RequestString) appendInt(i int64) {
	s.s += fmt.Sprint(i)
}

func (s *RequestString) appendComma() {
	s.s += ","
}

func (s *RequestString) wrap(l, r string) {
	s.s = l + s.s + r
}

func (s *RequestString) len() int {
	return len(s.s)
}

func (s *RequestString) isLastCharComma() bool {
	return s.s[s.len()-1:] == ","
}
