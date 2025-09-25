package security

type Role string

const Anon Role = "ANON"
const User Role = "USER"

func RoleFromString(role string) Role {
	switch role {
	case "ANON":
		return Anon
	case "USER":
		return User
	default:
		return ""
	}
}
