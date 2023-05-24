package models

// Roster Cache Keys
const (
	// RosterCacheKey Players Avg Skill
	rCK_PAS = "rck:plsAvgSkill"
	// RosterCacheKey Players Avg Morale
	rCK_PAM = "rck:plsAvgMorale"
)

type Roster struct {
	players     map[string]*Player
	cache       map[string]interface{}
	indexByRole map[Role]string
}

func NewRoster() *Roster {
	return &Roster{
		players:     map[string]*Player{},
		cache:       map[string]interface{}{},
		indexByRole: map[Role]string{},
	}
}

func (r *Roster) calculateAvgs() (float64, float64) {
	totS := 0
	totM := 0
	for _, p := range r.players {
		totS += p.Skill.Val()
		totM += p.Morale.Val()
	}

	valS := float64(totS) / float64(r.Len())
	r.cache[rCK_PAS] = valS
	valM := float64(totM) / float64(r.Len())
	r.cache[rCK_PAM] = valM

	return valS, valM
}

func (r *Roster) AvgSkill() float64 {
	if val, ok := r.cache[rCK_PAS]; ok {
		return val.(float64)
	}

	v, _ := r.calculateAvgs()

	return v
}

func (r *Roster) AddPlayer(player *Player) {
	r.players[player.Id] = player
	r.calculateAvgs()
}

func (r *Roster) AddPlayers(players []*Player) {
	for _, p := range players {
		r.players[p.Id] = p
	}
	r.calculateAvgs()
}

func (r *Roster) Len() int {
	return len(r.players)
}
