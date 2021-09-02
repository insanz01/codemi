package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var isRun bool
var n_loker int

type Loker struct {
	Nomor, Tipe_ID, Nomor_ID, Status string
}

var loker []Loker

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func isEmpty() bool {
	if n_loker == 0 {
		fmt.Println("n loker masih kosong. (harus init value)")

		return true
	}

	return false

}

func initialize(n int) string {

	if n_loker == 0 {
		return "Harus lebih dari 0"
	}

	n_loker = n

	for i := 0; i < n_loker; i++ {
		loker = append(loker, Loker{
			strconv.Itoa(i + 1),
			"-",
			"-",
			"kosong",
		})
	}

	return "Berhasil membuat loker dengan jumlah " + strconv.Itoa(n)
}

func stats() string {
	fmt.Println("No Loker\tTipe Identitas\tNo Identitas")

	result := ""

	for i := 0; i < n_loker; i++ {
		// fmt.Print(loker[i].Nomor + "\t\t")
		result += loker[i].Nomor + "\t\t"
		// fmt.Print(loker[i].Tipe_ID + "\t\t")
		result += loker[i].Tipe_ID + "\t\t"
		// fmt.Print(loker[i].Nomor_ID + "\n")
		result += loker[i].Nomor_ID + "\n"
	}

	return result
}

func input(Tipe_ID string, Nomor_ID string) string {
	ketemu := false
	lokasi_loker := 0

	for i := 0; i < n_loker; i++ {
		if loker[i].Status == "kosong" {

			lokasi_loker = i + 1
			loker[i].Tipe_ID = Tipe_ID
			loker[i].Nomor_ID = Nomor_ID
			loker[i].Status = "ada"

			ketemu = true
			break
		}
	}

	if !ketemu {
		return "Maaf loker sudah penuh"
	}

	return "Kartu identitas berhasil disimpan di loker " + strconv.Itoa(lokasi_loker)
}

func leave(nomor_loker string) string {
	ketemu := false

	for i := 0; i < n_loker; i++ {
		if loker[i].Nomor == nomor_loker {
			loker[i].Tipe_ID = "-"
			loker[i].Nomor_ID = "-"
			loker[i].Status = "kosong"

			ketemu = true
			break
		}
	}

	if !ketemu {
		return "Maaf loker tidak ditemukan"
	}

	return "Loker nomor " + nomor_loker + " berhasil dikosongkan"
}

func find(Nomor_ID string) string {

	for i := 0; i < n_loker; i++ {
		if loker[i].Nomor_ID == Nomor_ID {
			return "Kartu identitas tersebut berada di loker nomor " + strconv.Itoa(i+1)
		}
	}

	return "Nomor identitas tidak ditemukan"
}

func search(Tipe_ID string) string {
	result := ""
	for i := 0; i < n_loker; i++ {
		if loker[i].Tipe_ID == Tipe_ID {
			result += loker[i].Nomor_ID + ","
		}
	}

	return TrimSuffix(result, ",")
}

func exit() {
	isRun = false
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func app() string {

	isRun = true
	n_loker = 0
	// command := make([]string, 3, 3)
	for isRun == true {
		command := []string{}

		var action string
		argumen := []string{}

		var line string

		Block{
			Try: func() {
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					line = scanner.Text()
					// fmt.Printf("Input was: %q\n", line)

					testArray := strings.Fields(line)
					for _, v := range testArray {
						command = append(command, v)
					}
				}

				if len(command) > 2 {
					action = command[0]
					argumen = append(argumen, command[1])
					argumen = append(argumen, command[2])
				} else if len(command) > 1 {
					action = command[0]
					argumen = append(argumen, command[1])
					// argumen[0] = command[1]
				} else {
					action = command[0]
				}

				switch action {
				case "init":
					n, err := strconv.Atoi(argumen[0])
					// n := 3
					if err != nil {
						fmt.Println("Argumen invalid")
					} else {
						response := initialize(n)
						fmt.Println(response)
					}
				case "status":
					if len(argumen) > 0 {
						fmt.Println("Argumen invalid")
					} else if !isEmpty() {
						fmt.Println(stats())
					}

				case "input":
					if !isEmpty() {
						Tipe_ID := argumen[0]
						Nomor_ID := argumen[1]

						fmt.Println(input(Tipe_ID, Nomor_ID))
					}
				case "leave":
					if len(argumen) > 2 {
						fmt.Println("Argumen Invalid")
					} else if !isEmpty() {
						nomor_loker := argumen[0]

						fmt.Println(leave(nomor_loker))
					}
				case "find":
					if len(argumen) > 2 {
						fmt.Println("Argumen Invalid")
					} else if !isEmpty() {
						Nomor_ID := argumen[0]

						fmt.Println(find(Nomor_ID))
					}
				case "search":
					if len(argumen) > 2 {
						fmt.Println("Argumen Invalid")
					} else if !isEmpty() {
						Tipe_ID := argumen[0]

						fmt.Println(search(Tipe_ID))
					}
				case "exit":
					exit()
				default:
					fmt.Println("Perintah tidak dikenali")
				}

			},
			Catch: func(e Exception) {
				fmt.Println("Invalid Command")
			},
		}.Do()
	}

	return "Program berhenti"
}

func main() {
	app()
}
