package security

import "slices"

type Policy struct {
	AllowedRoles    []Role
	AllowedIPs      []string
	CustomValidator func() bool
}

func NewPolicy() *Policy {
	return &Policy{
		AllowedRoles:    []Role{},
		AllowedIPs:      []string{},
		CustomValidator: nil,
	}
}

func Anonymous() *Policy {
	return NewPolicy()
}

func UserOnly() *Policy {
	return NewPolicy().WithRoles(User)
}

func AnonymousOrUser() *Policy {
	return NewPolicy().WithRoles(Anon, User)
}

func AnonymousOnly() *Policy {
	return NewPolicy().WithRoles(Anon)
}
func WithRoles(roles ...Role) *Policy {
	return NewPolicy().WithRoles(roles...)
}

func WithIPs(ips ...string) *Policy {
	return NewPolicy().WithIPs(ips...)
}

func WithCustomValidator(validator func() bool) *Policy {
	return NewPolicy().WithCustomValidator(validator)
}

func (p *Policy) WithRoles(roles ...Role) *Policy {
	p.AllowedRoles = append(p.AllowedRoles, roles...)
	return p
}

func (p *Policy) WithIPs(ips ...string) *Policy {
	p.AllowedIPs = append(p.AllowedIPs, ips...)
	return p
}

func (p *Policy) WithCustomValidator(validator func() bool) *Policy {
	p.CustomValidator = validator
	return p
}

type Passport struct {
	UserID string
	Roles  []Role
	IP     string
}

func (p *Policy) Validate(passport *Passport) bool {
	if passport == nil {
		return false
	}

	if p.AllowedRoles != nil && len(p.AllowedRoles) > 0 {
		checkResult := false

		for _, role := range passport.Roles {
			if slices.Contains(p.AllowedRoles, role) {
				checkResult = true
			}
		}
		if !checkResult {
			return false
		}
	}

	if p.AllowedIPs != nil && len(p.AllowedIPs) > 0 {
		checkResult := false
		if slices.Contains(p.AllowedIPs, passport.IP) {
			checkResult = true
		}
		if !checkResult {
			return false
		}
	}

	if p.CustomValidator != nil {
		if !p.CustomValidator() {
			return false
		}
	}

	return true
}
