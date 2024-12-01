# card-validator-svc

This is a simple service that validates credit card numbers using the Luhn algorithm.

## How to run

### Using local environment

```bash
KV_VIPER_FILE=./config.yaml go run main.go -- run api
```

### Using Docker

```bash
docker build -t card-validator-svc .
docker run -p 8000:8000 card-validator-svc run api
```

## How to test

Run the following command in the terminal:

```bash
curl -X POST http://localhost:8000/v1/cards \
-H "Content-Type: application/json" \
-d '{
    "card_number": "4111111111111111",
    "expiration_month": "12",
    "expiration_year": "2028"
}'
```

### Valid cards

- 4111111111111111

    ```bash
    curl -X POST http://localhost:8000/v1/cards \
    -H "Content-Type: application/json" \
    -d '{
        "card_number": "4111111111111111",
        "expiration_month": "12",
        "expiration_year": "2028"
    }'
    ```

- 5555555555554444

    ```bash
    curl -X POST http://localhost:8000/v1/cards \
    -H "Content-Type: application/json" \
    -d '{
        "card_number": "5555555555554444",
        "expiration_month": "12",
        "expiration_year": "2028"
    }'
    ```

- 2223000048410010

    ```bash
    curl -X POST http://localhost:8000/v1/cards \
    -H "Content-Type: application/json" \
    -d '{
        "card_number": "2223000048410010",
        "expiration_month": "12",
        "expiration_year": "2028"
    }'
    ```

### Invalid cards

- 4111111111111111

    ```bash
    curl -X POST http://localhost:8000/v1/cards \
    -H "Content-Type: application/json" \
    -d '{
        "card_number": "2223000048410010",
        "expiration_month": "01",
        "expiration_year": "2021"
    }'
    ```

- 1111111111111

    ```bash
    curl -X POST http://localhost:8000/v1/cards \
    -H "Content-Type: application/json" \
    -d '{
        "card_number": "1111111111111",
        "expiration_month": "10",
        "expiration_year": "2028"
    }'
    ```

