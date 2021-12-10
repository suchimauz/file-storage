package main

import "github.com/suchimauz/file-storage/internal/app"

const configDirPath string = "configs"

func main() {
	app.Run(configDirPath)
}
