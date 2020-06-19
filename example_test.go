package time_test

import (
	"fmt"
	stdtime "time"

	"github.com/kelseyhightower/time"
)

func ExampleJuneteenth() {
	fmt.Println(time.Juneteenth())
}

func ExampleCPT() {
	_, err := time.CPT(stdtime.Now())
	if err != nil {
		fmt.Println(err.Error())
	}
}
