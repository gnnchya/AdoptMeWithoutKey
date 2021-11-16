package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

func Get() *TestStruct {
	appConfig := &TestStruct{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}

type TestStruct struct{
	Name string `env:"test_name" envDefault:"tn"`
	ID string `env:"test_id" envDefault:"ti"`
}

func main(){
	testtemp := Get()

	fmt.Println(testtemp.Name)
	fmt.Println(testtemp.ID)
}
