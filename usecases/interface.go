package usecases

type UC interface {
	ProcessTableAlias() UC
	Process() UC
	Export() string
}
