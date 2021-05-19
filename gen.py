import random
import string


def gen_data():
    accno = ''.join(random.sample(string.ascii_letters, 10))
    currency = 'cny'
    cuno = random.randint(0, 1000000000)
    return "'%s', '%s', %d" % (accno, currency, cuno)

if __name__ == "__main__":
    N = 1000000
    open("prep.csv", "w").write("\n".join([gen_data() for i in range(N)]))

