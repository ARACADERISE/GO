package src

type DirInfo struct {
	path	string
	file	string
	valid	bool
}

func gather() *DirInfo {
	return &DirInfo{ path: ".", file: "", valid: true };
}
