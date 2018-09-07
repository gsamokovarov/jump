package importer

// multiImporter tries to import configurations from multiple importers. If at
// least on of the importers succeed, no errors will be returned.
type multiImporter []Importer

func (mi multiImporter) Import(fn Callback) error {
	var lastErr error
	atLeastOneSucceeded := false

	for _, i := range mi {
		err := i.Import(fn)
		if err == NotFoundErr {
			continue
		}
		if err != nil {
			lastErr = err
		}

		atLeastOneSucceeded = true
	}

	if !atLeastOneSucceeded && lastErr != nil {
		return lastErr
	}

	return nil
}
