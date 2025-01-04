package logger

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	model "system-Info-collector/pkg/db/model"
	"time"

	"gorm.io/gorm"
)

var (
	LogEntries []*model.CbBatonLogTable
	mu         sync.Mutex
	DebugLevel = "DEBUG"
	InfoLevel  = "INFO"
	WarnLevel  = "WARN"
	ErrorLevel = "ERROR"
	FatalLevel = "FATAL"
)

//filelog, db log 남기기

/*
로그 테이블 정의
CREATE TABLE log_table (
    id INT AUTO_INCREMENT PRIMARY KEY,          -- 아이디: 고유 식별자
    log_level VARCHAR(255) NOT NULL,
		log_level ENUM('DEBUG', 'INFO', 'WARN', 'ERROR', 'FATAL') NOT NULL,  -- 로그 레벨: ENUM 타입으로 정의   ==> vachar 로 변경 요청
    target VARCHAR(255) NOT NULL,               -- 타겟: 문자열로 정의
    collector_name VARCHAR(255) NOT NULL,       -- collector_name: 문자열로 정의
    log_message TEXT NOT NULL,                   -- 로그 내용: 긴 문자열을 저장할 수 있도록 TEXT 타입으로 정의
    created_at TIMESTAMP -- 로그 생성 시간: 자동으로 현재 시간 저장
    insert_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- 로그 생성 시간: 자동으로 현재 시간 저장
);

로그 쿼리 형식 정의

INSERT INTO log_table (log_level, target, collector_name, log_message)
VALUES
    ('INFO', 'Target1', 'CollectorA', '로그 내용 1'),
    ('ERROR', 'Target2', 'CollectorB', '로그 내용 2'),
    ('WARN', 'Target3', 'CollectorC', '로그 내용 3'),
    ('DEBUG', 'Target4', 'CollectorD', '로그 내용 4'),
    ('FATAL', 'Target5', 'CollectorE', '로그 내용 5');

gorm model 정의

type LogEntry struct {
    ID            uint   `gorm:"primaryKey"`
    LogLevel      string `gorm:"type:enum('DEBUG', 'INFO', 'WARN', 'ERROR', 'FATAL');not null"`
    Target        string `gorm:"type:varchar(255);not null"`
    CollectorName string `gorm:"type:varchar(255);not null"`
    LogMessage    string `gorm:"type:text;not null"`
    CreatedAt     time.Time

	==> labels 추가
}

*/

func LogSetup(LogCollectTime, LogSaveTime, LogFileClearTime int, LogLevel string, db *gorm.DB) {
	go startLogging(LogCollectTime, db, LogLevel)
	go clearLog(LogSaveTime, db)
	go scheduleFolderDeletion(LogFileClearTime)
}

func Debug(target, logMessage string, logMessageV ...any) {
	pc, fileName, fileLine, _ := runtime.Caller(1)
	saveLog(DebugLevel, target, fmt.Sprintf("%s:%s:%d", runtime.FuncForPC(pc), fileName, fileLine), fmt.Sprintf(logMessage, logMessageV...))
}

func Info(target, logMessage string, logMessageV ...any) {
	pc, fileName, fileLine, _ := runtime.Caller(1)
	saveLog(InfoLevel, target, fmt.Sprintf("%s:%s:%d", runtime.FuncForPC(pc).Name(), fileName, fileLine), fmt.Sprintf(logMessage, logMessageV...))
}

func Warn(target, logMessage string, logMessageV ...any) {
	pc, fileName, fileLine, _ := runtime.Caller(1)
	saveLog(WarnLevel, target, fmt.Sprintf("%s:%s:%d", runtime.FuncForPC(pc).Name(), fileName, fileLine), fmt.Sprintf(logMessage, logMessageV...))
}

func Error(target, logMessage string, logMessageV ...any) {
	pc, fileName, fileLine, _ := runtime.Caller(1)
	saveLog(ErrorLevel, target, fmt.Sprintf("%s:%s:%d", runtime.FuncForPC(pc).Name(), fileName, fileLine), fmt.Sprintf(logMessage, logMessageV...))
}

func Fatal(target, logMessage string, logMessageV ...any) {
	pc, fileName, fileLine, _ := runtime.Caller(1)
	saveLog(FatalLevel, target, fmt.Sprintf("%s:%s:%d", runtime.FuncForPC(pc).Name(), fileName, fileLine), fmt.Sprintf(logMessage, logMessageV...))
}

// 로그 저장 함수
// SaveLog 함수: 로그 정보를 수집하고 저장
func saveLog(logLevel, target, detail, logMessage string) {
	log.Printf("LogLevel:(%s), Target:(%s), detail:(%s), LogMessage:(%s)", logLevel, target, detail, logMessage)
	mu.Lock() // 동기화 시작
	LogEntries = append(LogEntries, &model.CbBatonLogTable{
		LogLevel:   logLevel,
		Target:     target,
		Detail:     detail,
		LogMessage: logMessage,
		CreatedAt:  time.Now(),
	})
	mu.Unlock() // 동기화 종료
}

