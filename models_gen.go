// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package cfop

type Algorithm struct {
	ID         string `json:"id"`
	SubgroupID string `json:"subgroup_id"`
	Algorithm  string `json:"algorithm"`
}

type Group struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Subgroups   []Subgroup `json:"subgroups"`
}

type Subgroup struct {
	ID           string      `json:"id"`
	GroupID      string      `json:"group_id"`
	Type         string      `json:"type"`
	Name         string      `json:"name"`
	OptimalMoves string      `json:"optimal_moves"`
	ImageLink    string      `json:"image_link"`
	Algorithms   []Algorithm `json:"algorithms"`
}