package main

func validEnv(env string) bool {
	switch env {
	case "dev", "sit", "prod":
		return true
	default:
		return false
	}
}
