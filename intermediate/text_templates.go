package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {

	// tmpl := template.New("example")

	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you today?\n")
	// if err != nil {
	// 	panic(err)
	// }
	// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you today?\n"))

	// // define data for the welcom template
	// data := map[string]interface{}{
	// 	"name": "Alice",
	// }
	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader((os.Stdin))
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	// Define named templates for different types of
	templates := map[string]string{
		"welcome": "Welcome, {{.name}}! We're glad to join us today.",
		"notification": "{{.nm}}, you have a new notification waiting: {{.ntf}}",
		"error": "Oops, An error occured: {{.em}}",
	}

	// Perse and store templates
	perseTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		perseTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Show menu
		fmt.Println("\nMenu: ")
		fmt.Println("1. Join ")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1": 
			tmpl = perseTemplates["welcome"]
			data = map[string]interface{}{"name":name}
		case "2":
			fmt.Println("Enter notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = perseTemplates["notification"]
			data = map[string]interface{}{"nm": name, "ntf": notification}
		case "3":
			fmt.Println("Enter error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = perseTemplates["error"]
			data = map[string]interface{}{"nm": name, "em": errorMessage}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}
		// render and print template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	}

}