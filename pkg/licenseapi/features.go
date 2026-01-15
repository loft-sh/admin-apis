package licenseapi

// This code was generated. Change features.yaml to add, remove, or edit features.

// Features
const (
	VirtualCluster FeatureName = "vclusters" // Virtual Cluster Management

	VirtualClusterSleepMode FeatureName = "vcluster-sleep-mode" // Sleep Mode for Virtual Clusters

	VirtualClusterHostPathMapper FeatureName = "vcluster-host-path-mapper" // Central HostPath Mapper

	VirtualClusterEnterprisePlugins FeatureName = "vcluster-enterprise-plugins" // Enterprise Plugins

	VirtualClusterProDistroImage FeatureName = "vcp-distro-image" // Security-Hardened vCluster Image

	VirtualClusterProDistroBuiltInCoreDNS FeatureName = "vcp-distro-built-in-coredns" // Built-In CoreDNS

	VirtualClusterProDistroAdmissionControl FeatureName = "vcp-distro-admission-control" // Virtual Admission Control

	VirtualClusterProDistroSyncPatches FeatureName = "vcp-distro-sync-patches" // Sync Patches

	VirtualClusterProDistroEmbeddedEtcd FeatureName = "vcp-distro-embedded-etcd" // Embedded etcd

	VirtualClusterProDistroIsolatedControlPlane FeatureName = "vcp-distro-isolated-cp" // Isolated Control Plane

	VirtualClusterProDistroCentralizedAdmissionControl FeatureName = "vcp-distro-centralized-admission-control" // Centralized Admission Control

	VirtualClusterProDistroGenericSync FeatureName = "vcp-distro-generic-sync" // Generic Sync

	VirtualClusterProDistroTranslatePatches FeatureName = "vcp-distro-translate-patches" // Translate Patches

	VirtualClusterProDistroIntegrationsKubeVirt FeatureName = "vcp-distro-integrations-kube-virt" // KubeVirt Integration

	VirtualClusterProDistroIntegrationsExternalSecrets FeatureName = "vcp-distro-integrations-external-secrets" // External Secrets Integration

	VirtualClusterProDistroIntegrationsCertManager FeatureName = "vcp-distro-integrations-cert-manager" // Cert Manager Integration

	VirtualClusterProDistroFips FeatureName = "vcp-distro-fips" // FIPS

	VirtualClusterProDistroExternalDatabase FeatureName = "vcp-distro-external-database" // External Database

	VirtualClusterProDistroPrivateNodes FeatureName = "vcp-distro-private-nodes" // Private Nodes

	ExternalDatabaseRdsIam FeatureName = "external-database-rds-iam" // External Database RDS IAM Authentication

	ConnectorExternalDatabase FeatureName = "connector-external-database" // Database Connector

	ConnectorExternalDatabaseEksPodIdentity FeatureName = "connector-external-database-eks-pod-identity" // EKS Pod Identity for External Database Connections

	VirtualClusterProDistroSleepMode FeatureName = "vcp-distro-sleep-mode" // SleepMode

	Devpod FeatureName = "devpod" // Dev Environment Management

	Namespaces FeatureName = "namespaces" // Namespace Management

	NamespaceSleepMode FeatureName = "namespace-sleep-mode" // Sleep Mode for Namespaces

	ConnectedClusters FeatureName = "connected-clusters" // Connected Clusters

	ClusterAccess FeatureName = "cluster-access" // Cluster Access

	ClusterRoles FeatureName = "cluster-roles" // Cluster Role Management

	SSOAuth FeatureName = "sso-authentication" // Single Sign-On

	AuditLogging FeatureName = "audit-logging" // Audit Logging

	AutoIngressAuth FeatureName = "auto-ingress-authentication" // Automatic Auth For Ingresses

	OIDCProvider FeatureName = "oidc-provider" // Platform as OIDC Provider

	MultipleSSOProviders FeatureName = "multiple-sso-providers" // Multiple SSO Providers

	Apps FeatureName = "apps" // Apps

	TemplateVersioning FeatureName = "template-versioning" // Template Versioning

	ArgoIntegration FeatureName = "argo-integration" // Argo Integration

	RancherIntegration FeatureName = "rancher-integration" // Rancher Integration

	Secrets FeatureName = "secrets" // Secrets Sync

	SecretEncryption FeatureName = "secret-encryption" // Secrets Encryption

	VaultIntegration FeatureName = "vault-integration" // HashiCorp Vault Integration

	HighAvailabilityMode FeatureName = "ha-mode" // High-Availability Mode

	MultiRegionMode FeatureName = "multi-region-mode" // Multi-Region Mode

	AirGappedMode FeatureName = "air-gapped-mode" // Air-Gapped Mode

	CustomBranding FeatureName = "custom-branding" // Custom Branding

	AdvancedUICustomizations FeatureName = "advanced-ui-customizations" // Advanced UI Customizations

	VNodeRuntime FeatureName = "vnode-runtime" // vNode Runtime

	ProjectQuotas FeatureName = "project-quotas" // Project Quotas

	ResolveDns FeatureName = "resolve-dns" // Resolve DNS

	IstioIntegration FeatureName = "istio-integration" // Istio Integration

	HybridScheduling FeatureName = "hybrid-scheduling" // Hybrid Scheduling

	SyncNamespacesTohost FeatureName = "sync-namespaces-tohost" // Sync Namespaces toHost

	ScheduledSnapshots FeatureName = "scheduled-snapshots" // Auto Snapshots

	PrivateNodesVpn FeatureName = "private-nodes-vpn" // Private Nodes VPN

	PrivateNodesAutoNodes FeatureName = "private-nodes-auto-nodes" // Private Nodes Auto Nodes

	DisablePlatformDB FeatureName = "disable-platform-db" // Disable Platform Database

	Standalone FeatureName = "standalone" // Standalone

	Netris FeatureName = "netris" // Netris

	KubeVip FeatureName = "kube-vip" // Kube-vip Integration

	VirtualClusterProxyResources FeatureName = "vcluster-proxy-resources" // vCluster Proxy Resources

)

