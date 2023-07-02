package adapters

import "github.com/MikeMwita/procastination_proxy/internal/core/ports"

type ProxyService struct {
	domainBlocker        ports.DomainBlocker
	blocklistManager     ports.BlocklistManager
	userBlocklistManager ports.UserBlocklistManager
	globalBlocklist      []string
	officeHoursStartHour int
	officeHoursEndHour   int
}

func NewProxyService(domainBlocker ports.DomainBlocker, blocklistManager ports.BlocklistManager, userBlocklistManager ports.UserBlocklistManager) *ProxyService {
	return &ProxyService{
		domainBlocker:        domainBlocker,
		blocklistManager:     blocklistManager,
		userBlocklistManager: userBlocklistManager,
	}

}
