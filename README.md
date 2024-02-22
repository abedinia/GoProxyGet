# GoProxyGet

GoProxyGet is a command-line utility written in GoLang that simplifies the process of downloading files through SOCKS5 proxies. It provides a convenient way to access online resources securely and anonymously by routing your requests through SOCKS5 proxy servers.

### Features:

- **Effortless File Retrieval**: Downloading files through SOCKS5 proxies is made easy with GoProxyGet. It abstracts the complexities of proxy handling, allowing you to focus on downloading the data you need.

- **Enhanced Security**: Protect your online identity and maintain privacy while accessing online resources. GoProxyGet ensures that your requests are channeled through a SOCKS5 proxy, keeping your real IP address hidden.

- **Customizable Configuration**: Tailor your proxy settings to match your specific requirements. GoProxyGet supports various SOCKS5 proxy configurations, including authentication methods and proxy chaining.

### Reliable Performance**: Enjoy fast and reliable downloads, thanks to GoLang's efficient concurrency and network libraries. GoProxyGet maximizes your download speed while minimizing resource usage.

### Cross-Platform Compatibility**: Whether you're working on Windows, macOS, or Linux, GoProxyGet is designed to work seamlessly across different operating systems.

- **Open-Source**: GoProxyGet is an open-source project, welcoming contributions from the community. Feel free to enhance its features or customize it to meet your unique needs.

## Usage:

To use GoProxyGet, you can follow this simple command-line syntax:


```bash
go run main.go <proxy_address> <file_url>
```

- `<proxy_address>`: Specify the SOCKS5 proxy address and port you want to use in the format `hostname:port`.
- `<file_url>`: Provide the URL of the file you want to download.

Example:

```bash
go run main.go localhost:2080 https://example.com/sample-file.zip
```

This command will use the specified SOCKS5 proxy to download the file from the given URL.

## Installation

To get started with GoProxyGet, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/abedinia/GoProxyGet.git
```

2. Navigate to the project directory:

```bash
cd goproxyget
```

3. Run GoProxyGet with the appropriate command as described in the "Usage" section.

## Contributing

Contributions to GoProxyGet are welcome! Whether you want to fix a bug, add a feature, or improve documentation, please feel free to submit a pull request. Be sure to follow the project's code of conduct.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
