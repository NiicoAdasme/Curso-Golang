package main

import "fmt"

// Lista de tareas
type TaskList struct{
	tasks [] *Task
}

func (tl *TaskList) appendTask(t *Task){
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

// Tareas
type Task struct{
	name string
	desc string
	completed bool
}

// Metodos
func (t *Task) toPrint(){
	fmt.Printf("Nombre: %s\nDescripcion: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted(){
	t.completed = true
}

func main() {
	
	t1 := Task{
		name: "Curso Go",
		desc: "Completar curso de Golang",
		completed: false,
	}

	t2 := Task{
		name: "Curso Vue",
		desc: "Completar curso de Vue.js",
		completed: true,
	}

	
	lista := TaskList{}

	lista.appendTask(&t1)
	lista.appendTask(&t2)

	fmt.Println(lista)

	t3 := Task{
		name: "Curso React Native",
		desc: "Completar curso de React Native",
		completed: false,
	}

	lista.appendTask(&t3)

	// t1.toPrint()
	// t2.toPrint()

	fmt.Println(lista)

	lista.removeTask(1)

	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}
	

}