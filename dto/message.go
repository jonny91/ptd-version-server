package dto

type Message struct {
	Message string
}

type MissionResult struct {
	Time      string `gorm:"type:timestamp"`
	Platform  string `gorm:"type:varchar(100)"`
	MissionId string `gorm:"type:varchar(20)"`
	State     string `gorm:"type:varchar(10)"`
	Cards     string `gorm:"type:varchar(100)"`
}

func (MissionResult) TableName() string {
	return "mission_logs"
}
