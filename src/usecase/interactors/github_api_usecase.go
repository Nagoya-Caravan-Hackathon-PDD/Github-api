package interactors

import (
	"fmt"
	"math"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/input"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/datastructure/types"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/usecase/dai"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/usecase/ports"
)

type GithubAPIInteractor struct {
	store       dai.GithubAPIDai
	gitmonStore dai.GitmonDai
	outputport  ports.GithubAPIOutput
}

func NewGithubAPIInteractor(store dai.GithubAPIDai, gitmonStore dai.GitmonDai, outputport ports.GithubAPIOutput) *GithubAPIInteractor {
	return &GithubAPIInteractor{
		store,
		gitmonStore,
		outputport,
	}
}

func (i *GithubAPIInteractor) CreateGitmon(reqBody input.GithubAPIRequest) (int, *output.GithubAPIResponse) {
	var err error
	fmt.Println(reqBody)
	if reqBody.GithubID == "" {
		return i.outputport.GithubAPI(types.CreateGitmon{}, err)
	}

	status, err := i.store.GetStatus(reqBody.GithubID)
	if err != nil {
		return i.outputport.GithubAPI(types.CreateGitmon{}, err)
	}

	// ここでぎっともんの基礎ステータスを計算して作る
	exp := countExp(status)
	level := calcLevel(exp)
	langPercent := calcLangPercent(status)
	baseGitmonStatus := calcBaseGitmonStatus(exp, langPercent)
	fmt.Println(baseGitmonStatus)
	//1. status から必要情報抜き出して ぎっともんの基礎ステータスを計算して作る Done
	//2. そのデータを基にぎっともんをDBに保存する => dai  (gateways)
	//3. ぎっともんのデータを返す => outputport (presenter)
	// 															↓ ここにデータ入れる
	arg := types.CreateGitmon{
		Owner:   status.User.Login,
		Name:    status.User.Name,
		Level:   level,
		Exp:     exp,
		HP:      baseGitmonStatus.HP,
		Attack:  baseGitmonStatus.Attack,
		Defense: baseGitmonStatus.Defense,
		Speed:   baseGitmonStatus.Speed,
	}
	if err := i.gitmonStore.Create(arg); err != nil {
		return i.outputport.GithubAPI(arg, err)
	}

	return i.outputport.GithubAPI(arg, nil)
}

func countExp(status *types.GitHubStatusQuery) int {
	contributions := status.User.ContributionsCollection.ContributionCalendar.TotalContributions
	followers := status.User.Followers.TotalCount
	exp := contributions*5 + followers*10
	return exp
}
func calcLevel(exp int) int {
	var requiredExp, level int
	for exp >= requiredExp {
		level++
		switch {
		case level < 50:
			requiredExp = level*13 + int(math.Pow(float64(level), 2))*3
		case level < 85:
			requiredExp = level*13 + int(math.Pow(float64(level), 2))*3 + int(math.Pow(float64(level-49), 3))
		default:
			requiredExp = level*13 + int(math.Pow(float64(level), 2))*3 + int(math.Pow(float64(level-49), 3)) + int(math.Pow(float64(level-84), 5))
		}
	}
	return level
}

func calcLangPercent(status *types.GitHubStatusQuery) map[string]int {
	langCount := make(map[string]int)
	langSize := make(map[string]int)
	langPercent := make(map[string]int)
	totalCount := 0

	for _, repo := range status.User.Repositories.Nodes {
		for _, lang := range repo.Languages.Edges {
			// ここで言語比率を計算してぎっともんのステータスを計算して作る
			langName := lang.Node.Name
			langCount[langName]++
			langSize[langName] += lang.Size
			totalCount += lang.Size
		}
	}
	for lang, _ := range langCount {
		langPercent[lang] = langSize[lang] / totalCount * 100
	}
	return langPercent
}

func calcBaseGitmonStatus(exp int, langPercent map[string]int) *types.CreateGitmon {
	// ここでぎっともんのステータスを計算して作る
	var hp, attack, defense, speed int
	for lang, percent := range langPercent {
		switch lang {
		case "Python":
			hp += 800 * percent / 100
			attack += 385 * percent / 100
			defense += 150 * percent / 100
			speed += 220 * percent / 100
		case "C":
			hp += 1275 * percent / 100
			attack += 130 * percent / 100
			defense += 130 * percent / 100
			speed += 460 * percent / 100
		case "C++":
			hp += 1000 * percent / 100
			attack += 145 * percent / 100
			defense += 180 * percent / 100
			speed += 440 * percent / 100
		case "Java":
			hp += 674 * percent / 100
			attack += 205 * percent / 100
			defense += 230 * percent / 100
			speed += 360 * percent / 100
		case "C#":
			hp += 575 * percent / 100
			attack += 250 * percent / 100
			defense += 210 * percent / 100
			speed += 340 * percent / 100
		case "VisualBasic":
			hp += 800 * percent / 100
			attack += 235 * percent / 100
			defense += 80 * percent / 100
			speed += 140 * percent / 100
		case "JavaScript":
			hp += 700 * percent / 100
			attack += 325 * percent / 100
			defense += 160 * percent / 100
			speed += 300 * percent / 100
		case "SQL":
			hp += 1225 * percent / 100
			attack += 250 * percent / 100
			defense += 150 * percent / 100
			speed += 300 * percent / 100
		case "ASM":
			hp += 1900 * percent / 100
			attack += 115 * percent / 100
			defense += 70 * percent / 100
			speed += 480 * percent / 100
		case "PHP":
			hp += 700 * percent / 100
			attack += 340 * percent / 100
			defense += 120 * percent / 100
			speed += 260 * percent / 100
		case "R":
			hp += 750 * percent / 100
			attack += 295 * percent / 100
			defense += 110 * percent / 100
			speed += 200 * percent / 100
		case "Go":
			hp += 350 * percent / 100
			attack += 280 * percent / 100
			defense += 200 * percent / 100
			speed += 400 * percent / 100
		case "Swift":
			hp += 225 * percent / 100
			attack += 175 * percent / 100
			defense += 220 * percent / 100
			speed += 380 * percent / 100
		case "Ruby":
			hp += 700 * percent / 100
			attack += 370 * percent / 100
			defense += 140 * percent / 100
			speed += 240 * percent / 100
		case "Perl":
			hp += 900 * percent / 100
			attack += 310 * percent / 100
			defense += 90 * percent / 100
			speed += 180 * percent / 100
		case "Rust":
			hp += 325 * percent / 100
			attack += 160 * percent / 100
			defense += 240 * percent / 100
			speed += 420 * percent / 100
		case "dart":
			hp += 300 * percent / 100
			attack += 190 * percent / 100
			defense += 100 * percent / 100
			speed += 160 * percent / 100
		case "Kotlin":
			hp += 300 * percent / 100
			attack += 220 * percent / 100
			defense += 190 * percent / 100
			speed += 320 * percent / 100
		case "TypeScript":
			hp += 275 * percent / 100
			attack += 265 * percent / 100
			defense += 170 * percent / 100
			speed += 280 * percent / 100
		case "HTML":
			hp += 575 * percent / 100
			attack += 355 * percent / 100
			defense += 150 * percent / 100
			speed += 300 * percent / 100
		default:
			hp += 725 * percent / 100
			attack += 250 * percent / 100
			defense += 150 * percent / 100
			speed += 300 * percent / 100
		}
	}
	return &types.CreateGitmon{
		HP:      hp,
		Attack:  attack,
		Defense: defense,
		Speed:   speed,
	}

}