func GetFeatures() []FeatureName {
	return []FeatureName{
		VirtualCluster,
		VirtualClusterSleepMode,
		VirtualClusterHostPathMapper,
		VirtualClusterEnterprisePlugins,
		VirtualClusterProDistroImage,
		VirtualClusterProDistroBuiltInCoreDNS,
		VirtualClusterProDistroAdmissionControl,
		VirtualClusterProDistroSyncPatches,
		VirtualClusterProDistroEmbeddedEtcd,
		VirtualClusterProDistroIsolatedControlPlane,
		VirtualClusterProDistroCentralizedAdmissionControl,
		VirtualClusterProDistroGenericSync,
		VirtualClusterProDistroTranslatePatches,
		VirtualClusterProDistroIntegrationsKubeVirt,
		VirtualClusterProDistroIntegrationsExternalSecrets,
		VirtualClusterProDistroIntegrationsCertManager,
		VirtualClusterProDistroFips,
		VirtualClusterProDistroExternalDatabase,
		VirtualClusterProDistroPrivateNodes,
		ExternalDatabaseRdsIam,
		ConnectorExternalDatabase,
		ConnectorExternalDatabaseEksPodIdentity,
		VirtualClusterProDistroSleepMode,
		Devpod,
		Namespaces,
		NamespaceSleepMode,
		ConnectedClusters,
		ClusterAccess,
		ClusterRoles,
		SSOAuth,
		AuditLogging,
		AutoIngressAuth,
		OIDCProvider,
		MultipleSSOProviders,
		Apps,
		TemplateVersioning,
		ArgoIntegration,
		RancherIntegration,
		Secrets,
		SecretEncryption,
		VaultIntegration,
		HighAvailabilityMode,
		MultiRegionMode,
		AirGappedMode,
		CustomBranding,
		AdvancedUICustomizations,
		VNodeRuntime,
		ProjectQuotas,
		ResolveDns,
		IstioIntegration,
		HybridScheduling,
		SyncNamespacesTohost,
		ScheduledSnapshots,
		PrivateNodesVpn,
		PrivateNodesAutoNodes,
		DisablePlatformDB,
		Standalone,
		Netris,
		KubeVip,
		VirtualClusterProxyResources,
	}
}

