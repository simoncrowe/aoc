from collections.abc import Iterable
from functools import cmp_to_key
from itertools import starmap


def cmp(a, b):
    if a == b: return 0
    return -1 if a > b else 1


def cmp_packets(a, b):
    if isinstance(a, Iterable) and isinstance(b, Iterable):
        for a_val, b_val in zip(a, b):
            if (o := cmp_packets(a_val, b_val)) == 0:
                continue
            return o
        return cmp(len(a), len(b))
    elif not isinstance(a, Iterable) and not isinstance(b, Iterable):
        return cmp(a, b)
    elif isinstance(a, Iterable):
        return cmp_packets(a, [b])
    else:
        return cmp_packets([a], b)


with open('/home/sc/git/aoc/2022/input/13_input.txt') as fd:
    pairs = [tuple(map(eval, raw_pair.strip().split('\n')))
             for raw_pair in fd.read().split('\n\n')]
idx_count = sum(
    i for i, o in enumerate(starmap(cmp_packets, pairs), 1) if o > 0
)
print('Part 1', idx_count)


with open('/home/sc/git/aoc/2022/input/13_input.txt') as fd:
    pairs = [packet
             for raw_pair in fd.read().split('\n\n')
             for packet in map(eval, raw_pair.strip().split('\n'))]
div1, div2 = [[2]], [[6]]
ordered = sorted(
    pairs + [div1, div2], key=cmp_to_key(cmp_packets), reverse=True
)
print('Part 2:', (ordered.index(div1) + 1) * (ordered.index(div2) + 1))
