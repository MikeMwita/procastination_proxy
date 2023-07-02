package services

import "github.com/MikeMwita/procastination_proxy/internal/adapters"

type ProxyService struct {
	domainBlocker        adapters.DomainBlockerAdapter
	blocklistManager     adapters.BlocklistManagerAdapter
	userBlocklistManager adapters.UserBlocklistManagerAdapter
	globalBlocklist      []string
	officeHoursStartHour int
	officeHoursEndHour   int
}

func (ps *ProxyService) BlockDomain(domain string) error {
	// Implement the logic to block a domain using the domainBlocker
	err := ps.domainBlocker.BlockDomain(domain)
	if err != nil {
		return err
	}
	return ps.domainBlocker.BlockDomain(domain)
}

func (ps *ProxyService) UnblockDomain(domain string) error {
	// Implement the logic to unblock a domain using the domainBlocker
	return ps.domainBlocker.UnblockDomain(domain)
}

func (ps *ProxyService) GetBlocklist() ([]string, error) {
	// Implement the logic to retrieve the blocklist using the blocklistManager
	return ps.blocklistManager.GetBlocklist()
}

func (ps *ProxyService) AddToBlocklist(domain string) error {
	// Implement the logic to add a domain to the blocklist using the blocklistManager
	return ps.blocklistManager.AddToBlocklist(domain)
}

func (ps *ProxyService) RemoveFromBlocklist(domain string) error {
	// Implement the logic to remove a domain from the blocklist using the blocklistManager
	return ps.blocklistManager.RemoveFromBlocklist(domain)
}

func (ps *ProxyService) GetUserBlocklist(userID string) ([]string, error) {
	// Implement the logic to retrieve the user-specific blocklist using the userBlocklistManager
	return ps.userBlocklistManager.GetUserBlocklist(userID)
}

func (ps *ProxyService) AddToUserBlocklist(userID string, domain string) error {
	// Implement the logic to add a domain to the user-specific blocklist using the userBlocklistManager
	return ps.userBlocklistManager.AddToUserBlocklist(userID, domain)
}

func (ps *ProxyService) RemoveFromUserBlocklist(userID string, domain string) error {
	// Implement the logic to remove a domain from the user-specific blocklist using the userBlocklistManager
	return ps.userBlocklistManager.RemoveFromUserBlocklist(userID, domain)
}
