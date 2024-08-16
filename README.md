# Incident Status Prometheus Exporter

A Prometheus exporter  i
The **Incident Status Prometheus Exporter** is a Prometheus exporter designed to ArvanCloud ongoing and scheduled incidents broadcasting by fetching data from a specified API. This project is written in Golang and provides metrics that can be scraped by Prometheus to track the status of various incidents.

## Table of Contents
<!--ts-->
* [Installation](#installation)
* [Usage](#usage)
* [Features](#features)
* [Configuration](#configuration)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)
<!--te-->

## Installation

### Linux (manual installation)

1. **Clone the repository:**
    ```bash
    git clone https://github.com/r1w/arvanincidents-prometheus-exporter.git
    ```

2. **Navigate into the directory:**
    ```bash
    cd arvanstatus-prometheus-exporter
    ```

3. **Build the project:**
    ```bash
    go build -o incident-status-exporter
    ```

4. **Run the exporter:**
    ```bash
    ./incident-status-exporter
    ```

### Add to Prometheus configuration

Add the following job to your Prometheus `prometheus.yml` configuration file:

```yaml
scrape_configs:
   - job_name: 'incident_status'
     static_configs:
        - targets: ['0.0.0.0:8001']
          labels:
             container: 'incident_status'
```

Restart Prometheus:

    Restart Prometheus to apply the new configuration.

Usage
============
Once the exporter is running, Prometheus will start scraping metrics from it. The exporter listens on a specific port (default: :8080) and exposes metrics that include the status of various ArvanCloud services.

Example metrics include:

    api_incidents_details{end_time="2024-07-07T15:54:16",start_time="2024-07-07T15:12:00",title="Live Streaming Issue",type="major"} 1
    api_incidents_details{end_time="2024-07-07T21:15:26",start_time="2024-07-07T20:30:00",title="Live Streaming Essensial Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-08T19:25:37",start_time="2024-07-08T16:51:53",title="Ticketing Service Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-09T00:28:40",start_time="2024-07-08T23:05:00",title="Ticketing Service Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-09T13:31:49",start_time="2024-07-09T11:00:00",title="Live Streaming Issue",type="major"} 1
    api_incidents_details{end_time="2024-07-09T15:44:44",start_time="2024-07-09T12:00:00",title="CDN API",type="minor"} 1
    api_incidents_details{end_time="2024-07-10T11:20:00",start_time="2024-07-10T10:30:00",title="Scheduled maintenance on Managed Database service",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-11T10:15:47",start_time="2024-07-11T07:20:00",title="Cloud Server Issue",type="major"} 1
    api_incidents_details{end_time="2024-07-16T10:47:15",start_time="2024-07-16T10:16:00",title="CDN API Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-16T16:31:14",start_time="2024-07-16T15:47:00",title="Submitting New Changes in Container Service",type="minor"} 1
    api_incidents_details{end_time="2024-07-17T02:30:00",start_time="2024-07-16T20:31:00",title="CDN API Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-17T16:26:47",start_time="2024-07-17T15:31:37",title="CDN  Log Forwarding Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-21T03:01:00",start_time="2024-07-20T20:30:00",title="Video Platform Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-21T15:41:31",start_time="2024-07-21T14:30:00",title="Video Player Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-21T21:00:00",start_time="2024-07-21T20:00:00",title="Shahriar Scheduled Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-22T13:35:36",start_time="2024-07-22T13:20:00",title="Container Services Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-22T14:43:14",start_time="2024-07-22T14:04:00",title="Shahriar and Bamdad Network Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-25T14:56:16",start_time="2024-07-25T09:29:08",title="Notification Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-26T09:43:29",start_time="2024-07-23T20:31:00",title="Simin Data Center Scheduled Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-27T11:53:19",start_time="2024-07-27T11:16:14",title="Notification Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-28T02:30:00",start_time="2024-07-27T20:30:00",title="Container Services Maintenance ",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-29T04:30:00",start_time="2024-07-29T02:30:00",title="TIC Internet Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-29T22:00:00",start_time="2024-07-29T21:00:00",title="Simin Cloud Server Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-30T17:12:09",start_time="2024-07-30T15:56:00",title="Container Service Panel Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-31T02:30:00",start_time="2024-07-30T20:30:00",title="Container Services Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-07-31T07:34:27",start_time="2024-07-31T06:45:00",title="CDN Panel Issue",type="minor"} 1
    api_incidents_details{end_time="2024-07-31T16:46:58",start_time="2024-07-31T15:57:42",title="CDN Connectivity Issue",type="minor"} 1
    api_incidents_details{end_time="2024-08-01T06:14:12",start_time="2024-08-01T05:52:30",title="CDN API Issue",type="minor"} 1
    api_incidents_details{end_time="2024-08-01T11:00:36",start_time="2024-08-01T07:54:22",title="Ticketing Service Issue",type="minor"} 1
    api_incidents_details{end_time="2024-08-01T13:17:49",start_time="2024-08-01T12:02:00",title="Forough Bandwidtth Drop",type="minor"} 1
    api_incidents_details{end_time="2024-08-01T16:12:51",start_time="2024-08-01T15:10:00",title=" Forough Bandwidtth Drop ",type="minor"} 1
    api_incidents_details{end_time="2024-08-03T23:00:00",start_time="2024-08-03T21:00:00",title="Simin Object Storage Essential Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-04T01:30:00",start_time="2024-08-03T20:31:00",title="Scheduled maintenance on CDN Log Forwarding",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-04T02:30:00",start_time="2024-08-03T20:31:00",title="Container Services Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-05T04:30:00",start_time="2024-08-04T20:30:00",title="Schedule Maintenance on Video Platform",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-06T20:56:54",start_time="2024-08-06T19:02:00",title="Managed Database issue",type="minor"} 1
    api_incidents_details{end_time="2024-08-07T02:30:00",start_time="2024-08-06T20:30:00",title="Container Services Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-07T14:28:16",start_time="2024-08-07T13:00:00",title="Iaas Panel Issue",type="minor"} 1
    api_incidents_details{end_time="2024-08-08T02:30:00",start_time="2024-08-07T20:31:00",title="Container Services Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-08T20:30:00",start_time="2024-08-08T19:30:00",title="Panel and API Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-10T06:30:00",start_time="2024-08-09T22:30:00",title="TIC Scheduled Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-10T22:40:15",start_time="2024-08-10T20:31:00",title="Cloud Server API Scheduled Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-11T22:41:00",start_time="2024-08-11T20:31:00",title="Cloud Server API Scheduled Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-12T23:39:58",start_time="2024-08-12T20:31:00",title="Cloud Server API Scheduled Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-14T01:30:00",start_time="2024-08-13T20:31:00",title="Cloud Server Scheduled Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-14T02:30:00",start_time="2024-08-13T20:30:00",title="Container Services Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-15T02:30:00",start_time="2024-08-14T20:30:00",title="Container Services Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-19T02:30:00",start_time="2024-08-18T20:30:00",title="Container Services Maintenance",type="scheduled"} 1
    api_incidents_details{end_time="2024-08-21T02:30:00",start_time="2024-08-20T20:30:00",title="Container Services Maintenance",type="scheduled"} 1

## Incident Types

The exporter tracks incidents categorized into three types. These types describe the nature of the incident:

- **`type="major"`**: Significant incidents that have a major impact on the service. 
- **`type="minor"`**: Less severe incidents that have a minor impact on the service. 
- **`type="scheduled"`**: Planned maintenance or scheduled events. These are known in advance and are usually communicated to users before they occur.

Command-line Flags

    --port: Specify the port on which the exporter listens (default: 8000).
    --interval: Set the interval for fetching incident data (default: 60s).

Example:

```bash
./incident-status-exporter --port 9090 --interval 120s
```

Features

    Real-time Monitoring: Continuously fetches and exports the status of incidents.
    Customizable: Supports configuration of the scrape interval and listening port.
    Lightweight: Minimal resource usage, ideal for deployment alongside other Prometheus exporters.

Configuration
============

You can configure the exporter via command-line flags or environment variables:
Command-line Flags:

    --port: Port for the exporter to listen on.
    --interval: Interval for fetching the incident data.

Environment Variables:

    EXPORTER_PORT: Set the port number.
    SCRAPE_INTERVAL: Set the interval for scraping the incident data.

Example:

```bash
export EXPORTER_PORT=9090
export SCRAPE_INTERVAL=120s
./incident-status-exporter
```

Contributing
============

We welcome contributions to improve this project.

To contribute:

1. Fork the repository
2. Create a new branch (git checkout -b feature-branch)
3. Make your changes and commit them (git commit -m 'Add some feature')
4. Push to the branch (git push origin feature-branch)
5. Open a Pull Request

> [!IMPORTANT]
> Please follow the coding style and standards outlined in the repository.

License:
============

This project is licensed under the MIT License - see the LICENSE file for details.

Contact info:
============

[![LinkedIn](https://img.shields.io/badge/LinkedIn-blue?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/hamid-hadigol/)  
[![Email](https://img.shields.io/badge/Email-D14836?style=flat&logo=gmail&logoColor=white)](mailto:kurosch86@gmail.com)  
[![GitHub](https://img.shields.io/badge/GitHub-333?style=flat&logo=github&logoColor=white)](https://github.com/r1w/)