import json
from _ctypes import Union

from bs4 import BeautifulSoup
import re

from utils.common import date_str_to_timestamp, get_html


class ImbdEpisode:
    imbd_id: str
    image_url: Union[None, str]
    season: int
    episode: int
    air_date: Union[None, int]
    title: str
    imbd_rate: Union[None, float]
    imbd_rate_votes: Union[None, int]
    description: str

    def to_json(self):
        return json.dumps(self, default=lambda o: o.__dict__)


def scrap_season_imbd_html(html: BeautifulSoup) -> list[ImbdEpisode]:
    data = []
    episode_list = html.select("div.list.detail.eplist > div")
    for episode in episode_list:
        ep = ImbdEpisode()
        ep.imbd_id = episode.select_one("div.image > a > div").get("data-const")

        image_url_raw = episode.select_one("div.image > a > div > img")
        if image_url_raw:
            ep.image_url = image_url_raw.get("src")

        s, e = episode.select_one(
            "div.image > a > div > div").get_text().split(",")
        ep.season = int(s.replace("S", ""))
        ep.episode = int(e.replace("Ep", ""))

        air_date_raw = episode.select_one(
            "div.info > div.airdate")

        if air_date_raw:
            raw_air_date = air_date_raw.get_text().replace("\n", "").strip()
            ep.air_date = date_str_to_timestamp(raw_air_date)

        ep.title = episode.select_one("div.info > strong > a").get("title")

        imbd_rate_raw = episode.select_one(
            "div.info span.ipl-rating-star__rating")
        if imbd_rate_raw:
            ep.imbd_rate = float(imbd_rate_raw.get_text())

        imbd_rate_votes_raw = episode.select_one(
            "div.info span.ipl-rating-star__total-votes")
        if imbd_rate_votes_raw:
            ep.imbd_rate_votes = int(re.sub(
                r"([()])", "", imbd_rate_votes_raw.get_text()))

        description_raw = episode.select_one(
            "div.info > div[itemprop='description']")
        if description_raw:
            ep.description = description_raw.get_text().replace("\n", "")
        data.append(ep)

    return data


def scrap_imbd_all_seasons():
    imbd_base_url = "https://www.imdb.com/title/tt1086761"
    seasons = range(1, 21)
    errors = []
    data = []
    for s in seasons:
        season_url = f"{imbd_base_url}/episodes/season={s}"
        try:
            html = get_html(season_url)
            season_data = scrap_season_imbd_html(html)
            data += season_data
        except Exception as error:
            errors.append(error)
    return data, errors
