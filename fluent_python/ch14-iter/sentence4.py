import re
import reprlib

RE_WORD = re.compile('\w+')

class Sentence:
    def __init__(self, text):
        self.text = text

    def __iter__(self):
        # generator expression
        return (w.group() for w in RE_WORD.finditer(self.text))

        #for w in RE_WORD.finditer(self.text):
            #yield w.group()

    def __repr__(self):
        return 'Sentence(%s)' % reprlib.repr(self.text)

if __name__ == '__main__':
    s = Sentence("hello world, yes i love python3")
    for e in s:
        print(e)

