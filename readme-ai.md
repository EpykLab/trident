<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">TRIDENT</h1>
</p>
<p align="center">
    <em>Logs & Learning💾</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/epyklab/trident?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/epyklab/trident?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/epyklab/trident?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/epyklab/trident?style=default&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<br><!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary><br>

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Tests](#-tests)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)
</details>
<hr>

##  Overview

Trident Agent-An Open Source Honeypot Application Tridents Agent is an open-source honeypot application designed to enhance incident detection capabilities. It offers user-friendly landing pages, including a login interface for secure authentication and a search function to streamline inquiries. A log file stores system data with timestamps and filenames, promoting efficient diagnostics, tracing issues, and referencing logs for future purposes. The application communicates via GRPC protocol for logging events, leveraging gzip compression for efficient storage and transmission of data. Load tests are possible using the `load_test.sh` script, ensuring optimal performance under high traffic conditions. The log system employs centralized logging with structured logs for traceability, while a unit test suite validates server functionality.

---

##  Features



---

##  Repository Structure

```sh
└── trident/
    ├── .github
    │   ├── dependabot.yml
    │   └── workflows
    ├── README.md
    ├── agent
    │   ├── .goreleaser.yml
    │   ├── Dockerfile
    │   ├── README.Docker.md
    │   ├── README.md
    │   ├── Vagrantfile
    │   ├── build.sh
    │   ├── compose.yaml
    │   ├── config.yaml
    │   ├── go.mod
    │   ├── go.sum
    │   ├── logparser
    │   ├── logstream
    │   ├── main.go
    │   ├── malhandler
    │   ├── protobuff.sh
    │   ├── utils
    │   └── webserver
    ├── agent-versions
    │   └── Dockerfile.tml.yml
    ├── docker-compose.yaml
    └── server
        ├── .goreleaser.yml
        ├── Dockerfile
        ├── README.Docker.md
        ├── api
        ├── cmd
        ├── compose.yaml
        ├── go.mod
        ├── go.sum
        ├── internal
        ├── logstream
        ├── logstream.proto
        ├── pkg
        ├── scripts
        └── test
```

---

##  Modules

<details closed><summary>.</summary>

| File                                                                                      | Summary                                                                                                                                                                                                                                                                                                                                                                                      |
| ---                                                                                       | ---                                                                                                                                                                                                                                                                                                                                                                                          |
| [docker-compose.yaml](https://github.com/epyklab/trident/blob/master/docker-compose.yaml) | Orchestrates two services-Trident agent and server, within predefined IP address subnets for communication (172.16.238.x & 172.16.239.x).2. Employs custom Docker images from GitHub Container Registry (ghcr.io/epyklab) for each service, and opens essential ports (8080).3. Ensures always-on restart of services and elevated container capabilities to administer network (NET_ADMIN). |

</details>

<details closed><summary>.github</summary>

| File                                                                                    | Summary                                                                                                                                                                                                                                                  |
| ---                                                                                     | ---                                                                                                                                                                                                                                                      |
| [dependabot.yml](https://github.com/epyklab/trident/blob/master/.github/dependabot.yml) | Dependabot configuration file optimizes the software repository by automatically updating devcontainer packages on a weekly basis, fostering version consistency and promoting efficiency in package management within the projects container ecosystem. |

</details>

<details closed><summary>agent-versions</summary>

| File                                                                                                   | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| ---                                                                                                    | ---                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| [Dockerfile.tml.yml](https://github.com/epyklab/trident/blob/master/agent-versions/Dockerfile.tml.yml) | Dockerfile.tml.yml prepares the agents Go binaries for efficient execution by defining the required GO_VERSION and target architecture (GOARCH).2. Streamline: The script optimizes the build process, using caching and mounting essential directories to reduce compilation time and dependency management.3. Deploy: Finally, the Dockerfile sets up a production environment based on Ubuntu, ensuring secure and seamless operation of the Trident agent by installing required packages, managing log files, and configuring OpenSSH server. |

</details>

<details closed><summary>agent</summary>

| File                                                                                    | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| ---                                                                                     | ---                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [.goreleaser.yml](https://github.com/epyklab/trident/blob/master/agent/.goreleaser.yml) | The `.goreleaser.yml` file in Tridents agent directory configures its continuous delivery pipeline, automating the Linux binary release process for the Go-based agent software. This setup includes signing with Cosign for secure distribution, and broadcasting updates via Telegram when a new version is available.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| [Dockerfile](https://github.com/epyklab/trident/blob/master/agent/Dockerfile)           | The Dockerfile within `trident/agent` instructs a multi-stage container to build an executable for the application called trident." It optimizes the process by utilizing Docker's caching mechanisms and leverages separate stages for building and running the app on a Ubuntu platform.2. **Deploy & Run:** Once the trident binary is built, the final container stage is launched, with its primary responsibilities being installing necessary runtime dependencies and creating a non-privileged user to run the application. The resulting executable is copied from the build stage and executed when the container starts.3. **Ports Exposure:** The container exposes ports 5555 and 8080, allowing communication with the running application instance in other environments or services.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| [Vagrantfile](https://github.com/epyklab/trident/blob/master/agent/Vagrantfile)         | Configures virtual machine setup for Tridents agent software via Vagrant (Ruby script), utilizing Ubuntu 22.04 as its base image.2. Deployment Flexibility: Allows developers to manage resources like memory, CPUs, and storage pool during deployment in Libvirt provider.3. File Synchronization: Enables the synchronization of current working directory with the guest systems `/home/vangrant` folder during the VM setup.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| [build.sh](https://github.com/epyklab/trident/blob/master/agent/build.sh)               | Config.yaml`, directory structure (e.g., /etc/trident, /var/log/trident.log), ensuring vital agent settings are in place to maintain a smooth functioning of the Trident ecosystem.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| [compose.yaml](https://github.com/epyklab/trident/blob/master/agent/compose.yaml)       | The configuration includes defining a PostgreSQL database (commented out) that could potentially be utilized by the application, demonstrating modularity.4. The containerized services depend on each other in a coordinated manner, ensuring seamless operation upon startup.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| [config.yaml](https://github.com/epyklab/trident/blob/master/agent/config.yaml)         | This YAML file tailors the Trident Agent to monitor specified system logs.2. Define: It selects services (auditd, ssh, webserver) for monitoring, while ignoring others (faillog by default).3. Stream: It sets the URL for the log stream and specifies log files (/var/log/auth.log by default).4. Capture": Enables uploads of captured logs to the centralized platform and offers customization for web server templates.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [go.mod](https://github.com/epyklab/trident/blob/master/agent/go.mod)                   | The go.mod file anchors dependencies for the agent module in Tridents architecture. It orchestrates the integration of external libraries such as gin-gonic/gin, google/protobuf, and fsnotify, facilitating seamless communication, data handling, and configuration for the agent component in the system.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [go.sum](https://github.com/epyklab/trident/blob/master/agent/go.sum)                   | TridentThe provided code file resides within the `trident/agent` directory of this open-source project, indicating that it is part of the Agent component of the Trident system.The key purpose of the agent code in relation to the repository architecture is to:1. **Build and manage a containerized applicationThe Dockerfile, one of the critical artifacts within the agent directory, serves as instructions for creating a docker image that can contain the necessary dependencies and run the agent.2. **Automate release processesWith.goreleaser.yml configuration file, this code aims to automate the process of creating and managing releases by handling tasks like versioning, packaging, and distributing binaries for various platforms.3. **Implement continuous integration/continuous deployment (CI/CD)The workflows defined in the `trident/.github` folder enable automating the building, testing, and deployment of changes made to the agent component. This promotes faster feedback loops and ensures that changes are consistently delivered efficiently.4. **Support a maintainable open-source projectBy incorporating dependabot.yml within the.github directory, Tridents maintenance team can automatically receive and manage updates for its dependencies, promoting best practices around software maintenance and security. |
| [main.go](https://github.com/epyklab/trident/blob/master/agent/main.go)                 | The `main.go` script initiates and coordinates essential services within the agent directory of Trident repository, such as log parsing, file monitoring, logging streaming, malware handling, and web server functionality.2. Facilitate: The agent orchestrated in this code collects, processes, and transmits log data to ensure comprehensive log management within the Trident application.3. Manage: The main function of `main.go` initiates and spawns go routines for various modules like log parsing, file watcher, logging shipper, malware handler, and web server, facilitating efficient and effective Trident system performance.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [protobuff.sh](https://github.com/epyklab/trident/blob/master/agent/protobuff.sh)       | Configure Protocol Buffers setup in agent folder, ensuring efficient message interchange between microservices.2. Enable: Installs necessary compilers (protobuf-compiler), Go plugins (protoc-gen-go and protoc-gen-go-grpc), and updates PATH for smooth compilation.3. Accomplish: Verifies correct installation and provides paths for the installed components, facilitating protocol buffer usage in the applications architecture.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |

</details>

<details closed><summary>server</summary>

| File                                                                                     | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ---                                                                                      | ---                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| [.goreleaser.yml](https://github.com/epyklab/trident/blob/master/server/.goreleaser.yml) | This `.goreleaser.yml` file within the Trident servers architecture sets up automated releases for the application on Linux platforms. It invokes CGo-free builds with LD flags (-s-w) and applies cosign digital signatures to ensure authenticity before announcements via Telegram. This process streamlines deployment workflows, ensuring a seamless update experience.                                                             |
| [Dockerfile](https://github.com/epyklab/trident/blob/master/server/Dockerfile)           | This Dockerfile within the Trident repository constructs an Alpine Linux-based container for the server application. It downloads dependencies and builds the Go application using specified parameters. Upon successful construction, it installs minimal runtime requirements and copies the built executable into the new stage. The final container starts the server on port 5555 with the specified user credentials when invoked. |
| [compose.yaml](https://github.com/epyklab/trident/blob/master/server/compose.yaml)       | Navigates and configures the server service in this Docker Compose environment. It directs the build of the application using the local Dockerfile and opens a specific port for external access, enabling interaction between services. Additional service dependencies or configuration are optional and can be customized as needed.                                                                                                  |
| [go.mod](https://github.com/epyklab/trident/blob/master/server/go.mod)                   | The `server/go.mod` file defines the modular structure for a core service within the Trident project, ensuring interoperability with external libraries and Google-specific APIs for robustness and scalability.2. **Crucial ComponentsIt facilitates dependency management using Go Modules (`google.golang.org/grpc`, `google.golang.org/protobuf`) as well as critical indirect dependencies, including Golangs x-package.            |
| [go.sum](https://github.com/epyklab/trident/blob/master/server/go.sum)                   | This file records dependencies cryptographic hash sums for the Trident server project, ensuring the correct version is installed when modules are downloaded, preventing potential security threats or software errors due to wrong package versions.                                                                                                                                                                                    |
| [logstream.proto](https://github.com/epyklab/trident/blob/master/server/logstream.proto) | Implement LogService protocol in protobuf for seamless log communication.2. **ExtendFacilitates the exchange of structured log data between components, using a gRPC interface.3. **IntegrateEssential part of Trident servers microservices architecture, handling logging operations efficiently.                                                                                                                                      |

</details>

<details closed><summary>.github.workflows</summary>

| File                                                                                                                          | Summary                                                                                                                                                                                                                                                                                                                                                                                                     |
| ---                                                                                                                           | ---                                                                                                                                                                                                                                                                                                                                                                                                         |
| [agent-container-generic.yaml](https://github.com/epyklab/trident/blob/master/.github/workflows/agent-container-generic.yaml) | The `agent-container-generic.yaml` workflow script initiates and manages containerized Trident Agent deployments across various environments, ensuring smooth integration and scalability. It utilizes GitHub Actions for continuous deployment strategies, enhancing the repositorys infrastructure and reliability.                                                                                       |
| [server-container.yaml](https://github.com/epyklab/trident/blob/master/.github/workflows/server-container.yaml)               | GitHub Actions workflow for deploying server container.2. Targets: Docker-compose configuration to launch the Trident server instance.3. Goal: Enables seamless, automated container deployment with version updates.4. Facilitates: Continuous integration and delivery, ensuring robust server functionality.5. Maintains: Consistency in deploying server containers across various GitHub environments. |

</details>

<details closed><summary>agent.logparser</summary>

| File                                                                                    | Summary                                                                                                                                                                                                                                                                                                                                                             |
| ---                                                                                     | ---                                                                                                                                                                                                                                                                                                                                                                 |
| [generic.go](https://github.com/epyklab/trident/blob/master/agent/logparser/generic.go) | Tail logs from specified paths like /var/log/auth.log, converts the lines into structured JSON format and passes them to a channel. Enhances logging consistency for Tridents agent component.                                                                                                                                                                      |
| [systemd.go](https://github.com/epyklab/trident/blob/master/agent/logparser/systemd.go) | Logparser in the Trident agent processes Systemd journal logs, extracting critical insights that enhance system monitoring for superior performance analysis.2. Connect": This logparser module in the Trident agent integrates with Systemd journal to provide continuous system health checks across diverse Linux environments within our platform architecture. |

</details>

<details closed><summary>agent.logstream</summary>

| File                                                                                              | Summary                                                                                                                                                                                                                                                                                                                                                                      |
| ---                                                                                               | ---                                                                                                                                                                                                                                                                                                                                                                          |
| [logstream.proto](https://github.com/epyklab/trident/blob/master/agent/logstream/logstream.proto) | Remote logging protocol between the Trident Agent and Server using protobufs.2. Structures: Log entries as messages, providing a unified format for logging data exchange.3. Facilitates: Efficient bi-directional communication for log management, ensuring seamless system operations.                                                                                    |
| [stream.go](https://github.com/epyklab/trident/blob/master/agent/logstream/stream.go)             | The ShipLogs' function, found in logstream/stream.go, acts as the data sender within Trident's architecture.2. It ships logs (events, status updates) to a specified server, defined by the TRIDENT_SERVER_ADDRESS environment variable.3. Utilizing gRPC and Go language, it streams incoming data via channel lineChan', ensuring real-time logging to a central location. |

</details>

<details closed><summary>agent.malhandler</summary>

| File                                                                                                 | Summary                                                                                                                                                                                                                                                                                                                                                                                                     |
| ---                                                                                                  | ---                                                                                                                                                                                                                                                                                                                                                                                                         |
| [handleuploads.go](https://github.com/epyklab/trident/blob/master/agent/malhandler/handleuploads.go) | Monitoring component monitors files uploaded into a specified directory. When new files are created, this file-watching service base64 encodes their content and logs the events. Subsequently, it removes these processed files from the system, with an option to hash them before deletion. This ensures timely logging of file upload activities for effective monitoring within Tridents agent module. |

</details>

<details closed><summary>agent.utils</summary>

| File                                                                                    | Summary                                                                                                                                                                                                                       |
| ---                                                                                     | ---                                                                                                                                                                                                                           |
| [confparse.go](https://github.com/epyklab/trident/blob/master/agent/utils/confparse.go) | Manages services, stream URL, log files, file uploads, and webfiles. Written in Golang using YAML parsing to facilitate easy customization of agent behavior for diverse use cases within the parent repository architecture. |

</details>

<details closed><summary>agent.webserver</summary>

| File                                                                                        | Summary                                                                                                                                                                                                                                                                                                                                                                        |
| ---                                                                                         | ---                                                                                                                                                                                                                                                                                                                                                                            |
| [webserver.go](https://github.com/epyklab/trident/blob/master/agent/webserver/webserver.go) | Manages web server operations for Trident Agent, providing route handling, HTML rendering, and POST data processing. Implemented as Gin framework-based application with middleware logging functionality for tracking user activities such as login attempts and search queries. Webserver also facilitates file uploads by saving user submissions to specified directories. |

</details>

<details closed><summary>agent.logstream.rpcservice</summary>

| File                                                                                               | Summary                                                                                                                                                                                                                                                                                          |
| ---                                                                                                | ---                                                                                                                                                                                                                                                                                              |
| [sendrpc.go](https://github.com/epyklab/trident/blob/master/agent/logstream/rpcservice/sendrpc.go) | Simplifies logging within the Trident ecosystem by providing an RPC access to logging functionality.2. Key Feature: LogService-Exposes logMessage method enabling remote systems to submit logs for processing in real-time, increasing centralization and streamlining system management tasks. |

</details>

<details closed><summary>agent.logstream.rpcstream</summary>

| File                                                                                                                  | Summary                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ---                                                                                                                   | ---                                                                                                                                                                                                                                                                                                                                                                                                                          |
| [logstream.pb.go](https://github.com/epyklab/trident/blob/master/agent/logstream/rpcstream/logstream.pb.go)           | LogEntry and LogResponse to handle log stream communication.3. Establish dependencies between defined types based on input/output requirements.4. Create a custom exporter function for handling access to message fields while preserving struct integrity.5. Perform type validation during initialization with an option for unsafe enabled mode.6. Build protocol data structure with Go-specific types and definitions. |
| [logstream_grpc.pb.go](https://github.com/epyklab/trident/blob/master/agent/logstream/rpcstream/logstream_grpc.pb.go) | Generates RPC-based client and server implementations for logging messages in the Trident architecture. Leverages protocol buffer files to define log entry format and gRPC for communication. This enables the exchange of structured log data between the Agent and Server.                                                                                                                                                |

</details>

<details closed><summary>agent.utils.logging</summary>

| File                                                                                      | Summary                                                                                                                                                                                                                                                                             |
| ---                                                                                       | ---                                                                                                                                                                                                                                                                                 |
| [format.go](https://github.com/epyklab/trident/blob/master/agent/utils/logging/format.go) | The `format.go` file creates structured log data within the utils/logging package, forming human-readable and JSON-serializable log format. It provides customization for log timestamp, UUID, source, and message fields.                                                          |
| [logger.go](https://github.com/epyklab/trident/blob/master/agent/utils/logging/logger.go) | Logger utility.In Tridents agent module, this logger initializes a file named trident.log. This log file stores application data with timestamp and file-name for ease of tracking issues, enhancing the system's diagnostic capabilities and providing a log for future reference. |

</details>

<details closed><summary>agent.webserver.templates</summary>

| File                                                                                                | Summary                                                                                                                                                                                                                                                                                                            |
| ---                                                                                                 | ---                                                                                                                                                                                                                                                                                                                |
| [home.html](https://github.com/epyklab/trident/blob/master/agent/webserver/templates/home.html)     | The home.html template defines a user-friendly landing page for our honeypot agent application.2. Functionality: It welcomes users, offers login and search options, and provides an upload form for suspicious files, enhancing incident detection capabilities.                                                  |
| [login.html](https://github.com/epyklab/trident/blob/master/agent/webserver/templates/login.html)   | User Login System (agent/webserver/templates/login.html)The provided HTML template defines a user login interface within the Trident agent software architecture. Users are invited to input their username and password for secure authentication, facilitating seamless access to the platforms functionalities. |
| [search.html](https://github.com/epyklab/trident/blob/master/agent/webserver/templates/search.html) | Agent/webserver/templates/search.html In this repository, we provide an interactive search function for users. The HTML template renders a simple form that accepts text input and initiates searches through the designated API endpoint /search when submitted.                                                  |

</details>

<details closed><summary>server.logstream</summary>

| File                                                                                                         | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| ---                                                                                                          | ---                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| [logstream.pb.go](https://github.com/epyklab/trident/blob/master/server/logstream/logstream.pb.go)           | Import necessary packages and define your file structs using the given raw data, applying appropriate message info.2. Utilize gzip compression to compress your file's raw data for efficient transmission and storage.3. Initialize your protobuf definitions by building a ProtoFile object based on defined types and dependencies.4. Store this newly initialized File_logstream_proto for later use in other modules or functions as required. |
| [logstream_grpc.pb.go](https://github.com/epyklab/trident/blob/master/server/logstream/logstream_grpc.pb.go) | Generate protobuf-generated Go code for a gRPC LogService, enabling communication between server and client for logging events. The provided file, `server/logstream/logstream_grpc.pb.go`, serves as the foundation for defining log stream communication in the Trident repository architecture.                                                                                                                                                  |

</details>

<details closed><summary>server.scripts</summary>

| File                                                                                       | Summary                                                                                                                                                                                                                                                                                             |
| ---                                                                                        | ---                                                                                                                                                                                                                                                                                                 |
| [load_test.sh](https://github.com/epyklab/trident/blob/master/server/scripts/load_test.sh) | The script `load_test.sh`, situated within `server/scripts` folder, serves as a placeholder for conducting load tests on the server. Users can employ tools such as `hey`, `ab`, or custom scripts to assess the servers load capacity, ensuring optimal performance under high traffic conditions. |

</details>

<details closed><summary>server.api.logstream</summary>

| File                                                                                                   | Summary                                                                                                                                                                                                                               |
| ---                                                                                                    | ---                                                                                                                                                                                                                                   |
| [logstream.proto](https://github.com/epyklab/trident/blob/master/server/api/logstream/logstream.proto) | The logstream.proto file in tridents server architecture defines a Protocol Buffers service named LogService', enabling seamless communication between services by exchanging LogEntry and LogResponse messages for logging purposes. |

</details>

<details closed><summary>server.cmd.server</summary>

| File                                                                                | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| ---                                                                                 | ---                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| [main.go](https://github.com/epyklab/trident/blob/master/server/cmd/server/main.go) | The main.go file within server/cmd/server initiates a TCP listener on port 5555 for the Trident server, accepting incoming GRPC requests.2. **GRPC Service ImplementationIt sets up a LogServiceServer and defines a single method `LogMessage`, which accepts GRPC messages, logs them, and sends a success response.3. **Connecting ComponentsThe code establishes connections between the external client and the internal logic of Tridents server, handling incoming logs via GRPC communication. |

</details>

<details closed><summary>server.internal.logservice</summary>

| File                                                                                               | Summary                                                                                                                                                                                                                                                                                                                                                                            |
| ---                                                                                                | ---                                                                                                                                                                                                                                                                                                                                                                                |
| [service.go](https://github.com/epyklab/trident/blob/master/server/internal/logservice/service.go) | This file within the Trident servers internal logservice package defines a LogServiceServer structure that handles incoming LogMessage requests from the protobuf API. The server returns success confirmation upon receiving and processing the log entry. This crucial component is an essential part of the Trident server architecture, managing logs for seamless operations. |

</details>

<details closed><summary>server.internal.server</summary>

| File                                                                                         | Summary                                                                                                                                                                                                                                                                                                             |
| ---                                                                                          | ---                                                                                                                                                                                                                                                                                                                 |
| [server.go](https://github.com/epyklab/trident/blob/master/server/internal/server/server.go) | This `server.go` script initiates a gRPC server for Tridents log management service.2. It registers the logService implementation as a server-side endpoint within the provided grpc instance.3. Interacts with the defined Protobuf protocol and dependencies to manage logs efficiently in Tridents architecture. |

</details>

<details closed><summary>server.pkg.logger</summary>

| File                                                                                      | Summary                                                                                                                                                                                                                                                                                                                                                                                                  |
| ---                                                                                       | ---                                                                                                                                                                                                                                                                                                                                                                                                      |
| [loggger.go](https://github.com/epyklab/trident/blob/master/server/pkg/logger/loggger.go) | A module within the Trident server package introduces the logger, a crucial system for managing logs. It exports InfoLogger and ErrorLogger, creating structured logging for better traceability and error handling. By initializing loggers in a central location, it simplifies and standardizes log management across the application, promoting maintainable code and efficient debugging processes. |

</details>

<details closed><summary>server.pkg.protobuf</summary>

| File                                                                                                            | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| ---                                                                                                             | ---                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| [logstream.pb.go](https://github.com/epyklab/trident/blob/master/server/pkg/protobuf/logstream.pb.go)           | Import the generated protocol file for the `logstream` service (file_logstream_proto) to start using the defined message types, LogEntry and LogResponse.2. Implement the LogService method by defining input and output data structures in accordance with the specified protocol for handling log messages.3. Create a new goroutine or function that continuously reads from an appropriate data source (such as a log file) to generate the desired log messages, sending them to the provided LogService client using the generated interface.4. To compress the logstream data during transfer, make use of the provided `file_logstream_proto_rawDescGZIP()` method that will automatically compress and decompress data as necessary for transmission.5. Regularly inspect logs from the receiving side to ensure correct functioning and adapt code if issues arise during the communication process with the logstream service. |
| [logstream_grpc.pb.go](https://github.com/epyklab/trident/blob/master/server/pkg/protobuf/logstream_grpc.pb.go) | Generate Go code for gRPC service definitions (specifically `logstream_grpc.pb.go`) from `logstream.proto`.2. Role: Key component in the parent repository's server, providing gRPC communication between LogService components via methods like `LogMessage` and defining corresponding RPC handlers.3. Usage: Enables bi-directional communication between log stream agents and the server to exchange log entries in the specified architecture.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |

</details>

<details closed><summary>server.test.integration</summary>

| File                                                                                                    | Summary                                                                                                                                                                                                                                                                                             |
| ---                                                                                                     | ---                                                                                                                                                                                                                                                                                                 |
| [server_test.go](https://github.com/epyklab/trident/blob/master/server/test/integration/server_test.go) | Server/test/integration/server_test.go contains the Go-implemented TestServer function, which validates the servers correct operation when integrated with other components within the Trident system. This ensures reliable communication and data processing throughout our open-source platform. |

</details>

<details closed><summary>server.test.unit</summary>

| File                                                                                               | Summary                                                                                                                                                                                              |
| ---                                                                                                | ---                                                                                                                                                                                                  |
| [service_test.go](https://github.com/epyklab/trident/blob/master/server/test/unit/service_test.go) | Test unit functions for the LogService within the servers test suite, validating its functionality by simulating various input scenarios, ensuring error-free data handling and optimal performance. |

</details>

---

##  Getting Started

**System Requirements:**

* **Go**: `version x.y.z`

###  Installation

<h4>From <code>source</code></h4>

> 1. Clone the trident repository:
>
> ```console
> $ git clone https://github.com/epyklab/trident
> ```
>
> 2. Change to the project directory:
> ```console
> $ cd trident
> ```
>
> 3. Install the dependencies:
> ```console
> $ go build -o myapp
> ```

###  Usage

<h4>From <code>source</code></h4>

> Run trident using the command below:
> ```console
> $ ./myapp
> ```

###  Tests

> Run the test suite using the command below:
> ```console
> $ go test
> ```

---

##  Project Roadmap

- [X] `► INSERT-TASK-1`
- [ ] `► INSERT-TASK-2`
- [ ] `► ...`

---

##  Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/epyklab/trident/issues)**: Submit bugs found or log feature requests for the `trident` project.
- **[Submit Pull Requests](https://github.com/epyklab/trident/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/epyklab/trident/discussions)**: Share your insights, provide feedback, or ask questions.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/epyklab/trident
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="center">
   <a href="https://github.com{/epyklab/trident/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=epyklab/trident">
   </a>
</p>
</details>

---

##  License

This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

##  Acknowledgments

- List any resources, contributors, inspiration, etc. here.

[**Return**](#-overview)

---
