//go:generate go run ./scripts/gqlgen.go

package cfop

import (
	"context"
	"fmt"

	DB "github.com/aduryagin/cfop-backend/db"
)

// Normalizers
func normalizeSubgroups(subgroups []DB.Subgroup) []Subgroup {
	normalizedSubgroups := make([]Subgroup, len(subgroups))

	for i := 0; i < len(subgroups); i++ {
		normalizedAlgorithms := make([]Algorithm, len(subgroups[i].Algorithms))

		for q := 0; q < len(subgroups[i].Algorithms); q++ {
			normalizedAlgorithms[q] = Algorithm{
				ID:        fmt.Sprint(subgroups[i].Algorithms[q].ID),
				Algorithm: subgroups[i].Algorithms[q].Algorithm,
			}
		}

		normalizedSubgroups[i] = Subgroup{
			ID:           fmt.Sprint(subgroups[i].ID),
			Name:         subgroups[i].Name,
			OptimalMoves: subgroups[i].OptimalMoves,
			ImageLink:    subgroups[i].ImageLink,
			Algorithms:   normalizedAlgorithms,
		}
	}

	return normalizedSubgroups
}

// Resolver type
type Resolver struct{}

// Query function
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Groups(ctx context.Context) ([]Group, error) {
	var groups = []DB.Group{}
	if err := DB.Instance.Find(&groups).Error; err != nil {
		fmt.Println(err)
	}

	normalizedGroups := make([]Group, len(groups))

	for i := 0; i < len(groups); i++ {
		normalizedGroups[i] = Group{
			ID:          fmt.Sprint(groups[i].ID),
			Title:       groups[i].Title,
			Description: groups[i].Description,
		}
	}

	return normalizedGroups, nil
}

func (r *queryResolver) Group(ctx context.Context, groupID string) (*Group, error) {
	var group = DB.Group{}

	if err := DB.Instance.Where("id = ?", groupID).Preload("Subgroups").Preload("Subgroups.Algorithms").Find(&group).Error; err != nil {
		fmt.Println(err)
	}

	return &Group{
		ID:          fmt.Sprint(group.ID),
		Title:       group.Title,
		Description: group.Description,
		Subgroups:   normalizeSubgroups(group.Subgroups),
	}, nil
}

func (r *queryResolver) Subgroups(ctx context.Context, groupID string) ([]Subgroup, error) {
	var subgroups = []DB.Subgroup{}

	if err := DB.Instance.Where("group_id = ?", groupID).Preload("Algorithms").Find(&subgroups).Error; err != nil {
		fmt.Println(err)
	}

	normalizedSubgroups := normalizeSubgroups(subgroups)
	return normalizedSubgroups, nil
}

func (r *queryResolver) Favorites(ctx context.Context, algorithmsIds []string) ([]Subgroup, error) {
	type Favorites struct {
		SubgroupID  string
		AlgorithmID string
		Subgroup
		Algorithm
	}
	var favorites []Favorites

	if err := DB.Instance.Table("algorithms").Where("algorithms.id in (?)", algorithmsIds).Joins("left join subgroups on algorithms.subgroup_id = subgroups.id").Select("subgroups.*, algorithms.*, subgroups.id as subgroup_id, algorithms.id as algorithm_id").Find(&favorites).Error; err != nil {
		fmt.Println(err)
	}

	var subgroups []Subgroup

	for i := 0; i < len(favorites); i++ {
		var normalizedAlgorithm Algorithm

		// normalize algorithm

		normalizedAlgorithm = Algorithm{
			ID:         favorites[i].AlgorithmID,
			SubgroupID: favorites[i].SubgroupID,
			Algorithm:  favorites[i].Algorithm.Algorithm,
		}

		// check that subgroup exists in favorites array

		subgroupExists := false
		subgroupIndex := -1
		for q, subgroup := range subgroups {
			if fmt.Sprint(favorites[i].SubgroupID) == subgroup.ID {
				subgroupExists = true
				subgroupIndex = q
			}
		}

		if subgroupExists {
			subgroups[subgroupIndex].Algorithms = append(subgroups[subgroupIndex].Algorithms, normalizedAlgorithm)
		} else {
			subgroups = append(subgroups, Subgroup{
				ID:           fmt.Sprint(favorites[i].SubgroupID),
				Name:         favorites[i].Name,
				OptimalMoves: favorites[i].OptimalMoves,
				ImageLink:    favorites[i].ImageLink,
				Algorithms:   []Algorithm{normalizedAlgorithm},
			})
		}

	}

	return subgroups, nil
}
