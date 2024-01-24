package tasks

import "fmt"
import "encoding/json"
import "os"
import "io"
import "github.com/fatih/color"

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTask() {

	task, err := loadTasks()
	if err != nil {
		panic(err)
	}

	if len(task) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, t := range task {

		status := " "
		if t.Complete {
			status = color.BlueString("✓")
		}

		fmt.Printf("[%s] %d %s\n", status, t.ID, t.Name)
	}
}

func Addtask(name string) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}

	newTask := Task{
		ID:       len(tasks) + 1,
		Name:     name,
		Complete: false,
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error al guardar las tareas", err)
		return
	}
	fmt.Printf("Tarea agregada correctamente")
}

func CompleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Complete = true
		}
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error al guardar las tareas", err)
		return
	}
	fmt.Printf("Tarea completada correctamente")
}

func DeleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	// Reasignar IDs para que sean consecutivos
	for i := range tasks {
		fmt.Println(i)
		tasks[i].ID = i + 1
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error al guardar las tareas", err)
		return
	}
	fmt.Printf("Tarea eliminada correctamente")
}

func loadTasks() ([]Task, error) {
	var tasks []Task

	// Leer el contenido del archivo JSON
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		/*
		Esta función toma una cadena de texto JSON y la convierte de nuevo 
		en unaestructura de datos de Go. Por ejemplo, toma la cadena de texto 
		JSON que representa una tarea y la convierte en una instancia de 
		la estructura Task.
		*/

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []Task{}
	}

	return tasks, nil
}

func saveTasks(tareas []Task) error {
	/* Esta función toma una estructura de datos de Go y 
	la convierte en una representación JSON. Por ejemplo, 
	toma una estructura Task y produce una cadena de texto 
	JSON que representa esa tarea en formato JSON.*/
	jsonData, err := json.Marshal(tareas)
	if err != nil {
		return err
	}

	// Escribir el contenido JSON en el archivo
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
