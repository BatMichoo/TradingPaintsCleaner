# iRacing Paint File Cleaner

This Go program helps you manage your iRacing paint files by automatically deleting files that do **not** contain specific identifiers defined in a `.env` file. This is useful for keeping only your favorite or frequently used paint schemes and freeing up disk space.

## Features

* **Configuration via `.env`:** Define the ID of your iRacing profile and/or team IDs you want to keep in a simple `.env` file.
* **Identifier-Based Deletion:** Files whose names do not contain any of the specified identifiers will be deleted.
* **Disk Space Reporting:** After the process, the program reports the number of files deleted and the amount of disk space freed.

## Getting Started

### Prerequisites

* **Go Installation:** Make sure you have Go installed on your system. You can download it from the [official Go website](https://go.dev/dl/).
* **`godotenv` Library:** This project uses the `github.com/joho/godotenv` library to load environment variables from a `.env` file. You can install it using Go modules:

    ```bash
    go get [github.com/joho/godotenv](https://github.com/joho/godotenv)
    ```

### Setup

1.  **Clone the Repository (or create the Go file):** If you have the code in a Git repository, clone it to your local machine. Otherwise, create a new Go file (e.g., `paint_cleaner.go`) and paste the provided code into it.

2.  **Create a `.env` File:** In the same directory as your Go source file, create a file named `.env`.

3.  **Define Your Identifiers:** Open the `.env` file and add a line defining the `IDS` variable. This variable should contain a comma-separated list of strings that you want to **keep** in your iRacing `paint` folder. For example your `.env` file might look like this:

    ```
    IDS=123456,234567
    ```

    **Important:** Only files whose names *do not* contain any of these identifiers will be deleted. Be specific with your identifiers to avoid accidentally deleting files you want to keep.

### Running the Program

1.  **Navigate to the Project Directory:** Open your terminal or command prompt and navigate to the directory where you saved the Go file and the `.env` file.

2.  **Run the Go Program:** Execute the following command:

    ```bash
    go run paint_cleaner.go
    ```

    The program will then:
    * Attempt to load the `.env` file.
    * Read the `IDS` environment variable.
    * Locate your iRacing `paint` folder (usually in `~/Documents/iRacing/paint`).
    * Iterate through the files and subdirectories within the `paint` folder.
    * Delete any files whose names do not contain any of the identifiers specified in the `IDS` variable.
    * Print the number of deleted files and the amount of disk space freed.

## Important Notes

* **No Undo Functionality:** This program permanently deletes files. There is no built-in undo functionality. Trading Paints does download all the missing paints when joining a server.

## Contributing

If you find any issues or have suggestions for improvements, feel free to open a pull request or submit an issue on the repository.
