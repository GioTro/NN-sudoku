import torch
from typing import Tuple
from pathlib import Path
from torch.utils.data import DataLoader, Dataset
from utils import Utils

class Dataset(Dataset):
    def __init__(self, dataset : iter):
        self.dataset = dataset

    def __getitem__(self, idx :int) -> Tuple[torch.tensor, torch.tensor]:
        X, y = Utils.preprocess(self.dataset[idx])
        return X, y
    
    def __len__(self):
        return len(self.dataset)

def return_loaders(**kwargs) -> dict:
    """ 
        **kwargs
        num_workers, batch_size : int
        shuffle : bool
    """
    def prepare(path, **kwargs) -> DataLoader:
        X = Utils.load_file(path)
        return DataLoader(
            Dataset(X),
            drop_last=False,
            **kwargs
        )
    root = Path(__file__).parent.absolute()

    paths = [
        root / ".." / ".." / "dataset" / "test" / "test_set.txt",
        root /".." / ".." / "dataset" / "train" / "train_set.txt",
        root / ".." / ".." / "dataset" / "val" / "val_set.txt",
    ]

    dataset = {}
    dataset["test"], dataset["train"], dataset["validation"] = map(
        lambda x: prepare(x, **kwargs), map(Path, paths)
    )

    return dataset
