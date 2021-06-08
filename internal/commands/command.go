package commands

type Command interface {
	Invokes() []string
	Description() string
	//AdminRequired() bool
	PermissionsRequired() (bool, uint)
	Exec(ctx *Context) error
}
