package ports

//this interfaces will be implemeted in the adapters

// DomainBlocker defines the interface for blocking and unblocking domains.
type DomainBlocker interface {
	BlockDomain(domain string) error
	UnblockDomain(domain string) error
}

// BlocklistManager defines the interface for managing the blocklist.
type BlocklistManager interface {
	GetBlocklist() ([]string, error)
	AddToBlocklist(domain string) error
	RemoveFromBlocklist(domain string) error
}

// UserBlocklistManager defines the interface for managing user-specific blocklists.
type UserBlocklistManager interface {
	GetUserBlocklist(userID string) ([]string, error)
	AddToUserBlocklist(userID string, domain string) error
	RemoveFromUserBlocklist(userID string, domain string) error
}
