from datetime import datetime
import re
from typing import TypedDict

from utils.common import get_html
from unidecode import unidecode

wiki_dtypes = {'title': str,
               'episode': int,
               'season': int,
               'episode_overall': int,
               'air_date': int,
               'us_viewers': float,
               'special_episode': bool}


class WikiEpisode(TypedDict):
    title: str
    episode: int
    season: int
    episode_overall: int
    air_date: int
    us_viewers: float
    special_episode: bool


def parse_date(date: str):
    return int(datetime.fromisoformat(re.sub(r"([()])", "", re.search(r"\((.*?)\)", date).group())).timestamp())


def clean(string: str):
    return re.sub(r"\[[^)]*]", "", string).replace('″', "")


def scrap_wiki_all_episodes() -> list[WikiEpisode]:
    wiki_base_url = "https://en.wikipedia.org/wiki/List_of_Keeping_Up_with_the_Kardashians_episodes"

    html = get_html(wiki_base_url)
    episodes = html.select("tr.vevent")
    data = []
    season_count = 1
    special_count = 1
    for e in episodes:
        ep = WikiEpisode()
        ep['episode_overall'] = int(e.select_one("tr > th:first-child").get_text())
        ep['special_episode'] = False
        ep['season'] = season_count
        episode_data = [i.get_text() for i in e.select("tr > td")]
        ep['title'] = unidecode(episode_data[1]).replace('"', "").replace("\\", "")

        if len(episode_data) < 5:
            ep['episode'] = int(episode_data[0])
            ep['air_date'] = parse_date(unidecode(episode_data[2]))

            if len(episode_data) > 3 and episode_data[3] != "N/A":
                ep['us_viewers'] = float(clean(episode_data[3]))

        elif len(episode_data) == 5:
            ep['episode_overall'] = 0
            ep['special_episode'] = True
            ep['episode'] = special_count
            ep['season'] = int(episode_data[0])
            ep['air_date'] = parse_date(episode_data[3])
            if episode_data[4] != "N/A":
                ep['us_viewers'] = float(clean(episode_data[4]))

            special_count += 1

        if not ep['special_episode'] and ep['episode'] == 1 and ep['episode_overall'] != 1:
            season_count += 1
        data.append(ep)
    return data
