package router

import (
	"fmt"
	"net/http"
	"sort"
	"sync"

	"bitbucket.org/kyicy/seifer/app/api"
	"bitbucket.org/kyicy/seifer/app/model"
	"github.com/labstack/echo/v4"
	"gonum.org/v1/gonum/mat"
)

const DIM = 512

type createUserStoryBody struct {
	Body          string `json:"body" form:"body"`
	Capability    string `json:"capability" from:"capability"`
	SubCapability string `json:"subcapability" from:"subcapability"`
	Epic          string `json:"epic" from:"epic"`
}

// CreateUserStory handler
func CreateUserStory(c echo.Context) error {
	body := new(createUserStoryBody)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	userStory := new(model.UserStory)
	userStory.Body = body.Body

	bodyVector := titleAndBodyToVectors(body.Body)

	userStory.Body = body.Body
	userStory.BodyVector = bodyVector
	userStory.Capability = body.Capability
	userStory.SubCapability = body.SubCapability
	userStory.Epic = body.Epic
	userStory.TagID = 1

	model.DB.Save(userStory)

	return c.NoContent(http.StatusOK)
}

type similarUserStoriesBody struct {
	Body  string `json:"body" form:"body"`
	TagID int    `json:"tagId" form:"tagId"`
}

// SimilarUserStories handler
func SimilarUserStories(c echo.Context) error {
	body := new(similarUserStoriesBody)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	bodyVector := titleAndBodyToVectors(body.Body)

	bodyDense := mat.NewDense(1, DIM, bodyVector)

	var userStories []model.UserStory

	model.DB.Where("tag_id = ?", body.TagID).Find(&userStories)

	// var ids []uint
	// var titleVecs, bodyVecs []float64
	var bodyVecs []float64
	for _, userStory := range userStories {
		// ids = append(ids, userStory.ID)
		// titleVecs = append(titleVecs, userStory.TitleVector...)
		bodyVecs = append(bodyVecs, userStory.BodyVector...)
	}

	bodyMatrix := mat.NewDense(len(userStories), DIM, bodyVecs)
	bodyMatrixT := bodyMatrix.T()

	var bodyRank mat.Dense
	bodyRank.Mul(bodyDense, bodyMatrixT)

	Score := mat.Row(nil, 0, &bodyRank)

	slice := NewSlice(Score...)

	sort.Sort(sort.Reverse(slice))

	var userStoryResult []*model.UserStory
	for i, index := range slice.idx {
		userStories[index].Score = Score[i]
		userStoryResult = append(userStoryResult, &userStories[index])
	}
	return c.JSON(http.StatusOK, userStoryResult)
}

// func titleAndBodyToVectors(title, body string) (titleVector []float64, bodyVector []float64) {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	// process title
// 	go func() {
// 		defer wg.Done()
// 		titleVector = api.SentenceToVector(title)
// 	}()

// 	// process body
// 	go func() {
// 		defer wg.Done()
// 		bodyVector = api.SentenceToVector(body)
// 	}()
// 	wg.Wait()

// 	return
// }

func titleAndBodyToVectors(body string) (bodyVector []float64) {
	var wg sync.WaitGroup
	wg.Add(1)
	// process body
	go func() {
		defer wg.Done()
		bodyVector = api.SentenceToVector(body)
	}()
	wg.Wait()

	return
}

// Slice struct
type Slice struct {
	sort.Float64Slice
	idx []int
}

// Swap for order
func (s Slice) Swap(i, j int) {
	s.Float64Slice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

// NewSlice helper
func NewSlice(n ...float64) *Slice {
	s := &Slice{Float64Slice: sort.Float64Slice(n), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

type createUserStoryBodyExpand struct {
	Body string `json:"body" form:"body"`
}

// CreateUserStoryExpand handler
func CreateUserStoryExpand(c echo.Context) error {
	body := new(createUserStoryBodyExpand)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	userStoryExpands := new(model.UserStoryExpand)
	userStoryExpands.Body = body.Body

	bodyVector := titleAndBodyToVectors(body.Body)

	userStoryExpands.Body = body.Body
	userStoryExpands.BodyVector = bodyVector

	model.DB.Save(userStoryExpands)

	return c.NoContent(http.StatusOK)
}

type similarUserStoriesBodyExpand struct {
	Body string `json:"body" form:"body"`
}

// SimilarUserStoriesExpand handler
func SimilarUserStoriesExpand(c echo.Context) error {
	body := new(similarUserStoriesBodyExpand)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	bodyVector := titleAndBodyToVectors(body.Body)

	bodyDense := mat.NewDense(1, DIM, bodyVector)

	var userStoryExpands []model.UserStoryExpand

	model.DB.Find(&userStoryExpands)

	var bodyVecs []float64
	for _, userStory := range userStoryExpands {
		bodyVecs = append(bodyVecs, userStory.BodyVector...)
	}

	bodyMatrix := mat.NewDense(len(userStoryExpands), DIM, bodyVecs)
	bodyMatrixT := bodyMatrix.T()

	var bodyRank mat.Dense
	bodyRank.Mul(bodyDense, bodyMatrixT)

	Score := mat.Row(nil, 0, &bodyRank)

	slice := NewSlice(Score...)

	sort.Sort(sort.Reverse(slice))

	var userStoryResult []*model.UserStoryExpand
outloop:
	for i, index := range slice.idx {
		userStoryExpands[index].Score = Score[i]
		userStoryResult = append(userStoryResult, &userStoryExpands[index])
		fmt.Print(userStoryResult)
		if i >= 4 {
			break outloop
		}
	}
	return c.JSON(http.StatusOK, userStoryResult)
}
