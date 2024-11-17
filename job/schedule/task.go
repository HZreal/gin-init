package schedule

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: task.go
 * @Description:
 */

import (
	"log"
	"time"
)

func CheckDevice() {
	log.Println("任务运行于 ---------------------------->", time.Now())
}
