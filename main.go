package main

import (
	"io/ioutil"
)




/*
func main() {
	print("Enter a path directory: ")
	var path string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		path = scanner.Text()
	}

	if path[len(path)-1:] != "/" {
		path = path + "/"
	}

	files := listFiles(path)
	printFiles(files)
	println(len(files), "files found")

}
*/

func padding(text string, blockSize int) []string {
	array := make([]string, 0)

	i := 0

	for i < len(text) {

		end := i + blockSize
		if end > len(text) - 1 {
			end = len(text) - 1
		}

		slice := text[i:end]
		array = append(array, slice)
		i += blockSize
	}

	return array
}

func cipher(text string, direction int) string {
	shift, offset := rune(3), rune(26)

	runes := []rune(text)

	for index, char := range runes {
		switch direction {
		case -1:
			if char >= 'a'+shift && char <= 'z' || char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift || char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case +1:
			if char >= 'a' && char <= 'z'-shift || char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' || char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		runes[index] = char
	}

	return string(runes)
}



func main() {
	//key := "1234"
	text := "Bonjour encule de ta race il n y a pas assez de caracteres donc j en rajoute car tu es un sale connard espece de Kevin de merde est-ce que t entends ca, ca depasse l'entendement fils de pute."

	/*padding := padding(text, 4)

	for index, txt := range padding {
		println(index, txt)
	}*/

	encoded := cipher(text, -1)
	decoded := cipher(encoded, +1)

	println(encoded)
	println(decoded)
}

func listFiles(path string) ([]string) {
	files := make([]string, 0)

	directoryContent, err := ioutil.ReadDir(path)
	if err != nil {
		println(err.Error())
		return files
	}

	for _, fileInDirectory := range directoryContent {
		if !fileInDirectory.IsDir() {
			files = append(files, fileInDirectory.Name())
		} else {
			newPath := path + "/" + fileInDirectory.Name()
			files = append(files, listFiles(newPath)...)
		}
	}

	return files
}

func printFiles(files []string) {
	for _, file := range files {
		println(file)
	}
}