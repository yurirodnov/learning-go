package main

import "fmt"

func main() {

	var m = make(map[int]string)

	var counts = [11]string{"Ноль", "Один", "Два", "Три", "Четыре", "Пять",
		"Шесть", "Семь", "Восемь", "Девять", "Десять"}

	var count_len = len(counts);

	for i := 0; i < count_len; i++ {
		m[i] = counts[i];
	}

	var first, second, third int

	fmt.Scan(&first, &second, &third)
	fmt.Println(m[first])
	fmt.Println(m[second])
	fmt.Println(m[third])

}