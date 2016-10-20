package safepath

import ()

// Like filepath.Abs but does not remove ".."
func Abs(file string) (string, error) {
	dir, err := EnvWd()
	if err != nil {
		return "", err
	}
	return AbsPath(dir, file)
}
