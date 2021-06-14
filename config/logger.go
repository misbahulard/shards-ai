package config

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/misbahulard/shards-ai/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ConfigureLogger() error {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := f.File[strings.LastIndex(f.File, "/")+1:] + ":" + strconv.Itoa(f.Line)
			fnName := f.Function[strings.LastIndex(f.Function, ".")+1:]
			return fnName, fileName
		},
	})

	if viper.GetBool("log.debug") {
		log.SetLevel(log.DebugLevel)
	}

	// check if we need to store log to file
	if viper.GetBool("log.file.enable") {
		if viper.GetString("log.file.path") == "" {
			fmt.Println("You enable the file logging, but not define the log path")
			os.Exit(1)
		}

		err := util.CreateDirectoryByFile(viper.GetString("log.file.path"))
		if err != nil {
			return err
		}

		file, err := os.OpenFile(viper.GetString("log.file.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		log.SetOutput(file)
	}

	log.Info("Logger: ok")
	return nil
}
