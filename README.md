
# **Tracker**  
A **command-line accounting application** built in **Go** that simplifies **double-entry bookkeeping** principles into an easy-to-use tool for managing **accounts** and **transactions**.

---

## **Features**  

### **Accounts Management**  
- **Create Accounts** with specific types (ASSET, LIABILITY, EQUITY, EXPENSE, REVENUE).  
- **View Account Details**, including code, type, and balance.  
- **Update Account Names** without changing codes.  
- **Deactivate Accounts** to prevent further transactions.  
- Generate a **Chart of Accounts** displaying all active accounts and their details.  

### **Ledger and Transactions**  
- **Create Transactions** using double-entry bookkeeping (debit and credit).  
- View **T-Account Reports** for any account, displaying entries in a **ledger-style table**.  
- **Reverse Transactions** to correct errors while maintaining an **audit trail**.  
- Dynamic **balance calculation** for each account based on transaction history.  

### **CLI Interface**  
- **Interactive Menus** for managing accounts and transactions.  
- **Validation Handling** for inputs to prevent invalid data entry.  
- Built-in **table formatting** for reports using the **`tablewriter`** package.  

---

## **Getting Started**  

### **Prerequisites**  
- **Go 1.20+** installed.  
- **Git** (optional for cloning).  

### **Installation**  

1. Clone the repository:  
   ```bash
   git clone https://github.com/yourusername/tracker-go.git
   cd tracker-go
   ```

2. Install dependencies:  
   ```bash
   go mod tidy
   ```

3. Build the application:  
   ```bash
   go build -o trackergo cmd/main.go
   ```

4. Run the application:  
   ```bash
   go run cmd/main.go
   ```

---

## **Usage**  

### **1. Accounts Menu**  

| Option                      | Description                                               |
|-----------------------------|-----------------------------------------------------------|
| **1. Create Account**        | Adds a new account by specifying name and type.           |
| **2. View Account**          | Displays account details, including code and balance.     |
| **3. Update Account**        | Changes the name of an existing account.                  |
| **4. Chart Of Accounts**     | Lists all accounts along with codes, names, and types.    |
| **5. Deactivate Account**    | Marks an account as inactive, preventing new transactions.|
| **6. Exit**                  | Exits the program.                                        |
| **7. Main Menu**             | Returns to the main menu.                                 |

#### **Example: Create an Account**
```
Choose an option: 1
Enter account name: Rent
Enter account type (ASSET, LIABILITY, EQUITY, EXPENSE, REVENUE): EXPENSE
Account 'RENT - 500' has been created
```

---

### **2. Transactions Menu**  

| Option                      | Description                                               |
|-----------------------------|-----------------------------------------------------------|
| **1. Create Entry**          | Adds a new transaction (debit and credit entries).         |
| **2. View T-Account**        | Displays ledger-style transactions for an account.        |
| **3. Reverse Entry**         | Reverses a transaction while keeping an audit trail.      |
| **4. Exit**                  | Exits the program.                                        |
| **5. Main Menu**             | Returns to the main menu.                                 |

#### **Example: Create a Transaction**
```
Choose an option: 1
Enter an account to debit: RENT
Enter an account to credit: CASH
Enter an amount: 500
Enter a description: Rent Payment
Transactions processed: Debit Entry, Credit Entry
```

---

### **3. T-Account View**  
Displays **ledger entries** for any account in **table format**:

```
====================================================================================================
                                                   CASH                                              
====================================================================================================
DATE         TXN ID                               DESCRIPTION     DEBIT       CREDIT      BALANCE
----------------------------------------------------------------------------------------------------
2024-12-31   78ae33c5-2fc4-4d4d-a923-df7bd834d5f1  Rent Payment    0.00        500.00      -500.00
----------------------------------------------------------------------------------------------------
```

---

## **Key Design Concepts**  

1. **Double-Entry Bookkeeping**:  
   - Each transaction creates a **debit and credit entry**, ensuring balanced records.  

2. **Dynamic Account Codes**:  
   - Account codes are automatically generated based on **account type** (e.g., 100 for ASSET).  

3. **Audit Trails**:  
   - Transactions cannot be deleted—reversals provide a **historical record** for changes.  

4. **Scalable Design**:  
   - Modular services for accounts and ledger, enabling **future extensions** like **multi-currency support** and **databases**.  

---

## **Current Limitations**  

- **No Persistent Storage**:  
  - Data is stored **in-memory**, and all information resets when the program restarts.  
  - Future plans include adding **PostgreSQL or SQLite** for persistence.  

- **No Multi-Currency Support Yet**:  
  - Currently assumes all entries use a **single currency**.  
  - Planned feature to handle **currencies with different decimal scales**.  

- **Limited Error Handling**:  
  - Basic checks for invalid inputs but lacks **advanced validation** for edge cases.  

---

## **Future Improvements**  

1. **Persistence with PostgreSQL/SQLite**:  
   - Migrate from in-memory storage to a **database-backed model**.  

2. **Multi-Currency Support**:  
   - Add currency codes and precision handling for **global usage**.  

3. **Web or GUI Interface**:  
   - Replace the CLI with a **web-based dashboard** using a **React frontend**.  

4. **Reports and Analytics**:  
   - Add more reports like **income statements**, **balance sheets**, and **cash flow summaries**.  

5. **Export Features**:  
   - Allow exporting data to **CSV** or **PDF**.  

---

## **Tech Stack**  

- **Language**: Go  
- **CLI Library**: Built-in Go `fmt` and enhanced with `tablewriter` for tables.  
- **UUID Generation**: Google’s `uuid` package for unique IDs.  
- **Concurrency**: Mutex locks to handle **concurrent data operations** safely.  

---

## **Contributing**  
Feel free to **fork** the repository, submit **issues**, or create **pull requests** if you’d like to contribute to the project’s development.

---

## **License**  
This project is licensed under the **MIT License**—see the [LICENSE](./LICENSE) file for details.
