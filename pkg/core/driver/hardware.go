package driver

func findCUDA() string {
	return ""
}

func findOpenCl() string {
	return ""
}

func UseGpu() string {
	var gpuInfo = []string{findCUDA(), findOpenCl()}
	for _, i := range gpuInfo {
		if i != "" {
			return i
		}
	}
	return "native"
}
