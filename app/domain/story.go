package domain

type Story struct {
	id         string
	cocktailID string
	trivia     string
	day        string
}

type StoryRepository interface {
	FindByAll() (map[string]Cocktail, error)       // 引数はなしで、全ての日付で作るカクテル構造体を取得する
	FindByDay(string) (map[string]Cocktail, error) // 引数は日付で、その日付に作るカクテルの構造体を取得
}

func NewStory(
	id string,
	cocktailID string,
	trivia string,
	day string,
) Story {
	return Story{
		id:         id,
		cocktailID: cocktailID,
		trivia:     trivia,
		day:        day,
	}
}

func (story Story) ID() string {
	return story.id
}

func (story Story) CocktailID() string {
	return story.cocktailID
}

func (story Story) Trivia() string {
	return story.trivia
}

func (story Story) Day() string {
	return story.day
}
