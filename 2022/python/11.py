from functools import cached_property, reduce
from typing import Callable, Iterable
import dataclasses
import copy
import math
import operator
import re


@dataclasses.dataclass
class Item:
    worry: int
    starting_worry: int 
    
    @cached_property
    def starting_worry_is_prime(self):
        return math.factorial(x - 1)  % x == x - 1


@dataclasses.dataclass
class Monkey:
    id_: int
    items: list[Item]
    worry_operator: Callable
    worry_operand: int | object 
    worry_divisor: int
    target_id_if_divisible: int
    target_id_if_not_dividible: int
    inspection_count: int = 0 


_operators = {"*": operator.mul, "+": operator.add}
OPERAND_OLD = object()


def parse_monkey(data: str) -> Monkey:
    operation_matches = re.search(r"Operation: new = old ([\*+]) (\w+)", data)
    parsed_operator = _operators[operation_matches.group(1)]
    if (operand := operation_matches.group(2)) == "old":
        parsed_operand = OPERAND_OLD 
    else:
        parsed_operand = int(operand)
    
    return Monkey(
        id_ = int(re.search(r"Monkey ([0-9])+:", data).group(1)),
        items = [
            Item(worry=int(item), starting_worry=(int(item)))
            for item in re.search(r"items: ([0-9 ,]+)", data).group(1).split(", ")
        ],
        worry_operator=parsed_operator,
        worry_operand=parsed_operand,
        worry_divisor=int(re.search(r"divisible by ([0-9]+)", data).group(1)),
        target_id_if_divisible=int(
            re.search(r"If true: throw to monkey ([0-9]+)", data).group(1)
        ),
        target_id_if_not_dividible=int(
            re.search(r"If false: throw to monkey ([0-9]+)", data).group(1)
        )
    )
        

def parse_monkeys(data_path):
    with open(data_path) as fileobj:
        for monkey_data in fileobj.read().split("\n\n"):
            yield parse_monkey(monkey_data)


def solve_monkey_business(monkey_iter: Iterable[Monkey], rounds=20) -> list[Monkey]:
    monkeys = sorted(monkey_iter, key=operator.attrgetter("id_")) 
    lookup = {monkey.id_: monkey for monkey in monkeys}
    for i in range(rounds):
        for monkey in monkeys:
            while monkey.items:
                item = monkey.items.pop()
                if monkey.worry_operand is OPERAND_OLD:
                    item.worry = monkey.worry_operator(item.worry, item.worry)
                else:
                    item.worry = monkey.worry_operator(item.worry,
                                                       monkey.worry_operand)
                item.worry //= 3 

                if item.worry % monkey.worry_divisor == 0:
                    lookup[monkey.target_id_if_divisible].items.append(item)
                else:
                    lookup[monkey.target_id_if_not_dividible].items.append(item)

                monkey.inspection_count += 1
    return monkeys 


def solve_relentless_monkey_business(monkey_iter: Iterable[Monkey], rounds=20) -> list[Monkey]:
    monkeys = sorted(monkey_iter, key=operator.attrgetter("id_")) 
    least_commmon_multiple = reduce(operator.mul, (monkey.worry_divisor for monkey in monkeys))
    lookup = {monkey.id_: monkey for monkey in monkeys}
    for i in range(rounds):
        for monkey in monkeys:
            while monkey.items:
                item = monkey.items.pop()
                if monkey.worry_operand is OPERAND_OLD:
                    item.worry = monkey.worry_operator(item.worry, item.worry)
                else:
                    item.worry = monkey.worry_operator(item.worry,
                                                       monkey.worry_operand)
                item.worry %= least_commmon_multiple 

                if item.worry % monkey.worry_divisor == 0:
                    lookup[monkey.target_id_if_divisible].items.append(item)
                else:
                    lookup[monkey.target_id_if_not_dividible].items.append(item)

                monkey.inspection_count += 1
    return monkeys 

if __name__ == "__main__":
    parsed = parse_monkeys("/home/sc/git/aoc/2022/input/11_input.txt")
    solved = solve_monkey_business(parsed)
    first, second = sorted(monkey.inspection_count for monkey in solved)[-2:]
    print(f"Answer to part one: {first * second}")


    parsed = parse_monkeys("/home/sc/git/aoc/2022/input/11_input.txt")
    solved = solve_relentless_monkey_business(parsed, rounds=10_000)
    first, second = sorted(monkey.inspection_count for monkey in solved)[-2:]
    print(f"Answer to part two: {first * second}")


def test_part_one():
    parsed = parse_monkeys("/home/sc/git/aoc/2022/input/11_test_input.txt")
    solved = solve_monkey_business(parsed)
    
    first, second = sorted(monkey.inspection_count for monkey in solved)[-2:]
    assert first * second == 10605

def test_part_two():
    parsed = parse_monkeys("/home/sc/git/aoc/2022/input/11_test_input.txt")
    solved = solve_relentless_monkey_business(parsed, rounds=10_000)
    
    first, second = sorted(monkey.inspection_count for monkey in solved)[-2:]
    assert first * second == 2713310158
