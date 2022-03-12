package main

import (
	"encoding/json"
	"fmt"
	"go-scrape/profile"
	"os"
	"strings"
)

func main() {
	//d := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{
	//	browser.SizeBy(800, 800),
	//}))
	//
	//if err := d.Start(); err != nil {
	//	panic(err)
	//}
	//defer d.Stop()
	//
	//page, err := d.NewPage(agouti.Browser("chrome"))
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(parse("profile.json")[0])

	//if err := page.Navigate("https://google.com"); err != nil {
	//	panic(err)
	//}
	//<-time.Tick(15 * time.Second)
	//page.Screenshot("sample.png")
}

func init() {
	if err := setDriverPath(); err != nil {
		panic(err)
	}
}

func setDriverPath() error {
	current, err := os.Getwd()
	if err != nil {
		return err
	}

	pathEnv := []string{os.Getenv("PATH"), fmt.Sprintf("%s\\drivers", current)}
	return os.Setenv("PATH", strings.Join(pathEnv, ";"))
}

func loadJSON(path string) (interface{}, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func parse(path string) profile.Profiles {
	bytes, _ := os.ReadFile(path)
	var profiles profile.Profiles
	if err := json.Unmarshal(bytes, &profiles); err != nil {
		panic(err)
	}
	return profiles
}

//func loadJSON(path string, obj interface{}) (interface{}, error) {
//	bytes, err := os.ReadFile(path)
//	if err != nil {
//		return nil, err
//	}
//	//var result interface{}
//	if err := json.Unmarshal(bytes, &obj); err != nil {
//		return nil, err
//	}
//	return obj, nil
//}