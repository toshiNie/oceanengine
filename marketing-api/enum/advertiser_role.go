package enum

// AdvertiserRole 广告主角色
type AdvertiserRole string

const (
	// ROLE_ADVERTISER 普通广告主（直客）
	ROLE_ADVERTISER AdvertiserRole = "ROLE_ADVERTISER"
	// ROLE_CHILD_ADVERTISER 普通广告主（代理商子客户）
	ROLE_CHILD_ADVERTISER AdvertiserRole = "ROLE_CHILD_ADVERTISER"
	// ROLE_CHILD_AGENT 二级代理商
	ROLE_CHILD_AGENT AdvertiserRole = "ROLE_CHILD_AGENT"
	// ROLE_AGENT 一级代理商
	ROLE_AGENT AdvertiserRole = "ROLE_AGENT"
	// ROLE_ECP_VIRTUAL_ADVERTISER 千川虚拟广告主
	ROLE_ECP_VIRTUAL_ADVERTISER AdvertiserRole = "ROLE_ECP_VIRTUAL_ADVERTISER"
	// ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER 本地虚拟广告主
	ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER AdvertiserRole = "ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER"
	// ROLE_VIRTAUL_ADVERTISER 虚拟广告主
	ROLE_VIRTAUL_ADVERTISER AdvertiserRole = "ROLE_VIRTAUL_ADVERTISER"
)
