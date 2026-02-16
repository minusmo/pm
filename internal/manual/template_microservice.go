package manual

const serviceDependenciesTmpl = `---
title: Service Dependencies
description: Upstream and downstream service dependencies
tags: dependencies, services, integration
---

# Service Dependencies

## Dependency Map

<!-- TODO: Add a diagram or description of service relationships -->

## Upstream Dependencies

| Service | Purpose | Protocol | SLA | Owner |
|---------|---------|----------|-----|-------|
|         |         |          |     |       |

## Downstream Dependents

| Service | Purpose | Protocol | Owner |
|---------|---------|----------|-------|
|         |         |          |       |

## External Services

| Service | Purpose | Docs URL | Rate Limits |
|---------|---------|----------|-------------|
|         |         |          |             |

## Failure Modes

| Dependency | Impact if Down | Fallback Strategy |
|------------|---------------|-------------------|
|            |               |                   |
`

const apiContractsTmpl = `---
title: API Contracts
description: API endpoints, schemas, and contract specifications
tags: api, contracts, schema
---

# API Contracts

## Base URL

| Environment | Base URL |
|-------------|----------|
| Production  |          |
| Staging     |          |

## Authentication

<!-- TODO: Describe authentication method (API key, OAuth, JWT, etc.) -->

## Endpoints

### GET /health

**Description:** Health check endpoint

**Response:**
` + "```json" + `
{
  "status": "ok"
}
` + "```" + `

### POST /api/v1/resource

**Description:** <!-- TODO: Describe endpoint -->

**Request:**
` + "```json" + `
{
  // TODO: Add request schema
}
` + "```" + `

**Response:**
` + "```json" + `
{
  // TODO: Add response schema
}
` + "```" + `

## Error Codes

| Code | Meaning | Retry? |
|------|---------|--------|
| 400  | Bad Request | No |
| 401  | Unauthorized | No |
| 429  | Rate Limited | Yes |
| 500  | Internal Error | Yes |
`

const healthChecksTmpl = `---
title: Health Checks
description: Service health check endpoints and liveness/readiness probes
tags: health, probes, liveness, readiness
---

# Health Checks

## Endpoints

| Path | Type | Purpose |
|------|------|---------|
| /health | Liveness | Basic process alive check |
| /ready  | Readiness | Dependencies connected |

## Liveness Check

` + "```bash" + `
# TODO: Add liveness check command
` + "```" + `

**Expected response:** HTTP 200

**What it verifies:**
- Process is running
- Not deadlocked

## Readiness Check

` + "```bash" + `
# TODO: Add readiness check command
` + "```" + `

**Expected response:** HTTP 200

**What it verifies:**
- Database connection is alive
- Cache is reachable
- Downstream services are available

## Kubernetes Probes

` + "```yaml" + `
# TODO: Add probe configuration
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 15
readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
` + "```" + `
`

const scalingTmpl = `---
title: Scaling Guide
description: Horizontal and vertical scaling strategies
tags: scaling, performance, capacity
---

# Scaling Guide

## Current Configuration

| Resource | Min | Max | CPU Request | Memory Request |
|----------|-----|-----|-------------|----------------|
|          |     |     |             |                |

## Horizontal Scaling

### Auto-scaling Rules

| Metric | Target | Scale Up | Scale Down |
|--------|--------|----------|------------|
| CPU    |        |          |            |
| Memory |        |          |            |
| RPS    |        |          |            |

### Manual Scaling

` + "```bash" + `
# TODO: Add manual scaling commands
` + "```" + `

## Vertical Scaling

<!-- TODO: Document when and how to vertically scale -->

## Bottlenecks & Limits

| Component | Limit | Mitigation |
|-----------|-------|------------|
| Database connections |  |    |
| API rate limits      |  |    |
| Memory per instance  |  |    |

## Load Testing

` + "```bash" + `
# TODO: Add load testing commands
` + "```" + `
`
