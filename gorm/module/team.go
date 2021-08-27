package module

// HeroTeam 英雄队伍表
type HeroTeam struct {
	BaseModule
	Name string `gorm:"size:16;unique;not null"` // 队伍名称

	Heroes []Hero // 队伍里的英雄
}
