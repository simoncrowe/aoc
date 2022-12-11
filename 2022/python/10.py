def x_vals(input_path):
    x = 1
    with open(input_path) as fileobj:
        for line in fileobj:
            match line.strip().split(" "):
                case ["noop"]:
                    yield x
                case ["addx", addend]:
                    yield x
                    yield x
                    x += int(addend)


x = list(x_vals("/home/sc/git/aoc/2022/input/10_input.txt"))
strengths = sum(cycle * x[cycle - 1] for cycle in (20, 60, 100, 140, 180, 220))
print(f"Answer to part one: {strengths}")


def pixels(sprite_positions):
    for idx, sprite_center in enumerate(sprite_positions):
        pixel = idx % 40
        if sprite_center - 1 <= pixel <= sprite_center + 1:
            yield "#"
        else:
            yield "."
        if pixel == 39:
            yield "\n"


print("Answer to part two:", "".join(pixels(x)), sep="\n")


def test_part_one():
    x = list(x_vals("/home/sc/git/aoc/2022/input/10_test_input.txt"))
    
    assert x[19] * 20 == 420
    assert x[59] * 60 == 1140
    assert x[99] * 100 == 1800
    assert x[139] * 140  == 2940
    assert x[179] * 180  == 2880
    assert x[219] * 220  == 3960


def test_part_two():
    expected = """
##..##..##..##..##..##..##..##..##..##..\n
###...###...###...###...###...###...###.\n
####....####....####....####....####....\n
#####.....#####.....#####.....#####.....\n
######......######......######......####\n
#######.......#######.......#######.....\n
""".strip()

    x = x_vals("/home/sc/git/aoc/2022/input/10_test_input.txt")
    
    assert "".joint(draw(x)) == expected
