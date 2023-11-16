package interactors

import (
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

func (i *GithubAPIInteractor) GetGithubStatus(reqQuery input.GithubAPIRequest) (int, *output.GithubAPIResponse) {
	var err error

	if reqQuery.GithubID == "" {
		return i.outputport.GithubAPI(err)
	}

	status, err := i.store.GetStatus(reqQuery.GithubID)
	if err != nil {
		return i.outputport.GithubAPI(err)
	}

	// ここでぎっともんの基礎ステータスを計算して作る
	exp := countExp(status)
	level := calcLevel(exp)
	langPercent := calcLangPercent(status)

	//1. status から必要情報抜き出して ぎっともんの基礎ステータスを計算して作る
	//2. そのデータを基にぎっともんをDBに保存する => dai  (gateways)
	//3. ぎっともんのデータを返す => outputport (presenter)
	// 															↓ ここにデータ入れる
	if err := i.gitmonStore.Create(types.CreateGitmon{}); err != nil {
		return i.outputport.GithubAPI(err)
	}

	return i.outputport.GithubAPI(nil)
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
	for _, lang := range langCount {
		langPercent[lang] = langSize[lang] / totalCount * 100
	}
	return langPercent
}
