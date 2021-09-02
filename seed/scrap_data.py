from scripts.imbd_episode_credits import scrap_imbd_credits_for_episode, cast_dtypes
from scripts.imbd_episodes import scrap_imbd_all_seasons, imbd_dtypes
from scripts.wiki_episodes import scrap_wiki_all_episodes, wiki_dtypes
from utils.datasets import save_csv
import pandas as pd

err_save = []
imbd_ep, err = scrap_imbd_all_seasons()
if len(err) != 0:
    print(err)
err_save += save_csv(imbd_ep, "imbd_episodes", imbd_dtypes)

df_imbd_ep = pd.read_csv('source/imbd_episodes.csv')
imbd_id_list = df_imbd_ep['imbd_id'].unique()

imbd_cast = []
for e in imbd_id_list:
    data = scrap_imbd_credits_for_episode(e)
    imbd_cast += data
err_save += save_csv(imbd_cast, "imbd_cast", cast_dtypes)

wiki_ep = scrap_wiki_all_episodes()
err_save += save_csv(wiki_ep, "wiki_episodes", wiki_dtypes)

print(err_save)
