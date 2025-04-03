
# Release Notes

## v1.0.0

The first stable release of the IndianDataShop CLI tool. This version introduces core functionality for interacting with the IndianDataShop API, including configuration, credit checking, and search capabilities.

---

### Features

1. **Configuration**  
   - Easily configure the CLI with your API key and display preferences (`TABLE` or `PLAIN`).
   - Configuration is saved locally in the `INDIAN_DATA_SHOP_CONFIG` file.

2. **Check Credits**  
   - View the available credits associated with your API key using the `credits` action.

3. **Search Functionality**  
   - Perform searches using the `search` action with the following parameters:
     - `type`: Specify the type of search (e.g., `email`, `mobile`, `aadhar`).
     - `query`: Provide the search input.
     - `masked`: Optional. Set to `true` for masked results.

4. **Display Options**  
   - Results can be displayed in either a table format (`TABLE`) or a plain text format (`PLAIN`).

5. **Error Handling**  
   - Clear error messages for invalid actions, missing arguments, or configuration issues.

---

### Usage Examples

- Configure the CLI:
  ```bash
  ./indiandata configure
  ```

- Check available credits:
  ```bash
  ./indiandata credits
  ```

- Perform a search:
  ```bash
  ./indiandata search email example@example.com
  ```

---

### Bug Fixes

- Fixed issues with empty arguments causing the CLI to exit unexpectedly.
- Improved error messages for better user experience.

---

### Known Issues

- None reported for this release.

---

### Dependencies

- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) for table formatting.
- Standard Go libraries.

---

### License

This project is licensed under the MIT License.
