package types

type CreateGitmon struct {
	Owner          string
	Name           string
	Image          int
	Level          int
	Exp            int
	BaseHP         int
	CurrentHP      int
	BaseAttack     int
	CurrentAttack  int
	BaseDefense    int
	CurrentDefense int
	BaseSpeed      int
	CurrentSpeed   int
}
