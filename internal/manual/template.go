package manual

// DefaultTemplates maps section names to their default markdown content.
var DefaultTemplates = map[string]string{
	// Core (ops/SRE)
	"overview":     overviewTmpl,
	"deploy":       deployTmpl,
	"troubleshoot": troubleshootTmpl,
	"backup":       backupTmpl,
	"maintenance":  maintenanceTmpl,
	"monitoring":   monitoringTmpl,
	"contacts":     contactsTmpl,

	// Onboarding
	"setup-guide":          setupGuideTmpl,
	"codebase-walkthrough": codebaseWalkthroughTmpl,
	"dev-workflow":         devWorkflowTmpl,
	"coding-conventions":   codingConventionsTmpl,

	// Microservice
	"service-dependencies": serviceDependenciesTmpl,
	"api-contracts":        apiContractsTmpl,
	"health-checks":        healthChecksTmpl,
	"scaling":              scalingTmpl,

	// Library
	"api-reference":  apiReferenceTmpl,
	"usage-examples": usageExamplesTmpl,
	"versioning":     versioningTmpl,
	"publishing":     publishingTmpl,
	"contributing":   contributingTmpl,

	// Framework
	"getting-started": gettingStartedTmpl,
	"plugin-system":   pluginSystemTmpl,
	"migration-guide": migrationGuideTmpl,
}

// CoreSectionOrder defines the canonical ordering of core sections.
var CoreSectionOrder = []string{
	"overview",
	"deploy",
	"troubleshoot",
	"backup",
	"maintenance",
	"monitoring",
	"contacts",
}
