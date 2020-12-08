package dto

type Message struct {
	Message string
}

type MissionResult struct {
	Time        string `gorm:"type:timestamp"`
	RoleName    string `gorm:"type:varchar(255)"`
	Platform    string `gorm:"type:varchar(100)"`
	DeviceModel string `gorm:"type:varchar(255)"`
	MissionId   string `gorm:"type:varchar(20)"`
	State       string `gorm:"type:varchar(10)"`
	Cards       string `gorm:"type:varchar(100)"`
	Duration    int    //战斗时长
}

func (MissionResult) TableName() string {
	return "mission_logs"
}
