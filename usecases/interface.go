package usecases

type UC interface {
	ProcessTableAlias() UC
	Process() UC
	ProcessVRD() UC
	Export() string
}
