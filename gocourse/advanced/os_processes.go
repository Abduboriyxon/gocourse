package advanced

import (
	// "time"
	// "io"
	"fmt"
	"os/exec"
	// "strings"
)

func main() {

	// ===== Simple Command Execution =====
	// cmd := exec.Command("echo", "Hello, World!")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error executing command:", err)
	// 	return
	// }
	// fmt.Println(string(output))

	// ===== Command with Input =====
	// cmd := exec.Command("grep", "foo")
	// // set input for the command
	// cmd.Stdin = strings.NewReader("foo\nbar\nbaz\n")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error executing command:", err)
	// 	return
	// }
	// fmt.Println(string(output))

	// ===== Managing Process Lifecycle =====
	// cmd := exec.Command("sleep", "2")
	// // start the command
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("Error starting command:", err)
	// 	return
	// }
	
	// // waiting
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("Error waiting for command:", err)
	// 	return
	// }
	// fmt.Println("Command finished successfully")

	// time.Sleep(2 * time.Second)
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println("Error killing command:", err)
	// 	return
	// }
	// fmt.Println("Command killed")

	// ===== Advanced Usage =====
	// cmd := exec.Command("printenv", "SHELL")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error executing command:", err)
	// 	return
	// }
	// fmt.Println(string(output))

	// ===== Using Pipes =====
	// pr, pw := io.Pipe()
	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = pr
	// go func(){
	// 	defer pw.Close()
	// 	pw.Write([]byte("food is good\nbar\nbaz\n"))
	// }()

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error executing command:", err.Error())
	// 	return
	// }
	// fmt.Println(string(output))	

	// ===== Combined Output =====
	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println(string(output))
}