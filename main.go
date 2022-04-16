package main

import (
	"flag"
	"fmt"
	"github.com/DORE145/geobase/controllers"
	"github.com/DORE145/geobase/loader"
	"github.com/DORE145/geobase/service"
	"github.com/DORE145/geobase/storage/inmemory"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	verbosity string
	filePath  string
	port      int
)

func init() {
	flag.IntVar(&port, "port", 8080, "port to serve (default: 8080)")
	flag.StringVar(&verbosity, "verbosity", "debug", "logs verbosity (default: debug)")
	flag.StringVar(&filePath, "file", "geobase.dat", "database file (default: geobase.dat)")
}

func main() {
	logLevel, err := logrus.ParseLevel(verbosity)
	if err != nil {
		logrus.Warning("Failed to parse verbosity flag, failing back to DEBUG level")
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)

	// Data loading
	dataLoader, err := loader.NewBinaryDataLoader(filePath)
	if err != nil {
		logrus.Fatalf("Failed to create data dataLoader with error: %s", err)
	}

	ranges, err := dataLoader.LoadIPRanges()
	if err != nil {
		logrus.Fatalf("Failed to read data IP ranges data from file with error: %s", err)
	}

	locations, err := dataLoader.LoadLocations()
	if err != nil {
		logrus.Fatalf("Failed to read data locations data from file with error: %s", err)
	}

	locationsIndex, err := dataLoader.LoadLocationsCityIndex(locations)
	if err != nil {
		logrus.Fatalf("Failed to read data locations index data from file with error: %s", err)
	}

	// Creating models and routes specific entities
	locationsStorage, err := inmemory.NewLocationStorage(locations, locationsIndex)
	if err != nil {
		logrus.Fatalf("Failed to build inmemory locations storage with error: %s", err)
	}
	locationsService := service.NewLocationService(locationsStorage)
	locationsController := controllers.NewLocationController(*locationsService)

	ipRangeStorage := inmemory.NewIpRangeStorage(ranges)
	ipRangeService := service.NewIpRangeService(ipRangeStorage, locationsStorage)
	ipRangeController := controllers.NewIPRangeController(*ipRangeService)

	//Server initializing
	server := gin.Default()
	rootGroup := server.Group("")
	locationsController.RegisterLocationRouts(rootGroup)
	ipRangeController.RegisterIPRangeRoutes(rootGroup)
	address := fmt.Sprintf(":%d", port)

	err = server.Run(address)
	if err != nil {
		logrus.Fatalf("Server crashed with error: %s", err)
	}
}
