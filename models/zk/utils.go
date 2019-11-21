package zk

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
)

func LoadAtributes(node string, d interface{}) (ret error) {
	dval := reflect.ValueOf(d)
	if dval.Kind() != reflect.Ptr {
		ret = fmt.Errorf("The second paramter should be a pointer")
		return
	}

	var aval reflect.Value
	if dval.Elem().Kind() == reflect.Ptr {
		akind := dval.Elem().Type().Elem()
		if akind.Kind() != reflect.Struct {
			ret = fmt.Errorf("paramter d shall be a struct or pointer to a struct")
			return
		}

		dval.Set(reflect.New(akind))
		aval = dval.Elem()
	} else if dval.Elem().Kind() != reflect.Struct {
		ret = fmt.Errorf("paramter d shall be a struct or pointer to a struct")
		return
	} else {
		aval = dval.Elem()
	}

	atp := aval.Type()

	for i := 0; i < aval.NumField(); i++ {
		v := aval.Field(i)
		t := atp.Field(i)

		if path, suc := t.Tag.Lookup("zk"); suc {
			fpath := node + "/" + path
			content, _, err := ZKAgent.Get(fpath)
			if err != nil {
				log.Printf("Error while fetching node data for %s, some data will be missing, %v", fpath, err)
				continue
			}
			valStr := string(content)
			if len(valStr) == 0 {
				log.Printf("Empty field [%s@%s], ignore", path, fpath)
				continue
			}

			//log.Printf("tag = %s, kind = %s", path, t.Type.Kind())

			var cerr error
			switch t.Type.Kind() {
			case reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int:
				var r int64
				r, cerr = strconv.ParseInt(valStr, 10, 64)
				if cerr == nil {
					v.SetInt(r)
				}
			case reflect.Float32, reflect.Float64:
				var f float64
				f, cerr = strconv.ParseFloat(valStr, 64)
				if cerr == nil {
					v.SetFloat(f)
				}
			case reflect.Struct:
				if t.Type == reflect.ValueOf(time.Time{}).Type() {
					var r int64
					r, cerr = strconv.ParseInt(valStr, 10, 64)
					if cerr == nil {
						t := time.Unix(r, 0)
						v.Set(reflect.ValueOf(t))
					}
				} else {
					log.Printf("struct-typed field is not of type Time, ignore")
				}

			case reflect.String:
				v.SetString(valStr)
			}

			if cerr != nil {
				log.Printf("Error parsing filed %s, %v", fpath, cerr)
			}
		}
	}

	return
}

func ListAllNodes(root string) (ret []string, reterr error) {
	ret, _, reterr = ZKAgent.Children(root)
	return
}
