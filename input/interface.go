package input

type Input interface {
	ReadData(string) Input
	ReadKB(string) Input
	ReadVRDKB(string) Input
	Write(string) Input
	Export() (string, map[string]string, map[string]string)
}
