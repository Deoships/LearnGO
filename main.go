package main

func main() {
	todos := Todos{}
	todos.add("Купить хлеб")
	todos.add("Сделать зарядку")
	todos.toggle(0)
	todos.print()
}