func GetAllFeatures() []*Feature {
	return []*Feature{
 		{
			DisplayName: "Virtual Cluster Management",
			Name:        "vclusters",
			Module:      "virtual-clusters",
		},
		{
			DisplayName: "Sleep Mode for Virtual Clusters",
			Name:        "vcluster-sleep-mode",
			Module:      "virtual-clusters",
		},
		{
			DisplayName: "Central HostPath Mapper",
			Name:        "vcluster-host-path-mapper",
			Module:      "virtual-clusters",
		},
		{
			DisplayName: "Enterprise Plugins",
			Name:        "vcluster-enterprise-plugins",
			Module:      "virtual-clusters",
		},
		{
			DisplayName: "Security-Hardened vCluster Image",
			Name:        "vcp-distro-image",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Built-In CoreDNS",
			Name:        "vcp-distro-built-in-coredns",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Virtual Admission Control",
			Name:        "vcp-distro-admission-control",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Sync Patches",
			Name:        "vcp-distro-sync-patches",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Embedded etcd",
			Name:        "vcp-distro-embedded-etcd",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Isolated Control Plane",
			Name:        "vcp-distro-isolated-cp",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Centralized Admission Control",
			Name:        "vcp-distro-centralized-admission-control",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Generic Sync",
			Name:        "vcp-distro-generic-sync",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Translate Patches",
			Name:        "vcp-distro-translate-patches",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "KubeVirt Integration",
			Name:        "vcp-distro-integrations-kube-virt",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "External Secrets Integration",
			Name:        "vcp-distro-integrations-external-secrets",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Cert Manager Integration",
			Name:        "vcp-distro-integrations-cert-manager",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "FIPS",
			Name:        "vcp-distro-fips",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "External Database",
			Name:        "vcp-distro-external-database",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Private Nodes",
			Name:        "vcp-distro-private-nodes",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "External Database RDS IAM Authentication",
			Name:        "external-database-rds-iam",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Database Connector",
			Name:        "connector-external-database",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "EKS Pod Identity for External Database Connections",
			Name:        "connector-external-database-eks-pod-identity",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "SleepMode",
			Name:        "vcp-distro-sleep-mode",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Dev Environment Management",
			Name:        "devpod",
			Module:      "dev-environments",
		},
		{
			DisplayName: "Namespace Management",
			Name:        "namespaces",
			Module:      "kubernetes-namespaces",
		},
		{
			DisplayName: "Sleep Mode for Namespaces",
			Name:        "namespace-sleep-mode",
			Module:      "kubernetes-namespaces",
		},
		{
			DisplayName: "Connected Clusters",
			Name:        "connected-clusters",
			Module:      "kubernetes-clusters",
		},
		{
			DisplayName: "Cluster Access",
			Name:        "cluster-access",
			Module:      "kubernetes-clusters",
		},
		{
			DisplayName: "Cluster Role Management",
			Name:        "cluster-roles",
			Module:      "kubernetes-clusters",
		},
		{
			DisplayName: "Single Sign-On",
			Name:        "sso-authentication",
			Module:      "auth-audit-logging",
		},
		{
			DisplayName: "Audit Logging",
			Name:        "audit-logging",
			Module:      "auth-audit-logging",
		},
		{
			DisplayName: "Automatic Auth For Ingresses",
			Name:        "auto-ingress-authentication",
			Module:      "auth-audit-logging",
		},
		{
			DisplayName: "Platform as OIDC Provider",
			Name:        "oidc-provider",
			Module:      "auth-audit-logging",
		},
		{
			DisplayName: "Multiple SSO Providers",
			Name:        "multiple-sso-providers",
			Module:      "auth-audit-logging",
		},
		{
			DisplayName: "Apps",
			Name:        "apps",
			Module:      "templating-gitops",
		},
		{
			DisplayName: "Template Versioning",
			Name:        "template-versioning",
			Module:      "templating-gitops",
		},
		{
			DisplayName: "Argo Integration",
			Name:        "argo-integration",
			Module:      "templating-gitops",
		},
		{
			DisplayName: "Rancher Integration",
			Name:        "rancher-integration",
			Module:      "templating-gitops",
		},
		{
			DisplayName: "Secrets Sync",
			Name:        "secrets",
			Module:      "secrets-management",
		},
		{
			DisplayName: "Secrets Encryption",
			Name:        "secret-encryption",
			Module:      "secrets-management",
		},
		{
			DisplayName: "HashiCorp Vault Integration",
			Name:        "vault-integration",
			Module:      "secrets-management",
		},
		{
			DisplayName: "High-Availability Mode",
			Name:        "ha-mode",
			Module:      "deployment-modes",
		},
		{
			DisplayName: "Multi-Region Mode",
			Name:        "multi-region-mode",
			Module:      "deployment-modes",
		},
		{
			DisplayName: "Air-Gapped Mode",
			Name:        "air-gapped-mode",
			Module:      "deployment-modes",
		},
		{
			DisplayName: "Custom Branding",
			Name:        "custom-branding",
			Module:      "ui-customization",
		},
		{
			DisplayName: "Advanced UI Customizations",
			Name:        "advanced-ui-customizations",
			Module:      "ui-customization",
		},
		{
			DisplayName: "vNode Runtime",
			Name:        "vnode-runtime",
			Module:      "vnode",
		},
		{
			DisplayName: "Project Quotas",
			Name:        "project-quotas",
			Module:      "virtual-clusters",
		},
		{
			DisplayName: "Resolve DNS",
			Name:        "resolve-dns",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Istio Integration",
			Name:        "istio-integration",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Hybrid Scheduling",
			Name:        "hybrid-scheduling",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Sync Namespaces toHost",
			Name:        "sync-namespaces-tohost",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Auto Snapshots",
			Name:        "scheduled-snapshots",
			Module:      "virtual-clusters",
		},
		{
			DisplayName: "Private Nodes VPN",
			Name:        "private-nodes-vpn",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Private Nodes Auto Nodes",
			Name:        "private-nodes-auto-nodes",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Disable Platform Database",
			Name:        "disable-platform-db",
			Module:      "deployment-modes",
		},
		{
			DisplayName: "Standalone",
			Name:        "standalone",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Netris",
			Name:        "netris",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "Kube-vip Integration",
			Name:        "kube-vip",
			Module:      "vcluster-pro-distro",
		},
		{
			DisplayName: "vCluster Proxy Resources",
			Name:        "vcluster-proxy-resources",
			Module:      "vcluster-pro-distro",
		},
 	}
}
