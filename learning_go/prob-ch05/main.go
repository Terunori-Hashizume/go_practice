package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j != 0 {
		return i / j, nil
	} else {
		return 0, errors.New("0で割ることはできません")
	}
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func fileLen(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// ファイル情報を取得してサイズを返却
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return int(fileInfo.Size()), nil
}

func prefixer(prefix string) func(string) string {
	return func(target string) string {
		return prefix + " " + target
	}
}

func main() {
	fmt.Println("# 1")
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
		[]string{"2", "/", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 { // 演算子と被演算子の合計個数のチェック
			fmt.Print(expression, " -- 不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0]) // 1番目の被演算子のチェック
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1] // 演算子のチェック
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です: ", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2]) // 2番目の被演算子のチェック
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result, err := opFunc(p1, p2) // 実際の計算
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		fmt.Print(expression, " → ", result, "\n")
	}

	// ii
	fmt.Println("\n# 2")
	fmt.Println(fileLen("./main.go"))
	fmt.Println(fileLen("./go.mod"))
	fmt.Println(fileLen("./invalid.go"))

	// iii
	fmt.Println("\n# 3")
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
