# Wallet Manager CRUD APIs

## Call Flow

![Call Flow](./walletCRUD.png)


### 1. Create Wallet
- URL: http://127.0.0.1/wallet
- Method: POST
- Query Param: None
- Headers: 
  * content-type: application/json
- Body:
  - type: json
  - ```json
    {
        "mobile": "09129999999"
    }
    ```
- Response (Success):
  * * Status Code: 201
    * Status Code Name: CREATED
    * Body:
      - type: json
      - ```json
        {
            message : "wallet created",
            data: {
                "mobile": "09129999999",
                "balance": 0.0,
                "created_at": "2026-02-12 19:45:17",
                "updated_at": "2026-02-12 19:45:17"
                }
        }
        ```
- Response (failure)
  * * Status Code: any
    * Status Code Name: any
    * Body:
      - type: json
      - ```json
        {
            "message": "failure message",
            "error": "failure error",
        }
        ```

### 2. List Wallet
- URL: http://127.0.0.1/wallet
- Method: GET
- Query Param: None
- Headers: 
  * content-type: application/json
- Body:None
- Response (Success):
  * * Status Code: 200
    * Status Code Name: OK
    * Body:
      - type: json
      - ```json
        {
           data:
           [ 
                {
                "mobile": "09129999999",
                "balance": 0.0,
                "created_at": "2026-02-12 19:45:17",
                "updated_at": "2026-02-12 19:45:17"
                },
                {
                "mobile": "09128888888",
                "balance": 0.0,
                "created_at": "2026-02-12 19:45:17",
                "updated_at": "2026-02-12 19:45:17"
                },
            ]
        }
        ```
- Response (failure)
  * * Status Code: any
    * Status Code Name: any
    * Body:
      - type: json
      - ```json
        {
            "message": "failure message",
            "error": "failure error",
        }
        ```

### 3. Delete Wallet
- URL: http://127.0.0.1/wallet/:wallet-id
- Method: DELETE
- Query Param: 
  * name: wallet-id
  * type: string
- Headers: 
  * content-type: application/json
- Body: None
- Response (Success):
  * * Status Code: 200
    * Status Code Name: DELETED
    * Body: None
- Response (failure)
  * * Status Code: any
    * Status Code Name: any
    * Body:
      - type: json
      - ```json
        {
            "message": "failure message",
            "error": "failure error",
        }
        ```

### 4. Update Wallet
- URL: http://127.0.0.1/wallet/:wallet-id
- Method: PATCH
- Query Param:
  * Name: wallet-id
  * type: string
- Headers: 
  * content-type: application/json
- Body:
  - type: json
  - ```json
    {
        "amount": 1000.00
        "operation": "add" // valid values: add, subtract, set
    }
    ```
- Response (Success):
  * * Status Code: 202
    * Status Code Name: ACCEPTED
    * Body:
      - type: json
      - ```json
        {
            message : "wallet updated",
            data: {
                "mobile": "09129999999",
                "balance": 0.0,
                "created_at": "2026-02-12 19:45:17",
                "updated_at": "2026-02-12 19:45:17"
                }
        }
        ```
- Response (failure)
  * * Status Code: any
    * Status Code Name: any
    * Body:
      - type: json
      - ```json
        {
            "message": "failure message",
            "error": "failure error",
        }
        ```

### 5. Get Wallet Details
- URL: http://127.0.0.1/wallet/:wallet-id
- Method: GET
- Query Param:
  * name: wallet-id
  * type: string
- Headers: 
  * content-type: application/json
- Body: None
- Response (Success):
  * * Status Code: 200
    * Status Code Name: OK
    * Body:
      - type: json
      - ```json
        {
            message : "wallet listed",
            data: {
                "mobile": "09129999999",
                "balance": 0.0,
                "created_at": "2026-02-12 19:45:17",
                "updated_at": "2026-02-12 19:45:17"
                }
        }
        ```
- Response (failure)
  * * Status Code: any
    * Status Code Name: any
    * Body:
      - type: json
      - ```json
        {
            "message": "failure message",
            "error": "failure error",
        }
        ```
