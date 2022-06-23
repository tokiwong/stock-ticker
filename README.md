# stock-ticker

Stock Ticker is a Go webservice that leverages the [Alpha Vantage](https://www.alphavantage.co/) API to return stock price history for a company over a given number of days.

The latest docker image can be fetched with the following command:
```bash
docker pull docker.io/tokiwong/stock-ticker:latest
```

## Getting started

`stock-ticker` can be run locally and/or on Kubernetes

### Running locally
1. Ensure that your current dir is the project root
1. Run `make run`
1. Run `curl localhost:8080/api/daily` in a separate terminal

### Running on Kubernetes
1. Stand up your Kubernetes cluster. If you're using a kind cluster, you can go [here](https://kind.sigs.k8s.io/docs/user/quick-start)
1. Run `kubectl apply -k manifests/base` to deploy `stock-ticker`
1. Run `kubectl port-forward svc/stock-ticker 8080:8080` to expose the `stock-ticker` service
1. Run `curl localhost:8080/api/daily` in a separate terminal

#### Kustomize
Included are Kustomize overlays with different configurations to adjust replica count and the desired stock ticker

1. `kubectl apply -k manifests/overlays/t_stock` will apply configurations to gather data on `NYSE: T`
1. `kubectl apply -k manifests/overlays/f_stock` will apply configurations to gather data on `NYSE: F`

## Output
Running `curl localhost:8080/api/daily` should give you something like this:

```json
{
   "Meta Data": {
      "1. Information": "Daily Prices (open, high, low, close) and Volumes",
      "2. Symbol": "SPY",
      "3. Last Refreshed": "2022-06-22",
      "4. Output Size": "Compact",
      "5. Time Zone": "US/Eastern"
   },
   "Time Series (Daily)": {
      "2022-06-17": {
         "1. open": "365.5100",
         "2. high": "369.3800",
         "3. low": "362.1700",
         "4. close": "365.8600",
         "5. volume": "111113904"
      },
      "2022-06-21": {
         "1. open": "371.8900",
         "2. high": "376.5250",
         "3. low": "371.8100",
         "4. close": "375.0700",
         "5. volume": "76811861"
      },
      "2022-06-22": {
         "1. open": "370.6200",
         "2. high": "378.7200",
         "3. low": "370.1800",
         "4. close": "374.3900",
         "5. volume": "86575924"
      }
   },
   "Output Data": {
      "N Days": 3,
      "Average Close": "371.773336"
   }
}
```

## [Resilience](RESILIENCE.md)
