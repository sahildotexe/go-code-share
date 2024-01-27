# GoShare

GoShare is a command-line tool written in Golang that allows users to easily share file content with others using a temporary URL.



https://github.com/sahildotexe/go-code-share/assets/71642465/98b885cf-e907-4aac-a4ae-59928b14768a



## Usage

To share a file, use the following command:

```bash
goshare <file_path>
```

Replace `<file_path>` with the path to the file you want to share.

## Prerequisites

- Make sure you have [NGROK](https://ngrok.com/) installed on your system.
- Golang
- Git 

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/sahildotexe/go-code-share.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-code-share
   ```

3. Build the executable:

   ```bash
   go build -o goshare
   ```

4. Move the executable to a directory in your system's PATH:
   
   For MacOs/ Linux :
   
   ```bash
   sudo mv goshare /usr/local/bin/
   ```

## Example

```bash
goshare /folder/file.go
```

The tool will generate an NGROK URL that others can use to access the shared file.

## How It Works

1. The tool checks if the user provided a file path.
2. It verifies if the file exists.
3. Creates a temporary directory and a copy of the file to serve.
4. Spawns an NGROK process to expose a local HTTP server.
5. Retrieves the NGROK URL.
6. Prints the NGROK URL for sharing.

## Contributing

If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.
