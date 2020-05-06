from pykeyboard import PyKeyboard

k = PyKeyboard()


def click(key):
    k.press_key(key)


if __name__ == '__main__':
    click('H')
