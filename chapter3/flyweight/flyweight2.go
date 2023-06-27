package main

import "fmt"

//服装接口
type Dress interface {
	GetColor() string
}

//蓝队服装
type BlueTeamDress struct {
	color string
}

func (t *BlueTeamDress) GetColor() string {
	return t.color
}

func newBlueTeamDress() *BlueTeamDress {
	return &BlueTeamDress{color: "blue"}
}

//创建红队服装
type RedTeamDress struct {
	color string
}

func (c *RedTeamDress) GetColor() string {
	return c.color
}

func newRedTeamDress() *RedTeamDress {
	return &RedTeamDress{color: "red"}
}

const (
	//蓝队服装类型
	BlueTeamDressType = "Blue Dress"
	//红队服装类型
	RedTeamDressType = "Red Dress"
)

var (
	DressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]Dress),
	}
)

//享元服装工厂
type DressFactory struct {
	DressMap map[string]Dress
}

//获取服装类型
func (d *DressFactory) GetDressByType(DressType string) (Dress, error) {
	if d.DressMap[DressType] != nil {
		return d.DressMap[DressType], nil
	}

	if DressType == BlueTeamDressType {
		d.DressMap[DressType] = newBlueTeamDress()
		return d.DressMap[DressType], nil
	}
	if DressType == RedTeamDressType {
		d.DressMap[DressType] = newRedTeamDress()
		return d.DressMap[DressType], nil
	}

	return nil, fmt.Errorf("%s", "Wrong Dress type")
}

//获取服装工厂单例
func GetDressFactorySingleInstance() *DressFactory {
	return DressFactorySingleInstance
}

//队员类
type Player struct {
	Dress      Dress
	PlayerType string
	lat        int
	long       int
}

//创建一个队员
func NewPlayer(PlayerType, DressType string) *Player {
	Dress, _ := GetDressFactorySingleInstance().GetDressByType(DressType)
	return &Player{
		PlayerType: PlayerType,
		Dress:      Dress,
	}
}

//创建队员位置
func (p *Player) NewLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

//创建游戏
type NewGame struct {
}

//创建蓝队队员
func (ng *NewGame) AddBlueTeam(DressType string) *Player {
	return NewPlayer("terrorist", DressType)
}

//创建红队队员
func (ng *NewGame) AddRedTeam(DressType string) *Player {
	return NewPlayer("counterBlueTeam", DressType)
}

func main() {
	game := NewGame{}

	//创建红队
	game.AddBlueTeam(BlueTeamDressType)
	game.AddBlueTeam(BlueTeamDressType)
	game.AddBlueTeam(BlueTeamDressType)
	game.AddBlueTeam(BlueTeamDressType)

	//创建蓝队
	game.AddRedTeam(RedTeamDressType)
	game.AddRedTeam(RedTeamDressType)
	game.AddRedTeam(RedTeamDressType)

	DressFactoryInstance := GetDressFactorySingleInstance()

	for DressType, Dress := range DressFactoryInstance.DressMap {
		fmt.Printf("服装类型: %s\n服装颜色: %s\n", DressType, Dress.GetColor())
	}
}
