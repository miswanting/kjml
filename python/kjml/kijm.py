class KjmlFile:
    def __init__(self, file, mode="r"):
        if mode not in ('r', 'w', 'x', 'a'):
            raise ValueError("KjmlFile requires mode 'r', 'w', 'x', or 'a'")

    def __enter__(self):
        return self

    def __exit__(self, type, value, traceback):
        self.close()

    def close(self):
        pass


if __name__ == "__main__":
    with KjmlFile('sample.kjml') as k:
