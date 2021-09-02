import pandas as pd


def save_csv(data: list[dict], filename: str, dtypes: dict) -> Exception:
    try:
        df = pd.DataFrame(data)
        df.astype(dtypes)
        df.to_csv(f'source/{filename}.csv', index=False)
    except Exception as err:
        return err