// saveLogs 함수: 로그를 데이터베이스에 한 번에 insert
func saveLogs(db *gorm.DB, loglevel string) error {
	var logEntriesfilter []*model.CbBatonLogTable
	// GORM을 사용하여 다중 삽입
	switch loglevel {
	case "DEBUG":
		for i := range LogEntries {
			if LogEntries[i].LogLevel == "DEBUG" {
				logEntriesfilter = append(logEntriesfilter, LogEntries[i])
			}
		}
	case "INFO":
		for i := range LogEntries {
			if LogEntries[i].LogLevel == "INFO" {
				logEntriesfilter = append(logEntriesfilter, LogEntries[i])
			}
		}
	case "WARN":
		for i := range LogEntries {
			if LogEntries[i].LogLevel == "WARN" {
				logEntriesfilter = append(logEntriesfilter, LogEntries[i])
			}
		}
	case "ERROR":
		for i := range LogEntries {
			if LogEntries[i].LogLevel == "ERROR" {
				logEntriesfilter = append(logEntriesfilter, LogEntries[i])
			}
		}
	case "FATAL":
		for i := range LogEntries {
			if LogEntries[i].LogLevel == "FATAL" {
				logEntriesfilter = append(logEntriesfilter, LogEntries[i])
			}
		}
	case "ALL":
		logEntriesfilter = LogEntries
	}
	result := db.Create(&logEntriesfilter)
	if result.Error != nil {
		Error("LOG_MODULE", "Error inserting log entry: %v\n", result.Error)
	}
	return nil
}

// StartLogging 함수: t 분마다 로그 저장
func startLogging(t int, db *gorm.DB, loglevel string) {
	ticker := time.NewTicker(time.Duration(t) * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		mu.Lock() // 동기화 시작
		if len(LogEntries) > 0 {
			err := saveLogs(db, loglevel)
			if err != nil {
				Error("LOG_MODULE", "Error saving logs: %s", err.Error())
			}
			LogEntries = nil // 로그 저장 후 슬라이스 초기화
		}
		mu.Unlock() // 동기화 종료
	}
}

func clearLog(t int, db *gorm.DB) {
	ticker := time.NewTicker(time.Duration(t) * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		mu.Lock() // 동기화 시작

		// 1. DB에서 모든 로그 데이터 읽기
		var logs []*model.CbBatonLogTable
		if err := db.Find(&logs).Error; err != nil { // &logs로 포인터 전달
			Error("LOG_MODULE", "Error reading from database: %s", err.Error())
			mu.Unlock()
			continue
		}

		// 로그 폴더가 존재하지 않으면 생성
		if _, err := os.Stat("logs"); os.IsNotExist(err) {
			err := os.Mkdir("logs", os.ModePerm)
			if err != nil {
				Error("LOG_MODULE", "Error creating logs directory: %s", err.Error())
				mu.Unlock()
				continue
			}
		}

		// 2. 현재 날짜와 시간으로 파일 이름 생성
		timestamp := time.Now().Format("20060102_150405") // YYYYMMDD_HHMMSS 형식
		fileName := fmt.Sprintf("logs/logs_%s.csv", timestamp)

		// 파일에 로그 데이터 저장
		file, err := os.Create(fileName)
		if err != nil {
			Error("LOG_MODULE", "Error creating file: %s", err.Error())
			mu.Unlock()
			continue
		}

		// UTF-8 BOM 추가
		_, err = file.WriteString("\xEF\xBB\xBF")
		if err != nil {
			Error("LOG_MODULE", "Error writing BOM to file: %s", err.Error())
			file.Close()
			mu.Unlock()
			continue
		}

		writer := csv.NewWriter(file)

		// CSV 헤더 작성
		writer.Write([]string{"ID", "LogLevel", "Target", "Detail", "LogMessage", "CreatedAt", "InsertAt"})

		// 로그 데이터를 CSV 파일에 기록
		for _, log := range logs {
			record := []string{
				fmt.Sprint(log.ID),
				log.LogLevel,
				log.Target,
				log.Detail,
				log.LogMessage,
				log.CreatedAt.String(),
				log.InsertAt.String(),
			}
			if err := writer.Write(record); err != nil {
				Error("LOG_MODULE", "Error writing record to file: %s", err.Error())
			}
		}

		writer.Flush()

		// 파일 닫기
		if err := file.Close(); err != nil {
			Error("LOG_MODULE", "Error closing file: %s", err.Error())
		}

		// 3. DB에서 모든 로그 데이터 삭제
		if err := db.Exec("DELETE FROM cb_baton_log_table").Error; err != nil {
			Error("LOG_MODULE", "Error deleting from database: %s", err.Error())
		}
		mu.Unlock() // 동기화 종료
	}
}

// 특정 폴더를 삭제하는 함수
func scheduleFolderDeletion(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		err := os.RemoveAll("logs")
		if err != nil {
			Error("LOG_MODULE", "Error deleting folder: %s", err.Error())
		} else {
			Info("LOG_MODULE", "Successfully deleted folder: %s\n", "/logs")
		}
	}
}
