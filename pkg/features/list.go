package features

//lint:file-ignore U1000 we want to introduce this feature flag unused.

var (
	// csvExport enables CSV export of search results.
	csvExport = registerFeature("Enable CSV export of search results", "ROX_CSV_EXPORT", false)

	// NetworkDetectionBaselineSimulation enables new features related to the baseline simulation part of the network detection experience.
	NetworkDetectionBaselineSimulation = registerFeature("Enable network detection baseline simulation", "ROX_NETWORK_DETECTION_BASELINE_SIMULATION", true)

	// IntegrationsAsConfig enables loading integrations from config
	IntegrationsAsConfig = registerFeature("Enable loading integrations from config", "ROX_INTEGRATIONS_AS_CONFIG", false)

	// ComplianceOperatorCheckResults enables getting compliance results from the compliance operator
	ComplianceOperatorCheckResults = registerFeature("Enable fetching of compliance operator results", "ROX_COMPLIANCE_OPERATOR_INTEGRATION", true)

	// SystemHealthPatternFly enables the Pattern Fly version of System Health page. (used in the front-end app only)
	SystemHealthPatternFly = registerFeature("Enable Pattern Fly version of System Health page", "ROX_SYSTEM_HEALTH_PF", false)

	// NetworkPolicySystemPolicy enables two system policy fields (Missing (Ingress|Egress) Network Policy) to check deployments
	// against network policies applied in the secured cluster.
	NetworkPolicySystemPolicy = registerFeature("Enable NetworkPolicy-related system policy fields", "ROX_NETPOL_FIELDS", true)

	// NewPolicyCategories enables new policy categories as first-class entities.
	NewPolicyCategories = registerFeature("Enable new policy categories as first-class entities", "ROX_NEW_POLICY_CATEGORIES", false)

	// SearchPageUI enables search page instead of search modal in UI. Frontend only.
	SearchPageUI = registerFeature("Enable search page instead of search modal in UI", "ROX_SEARCH_PAGE_UI", false)

	// DecommissionedClusterRetention enables the setting in System Configuration.
	DecommissionedClusterRetention = registerFeature("Enable Decommissioned Cluster Retention in System Configuration", "ROX_DECOMMISSIONED_CLUSTER_RETENTION", true)

	// QuayRobotAccounts enables Robot accounts as credentials in Quay Image Integration.
	QuayRobotAccounts = registerFeature("Enable Robot accounts in Quay Image Integration", "ROX_QUAY_ROBOT_ACCOUNTS", true)

	// RoxctlNetpolGenerate enables 'roxctl netpol generate' command which integrates with NP-Guard
	RoxctlNetpolGenerate = registerFeature("Enable 'roxctl generate netpol' command", "ROX_ROXCTL_NETPOL_GENERATE", true)

	// ObjectCollections enables 'collection' entity APIs and Frontend collection pages
	ObjectCollections = registerFeature("Enable object collection entities", "ROX_OBJECT_COLLECTIONS", false)

	// NetworkGraphPatternFly enables the PatternFly version of NetworkGraph. (used in the front-end app only)
	NetworkGraphPatternFly = registerFeature("Enable PatternFly version of NetworkGraph", "ROX_NETWORK_GRAPH_PATTERNFLY", false)

	// RHCOSNodeScanning enables phase 1 functions of "Full host level vulnerability scanning for RHCOS nodes" (ROX-10818)
	RHCOSNodeScanning = registerFeature("Enable RHCOS node scanning of OS and installed packages", "ROX_RHCOS_NODE_SCANNING", false)

	// UseFakeNodeInventory tells compliance to use FakeNodeScanner with hardcoded data instead of calling scanner.Analyze()
	// TODO(ROX-13935): Remove this FF and the accompanying code
	UseFakeNodeInventory = registerFeature("Enables compliance to use FakeNodeScanner with hardcoded data instead of calling scanner.Analyze()", "ROX_RHCOS_FAKE_NODE_INVENTORY", false)

	// ProcessesListeningOnPort enables the NetworkFlow code to also update the processes that are listening on ports
	ProcessesListeningOnPort = registerFeature("Enable Processes Listening on Port", "ROX_PROCESSES_LISTENING_ON_PORT", false)

	// ResyncDisabled disables the resync behavior of the kubernetes listeners in sensor
	ResyncDisabled = registerFeature("Disable the re-sync", "ROX_RESYNC_DISABLED", false)

	// ClairV4Scanner enables Clair v4 as an Image Integration option
	ClairV4Scanner = registerFeature("Enable Clair v4 as an Image Integration option", "ROX_CLAIR_V4_SCANNING", false)

	// RoxSyslogExtraFields enables user to add additional key value pairs in syslog alert notification in cef format
	RoxSyslogExtraFields = registerFeature("Enable extra fields for syslog integration", "ROX_SYSLOG_EXTRA_FIELDS", false)

	// SourcedAutogeneratedIntegrations enables adding a "source" to autogenerated integrations.
	SourcedAutogeneratedIntegrations = registerFeature("Enable autogenerated integrations with cluster/namespace/secret source", "ROX_SOURCED_AUTOGENERATED_INTEGRATIONS", false)
)
