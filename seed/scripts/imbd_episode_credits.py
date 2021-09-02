from typing import TypedDict

from utils.common import get_html

cast_dtypes = {'imbd_id': str,
               'credit': str,
               'sub_credit': str,
               'name': str,
               }


class CreditImbdEpisode(TypedDict):
    imbd_id: str
    credit: str
    sub_credit: str
    name: str


class TableRawCredit(TypedDict):
    name: str
    sub_credit: str


class RawCredits(TypedDict):
    director: list[str]
    writer: list[str]
    cast: list[str]
    producer: list[TableRawCredit]
    composer: list[TableRawCredit]
    cinematographer: list[TableRawCredit]
    editor: list[TableRawCredit]
    make_up_department: list[TableRawCredit]
    production_manager: list[TableRawCredit]
    sound_department: list[TableRawCredit]
    visual_effects: list[TableRawCredit]
    camera_department: list[TableRawCredit]
    casting_department: list[TableRawCredit]
    editorial_department: list[TableRawCredit]
    music_department: list[TableRawCredit]
    miscellaneous: list[TableRawCredit]


def get_raw(html) -> RawCredits:
    rawCredit = RawCredits()
    rawCredit['director'] = [e.get_text().strip()
                             for e in html.select("h4[name='director'] + table a")]
    rawCredit['writer'] = [e.get_text().strip()
                           for e in html.select("h4[name='writer'] + table a")]
    rawCredit['cast'] = [e.get_text().strip() for e in html.select(
        "h4[name='cast'] + table tr > td:nth-of-type(2) a")]
    tables_names = ['producer',
                    'composer',
                    'cinematographer',
                    'editor',
                    'make_up_department',
                    'production_manager',
                    'sound_department',
                    'visual_effects',
                    'camera_department',
                    'casting_department',
                    'editorial_department',
                    'music_department',
                    'miscellaneous']
    for table in tables_names:
        raw = html.select(f"h4[name='{table}'] + table tr")
        rawCredit[table] = []
        for r in raw:
            name = r.select_one("td.name > a").get_text().strip()
            credit_raw = r.select_one("td.credit")
            if credit_raw:
                sub_credit = credit_raw.get_text().strip()
                rawCredit[table].append({"name": name, "sub_credit": sub_credit})
            else:
                rawCredit[table].append({"name": name})

    return rawCredit


def scrap_imbd_credits_for_episode(imbd_id: str) -> list[CreditImbdEpisode]:
    url = f"https://www.imdb.com/title/{imbd_id}/fullcredits"
    html = get_html(url)
    raw = get_raw(html)

    data = []
    for credit_name in raw:
        for d in raw[credit_name]:
            credit = CreditImbdEpisode()
            credit['imbd_id'] = imbd_id
            credit['credit'] = credit_name
            if type(d) == dict:
                if 'sub_credit' in d.keys():
                    credit['sub_credit'] = d['sub_credit']
                if 'name' in d.keys():
                    credit['name'] = d['name']
            else:
                credit['name'] = d
            data.append(credit)
    return data
