# ChatMate

ChatMate is an innovative two-person chatting application built using the Golang Gin framework and MongoDB, designed with Clean Architecture principles. The app aims to provide a feature-rich, attractive, and user-friendly experience for real-time communication between two users.

## Features

- **Real-time Messaging**: Enables two users to communicate instantly.
- **User Management**: Handles user registration, login, and authentication.
- **Innovative Frontend**: Colorful, attractive, and user-friendly interface with features like sender and receiver names displayed.
- **Clean Architecture**: Organized codebase following best practices for maintainability and scalability.
- **MongoDB Integration**: Efficient data storage and retrieval for messages and user data.

## Tech Stack

- **Backend**: Golang, Gin framework
- **Database**: MongoDB
- **Frontend**: HTML, CSS, JavaScript
- **Architecture**: Clean Architecture

## Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/yourusername/ChatMate.git
   cd ChatMate
   ```

2. **Install dependencies**:
   Ensure you have Golang and MongoDB installed on your machine. Then run:
   ```sh
   go mod download
   ```

3. **Set up MongoDB**:
   Start your MongoDB server and update the connection details in `pkg/Database/mongodb.go`.

4. **Run the application**:
   ```sh
   go run cmd/main.go
   ```

## Usage

1. Open your web browser and navigate to `http://localhost:8080`.
2. Register a new user or log in with existing credentials.
3. Start chatting with another user in real-time.

## Contributing

Contributions are welcome! Please follow these steps:

1. **Fork the repository**.
2. **Create a new branch**:
   ```sh
   git checkout -b feature/your-feature-name
   ```
3. **Make your changes**.
4. **Commit your changes**:
   ```sh
   git commit -m 'Add some feature'
   ```
5. **Push to the branch**:
   ```sh
   git push origin feature/your-feature-name
   ```
6. **Create a Pull Request**.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or issues, please contact [your-email@example.com].

---

Thank you for using ChatMate! We hope you enjoy your experience.
```

Feel free to customize the details such as the repository URL, email, and any other specific information about your project. This README provides an advanced and professional overview of your ChatMate project.
