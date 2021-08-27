package module

// Hero 英雄表
type Hero struct {
	BaseModule
	Name    string  `gorm:"size:16;unique;not null"` // 英雄名称
	Attack  float64 `gorm:"not null;default:0"`      // 英雄攻击力
	defense float64 `gorm:"not null;default:0"`      // 英雄防御力
	weapon  string  `gorm:"default:null"`            // 英雄武器

	// 多对多,一个英雄可以有多个超能力，一个超能力也可以属于多个英雄
	PowerList []SuperPower `gorm:"many2many:hero_power_relation;"`

	// 一对多,一个队伍可以有多个英雄，但每个英雄只能有一个队伍
	TeamID uint     // 所属的队伍ID
	Team   HeroTeam // 所属的队伍
}
