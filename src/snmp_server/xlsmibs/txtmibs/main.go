package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var genHandler = flag.Bool("handler", false, "-handler=true|false")
var mibsfile = flag.String("mibfile", "mibs.txt", "-mibfile=")

func main() {

	flag.Parse()
	template := `OIDAttr{Name: "%s", OID: "%s", Type: %s, ReadOnly: %t, ValidHander: %s},`

	handlerTemplate := `//%s %s %s
	func %s(value interface{}) error {
		//%s  
		return nil
	}`
	oids := make(map[string]bool)

	err := ReadLine(*mibsfile, func(row int, line string) {
		fileds := strings.Split(line, "\t")

		if len(fileds) < 4 {
			return
		}
		_, exist := oids[fileds[1]]
		if exist {
			fmt.Printf("%s 重复名称 \n", fileds[1])
			return
		}
		oids[fileds[1]] = true

		_, exist = oids[fileds[2]]
		if exist {
			fmt.Printf("%s oid 值 \n", fileds[2])
			return
		}
		oids[fileds[2]] = true

		if *genHandler {
			if strings.Index(fileds[5], "set") != -1 {
				data := fmt.Sprintf(handlerTemplate, toUpper(fileds[1]), fileds[1], fileds[3], toUpper(fileds[1]), fileds[4])
				fmt.Println(data)
			}

			return
		}

		xtype := ""
		if fileds[3] == "Octstring" {
			xtype = "gosnmp.OctetString"
		} else if fileds[3] == "int" {
			xtype = "gosnmp.Integer"
		} else if fileds[3] == "OBJECT_IDENTIFIER" {
			xtype = "gosnmp.ObjectIdentifier"
		} else if fileds[3] == "uint32" {
			xtype = "gosnmp.Uinteger32"
		} else {
			if fileds[2] == "" && fileds[3] == "" {
				return
			}

			fmt.Println("无法识别的数据类型:", fileds)

			return
		}
		readOnly := true
		if strings.Index(fileds[5], "set") != -1 {
			readOnly = false
		}
		hander := "defaultValidHander"
		if !readOnly {
			hander = toUpper(fileds[1])
		}
		data := fmt.Sprintf(template, fileds[1], fileds[2], xtype, readOnly, hander)
		fmt.Println(data)

	})
	if err != nil {
		fmt.Println("err:", err)
	}
}

//ReadLine readline from filename
func ReadLine(filename string, handler func(int, string)) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	number := 0
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(number, line)
		number++
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

}
func toUpper(name string) string {
	buf := bytes.NewBuffer(nil)
	underline := false
	for _, c := range name {
		if underline {
			if c == '_' {
				continue
			}
			underline = false
			buf.WriteString(strings.ToUpper(string(c)))
		} else if c == '_' {
			underline = true
		} else {
			buf.WriteString(string(c))
		}
	}
	return buf.String()
}
