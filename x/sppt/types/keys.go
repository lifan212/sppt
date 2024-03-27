package types

const (
	// ModuleName defines the module name
	ModuleName = "sppt"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sppt"
)

var (
	ParamsKey = []byte("p_sppt")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
