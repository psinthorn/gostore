package data_model

type Laptop struct {
	Id          string
	Brand       string
	Model       string
	CPU         cpu
	Gpu         gpu
	memory      memory
	screen      screen
	Resolution  resolution
	WeightKg    int
	PriceUsd    int
	ReleaseYear int
	CreatedAt   string
	UpdatedAt   string
}

type cpu struct {
	Brand         string
	Model         string
	NumberCores   int
	NumberThreads int
	MinGhz        float32
	MaxGhz        float32
}

type gpu struct {
	Brand  string
	Model  string
	MinGhz int
	MaxGhz int
	Memory memory
}

type memory struct {
	value string
	unit  string
}

type screen struct {
	SizeInch   int
	Panel      string
	Multitouch bool
}

type resolution struct {
	width  int
	height int
}

type storages struct {
	Driver string
	Memory memory
}
