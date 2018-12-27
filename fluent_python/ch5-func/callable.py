gen = (e for e in (abs, str, 15) if callable(e))
print(list(gen))
