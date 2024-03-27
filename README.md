# Password Generator

This project is a simple web application that allows users to generate random passwords and test their strength.

## Features

- **Password Generation**: Users can generate a random password by specifying the desired length. The generated password includes a mix of uppercase and lowercase letters, numbers, and special characters.

- **Password Strength Testing**: Users can test the strength of a password. The strength is determined based on several factors, including the length of the password, the types of characters used, and the predictability of the password.

## Usage

Visit the website [here](https://jwt2706.ca/PasswordGenerator).

## Development

This project uses just plain html/css/js for the frontend, and golang for the backend.

If you want to run the backend locally, you will need to have golang installed on your machine. You can then run the backend by executing `go run generate-password.go` and `go run strength-test.go` in the terminal respectively.

For the frontend, it's just `npx vite`.

## License

This project is licensed under the MIT License.

## Disclaimer

This project was more for a learning experiment for me, so any passwords generated or tests done should not be relied upon for securing sensitive information. Always use your best judgment when creating and managing passwords.
