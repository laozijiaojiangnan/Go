package module

// SuperPower 超能力表
type SuperPower struct {
	BaseModule
	Name  string `gorm:"size:16;unique;not null"` // 能力名称
	Level string `gorm:"size:8;default:D""`       // 能力级别 [S,A,B,C,D]
}
