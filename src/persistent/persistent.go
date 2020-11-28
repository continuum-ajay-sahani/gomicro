package persistent

// FileOffsetUser maintain offset of user
var FileOffsetUser map[string]int64

// Load init persistancy
func Load() {
	FileOffsetUser = make(map[string]int64)
}
