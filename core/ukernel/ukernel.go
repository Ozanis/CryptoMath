package ukernel

func findCUDA() bool {
	return false
}

func findOpenCl() bool {
	return false
}

func UseGpu() string {
	var gpuInfo = map[string]bool{
		"cuda":   findCUDA(),
		"opencl": findOpenCl(),
	}
	for i := range gpuInfo {
		if gpuInfo[i] {
			return i
		}
	}
	return "native"
}
