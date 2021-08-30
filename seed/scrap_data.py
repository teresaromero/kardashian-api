from scripts.imbd_episodes import scrap_imbd_all_seasons, imbd_dtypes
from utils.datasets import save_csv

imbd_ep, err = scrap_imbd_all_seasons()
if len(err) != 0:
    print(err)
print(f"Got {len(imbd_ep)} episodes")

err_save = save_csv(imbd_ep, "imbd_episodes", imbd_dtypes)
if err_save:
    print(f"Oopps... can't save the file {err_save}")
