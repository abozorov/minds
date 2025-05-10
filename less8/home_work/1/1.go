package main

import "fmt"

type StorageI interface {
	Save(data string)
	Load()
}

// Структура для сохранения файлов в массив файлов
type FileStorage struct {
	Files []string
}

func newFileStorage() *FileStorage {
	return &FileStorage{
		Files: make([]string, 0, 100),
	}
}

func (fs *FileStorage) Save(fileName string) {
	fs.Files = append(fs.Files, fileName)
	fmt.Printf("Файл \"%s\" добавлен в хранилище fileStr!\n", fileName)
}

func (fs FileStorage) Load() {
	if len(fs.Files) == 0 {
		fmt.Println("Хранилище пустое!")
		return
	}

	fmt.Println("Файлы в fileStr:")
	for _, v := range fs.Files {
		fmt.Printf(" - Файл: %s\n", v)
	}
}

// Структура для хранения длины названия файлов
type MemoryStorage struct {
	Files  []string
	Memory map[string]int
}

func newMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Files:  make([]string, 0, 100),
		Memory: make(map[string]int),
	}
}

func (ms *MemoryStorage) Save(fileName string) {
	ms.Files = append(ms.Files, fileName)
	ms.Memory[fileName] = len(fileName)
	fmt.Printf("Файл \"%s\" добавлен в хранилище memStr!\n", fileName)
}

func (ms MemoryStorage) Load() {

	if len(ms.Files) == 0 {
		fmt.Println("Хранилище пустое!")
		return
	}

	fmt.Println("Файлы в memStr:")
	for _, v := range ms.Files {
		fmt.Printf(" - Название файла \"%s\" и длина названия %d\n", v, ms.Memory[v])
	}
}


func main() {
	var storage StorageI

	storage = newFileStorage()
	storage.Save("<UNK> <UNK> <UNK>")
	storage.Save("asdas reawea")
	storage.Load()

	storage = newMemoryStorage()
	storage.Save("<UNK> <UNK> JHJHHJ")
	storage.Save("JSAGDHFL AWJEYWYYY")
	storage.Save("Aayuyu uyuyuy")
	storage.Load()
}
