<div align="center">
  <h1>🌱 AgriFlow</h1>
  <p><strong>A Cloud-Native Agricultural Market Data Pipeline</strong></p>
</div>

## 📋 Overview

AgriFlow is a modern, cloud-native microservices application designed to address the challenges faced by farmers and sellers in Sub-Saharan Africa and other developing regions. It ingests, processes, stores, and serves agricultural market data from multiple sources, enabling better pricing decisions and reducing market inefficiencies.

## 🌟 Features

- **Multi-source data collection** from APIs, web scraping, and SMS
- **Real-time data processing** via stream processing pipeline
- **Scalable architecture** built on cloud-native principles
- **Comprehensive APIs** for downstream applications
- **Notification system** for price alerts and market changes
- **Production-ready** with monitoring, logging, and tracing

## 🛠️ Tech Stack

<div align="center">
  <p>
    <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go">
    <img src="https://img.shields.io/badge/docker-%232496ED.svg?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
    <img src="https://img.shields.io/badge/kubernetes-%23326CE5.svg?style=for-the-badge&logo=kubernetes&logoColor=white" alt="Kubernetes">
    <img src="https://img.shields.io/badge/kafka-%23231F20.svg?style=for-the-badge&logo=apache-kafka&logoColor=white" alt="Kafka">
    <img src="https://img.shields.io/badge/postgres-%23336791.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL">
    <img src="https://img.shields.io/badge/prometheus-%23E6522C.svg?style=for-the-badge&logo=prometheus&logoColor=white" alt="Prometheus">
    <img src="https://img.shields.io/badge/grafana-%23F46800.svg?style=for-the-badge&logo=grafana&logoColor=white" alt="Grafana">
    <img src="https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white" alt="GitHub Actions">
    <img src="https://img.shields.io/badge/helm-%230F1689.svg?style=for-the-badge&logo=helm&logoColor=white" alt="Helm">
  </p>
</div>

## 🧱 Architecture

<div align="center">
  <img src="docs/architecture-diagram.png" alt="AgriFlow Architecture" width="800"/>
</div>

AgriFlow consists of the following microservices:

- **Ingestor Service**: Collects raw data from various market sources
- **ETL Service**: Cleans, transforms, and enriches the raw data
- **API Gateway**: Provides a unified API interface for clients
- **Notifier Service**: Sends alerts based on configured triggers


### Backend
- **Language**: Go
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **Message Bus**: Kafka/NATS

### Data Storage
- **Database**: PostgreSQL/ClickHouse
- **Caching**: Redis

### DevOps & Observability
- **CI/CD**: GitHub Actions
- **Metrics**: Prometheus + Grafana
- **Tracing**: OpenTelemetry + Jaeger
- **Logging**: ELK Stack/Loki

## 🚀 Getting Started

### Prerequisites
- Docker and Docker Compose
- Go 1.20 or later
- kubectl (for Kubernetes deployment)
- Helm (for Kubernetes deployment)

### Local Development

1. Clone the repository
```bash
git clone https://github.com/njange/AgriFlow.git
cd AgriFlow
```

2. Start the local development environment
```bash
cd deploy
docker-compose up -d
```

3. Access the services:
   - API Gateway: http://localhost:8080
   - Grafana: http://localhost:3000

## 📦 Project Structure

```
.
├── LICENSE
├── README.md
├── .github/workflows      # CI/CD pipeline configurations
├── api-gateway/           # API Gateway service
├── deploy/                # Deployment configurations
│   ├── docker-compose.yaml
│   └── k8s/
├── docs/                  # Documentation
├── etl/                   # ETL service
├── ingestor/              # Ingestor service
└── scripts/               # Utility scripts
```

## 🌍 Deployment

### Kubernetes Deployment
```bash
helm install agriflow ./deploy/k8s/helm/agriflow
```

## 📊 Monitoring and Observability

AgriFlow includes comprehensive observability:

- **Metrics**: Prometheus scrapes service metrics, visualized through Grafana dashboards
- **Logging**: Structured logs collected and searchable via ELK/Loki
- **Tracing**: Distributed request tracing with OpenTelemetry and Jaeger

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for more details.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🌟 Acknowledgements

- [OpenTelemetry](https://opentelemetry.io/) for observability tools
- [Kubernetes](https://kubernetes.io/) for orchestration capabilities
- All contributors who have helped shape this project

---

<div align="center">
  <p>Empowering agricultural markets with data-driven insights 🌱</p>
</div>