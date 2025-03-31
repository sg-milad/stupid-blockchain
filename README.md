# Blockchain in Go

there is bug in `SaveToFile` method (maybe you can solve it and send a PR???)

This project is a simplified implementation of a blockchain in Go, inspired by the series **"Building Blockchain in Go"** by Ivan Kuznetsov. The series provides a step-by-step guide to creating a basic blockchain prototype, incorporating features like Proof-of-Work, transactions, addresses, and network capabilities.

[Read the guide here](https://jeiwan.net/posts/building-blockchain-in-go-part-1/?utm_source=chatgpt.com)

## Features

- **Blockchain Structure**: A chain of blocks, each containing a timestamp, data, previous block hash, and its own hash.
- **Proof-of-Work**: A consensus mechanism to secure the blockchain by requiring computational work to add new blocks. ([Read more](https://jeiwan.net/posts/building-blockchain-in-go-part-2/?utm_source=chatgpt.com))
- **Transactions**: Mechanism to securely transfer data between parties, ensuring immutability and transparency. ([Read more](https://jeiwan.net/posts/building-blockchain-in-go-part-4/?utm_source=chatgpt.com))
- **Addresses**: Unique identifiers derived from public keys to receive transactions. ([Read more](https://jeiwan.net/posts/building-blockchain-in-go-part-5/?utm_source=chatgpt.com))
- **Networking**: P2P network implementation to synchronize the blockchain across multiple nodes. ([Read more](https://jeiwan.net/posts/building-blockchain-in-go-part-7/?utm_source=chatgpt.com))

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/sg-milad/stupid-blockchain.git
```

### Navigate to the Project Directory

```bash
cd stupid-blockchain
```

### Install Dependencies

```bash
go mod tidy
```

### Run the Application

```bash
go run main.go
```

## Contributing

Feel free to fork the repository, submit issues, and create pull requests. Ensure that your contributions align with the project's goals and maintain code quality.

## License

This project is licensed under the **MIT License**.
