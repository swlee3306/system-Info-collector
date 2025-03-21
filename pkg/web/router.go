package web

import (
	"system-Info-collector/internal/collector/cpu"
	"system-Info-collector/internal/collector/disk"
	"system-Info-collector/internal/collector/memory"
	"system-Info-collector/internal/collector/network"
	"system-Info-collector/pkg/logger"

	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	fnc := "StartWebServer"
	logger.Info("fnc", "Starting web server[%s]", fnc)

	r := gin.Default()

	r.Use(gin.Recovery())

	apiGroup := r.Group("/api/v1.0")
	{
		//cpu 정보를 가져와서 응답주는 API
		apiGroup.GET("/cpuinfo", func(c *gin.Context) {
			cpuInfos, err := cpu.GetCpuInfo()
			if err != nil {
				logger.Error("fnc", "Error getting CPU info: %v", err)
				c.JSON(500, gin.H{"error": "Failed to get CPU info"})
				return
			}
			c.JSON(200, cpuInfos)
		})

		//memory 정보를 가져와서 응답주는 API
		apiGroup.GET("/memoryinfo", func(c *gin.Context) {
			memoryInfos, err := memory.GetMemoryInfo()
			if err != nil {
				logger.Error("fnc", "Error getting memory info: %v", err)
				c.JSON(500, gin.H{"error": "Failed to get memory info"})
				return
			}
			c.JSON(200, memoryInfos)
		})

		//Disk 정보를 가져와서 응답주는 API

		apiGroup.GET("/diskinfo", func(c *gin.Context) {
			diskInfos, err := disk.GetDiskInfo()
			if err != nil {
				logger.Error("fnc", "Error getting disk info: %v", err)
				c.JSON(500, gin.H{"error": "Failed to get disk info"})
				return
			}
			c.JSON(200, diskInfos)
		})

		//Network Interface 정보를 가져와서 응답주는 API

		apiGroup.GET("/networkinfo", func(c *gin.Context) {
			networkInfos, err := network.GetNetworkInfo()
			if err != nil {
				logger.Error("fnc", "Error getting network info: %v", err)
				c.JSON(500, gin.H{"error": "Failed to get network info"})
				return
			}
			c.JSON(200, networkInfos)
		})

		err := r.Run(":8080")
		if err != nil {
			logger.Error(fnc, "%s error: %v", fnc, err)
		}
	}
}
