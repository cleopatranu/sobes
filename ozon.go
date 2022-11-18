package main

import "fmt"

func main() {
	//variables := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8}
	//variables := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//variables := []int{8, 7, 6, 5, 4, 3, 2, 1}

	variables := []int{1, 1, 2, 3, 4, 5, 1, 5, 6, 7, 8}

	db := check(variables)
	fmt.Println("variables: ", variables)
	fmt.Println("DB: ", db)
	fmt.Println("Результат:", checkDB(db))

}

//variables:  [1 1 2 3 4 5 1 5 6 7 8]
//res:  [2 2 2 2 1 2 2 2 2]
//Результат: Немонотонная функция

func check(data []int) []int {

	db := make([]int, len(data)-1)
	for i := range data {
		//fmt.Printf("%d) %d\n", i, data[i])
		if i != len(data) && i != len(data)-1 {
			db[i] = diff(data[i], data[i+1])
		}
	}
	var s []int
	for i := range db {
		if db[i] != 0 {
			s = append(s, db[i])
		}
	}

	return s
}

func diff(a int, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return 2
	}
	return 0
}

func checkDB(res []int) string {

	for i := range res {
		if i != len(res) && i != len(res)-1 {
			if res[i] != res[i+1] {
				return "Немонотонная функция"
			}
		}
	}
	return "Монотонная функция"
}
