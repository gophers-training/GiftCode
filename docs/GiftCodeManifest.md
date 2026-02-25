Imagine a system consisting of two services: Wallet and Discount Code.
- Wallet System: Responsible for maintaining user account balances and transaction history.

- Discount System: Responsible for managing discount codes and gift codes.

You are tasked with designing and implementing this system to support the following scenario:
During the halftime of the Iran vs. Brazil national team final, a gift discount code is announced. The first 1,000 users to enter this code will receive a 1,000,000 Toman credit in their wallets.
Requirements & Constraints

1. Real-time Reporting: A report of users who received the gift discount must be available in real-time.
2. Balance Check: Users must be able to view their wallet balance.
3. No Auth: No authentication system is required. Users provide their phone number when entering a code, which serves as their unique User ID.
4. No UI: No dashboard or frontend development is needed. Interactions (entering codes, checking balances, and fetching reports) occur strictly via APIs.

Decoupling: The Wallet and Discount services must be two separate projects/services.

Implementation Plan

- No Auth: Authentication will not be implemented.
- In-Memory Storage: No database will be used initially; data will be kept in memory (to be migrated to a database later if time permits).

- Version Control: Code will be stored on Git.

- Monorepo: Both services will be maintained in two separate directories within a single Git repository.

- Direct Communication: Requests will be sent directly to the relevant service; no intermediary service (like a Gateway) is required.

- Postman: API requests will be saved in Postman for easier testing.

- Documentation: All developed APIs must be documented using Swagger.

- Repo Documentation: All project documentation must be stored within the repository.

- Profile Management: Each discount code must be managed individually as a distinct profile/entity.

- Configurable Values: The credit amount added by a discount code must be definable.

- No Hard-coding: All settings must be adjustable via APIs or Configuration Files.

API Specifications
1. Wallet APIs
 * Create Wallet: Register a new user.
 * List Wallets: Retrieve all wallets.
 * Delete Wallet: Remove a wallet.
 * Update Wallet: Modify wallet details.
 * Get Wallet Info: Retrieve details for a specific wallet.

2. Discount Code APIs

 * Create Discount Code: Define a new code.
 * List Discount Codes: Retrieve all defined codes.
 * Get Discount Info: Retrieve details for a specific code.
 * Delete Discount Code: Remove a code.
 * Update Discount Code: (Not required for now, as codes should generally be immutable).

3. Core Processes (Main Logic)
 * Activate Discount Code: The primary endpoint for users to claim a gift code.
 * Get Code Transactions: A list of all successful applications of a specific code.
 * Get Code Status: Current status/stats of a specific discount code.