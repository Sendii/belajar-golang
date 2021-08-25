package pattern

type PQuery interface{
	// Read() bool
	Show() bool
	Create() bool
	Update() bool
	Delete() bool
}

