package Global

type ImagesInfo struct {
	Image1      string
	Image2      string
	Image3      string
	FolderId    string
	Instruction string
	Color       string
	BorderWidth string
}
type WorkCompletion struct {
	Err bool
	Msg string
	FolderId string
}
var TaskBeingDone = map[string]bool{}
var TaskCompleted = map[string]bool{}
var TaskError = map[string]bool{}