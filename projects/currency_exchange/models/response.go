package models

/*
// latest rate
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

// analyze
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
*/

type ResponseLatestRate struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type ResponseRateByDate = ResponseLatestRate

type ResponseRateAnalyze struct {
	Base  string               `json:"base"`
	Rates map[string]Analytics `json:"rates"`
}

type Analytics struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}
