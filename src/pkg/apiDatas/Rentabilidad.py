import yfinance as yf
import numpy as np
import pandas as pd

symbols = ["AAPL", "MSFT", "SPY"]

data = yf.download(
    symbols,
    start="2020-01-01",
    end="2025-07-31"
)["Close"]

log_returns = np.log(data / data.shift(1))

annual_returns = log_returns.mean() * 252
annual_volatility = log_returns.std() * np.sqrt(252)

comparison = pd.DataFrame({
    "Rentabilidad Anual (%)": annual_returns * 100,
    "Volatilidad Anual (%)": annual_volatility * 100
})

print(comparison)
