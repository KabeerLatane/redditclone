// internal/simulation/behavior.go
package simulation

import "math/rand"
import "reddit-clone/internal/common"

type UserPersona struct {
	Name             string
	PostProb         float64
	CommentProb      float64
	VoteProb         float64
	JoinProb         float64
	ActiveHourRanges [][2]int // Array of [start, end] hour ranges
}

var userPersonas = []UserPersona{
	{
		Name:             "Casual",
		PostProb:         0.1,
		CommentProb:      0.3,
		VoteProb:         0.5,
		JoinProb:         0.1,
		ActiveHourRanges: [][2]int{{9, 17}, {20, 22}},
	},
	{
		Name:             "PowerUser",
		PostProb:         0.3,
		CommentProb:      0.4,
		VoteProb:         0.2,
		JoinProb:         0.1,
		ActiveHourRanges: [][2]int{{7, 23}},
	},
	{
		Name:             "Lurker",
		PostProb:         0.05,
		CommentProb:      0.15,
		VoteProb:         0.7,
		JoinProb:         0.1,
		ActiveHourRanges: [][2]int{{12, 14}, {18, 23}},
	},
}

func GenerateUserBehavior() *common.ClientBehavior {
	persona := userPersonas[rand.Intn(len(userPersonas))]

	activeHours := make([]int, 0)
	for _, timeRange := range persona.ActiveHourRanges {
		for hour := timeRange[0]; hour <= timeRange[1]; hour++ {
			activeHours = append(activeHours, hour)
		}
	}

	return &common.ClientBehavior{
		PostProbability:    persona.PostProb,
		CommentProbability: persona.CommentProb,
		VoteProbability:    persona.VoteProb,
		JoinProbability:    persona.JoinProb,
		ActiveHours:        activeHours,
		Persona:            persona.Name,
	}
}
