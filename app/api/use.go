package api

import (
	"bitbucket.org/kyicy/seifer/config"
	"github.com/imroc/req"
)

//SentenceToVector function
func SentenceToVector(sentence string) []float64 {

	useaddr := config.Get().API.USE
	res, _ := req.Post(useaddr, req.BodyJSON(makeUseBody(sentence)))

	out := new(useOutput)
	res.ToJSON(out)

	return out.Outputs[0]
}

type useBody struct {
	Inputs []string `json:"inputs"`
}

type useOutput struct {
	Outputs [][]float64 `json:"outputs"`
}

func makeUseBody(sentence string) *useBody {
	return &useBody{
		Inputs: []string{sentence},
	}
}
