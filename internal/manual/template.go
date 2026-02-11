package manual

// DefaultTemplates maps core section names to their default markdown content.
var DefaultTemplates = map[string]string{
	"overview":     overviewTmpl,
	"deploy":       deployTmpl,
	"troubleshoot": troubleshootTmpl,
	"backup":       backupTmpl,
	"maintenance":  maintenanceTmpl,
	"monitoring":   monitoringTmpl,
	"contacts":     contactsTmpl,
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

const overviewTmpl = `---
title: Project Overview
description: High-level summary of this project
tags: overview, architecture
---

# Project Overview

## Summary

<!-- TODO: Describe the project purpose and key components -->

## Architecture

<!-- TODO: Add architecture diagram or description -->

| Component   | Technology | Notes |
|-------------|-----------|-------|
| Web Server  |           |       |
| Database    |           |       |
| Cache       |           |       |
| Task Queue  |           |       |

## Environments

| Environment | URL | Notes |
|-------------|-----|-------|
| Production  |     |       |
| Staging     |     |       |
| Development |     |       |

## Key Repositories

<!-- TODO: List related repositories -->
`

const deployTmpl = `---
title: Deployment Guide
description: Step-by-step deployment procedures
tags: deploy, release
---

# Deployment Guide

## Pre-deployment Checklist

- [ ] All tests passing
- [ ] Database migrations reviewed
- [ ] Configuration changes documented
- [ ] Rollback plan prepared

## Deployment Steps

<!-- TODO: Document your deployment procedure -->

### 1. Prepare Release

` + "```bash" + `
# TODO: Add release preparation commands
` + "```" + `

### 2. Deploy

` + "```bash" + `
# TODO: Add deployment commands
` + "```" + `

### 3. Post-deployment Verification

- [ ] Health check endpoints responding
- [ ] Key user flows verified
- [ ] Monitoring dashboards checked

## Rollback Procedure

<!-- TODO: Document rollback steps -->
`

const troubleshootTmpl = `---
title: Troubleshooting Guide
description: Common issues and resolution steps
tags: troubleshoot, incident, debug
---

# Troubleshooting Guide

## Common Issues

### Issue: High Response Time

**Symptoms:** API latency > 2s

**Diagnosis:**
` + "```bash" + `
# TODO: Add diagnostic commands
` + "```" + `

**Resolution:**
<!-- TODO: Document resolution steps -->

### Issue: Database Connection Errors

**Symptoms:** Connection pool exhaustion

**Diagnosis:**
` + "```bash" + `
# TODO: Add diagnostic commands
` + "```" + `

**Resolution:**
<!-- TODO: Document resolution steps -->

## Escalation Path

| Level | Contact | When |
|-------|---------|------|
| L1    |         | Initial triage |
| L2    |         | > 15 min unresolved |
| L3    |         | Critical / data loss |
`

const backupTmpl = `---
title: Backup & Recovery
description: Backup procedures and disaster recovery
tags: backup, recovery, disaster
---

# Backup & Recovery

## Backup Schedule

| Target    | Frequency | Retention | Location |
|-----------|-----------|-----------|----------|
| Database  |           |           |          |
| Files     |           |           |          |
| Config    |           |           |          |

## Backup Verification

<!-- TODO: How to verify backups are valid -->

## Recovery Procedures

### Full Database Restore

` + "```bash" + `
# TODO: Add restore commands
` + "```" + `

### Partial Data Recovery

<!-- TODO: Document partial recovery procedures -->
`

const maintenanceTmpl = `---
title: Maintenance Procedures
description: Routine maintenance tasks and schedules
tags: maintenance, cron, cleanup
---

# Maintenance Procedures

## Routine Tasks

| Task | Frequency | Command | Notes |
|------|-----------|---------|-------|
|      |           |         |       |

## Database Maintenance

` + "```bash" + `
# TODO: Add database maintenance commands (vacuum, reindex, etc.)
` + "```" + `

## Log Rotation

<!-- TODO: Document log rotation configuration -->

## Cache Management

` + "```bash" + `
# TODO: Add cache flush / warm-up commands
` + "```" + `
`

const monitoringTmpl = `---
title: Monitoring & Alerts
description: Monitoring setup, dashboards, and alert runbooks
tags: monitoring, alerts, metrics
---

# Monitoring & Alerts

## Dashboards

| Dashboard | URL | Description |
|-----------|-----|-------------|
|           |     |             |

## Key Metrics

| Metric | Normal Range | Alert Threshold |
|--------|-------------|-----------------|
| CPU    |             |                 |
| Memory |             |                 |
| Disk   |             |                 |
| Latency|             |                 |

## Alert Runbooks

### Alert: High Error Rate

**Condition:** Error rate > 5% for 5 minutes

**Steps:**
1. <!-- TODO: Add response steps -->

### Alert: Disk Usage Critical

**Condition:** Disk usage > 90%

**Steps:**
1. <!-- TODO: Add response steps -->
`

const contactsTmpl = `---
title: Contacts & Escalation
description: Team contacts and escalation procedures
tags: contacts, oncall, escalation
---

# Contacts & Escalation

## Team Contacts

| Role | Name | Contact | Availability |
|------|------|---------|-------------|
| Tech Lead |  |         |             |
| On-call    |  |         |             |
| DBA        |  |         |             |
| DevOps     |  |         |             |

## Escalation Matrix

| Severity | Response Time | Contact |
|----------|--------------|---------|
| P1 (Critical) | 15 min  |         |
| P2 (Major)    | 1 hour  |         |
| P3 (Minor)    | 4 hours |         |

## External Contacts

| Service | Support URL | Account Info |
|---------|------------|-------------|
| AWS     |            |             |
| CDN     |            |             |
`
