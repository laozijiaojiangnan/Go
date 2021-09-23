package module

// Human 人类表
type Human struct {
	BaseModule
	Name   string  `gorm:"size:16;unique;not null"` // 人类名称
	Weight float32 `gorm:"not null"`                // 人类体重
	Work   string  `gorm:"default:null"`            // 人类职业
	Slogan string  `gorm:"default:null"`            // 座右铭

	// 一对一， 一个人类对应一个英雄，一个英雄对应一个人类
	HeroID uint `gorm:"default:null"` // 扮演的英雄ID
	Hero   Hero // 扮演的英雄
}
