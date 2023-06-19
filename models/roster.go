package models

import (
	"fdsim/utils"
	"sort"

	"golang.org/x/exp/maps"
)

// Roster Cache Keys
const (
	// RosterCacheKey Players Avg Skill
	rCK_PAS = "rck:plsAvgSkill"
	// RosterCacheKey Players Avg Morale
	rCK_PAM = "rck:plsAvgMorale"
	// RosterCacheKey Players Avg Age
	rCK_PAA = "rck:plsAvgAge"

	// RosterCacheKey Players Wages
	rCK_PW = "rck:plsWage"
)

type PlayersMap map[string]*Player

type Roster struct {
	players     PlayersMap
	cache       map[string]interface{}
	indexByRole map[Role][]PPH
}

func NewRoster() *Roster {
	return &Roster{
		players:     PlayersMap{},
		cache:       map[string]interface{}{},
		indexByRole: NewRolePPHMap(),
	}
}

func (r *Roster) cacheValues() {
	totS := 0
	totM := 0
	totA := 0
	var totWages float64
	for _, p := range r.players {
		totS += p.Skill.Val()
		totM += p.Morale.Val()
		totA += p.Age
		totWages += p.Wage.Value()
	}

	valS := float64(totS) / float64(r.Len())
	r.cache[rCK_PAS] = valS
	valM := float64(totM) / float64(r.Len())
	r.cache[rCK_PAM] = valM
	valA := float64(totA) / float64(r.Len())
	r.cache[rCK_PAA] = valA

	r.cache[rCK_PW] = totWages
}

func (r *Roster) Wages() utils.Money {
	if val, ok := r.cache[rCK_PW]; ok {
		return utils.NewEurosFromF(val.(float64))
	}
	r.cacheValues()
	val := r.cache[rCK_PW].(float64)
	return utils.NewEurosFromF(val)
}

func (r *Roster) AvgSkill() float64 {
	if val, ok := r.cache[rCK_PAS]; ok {
		return val.(float64)
	}
	r.cacheValues()
	val := r.cache[rCK_PAS].(float64)
	return val
}

func (r *Roster) AvgAge() float64 {
	if val, ok := r.cache[rCK_PAA]; ok {
		return val.(float64)
	}

	r.cacheValues()
	val := r.cache[rCK_PAA].(float64)
	return val
}

func (r *Roster) AvgMorale() float64 {
	if val, ok := r.cache[rCK_PAM]; ok {
		return val.(float64)
	}

	r.cacheValues()
	val := r.cache[rCK_PAM].(float64)
	return val
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
	r.cacheValues()
}

func (r *Roster) AddPlayers(players []*Player) {
	for _, p := range players {
		r.add(p)

	}
	r.cacheValues()
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

func (r *Roster) Lineup(module Module) *Lineup {
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

	return NewLineup(module, lineup, TeamStats{r.AvgSkill(), r.AvgMorale(), r.AvgAge()})
}

func (r *Roster) Player(id string) (*Player, bool) {
	p, ok := r.players[id]
	return p, ok
}

func (r *Roster) Players() []*Player {
	return maps.Values(r.players)
}

func (r *Roster) PlayersByRole() []*Player {
	players := make([]*Player, r.Len())
	roles := AllPlayerRoles()
	count := 0
	for _, role := range roles {
		for _, pph := range r.indexByRole[role] {
			p, _ := r.Player(pph.Id)
			players[count] = p
			count++
		}
	}
	return players
}

type TeamStats struct {
	Skill  float64
	Morale float64
	Age    float64
}
