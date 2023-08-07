package executor

func stubAlpine() error {
	var err error
	if err = apkbuild(); err != nil {
		return err
	}

	if err = makefile(); err != nil {
		return err
	}
	return nil
}

func apkbuild() error {
	return nil
}

func makefile() error {
	return nil
}
