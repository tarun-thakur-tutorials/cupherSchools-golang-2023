Code a server using Golang.

## Note

If you have any other questions, please email to us for further clarification.

### Mandatory

1. Please ensure that you follow the best practice of Golang.
2. Please don't use ORM frameworks.
3. Please don't do DDD.
4. Zip your code and send back to us.
   Do not publish your work publicly.
5. Your code should be covered by tests.
6. A README which explain how to run the code.

### Bonus

1. Use only the standard library. It is OK if a library is needed for TDD or database connection.

## Tasks

### Task 1 - Load database

On startup, you server should download the [historical rates](https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml) and load it to a database.

### Task 2 - `GET /rates/latest`

Develop an endpoint that will return the latest exchange rate. Currency should be sorted in ascending order.

#### Sample Response

```json
{
  "base": "EUR",
  "rates": {
    "AUD": 1.5339,
    "BGN": 1.9558,
    "USD": 1.2023,
    "ZAR": 14.8845
  }
}
```

### Task 3 - `GET /rates/YYYY-MM-DD`

Develop an endpoint that will return the exchange rate on specific date. Currency should be sorted in ascending order. Response format is the same as task 2.

### Task 4 - `GET /rates/analyze`

Develop an endpoint that will return the minimum, maximum and average exchange rates based on the data loaded in task 1.

#### Sample Response

```json
{
  "base": "EUR",
  "rates_analyze": {
    "AUD": {
      "min": 1.4994,
      "max": 1.5693,
      "avg": 1.5340524590163933
    },
    "BGN": {
      "min": 1.9558,
      "max": 1.9558,
      "avg": 1.9557999999999973
    },
    "USD": {
      "min": 1.1562,
      "max": 1.2065,
      "avg": 1.1783852459016388
    },
    "ZAR": {
      "min": 14.7325,
      "max": 17.0212,
      "avg": 16.06074426229508
    }
  }
}
```
