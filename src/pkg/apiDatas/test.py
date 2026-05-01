import os
import requests
import pandas as pd
import pandas_datareader.data as web

def get_alpha_vantage(symbol: str, api_key: str) -> pd.DataFrame:
    url = "https://www.alphavantage.co/query"
    params = {
        "function": "TIME_SERIES_DAILY",
        "symbol": symbol,
        "outputsize": "compact",
        "apikey": api_key
    }

    try:
        r = requests.get(url, params=params, timeout=10)
        r.raise_for_status()
        data = r.json()
    except Exception as e:
        print(f"Error al conectar con Alpha Vantage: {e}")
        return pd.DataFrame()

    if "Note" in data or "Error Message" in data:
        print(f"Alpha Vantage error: {data}")
        return pd.DataFrame()

    if "Time Series (Daily)" not in data:
        print("No se encontró 'Time Series (Daily)' en la respuesta.")
        return pd.DataFrame()

    df = pd.DataFrame(data["Time Series (Daily)"]).T
    df = df.rename(columns={
        "1. open": "Open",
        "2. high": "High",
        "3. low": "Low",
        "4. close": "Close",
        "5. volume": "Volume"
    })

    df = df.apply(pd.to_numeric, errors="coerce")
    df.index = pd.to_datetime(df.index)
    df = df.sort_index()

    print(f"Alpha Vantage descargó {len(df)} registros para {symbol}")
    return df


def get_stooq(symbol: str) -> pd.DataFrame:
    try:
        df = web.DataReader(symbol, "stooq")
        df = df.sort_index()
        print(f"Stooq descargó {len(df)} registros para {symbol}")
        return df
    except Exception as e:
        print(f"Error en Stooq: {e}")
        return pd.DataFrame()


SYMBOL = "AAPL"
ALPHA_KEY = os.getenv("OOK66DTVIRZM7OY3")

alpha_df = get_alpha_vantage(SYMBOL, ALPHA_KEY)
stooq_df = get_stooq("AAPL.US")
