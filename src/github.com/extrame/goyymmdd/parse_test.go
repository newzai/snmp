package yymmdd

import (
	"fmt"
	"testing"
	"time"
)

func TestParseYY(t *testing.T) {
	_, tokens := Lexer(`yy-`)
	ds := Parse(tokens)
	fmt.Println(ds.Format(time.Now()))
}

func TestParseMM(t *testing.T) {
	_, tokens := Lexer(`yyyymm`)
	ds := Parse(tokens)
	fmt.Println(ds.Format(time.Now()))
}

func TestParseMM2(t *testing.T) {
	_, tokens := Lexer(`yyyy-mm`)
	ds := Parse(tokens)
	fmt.Println(ds.Format(time.Now()))
}

func TestParseDD(t *testing.T) {
	_, tokens := Lexer(`yyyy-mm-dd`)
	ds := Parse(tokens)
	fmt.Println(ds.Format(time.Now()))
}
