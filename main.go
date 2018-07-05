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

func main() {
	key := "abcdefghijklmnop"
	text := "Bonjour encule de ta race il n y a pas assez de caracteres donc j en rajoute car tu es un sale connard espece de Kevin de merde est ce que t entends ca, ca depasse l entendement fils de pute onche onche onche test de Jason qui pue des fesses et qui mange des zizis de maeva et en plus il est en couple avec elle xoxop"

	println("Taille texte de base", len(text))

	// Découpe en blocs
	paddingArray := padding(text, 16)

	// Création du tableau final
	encrypted := make([]string, 0)

	for index, txt := range paddingArray {
		println(index, txt)

		cryptedKey := cesar(key, index)
		encrypted = append(encrypted, xor(txt, cryptedKey))
	}

	finalString := ""

	for index, txt := range encrypted {
		println(index, txt)
		finalString += txt
	}

	println(finalString)
	println("Taille texte encrypté", len(finalString))

	////////// Déchiffrement

	decryptPadding := padding(finalString, 16)
	decrypted := make([]string, 0)

	for index, txt := range decryptPadding {

		cryptedKey := cesar(key, index)
		decrypted = append(decrypted, unxor(cryptedKey, txt))
	}

	decryptedString := ""
	for index, txt := range decrypted {
		println(index, txt)
		decryptedString += txt
	}

	println(decryptedString)
	println("Taille texte décrypté", len(decryptedString))
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