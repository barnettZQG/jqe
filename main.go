package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/urfave/cli"
)

//App jqe command app
var App *cli.App

func main() {
	App = cli.NewApp()
	App.Usage = "Simple JSON file editor command, extend the jq command."
	App.Version = "0.1"
	App.Flags = []cli.Flag{}
	App.Commands = []cli.Command{
		cli.Command{
			Name: "update",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file,f",
					Value: "",
					Usage: "edit json file path.",
				},
				cli.StringFlag{
					Name:  "type,t",
					Value: "string",
					Usage: "edit param value type",
				},
			},
			Action: func(ctx *cli.Context) error {
				file := ctx.String("file")
				if file == "" {
					log.Fatal("Please specified json file by -f")
				}
				jsonStruct, err := readFile(file)
				if err != nil {
					log.Fatal(err.Error())
					return err
				}
				allParam := append([]string{ctx.Args().First()}, ctx.Args().Tail()...)
				var changeParams = make(map[string]string, len(allParam))
				for _, param := range allParam {
					if strings.Contains(param, "=") {
						kv := strings.SplitN(param, "=", 2)
						changeParams[kv[0]] = kv[1]
					}
				}
				for k, v := range changeParams {
					var value interface{}
					switch ctx.String("t") {
					case "bool":
						value, err = strconv.ParseBool(v)
						if err != nil {
							log.Fatalf("param %s value failure %s", k, err.Error())
						}
					case "int":
						value, err = strconv.Atoi(v)
						if err != nil {
							log.Fatalf("param %s value failure %s", k, err.Error())
						}
					}
					jsonStruct.SetPath(strings.Split(k, "."), value)
				}

				if err := writeFile(file, jsonStruct); err != nil {
					log.Fatal(err.Error())
					return err
				}
				return nil
			},
		},
		cli.Command{
			Name: "get",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file,f",
					Value: "",
					Usage: "edit json file path.",
				},
			},
			Action: func(ctx *cli.Context) error {
				file := ctx.String("file")
				if file == "" {
					log.Fatal("Please specified json file by -f")
				}
				jsonStruct, err := readFile(file)
				if err != nil {
					log.Fatal(err.Error())
					return err
				}
				node := jsonStruct.GetPath(strings.Split(ctx.Args().First(), ".")...)
				if node != nil {
					fmt.Printf("%v", node.Interface())
					return nil
				}
				log.Fatal("param not found")
				return nil
			},
		},
	}
	App.Run(os.Args)
}

func readFile(filename string) (*simplejson.Json, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read json file failure %s", err.Error())
	}
	jsonStrct, err := simplejson.NewJson(body)
	if err != nil {
		return nil, fmt.Errorf("read json body failure %s", err.Error())
	}
	return jsonStrct, nil
}

func writeFile(filename string, jsonStruct *simplejson.Json) error {
	body, err := jsonStruct.Encode()
	if err != nil {
		return fmt.Errorf("json encoding failure %s", err.Error())
	}
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("open json file failure %s", err.Error())
	}
	defer f.Close()
	_, err = f.Write(body)
	if err != nil {
		return fmt.Errorf("write json file failure %s", err.Error())
	}
	return nil
}
