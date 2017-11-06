package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

type Animals []Animal

func main() {
	// Animal構造体初期化
	var inu Animal = Animal{
		Name: "わんわん",
		Age:  3,
	}

	// こっちの書き方でも初期化できる
	neko := Animal{
		Name: "にゃんにゃん",
		Age:  8,
	}

	// 配列宣言
	var animals Animals

	animals = append(animals, inu)
	animals = append(animals, neko)

	fmt.Println(animals[1].Name)
	// ⇛ [{わんわん 3} {にゃんにゃん 8}]

	// 配列をjsonに変換する
	b, _ := json.Marshal(animals)
	fmt.Printf("%s\n", b)
	// fmt.Printf(animals[1])
}
