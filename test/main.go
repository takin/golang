package main

import (
	"fmt"
	"time"

	"github.com/takin/golang/helpers"
)

func main() {
	d := &helpers.ConversionModel{
		Date: time.Now(),
	}

	res := d.Translate()
	fmt.Printf("%s, %d %s %d\n", res.Day, res.Date, res.Month, res.Year)
}
