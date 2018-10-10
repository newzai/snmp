package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"

	"github.com/extrame/xls"
)

var genHandler = flag.Bool("handler", false, "-handler=true|false")
var sheetIndex = flag.Int("index", 2, "-index=0|1")
var xlsfile = flag.String("xlsfile", "昆仑产品mib表-v3.0.xls", "-xlsfile=昆仑产品mib表-v2.8.xls")
var help = flag.Bool("help", false, "-help=true")

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

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	book, err := xls.Open(*xlsfile, "gbk")
	if err != nil {

		flag.Usage()
		return

	}

	sheet := book.GetSheet(*sheetIndex)
	if sheet == nil {
		flag.Usage()
		return
	}

	template := `OIDAttr{Name: "%s", OID: "%s", Type: %s, ReadOnly: %t, ValidHander: %s},`

	handlerTemplate := `//%s %s %s
	func %s(value interface{}) error {
		//%s  
		return nil
	}`
	oids := make(map[string]bool)
	fmt.Println(sheet.MaxRow)
	for i := 0; i <= int(sheet.MaxRow); i++ {
		row := sheet.Row(i)
		if i < 2 {
			continue
		}
		fmt.Println(i, " ", row.Col(1))
		if row != nil {
			if row.Col(1) == "" {
				continue
			}
			_, exist := oids[row.Col(1)]
			if exist {
				fmt.Printf("%s 重复名称 \n", row.Col(1))
				return
			}
			oids[row.Col(1)] = true

			_, exist = oids[row.Col(2)]
			if exist {
				fmt.Printf("%s oid 值 \n", row.Col(2))
				return
			}
			oids[row.Col(2)] = true

			if *genHandler {
				if strings.Index(row.Col(5), "set") != -1 {
					data := fmt.Sprintf(handlerTemplate, toUpper(row.Col(1)), row.Col(1), row.Col(3), toUpper(row.Col(1)), row.Col(4))
					fmt.Println(data)
				}

				continue
			}

			xtype := ""
			if row.Col(3) == "Octstring" {
				xtype = "gosnmp.OctetString"
			} else if row.Col(3) == "int" {
				xtype = "gosnmp.Integer"
			} else if row.Col(3) == "OBJECT_IDENTIFIER" {
				xtype = "gosnmp.ObjectIdentifier"
			} else if row.Col(3) == "uint32" {
				xtype = "gosnmp.Uinteger32"
			} else {
				if row.Col(2) == "" && row.Col(3) == "" {
					continue
				}

				fmt.Println("无法识别的数据类型:", row.Col(3))
				for j := row.FirstCol(); j < row.LastCol(); j++ {
					fmt.Printf("%s,", row.Col(j))
				}
				fmt.Println()

				break
			}
			readOnly := true
			if strings.Index(row.Col(5), "set") != -1 {
				readOnly = false
			}
			hander := "defaultValidHander"
			if !readOnly {
				hander = toUpper(row.Col(1))
			}
			data := fmt.Sprintf(template, row.Col(1), row.Col(2), xtype, readOnly, hander)
			fmt.Println(data)
		}

	}
}
