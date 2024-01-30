package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["NJ"] = "Node.Js"
	languages["PY"] = "Python"
	languages["GO"] = "GO Lang"
	languages["JV"] = "Java"

	fmt.Println("ALL LANGUAGES -", languages)
	fmt.Println("PY stands for -", languages["PY"])

	delete(languages, "JS")
	fmt.Println("ALL LANGUAGES -", languages)

	//loops
	for key, value := range languages {
		fmt.Println(key, "stands for", value)
	}
}
