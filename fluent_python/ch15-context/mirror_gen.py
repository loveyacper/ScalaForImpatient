import contextlib

def myctxmgr(func, *args, **kwargs): # pack args into tuple, dict
    class Deco:
        def __init__(self, *args, **kwargs):
            self.gen = func(*args, **kwargs) # unpack args from tuple, dict

        def __enter__(self):
            try:
                print("enter!")
                return next(self.gen)
            except StopIteration:
                raise RuntimeError("generator didn't yield") from None;

        def __exit__(self, etype, evalue, traceback):
            print("exit!")
            if etype is None:
                try:
                    next(self.gen)
                except:
                    return
                else:
                    raise RuntimeError("generator didn't stop");
            else:
                if evalue is None:
                    evalue = etype()

                try:
                    print("!throw " + str(evalue))
                    self.gen.throw(etype, evalue, traceback)
                    raise RuntimeError("generator didn't stop after throw");
                except StopIteration as exc:
                    print("!except StopIteration")
                    return exc is not evalue
                except RuntimeError as exc:
                    print("!except RuntimeError")
                    if exc is evalue:
                        return False

                    if exc.__cause__ is evalue:
                        return False

                    raise
                except:
                    import sys
                    if sys.exc_info()[1] is not evalue:
                        raise

        def __call__(self):
            return next(self.gen)

    return Deco

#@contextlib.contextmanager 
@myctxmgr
def looking_glass(name, age):

    import sys
    original_write = sys.stdout.write 

    def reverse_write(text): 
        original_write(text[::-1])

    sys.stdout.write = reverse_write 

    try:
        yield "HELLO:" + str(name) + ", age = " + str(age)
    except:
        msg = 'Please DO NOT divide by zero'
    finally:
        sys.stdout.write = original_write
        if msg is not None:
            print(msg)

