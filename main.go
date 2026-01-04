package main

import (
	"fmt"
	"net/http"
	"strings"
)

type USSDBody struct {
	SessionId   string
	PhoneNumber string
	Text        string
}

type USSDState struct {
	State string
}

var ussdState = USSDState{
	State: "page1",
}

var sessionStore = make(map[string]string)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := ""

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Unexpected error occured : %v", http.StatusBadRequest)
			return
		}
		ussdBody := USSDBody{
			SessionId:   r.FormValue("sessionId"),
			PhoneNumber: r.FormValue("phoneNumber"),
			Text:        r.FormValue("text"),
		}

		arrCommand := strings.Split(ussdBody.Text, "*")
		command := arrCommand[len(arrCommand)-1]

		switch command {
		case "98":
			switch ussdState.State {
			case "page1":
				ussdState.State = "page2"
			case "page2":
				ussdState.State = "page3"
			case "pageGenZ":
				ussdState.State = "pageGenZNext"
			default:
				response = "Invalid menu try again!\n" + response
			}
		case "0":
			switch ussdState.State {
			case "page3":
				ussdState.State = "page2"
			case "page2":
				ussdState.State = "page1"
			case "page1":
				ussdState.State = "pageGenZ"
			case "pageGenZ":
				ussdState.State = "page1"
			case "pageGenZNext":
				ussdState.State = "pageGenZ"
			case "socialBundles":
				ussdState.State = "page1"
			case "youtube1GB":
				ussdState.State = "page1"
			default:
				response = "Invalid menu try again!\n" + response
			}
		case "00":
			if ussdState.State == "pageGenZ" ||
				ussdState.State == "pageGenZNext" ||
				ussdState.State == "socialBundles" ||
				ussdState.State == "youtube1GB" {
				ussdState.State = "page1"
			} else {
				response = "Invalid menu try again!\n" + response
			}
		case "1":
			switch ussdState.State {
			case "page1":
				ussdState.State = "socialBundles"
			case "dataBlive":
				ussdState.State = "dataBlivePage1"
			default:
				response = "Invalid menu try again!\n" + response
			}
		case "2":
			switch ussdState.State {
			case "page1":
				ussdState.State = "dataBlive"
			case "dataBlive":
				ussdState.State = "dataBlivePage2"
			default:
				response = "Invalid menu try again!\n" + response
			}
		case "3":
			switch ussdState.State {
			case "page1":
				ussdState.State = "youtube1GB"
			default:
				response = "Invalid menu try again!\n" + response
			}

		default:
			response = "Invalid menu try again!\n" + response
		}

		switch ussdState.State {
		case "page1":
			response = GetMenu("page1")
		case "page2":
			response = GetMenu("page2")
		case "page3":
			response = GetMenu("page3")
		case "pageGenZ":
			response = GetMenu("pageGenZ")
		case "pageGenZNext":
			response = GetMenu("pageGenZNext")
		case "socialBundles":
			response = GetMenu("socialBundles")
		case "dataBlive":
			response = GetMenu("dataBlive")
		case "dataBlivePage1":
			response = GetMenu("dataBlivePage1")
		case "dataBlivePage2":
			response = GetMenu("dataBlivePage2")
		case "youtube1GB":
			response = GetMenu("youtube1GB")
		default:
			response = `END Unkown Error Occurred!`
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%v", response)
	})

	fmt.Println("Server listening on http://localhost:3000...")
	http.ListenAndServe(":3000", nil)
}
