package internal

import (
	"github.com/basterrus/Golang_Level_2/Lesson_8/config"
	"os"
	"sync"
	"time"
)

// workerPool управляющая структура
type workerPool struct {
	wg          sync.WaitGroup
	resultCh    chan FileEntity
	semaphoreCh chan struct{}
	mutex       sync.Mutex
}

// newWorkerPool метод инициализирует новый workerPool и возвращает указатель на него
func newWorkerPool(N int) *workerPool {
	return &workerPool{
		wg:          sync.WaitGroup{},
		resultCh:    make(chan FileEntity, N),
		semaphoreCh: make(chan struct{}, N),
	}
}

// filesInfo структура для сохранения промежуточного результата работы
type filesInfo struct {
	allFilesList       []FileEntity // list with all files
	duplicateFilesList []FileEntity // list with duplicate files
	deleteFilesList    []FileEntity // list with deleted files
	randomFilesList    []FileEntity // list with random create files
	directoryList      []string
}

// newFileEntity метод инициализирует новый FileEntity, и возвращает указатель на него
func newFileEntity() *FileEntity {
	return &FileEntity{
		Create: time.Now(),
		Name:   "",
		Path:   "",
		Hash:   "",
		Size:   0,
	}
}

// FileEntity структура c информацией о файле
type FileEntity struct {
	OriginalFile *FileEntity // pointer to original file
	Create       time.Time   // time of create file
	Name         string      // name of file
	Path         string      // path to file with name of file
	Hash         string      // hash of file
	Size         int64       // size of file
}

// findAllFilesInDir функция находит все файлы в исходном каталоге без каталогов, сохраняет информацию о файлах в структуре filesInfo, возвращает ошибку
func findAllFilesInDir(cfg *config.Config, fInfo *filesInfo) error {
	wp := newWorkerPool(cfg.CountGoroutine)
	defer wp.wg.Wait()

	wp.wg.Add(1)
	lsFiles(cfg.DirectoryPath, cfg, wp, fInfo)

	return nil
}

// lsFiles рекурсивный поиск файлов в каталогах/подкаталогах
func lsFiles(dir string, cfg *config.Config, wp *workerPool, fInfo *filesInfo) {

	wp.semaphoreCh <- struct{}{}

	go func() {
		defer func() {
			wp.mutex.Unlock()
			<-wp.semaphoreCh
			wp.wg.Done()
		}()

		wp.mutex.Lock()
		file, err := os.Open(dir)
		if err != nil {
			cfg.ErrorLogger.Printf("error opening directory: %s\n", err)
		}

		defer fileClose(cfg, file)

		files, err := file.Readdir(-1)
		if err != nil {
			cfg.ErrorLogger.Printf("error reading directory: %s\n", err)
		}

		for _, f := range files {
			path := dir + "/" + f.Name()
			if f.IsDir() {
				fInfo.directoryList = append(fInfo.directoryList, path)
				wp.wg.Add(1)
				go lsFiles(path, cfg, wp, fInfo)
			} else {
				fe := newFileEntity()
				fe.Name = f.Name()
				fe.Path = path
				fe.Create = f.ModTime()
				//// get hash of file
				//if err = fe.getHashOfFile(cfg); err != nil {
				//	cfg.ErrorLogger.Printf("can't get hash of file %s: %s\n", path, err)
				//}
				fe.Size = f.Size()
				fInfo.allFilesList = append(fInfo.allFilesList, *fe)
			}
		}
	}()
}

// fileClose function for defer close file
func fileClose(conf *config.Config, file *os.File) {
	err := file.Close()
	if err != nil {
		conf.ErrorLogger.Printf("error on defer close file %s: %s\n", file.Name(), err)
	}
}

// contains method check contains file in slice of file from i position, return bool
func (f *FileEntity) contains(fl []FileEntity, it int) bool {
	for i := it + 1; i < len(fl); i++ {
		if fl[i].Hash == f.Hash {
			f.OriginalFile = &fl[i]
			return true
		}
	}
	return false
}
