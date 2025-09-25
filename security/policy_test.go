package security

import "testing"

func TestPolicy_Validate(t *testing.T) {
	policy := NewPolicy()

	internalIP := "192.168.1.1"
	externalIP := "192.168.1.2"
	validValidator := func() bool { return true }
	invalidValidator := func() bool { return false }

	anonPassport := &Passport{
		UserID: "",
		Roles:  []Role{Anon},
		IP:     externalIP,
	}

	userPassport := &Passport{
		UserID: "user123",
		Roles:  []Role{User},
		IP:     internalIP,
	}

	tests := map[string]struct {
		AllowedRoles []Role
		AllowedIPs   []string
		Validator    func() bool
		passport     *Passport
		want         bool
	}{
		"nil passport": {
			AllowedRoles: nil,
			AllowedIPs:   nil,
			passport:     nil,
			want:         false,
		},
		"no restrictions, anonymous": {
			AllowedRoles: nil,
			AllowedIPs:   nil,
			passport:     anonPassport,
			want:         true,
		},
		"no restrictions, user": {
			AllowedRoles: nil,
			AllowedIPs:   nil,
			passport:     userPassport,
			want:         true,
		},
		"role allowed, anonymous": {
			AllowedRoles: []Role{Anon},
			AllowedIPs:   nil,
			passport:     anonPassport,
			want:         true,
		},
		"role not allowed, anonymous": {
			AllowedRoles: []Role{User},
			AllowedIPs:   nil,
			passport:     anonPassport,
			want:         false,
		},
		"role allowed, user": {
			AllowedRoles: []Role{User},
			AllowedIPs:   nil,
			passport:     userPassport,
			want:         true,
		},
		"role not allowed, user": {
			AllowedRoles: []Role{Anon},
			AllowedIPs:   nil,
			passport:     userPassport,
			want:         false,
		},
		"IP allowed, anonymous": {
			AllowedRoles: nil,
			AllowedIPs:   []string{externalIP},
			passport:     anonPassport,
			want:         true,
		},
		"IP not allowed, anonymous": {
			AllowedRoles: nil,
			AllowedIPs:   []string{internalIP},
			passport:     anonPassport,
			want:         false,
		},
		"IP allowed, user": {
			AllowedRoles: nil,
			AllowedIPs:   []string{internalIP},
			passport:     userPassport,
			want:         true,
		},
		"IP not allowed, user": {
			AllowedRoles: nil,
			AllowedIPs:   []string{externalIP},
			passport:     userPassport,
			want:         false,
		},
		"both role and IP allowed, user": {
			AllowedRoles: []Role{User},
			AllowedIPs:   []string{internalIP},
			passport:     userPassport,
			want:         true,
		},
		"role allowed but IP not allowed, user": {
			AllowedRoles: []Role{User},
			AllowedIPs:   []string{externalIP},
			passport:     userPassport,
			want:         false,
		},
		"role not allowed but IP allowed, user": {
			AllowedRoles: []Role{Anon},
			AllowedIPs:   []string{internalIP},
			passport:     userPassport,
			want:         false,
		},
		"both role and IP not allowed, user": {
			AllowedRoles: []Role{Anon},
			AllowedIPs:   []string{externalIP},
			passport:     userPassport,
			want:         false,
		},
		"custom validator passes": {
			AllowedRoles: nil,
			AllowedIPs:   nil,
			Validator:    validValidator,
			passport:     userPassport,
			want:         true,
		},
		"custom validator fails": {
			AllowedRoles: nil,
			AllowedIPs:   nil,
			Validator:    invalidValidator,
			passport:     userPassport,
			want:         false,
		},
		"all checks pass": {
			AllowedRoles: []Role{User},
			AllowedIPs:   []string{internalIP},
			Validator:    validValidator,
			passport:     userPassport,
			want:         true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			policy = NewPolicy().
				WithRoles(tt.AllowedRoles...).
				WithIPs(tt.AllowedIPs...).
				WithCustomValidator(tt.Validator)

			if got := policy.Validate(tt.passport); got != tt.want {
				t.Errorf("Policy.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
