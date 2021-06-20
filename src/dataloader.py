import torch
from typing import Tuple
from pathlib import Path
from torch.utils.data import DataLoader, Dataset


class Dataset(Dataset):
    def __init__(self, dataset : dict):
        self.dataset = dataset

    def __getitem__(self, idx : int) -> Tuple[torch.tensor, torch.tensor]:
        return

    def __len__(self) -> int:
        return len(self.dataset)

def return_loaders(**kwargs) -> dict:
    def prepare(path, **kwargs) -> DataLoader:
        X = list(path.glob("**/*.txt"))
        return DataLoader(
            Dataset(X),
            drop_last=False,
            **kwargs
        )

    root = Path(__file__).parent.absolute()

    paths = [
        root / ".." / "dataset" / "test",
        root / ".." / "dataset" / "train",
        root / ".." / "dataset" / "val",
    ]

    dataset = {}

    dataset["test"], dataset["train"], dataset["validation"] = map(
        lambda x: prepare(x, **kwargs), map(Path, paths)
    )

    return dataset