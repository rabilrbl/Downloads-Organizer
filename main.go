package main

import (
	"log"
	"os"
	"strings"
	"time"
)

// Function that creates Folders if it doesn't exist
func createFolder(dir string, path string) {
	if _, err := os.Stat(path + "/" + dir); os.IsNotExist(err) {
		os.Mkdir(path+"/"+"/"+dir, 0755)
	}
}

// Function that will create folders if they don't exist
// Documents
// Text Files
// Pictures
// Music
// Videos
// Compressed
// Programs
// Others
// Folders
func createFolders(dir string) {
	createFolder("Documents", dir)
	createFolder("Text Files", dir)
	createFolder("Pictures", dir)
	createFolder("Music", dir)
	createFolder("Videos", dir)
	createFolder("Compressed", dir)
	createFolder("Programs", dir)
	createFolder("Others", dir)
	createFolder("Folders", dir)
}

// Function to move a file
func moveFile(src string, dest string) {
	err := os.Rename(src, dest)
	if err != nil {
		log.Println(err)
	}
}

// Function that will move files to their respective folders based on their extension
func sortFilesToFolders(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		extension := strings.Split(file.Name(), ".")
		size := len(extension)
		if size > 1 && file.Type().IsRegular() {
			fileName := extension[size-1]
			file := file.Name()
			switch fileName {
			case ".tmp":
				continue
			case "pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx":
				log.Println("Moving " + file + " to Documents")
				moveFile(path+"/"+file, path+"/"+"Documents/"+file)
			case "txt", "log", "ini", "conf", "csv", "json", "xml", "yml", "yaml", "rtf":
				if file == "sort_folder.log" {
					continue
				}
				log.Println("Moving " + file + " to Text Files")
				moveFile(path+"/"+file, path+"/"+"Text Files/"+file)
			case "jpg", "jpeg", "png", "gif", "bmp":
				log.Println("Moving " + file + " to Pictures")
				moveFile(path+"/"+file, path+"/"+"Pictures/"+file)
			case "mp3", "wav", "flac", "aac", "wma", "m4a":
				log.Println("Moving " + file + " to Music")
				moveFile(path+"/"+file, path+"/"+"Music/"+file)
			case "mp4", "avi", "flv", "mov", "mkv":
				log.Println("Moving " + file + " to Videos")
				moveFile(path+"/"+file, path+"/"+"Videos/"+file)
			case "zip", "rar", "7z":
				log.Println("Moving " + file + " to Compressed")
				moveFile(path+"/"+file, path+"/"+"Compressed/"+file)
			case "html", "css", "js", "php", "py", "go", "c", "cpp", "java", "csharp", "sql":
				log.Println("Moving " + file + " to Text Files")
				moveFile(path+"/"+file, path+"/"+"Text Files/"+file)
			case "exe", "msi", "bat", "sh":
				log.Println("Moving " + file + " to Programs")
				moveFile(path+"/"+file, path+"/"+"Programs/"+file)
			default:
				dest := path + "/" + "Others"
				createFolder(fileName, dest)
				moveFile(path+"/"+file, dest+"/"+fileName+"/"+file)
				log.Println("Moving " + file + " to Others/" + fileName)
			}
		} else {
			if file.Type().IsRegular() {
				log.Println("Moving " + file.Name() + " to Others")
				moveFile(path+"/"+file.Name(), path+"/"+"Others/"+file.Name())
			} else if file.IsDir() {
				switch file.Name() {
				case "Documents":
				case "Text Files":
				case "Pictures":
				case "Music":
				case "Videos":
				case "Compressed":
				case "Programs":
				case "Others":
				case "Folders":
				default:
					log.Println("Moving " + file.Name() + " to Folders")
					moveFile(path+"/"+file.Name(), path+"/"+"Folders/"+file.Name())
				}
			}
		}
	}
}

func main() {
	dest := os.Getenv("SORT_FOLDER_DESTINATION")
	if dest == "" {
		log.Default().Fatalln("Please set the SORT_FOLDER_DESTINATION environment variable")
	}
	// save all logs to a file
	f, err := os.OpenFile(dest+"/sort_folder.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	createFolders(dest)
	for {
		sortFilesToFolders(dest)
		time.Sleep(time.Minute * 5)
	}
}

