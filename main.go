package main

import (
	"io/ioutil"
	"bufio"
	"os"
)

func main() {
	print("Enter a path directory: ")
	var path, command string
	var files []string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		path = scanner.Text()
	}

	if path[len(path)-1:] != "/" {
		path = path + "/"
	}

	print("Crypt or decrypt? ")

	if scanner.Scan() {
		command = scanner.Text()
	}

	if command == "crypt" {
		files = listFiles(path, true)
		cryptFiles(files)
	} else if command == "decrypt" {
		files = listFiles(path, false)
		decryptFiles(files)
	} else {
		println("Wrong command ! Type \"crypt\" or \"decrypt\".")
	}

	println(len(files), "files found")
}

func padding(text string, blockSize int) []string {
	array := make([]string, 0)

	i := 0

	for i < len(text) {

		end := i + blockSize
		if end > len(text) - 1 {
			end = len(text)
		}

		slice := text[i:end]
		array = append(array, slice)
		i += blockSize
	}

	return array
}

func abs(number int) (int) {
	if number < 0 {
		return -number
	}
	return number
}

func cesar(text string, shiftNumber int) string {
	shift, offset := rune(abs(shiftNumber) % 26), rune(26)

	runes := []rune(text)

	for index, char := range runes {

		if shiftNumber < 0 {
			if char >= 'a'+shift && char <= 'z' || char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift || char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		} else {
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

func xor(first string, second string) (string) {
	var length int

	if len(first) > len(second) {
		length = len(second)
	} else {
		length = len(first)
	}

	bytes := make([]byte, length)

	for i := 0; i < length; i++ {
		bytes[i] = first[i] ^ second[i]
	}

	return string(bytes)
}

func unxor(first string, second string) (string) {
	var length int

	if len(first) > len(second) {
		length = len(second)
	} else {
		length = len(first)
	}

	bytes := make([]byte, length)

	for i := 0; i < length; i++ {
		bytes[i] = (^first[i] & second[i]) | (first[i] & ^second[i])
	}

	return string(bytes)

	/*
&   bitwise AND
|   bitwise OR
^   bitwise XOR
&^   AND NOT
<<   left shift
>>   right shift
 */
}

func encrypt(text string, key string) (string) {
	// Découpe en blocs
	paddingArray := padding(text, 16)

	// Création du tableau final
	encrypted := make([]string, 0)

	for index, txt := range paddingArray {
		cryptedKey := cesar(key, index)
		encrypted = append(encrypted, xor(txt, cryptedKey))
	}

	result := ""

	for _, txt := range encrypted {
		result += txt
	}

	return result
}

func decrypt(text string, key string) (string) {
	// Découpe en blocs
	paddingArray := padding(text, 16)

	// Création du tableau final
	decrypted := make([]string, 0)

	for index, txt := range paddingArray {
		cryptedKey := cesar(key, index)
		decrypted = append(decrypted, unxor(cryptedKey, txt))
	}

	result := ""

	for _, txt := range decrypted {
		result += txt
	}

	return result
}

const key = "abcdefghijklmnop"

/*func main() {
	text := "Bonjour encule de ta race il n y a pas assez de caracteres donc j en rajoute car tu es un sale connard espece de Kevin de merde est ce que t entends ca, ca depasse l entendement fils de pute onche onche onche test de Jason qui pue des fesses et qui mange des zizis de maeva et en plus il est en couple avec elle xoxop"

	encrypted := encrypt(text, key)
	decrypted := decrypt(encrypted, key)

	println(decrypted)
} */

func listFiles(path string, crypt bool) ([]string) {
	files := make([]string, 0)

	directoryContent, err := ioutil.ReadDir(path)
	if err != nil {
		println(err.Error())
		return files
	}

	for _, fileInDirectory := range directoryContent {
		if !fileInDirectory.IsDir() {

			filename := fileInDirectory.Name()

			if crypt && filename[len(filename)-1:] != "_" || !crypt && filename[len(filename)-1:] == "_" {
				files = append(files, path + fileInDirectory.Name())
			}
		} else {
			newPath := path + fileInDirectory.Name() + "/"
			files = append(files, listFiles(newPath, crypt)...)
		}
	}

	return files
}

func printFiles(files []string) {
	for _, file := range files {
		println(file)
	}
}

func cryptFiles(files []string) {
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			println(err.Error())
			continue
		}

		encryptedContent := encrypt(string(content), key)

		err = ioutil.WriteFile(file + "_", []byte(encryptedContent), 0644)
		if err != nil {
			println(err.Error())
			continue
		}

		err = os.Remove(file)
		if err != nil {
			println(err.Error())
			continue
		}
	}
}

func decryptFiles(files []string) {
	for _, file := range files {
		content, err := ioutil.ReadFile( file)
		if err != nil {
			println(err.Error())
			continue
		}

		decryptedContent := decrypt(string(content), key)

		err = ioutil.WriteFile(file[0:len(file)-1], []byte(decryptedContent), 0644)
		if err != nil {
			println(err.Error())
			continue
		}

		err = os.Remove(file)
		if err != nil {
			println(err.Error())
			continue
		}
	}
}