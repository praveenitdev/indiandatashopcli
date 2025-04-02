
# IndianDataShop CLI

A command-line interface (CLI) tool to interact with the IndianDataShop API for searching and managing data.

## Usage

```bash
./indiandata <action> [options]
```

### Actions

1. **configure**  
   Configure the CLI with your API key and display preferences.

   ```bash
   ./indiandata configure
   ```

2. **credits**  
   Check the available credits for your API key.

   ```bash
   ./indiandata credits
   ```

3. **search**  
   Perform a search query.

   ```bash
   ./indiandata search <type> <query> [masked]
   ```

   - `<type>`: The type of search (e.g., `email`, `mobile`, `aadhar`).
   - `<query>`: The search input.
   - `[masked]`: Optional. Set to `true` for masked results.

### Examples

- Configure the CLI:
  ```bash
  ./indiandata configure
  ```

- Check available credits:
  ```bash
  ./indiandata credits
  ```

- Search by email:
  ```bash
  ./indiandata search email example@example.com
  ```

- Search by mobile with masked results:
  ```bash
  ./indiandata search mobile 1234567890 true
  ```

## Configuration

The configuration is stored in a file named `INDIAN_DATA_SHOP_CONFIG`. It includes:
- `APIKey`: Your API key for authentication.
- `DisplayType`: The format for displaying results (`TABLE` or `PLAIN`).

## Dependencies

- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) for table formatting.
- Standard Go libraries.

## License

This project is licensed under the MIT License.
