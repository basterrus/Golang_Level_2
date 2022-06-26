package internal

import (
	"fmt"
	"github.com/basterrus/Golang_Level_2/Lesson_8/config"
	"os"
	"sort"
	"strings"
)

func SearchFiles(conf *config.Config) error {
	fInfo := filesInfo{}
	fInfo.directoryList = append(fInfo.directoryList, conf.DirectoryPath)

	err := findAllFilesInDir(conf, &fInfo)
	if err != nil {
		conf.ErrorLogger.Printf("error on find all files in source path: %s", err)
		return err
	}

	// сортируем файлы, файлы в приоритете корневого каталога считаются как исходные
	sort.Slice(fInfo.allFilesList, func(i, j int) bool {
		return fInfo.allFilesList[i].Path < fInfo.allFilesList[j].Path
	})
	// сравниваем файлы
	for i, file := range fInfo.allFilesList {
		if file.contains(fInfo.allFilesList, i) {
			fmt.Printf("Duplicate file: %s	Original file: %s\n", file.Path, file.OriginalFile.Path)
			fInfo.duplicateFilesList = append(fInfo.duplicateFilesList, file)
		}
	}
	fmt.Printf("Total files: %d\n", len(fInfo.allFilesList))
	fmt.Printf("Duplicate files (without original file): %d\n", len(fInfo.duplicateFilesList))

	// удаляем файлы, если получаем флаг, получаем одобрение от пользователя
	if conf.FlagDelete || conf.RunInTest {
		if len(fInfo.duplicateFilesList) == 0 {
			fmt.Println("No files for delete!")
		} else {
			var confirm string
			if !conf.RunInTest {
				for strings.ToUpper(confirm) != "Y" && strings.ToUpper(confirm) != "N" {
					fmt.Print("Delete this duplicate files? (Y/N): ")
					_, err = fmt.Fscan(os.Stdin, &confirm)
					if err != nil {
						conf.ErrorLogger.Printf("error on get approval from console: %s\n", err)
						return err
					}
				}
			}
			if strings.ToUpper(confirm) == "Y" || conf.RunInTest {
				err = deleteFiles(conf, &fInfo)
				if err != nil {
					conf.ErrorLogger.Printf("error on delete files: %s\n", err)
					return err
				}
				fmt.Println("Files deleted!")
			}
		}
	}

	return nil
}

func deleteFiles(conf *config.Config, fInfo *filesInfo) error {
	wp := newWorkerPool(conf.CountGoroutine)
	defer wp.wg.Wait()

	for _, file := range fInfo.duplicateFilesList {
		wp.wg.Add(1)
		go func(file FileEntity) {
			defer func() {
				wp.mutex.Unlock()
				// read to release a slot
				<-wp.semaphoreCh
				fInfo.deleteFilesList = append(fInfo.deleteFilesList, file)
				wp.wg.Done()
			}()
			// block while full
			wp.semaphoreCh <- struct{}{}
			wp.mutex.Lock()
			if err := os.Remove(file.Path); err != nil {
				conf.ErrorLogger.Printf("error on delete file %s: %s\n", file.Path, err)
			}
		}(file)
	}

	return nil
}
