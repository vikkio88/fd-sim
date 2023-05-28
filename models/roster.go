package models

import "sort"

// Roster Cache Keys
const (
	// RosterCacheKey Players Avg Skill
	rCK_PAS = "rck:plsAvgSkill"
	// RosterCacheKey Players Avg Morale
	rCK_PAM = "rck:plsAvgMorale"
	// RosterCacheKey Players Avg Age
	rCK_PAA = "rck:plsAvgAge"
)

type Roster struct {
	players     map[string]*Player
	cache       map[string]interface{}
	indexByRole map[Role][]PPH
}

func NewRoster() *Roster {
	return &Roster{
		players:     map[string]*Player{},
		cache:       map[string]interface{}{},
		indexByRole: NewRolePPHMap(),
	}
}

func (r *Roster) calculateAvgs() (float64, float64, float64) {
	totS := 0
	totM := 0
	totA := 0
	for _, p := range r.players {
		totS += p.Skill.Val()
		totM += p.Morale.Val()
		totA += p.Age
	}

	valS := float64(totS) / float64(r.Len())
	r.cache[rCK_PAS] = valS
	valM := float64(totM) / float64(r.Len())
	r.cache[rCK_PAM] = valM
	valA := float64(totA) / float64(r.Len())
	r.cache[rCK_PAA] = valA

	return valS, valM, valA
}

func (r *Roster) AvgSkill() float64 {
	if val, ok := r.cache[rCK_PAS]; ok {
		return val.(float64)
	}

	v, _, _ := r.calculateAvgs()

	return v
}

func (r *Roster) AvgAge() float64 {
	if val, ok := r.cache[rCK_PAA]; ok {
		return val.(float64)
	}

	_, _, v := r.calculateAvgs()

	return v
}

func (r *Roster) add(player *Player) {
	r.players[player.Id] = player
	ps := player.PH()
	r.indexByRole[player.Role] = append(r.indexByRole[player.Role], ps)

	// Sorting players by Skill and Morale on insert
	sort.Slice(r.indexByRole[player.Role],
		func(i, j int) bool {
			if r.indexByRole[player.Role][i].Skill == r.indexByRole[player.Role][j].Skill {
				return r.indexByRole[player.Role][i].Morale > r.indexByRole[player.Role][j].Morale
			}

			return r.indexByRole[player.Role][i].Skill > r.indexByRole[player.Role][j].Skill
		})
}

func (r *Roster) AddPlayer(player *Player) {
	r.add(player)
	r.calculateAvgs()
}

func (r *Roster) AddPlayers(players []*Player) {
	for _, p := range players {
		r.add(p)

	}
	r.calculateAvgs()
}

func (r *Roster) Len() int {
	return len(r.players)
}

func (r *Roster) IdsInRole(role Role) []string {
	if v, ok := r.indexByRole[role]; ok {
		players := make([]string, len(v))
		for i, pph := range v {
			players[i] = pph.Id
		}

		return players
	}
	return []string{}
}

func (r *Roster) InRole(role Role) []*Player {
	if v, ok := r.indexByRole[role]; ok {
		players := make([]*Player, len(v))
		for i, pph := range v {
			players[i] = r.players[pph.Id]
		}

		return players
	}
	return []*Player{}
}

func (r *Roster) Lineup(module Module) Lineup {
	conf := module.Conf()
	lineup := NewRolePPHMap()
	//bench := NewRolePPHMap()
	missing := NewEmptyRoleCounter()
	for role, count := range conf {
		pRoleInRoster := len(r.indexByRole[role])
		if pRoleInRoster >= count {
			lineup[role] = r.indexByRole[role][0:count]
		} else {
			missing[role] = count - pRoleInRoster
			lineup[role] = r.indexByRole[role][0:pRoleInRoster]
		}
	}

	return NewLineup(module, lineup, TeamStats{})
}

func (r *Roster) Player(id string) (*Player, bool) {
	p, ok := r.players[id]
	return p, ok
}

type TeamStats struct {
	Skill  float64
	Morale float64
	Age    float64
}
