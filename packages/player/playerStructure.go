package player

/*
You need to implement a structure with fields "On", "Ammo" and "Power", with types bool, int, int, respectively.
This structure should have methods: Shoot and RideBike, which take no arguments but return a bool value.
If "On" == false, then both methods must return false.
You can do Shoot only if Ammo is present (then Ammo is reduced by one, and the method returns true),
otherwise the method must return false. The RideBike method logic are the same, but depends on the Power property.
*/

type Player struct {
	On    bool
	Ammo  int
	Power int
}

func (p *Player) Shoot() bool {
	return p.do(&p.Ammo)
}

func (p *Player) RideBike() bool {
	return p.do(&p.Power)
}

func (p *Player) do(observed *int) bool {
	if !p.On || *observed <= 0 {
		return false
	} else {
		*observed--
		return true
	}
}
