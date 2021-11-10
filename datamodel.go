package getinge

type Group struct {
	ID   int32
	Name string
}

type User struct {
	ID       int32
	Name     string
	Email    string
	Password string
	Groups   Group
}
