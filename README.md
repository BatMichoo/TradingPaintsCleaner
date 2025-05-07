# iRacing Paint File Cleaner

This Go program helps you manage your iRacing paint files by automatically deleting files that do **not** contain specific identifiers. It reads these identifiers from an environment variable, allowing you to keep only your favorite or frequently used paint schemes and free up disk space.

## Features

- **Configuration via Environment Variable:** The program uses an environment variable to determine which paint files to keep.
- **Identifier-Based Deletion:** Files whose names do not contain any of the specified identifiers will be deleted.
- **Disk Space Reporting:** After the process, the program reports the number of files deleted and the amount of disk space freed.

## Getting Started

### Setup

1.  **Download the latest release or clone the Repository (or create the Go file):** If you have the code in a Git repository, clone it to your local machine. Otherwise, create a new Go file (e.g., `paint_cleaner.go`) and paste the provided code into it.

2.  **Set the `iRacing_IDS` Environment Variable:** The program reads a comma-separated list of identifiers from the `iRacing_IDS` environment variable. You need to set this variable on your system.

    You can set environment variables using the Command Prompt (for the current session) or through the `setx` command or System Properties (for persistent settings).

    - **Using Command Prompt (temporary):**
      Open Command Prompt and run:

      ```cmd
      set iRacing_IDS=123456,234567
      ```

      This setting will only last for the current command prompt session.

    - **Using `setx` (persistent user or system):**
      Open Command Prompt. For user variables:

      ```cmd
      setx iRacing_IDS "123456,234567"
      ```

      For system variables (requires Administrator privileges):

      ```cmd
      setx /m iRacing_IDS "123456,234567"
      ```

      Changes to user variables will apply to future command prompt/PowerShell sessions and applications. System variable changes might require a system restart.

    - **Using System Properties (permanent):**
      1.  Press `Win + R`, type `sysdm.cpl`, and press Enter.
      2.  Go to the "Advanced" tab and click "Environment Variables...".
      3.  In the "User variables" or "System variables" section, click "New...".
      4.  Enter `iRacing_IDS` as the "Variable name" and your comma-separated list of identifiers (e.g., `123456,234567`) as the "Variable value".
      5.  Click "OK" on all dialogs. You might need to restart your terminal or applications for the changes to take effect.

    **Important:** Only files whose names _do not_ contain any of these identifiers will be deleted. Be specific with your identifiers to avoid accidentally deleting files you want to keep.

### Running the Program

1.  **Navigate to the Project Directory:** Open your terminal or command prompt and navigate to the directory where you saved the Go file.

2.  **Run the Go Program:** Execute the following command:

    ```bash
    go run paint_cleaner.go
    ```

    The program will then:

    - Read the `iRacing_IDS` environment variable.
    - Locate your iRacing `paint` folder (usually in `~/Documents/iRacing/paint`).
    - Iterate through the files and subdirectories within the `paint` folder.
    - Delete any files whose names do not contain any of the identifiers specified in the `IDS` variable.
    - Print the number of deleted files and the amount of disk space freed.

## Important Notes

- **No Undo Functionality:** This program permanently deletes files. There is no built-in undo functionality. Trading Paints does download all the missing paints when joining a server.

## Contributing

If you find any issues or have suggestions for improvements, feel free to open a pull request or submit an issue on the repository.
