# Countula

Countula is a powerful and flexible command-line tool written in Go that helps you analyze the codebase in your project directories with a vampiric flair. It scans all the files in a specified folder and counts the lines that contain more than two non-whitespace characters, sinking its teeth into your code to provide a detailed analysis. The results are organized either by file extensions or by subfolders, offering a glimpse into the very vein of your project's structure and content.

## Features

- **Customizable Path**: Specify the path of the folder you want to scan. The default path is the current directory, where Countula lurks by default.
- **Extension Filtering**: Choose specific file extensions to include in the scan, allowing Countula to focus its bite on the files that matter most to you.
- **Exclusion Patterns**: Exclude specific folders or files from the scan, keeping Countula away from the areas you want to protect.
- **Gitignore Support**: Optionally use `.gitignore` to exclude files from the scan, like garlic warding off a vampire from unwanted areas.
- **.git Folder Exclusion**: The `.git` folder is always excluded from the scan, as even Countula avoids meddling with the sanctity of version control files.

## Installation

Before you can use Countula, you need to have [Go](https://golang.org/) installed on your system. Once Go is installed, you can clone the repository and install the tool using the following commands:

```sh
git clone https://github.com/yourusername/countula.git
cd countula
go install
```

This will install Countula globally, allowing you to run it from anywhere like a vampire in the night.

## Usage

To use Countula, you can run it with the necessary flags from the command line, invoking its dark powers to analyze your code. Here is the basic usage syntax:

```sh
countula -path="<path_to_folder>" -extensions="<list_of_extensions>" -excludes="<list_of_exclusions>" -gitignore=<true_or_false>
```

### Flags

- `-path`: Specifies the path of the folder to scan, guiding Countula to its next victim. The default value is the current directory ("."), where it patiently waits in the shadows.

  Example: `-path="./my_project"`

- `-extensions`: A comma-separated list of file extensions to include in the scan, allowing Countula to focus its thirst on specific file types.

  Example: `-extensions=".go,.js,.css"`

- `-excludes`: A comma-separated list of folders or files to exclude from the scan, creating a barrier that even Countula cannot cross.

  Example: `-excludes="temp,backup"`

- `-gitignore`: A boolean flag indicating whether to use the `.gitignore` file to exclude files from the scan, like a protective charm. The default value is `true`.

  Example: `-gitignore=false`

### Examples

Here are some examples of how to use Countula:

1. To scan the current directory and include only `.go` and `.js` files, directing Countula's hunger towards specific file types:

   ```sh
   countula -extensions=".go,.js"
   ```

2. To scan a specific directory and exclude `temp` and `backup` folders, keeping Countula at bay from certain areas:

   ```sh
   countula -path="./my_project" -excludes="temp,backup"
   ```

3. To scan a directory without using the `.gitignore` file, removing the garlic and inviting Countula in:

   ```sh
   countula -path="./my_project" -gitignore=false
   ```

## License

Countula is open-source software licensed under the MIT license. Please see the [LICENSE](LICENSE) file for details, where the terms are laid out in black and white, much like a vampire's contract.

## Support

If you encounter any issues or have questions about Countula, please open an issue on the GitHub repository. We're here to help you, even in the dead of night.

Thank you for using Countula, the tool that brings a bite to code analysis!