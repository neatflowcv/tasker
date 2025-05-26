package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Content   string
	Completed bool
}

var tasks []Task
var lastID int

func addTask(content string) {
	lastID++
	task := Task{
		ID:        lastID,
		Content:   content,
		Completed: false,
	}
	tasks = append(tasks, task)
	fmt.Printf("Task 추가됨: [%d] %s\n", task.ID, task.Content)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("등록된 Task가 없습니다.")
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("[%d] [%s] %s\n", task.ID, status, task.Content)
	}
}

func completeTask(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Printf("Task 완료 처리됨: [%d] %s\n", id, tasks[i].Content)
			return true
		}
	}
	fmt.Printf("ID %d인 Task를 찾을 수 없습니다.\n", id)
	return false
}

func deleteTask(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task 삭제됨: [%d]\n", id)
			return true
		}
	}
	fmt.Printf("ID %d인 Task를 찾을 수 없습니다.\n", id)
	return false
}

func printHelp() {
	fmt.Println("사용 가능한 명령어:")
	fmt.Println("add [내용] - 새로운 Task 추가")
	fmt.Println("list - Task 목록 조회")
	fmt.Println("complete [ID] - Task 완료 처리")
	fmt.Println("delete [ID] - Task 삭제")
	fmt.Println("help - 도움말 표시")
	fmt.Println("exit - 프로그램 종료")
}

func main() {
	fmt.Println("Task 관리 프로그램 시작")
	fmt.Println("도움말을 보려면 'help'를 입력하세요")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]
		switch command {
		case "add":
			if len(args) < 2 {
				fmt.Println("사용법: add [내용]")
				continue
			}
			content := strings.Join(args[1:], " ")
			addTask(content)

		case "list":
			listTasks()

		case "complete":
			if len(args) != 2 {
				fmt.Println("사용법: complete [ID]")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("올바른 ID를 입력하세요")
				continue
			}
			completeTask(id)

		case "delete":
			if len(args) != 2 {
				fmt.Println("사용법: delete [ID]")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("올바른 ID를 입력하세요")
				continue
			}
			deleteTask(id)

		case "help":
			printHelp()

		case "exit":
			fmt.Println("프로그램을 종료합니다")
			return

		default:
			fmt.Println("알 수 없는 명령어입니다. 'help'를 입력하여 도움말을 확인하세요")
		}
	}
}
