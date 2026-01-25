import requests
import pandas as pd
import matplotlib.pyplot as plt
import pandas_datareader.data as web

def get_alpha_vantage(symbol: str, api_key: str):
    url = "https://www.alphavantage.co/query"
    params = {
        "function": "TIME_SERIES_DAILY",
        "symbol": symbol,
        "outputsize": "compact",
        "apikey": api_key
    }

    r = requests.get(url, params=params)
    data = r.json()

    if "Time Series (Daily)" not in data:
        print("No se encontró 'Time Series (Daily)' en la respuesta de Alpha Vantage.")
        return pd.DataFrame()

    df = pd.DataFrame(data["Time Series (Daily)"]).T
    df = df.rename(columns={
        "1. open": "Open",
        "2. high": "High",
        "3. low": "Low",
        "4. close": "Close",
        "5. volume": "Volume"
    })

    df = df.astype(float)
    df.index = pd.to_datetime(df.index)
    df = df.sort_index()

    print(f"Alpha Vantage descargó {len(df)} registros para {symbol}")
    return df

def get_stooq(symbol: str):
    try:
        df = web.DataReader(symbol, "stooq")
        df = df.sort_index()
        print(f"Stooq descargó {len(df)} registros para {symbol}")
        return df
    except Exception as e:
        print(f"Error en Stooq: {e}")
        return pd.DataFrame()

SYMBOL = "AAPL"
ALPHA_KEY = "OOK66DTVIRZM7OY3"

alpha_df = get_alpha_vantage(SYMBOL, ALPHA_KEY)
stooq_df = get_stooq(SYMBOL)

plt.figure(figsize=(10, 5))

if not alpha_df.empty:
    plt.plot(alpha_df.index, alpha_df["Close"], label="Alpha Vantage", linewidth=1.5)

if not stooq_df.empty:
    plt.plot(stooq_df.index, stooq_df["Close"], label="Stooq", linewidth=1.2, alpha=0.8)

plt.legend()
plt.title(f"Comparación de fuentes de datos - {SYMBOL}")
plt.xlabel("Fecha")
plt.ylabel("Precio de cierre (USD)")
plt.grid(True, alpha=0.3)
plt.show()
