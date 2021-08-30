import json

import requests
from bs4 import BeautifulSoup
from datetime import datetime


def get_html(url: str):
    response = requests.get(url)
    if response.status_code == 200:
        return BeautifulSoup(response.text, features="lxml")
    raise Exception(response)


def date_str_to_timestamp(date: str) -> int:
    try:
        date = datetime.strptime(date, '%d %b. %Y')
    except ValueError:
        date = datetime.strptime(date, '%d %b %Y')
    finally:
        return int(datetime.timestamp(date))


class Document:
    def to_json(self):
        return json.dumps(self, default=lambda o: o.__dict__)
