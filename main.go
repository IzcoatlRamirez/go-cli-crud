package main

import (
	"fmt"
	"os"
	 task "github.com/IzcoatlRam/go-cli-crud/tasks"
	"strconv"
)

func main(){

	if(len(os.Args) < 2){
		printUsage()
		return	
	}

	switch os.Args[1]{
	case "list": 
		task.ListTask()
	case "add": 
		task.Addtask(os.Args[2])
	case "complete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El id debe ser un numero")
			return
		}
		task.CompleteTask(id)
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El id debe ser un numero")
			return
		}
		task.DeleteTask(id)
	}

}

func printUsage(){
	fmt.Println("Uso: go-cli-crud [list|add|complete|delete]")
}