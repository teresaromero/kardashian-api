import requests
from bs4 import BeautifulSoup
from datetime import datetime


def get_html(url: str):
    response = requests.get(url)
    if response.status_code == 200:
        return BeautifulSoup(response.text)
    raise Exception(response)


def date_str_to_timestamp(date: str) -> datetime.timestamp:
    date = datetime.strptime(date, '%d %b. %Y')
    return datetime.timestamp(date)
