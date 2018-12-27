def tag(name, *content, cls=None, **attrs):
    print('-------------------')
    print('name: ' + str(name))
    print('content: ' + ','.join(c for c in content))
    print('cls: ' + str(cls))
    print('attrs: ' + str(attrs))

tag('a', 'b')
tag('a', 'z', cls = 'b')
tag('a', 'z', hello = 'world', cls = 'b')
tag('bert', content = 'keyargs')
tag(content = 'keyargs', name = 'young')
