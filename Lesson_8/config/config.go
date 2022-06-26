package config

import (
	"flag"
	"github.com/kkyr/fig"
	"log"
	"path/filepath"
)

const (
	usagePath              = "Путь к директории в которой будет произведен поиск дубликатов"
	errorConvertPathToABS  = "Ошибка при получении абсолютного пути из переданного %q: %v\n"
	ErrorLoadConfiguration = "Ошибка при загрузке файла конфигурации %s\n"
)

type Config struct {
	DirectoryPath    string `fig:"sourcePath" default:"."`
	ErrorLogger      *log.Logger
	CountGoroutine   int `fig:"countGoroutine" default:"10"`
	CountRndCopyIter int `fig:"countRndCopyIter" default:"10"`
	SizeCopyBuffer   int `fig:"sizeCopyBuffer" default:"512"`
	FlagDelete       bool
	RunInTest        bool `fig:"runInTest"`
}

func Init() (*Config, error) {
	var conf = Config{}
	err := fig.Load(&conf, fig.Dirs("../", "./", "./..."), fig.File("config/config.yaml"))
	if err != nil {
		log.Fatalf(ErrorLoadConfiguration, err)
	}
	return &conf, err
}

//InitFlags метод инициализирует флаги и перобразует переданный пусть в абсолютный
func (conf *Config) InitFlags() error {
	flag.StringVar(&conf.DirectoryPath, "path", conf.DirectoryPath, usagePath)

	flag.Parse()
	if err := conf.setABSPath(); err != nil {
		conf.ErrorLogger.Fatalf(errorConvertPathToABS, conf.DirectoryPath, err)
		return err
	}
	return nil
}

//setABSPath метод готовит абсолютный путь из переданного в него пути
func (conf *Config) setABSPath() error {
	directoryPath, err := filepath.Abs(conf.DirectoryPath)
	if err != nil {
		return err
	}
	conf.DirectoryPath = directoryPath
	return nil
}
