package main

import (
	"fmt"
	"github.com/SadikSunbul/IoTGo-Iot_Device_Simulation_Services/internal/server/http"
	"github.com/SadikSunbul/IoTGo-Iot_Device_Simulation_Services/pkg/config"
	"github.com/SadikSunbul/IoTGo-Iot_Device_Simulation_Services/pkg/loggerService"
)

func main() {
	// Config yükle
	cfg := config.LoadConfig()

	err := loggerService.InitLogger(cfg.LoogerLevel, cfg.LoggerOutputPath)
	if err != nil {
		fmt.Errorf("Error Logger : ", err.Error())
		return
	}

	httpSvr := http.NewServer()
	if err = httpSvr.Run(); err != nil {
		loggerService.Log.Fatal(err)
	}

}
