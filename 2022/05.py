import copy
import itertools


with open('05_input.txt') as fd:
    raw_state, raw_commands = fd.read().split('\n\n')


def parse_row(row):
    for idx in range(0, len(row), 4):
        cell = row[idx : idx + 4]
        yield cell[1] if not cell.isspace() else None


initial_state = tuple(
    [val for val in col if val is not None]
    for col in itertools.zip_longest(
        *(parse_row(ln) for ln in raw_state.split('\n')[-2::-1])
    )
)


def parse_command(line):
    _move, count, _from, source, _to, destination = line.split(' ')
    return int(count), int(source) - 1, int(destination) - 1


commands = [parse_command(line) for line in raw_commands.strip().split('\n')]

stacks = copy.deepcopy(initial_state)
for count, source, dest in commands:
    for _ in range(count):
        stacks[dest].append(stacks[source].pop())

part_one_answer = ''.join(stack[-1] for stack in stacks)
print(f'Answer to part 1: {part_one_answer}')

stacks = list(copy.deepcopy(initial_state))
for count, source, dest in commands:
    stacks[dest].extend(stacks[source][-count:])
    stacks[source] = stacks[source][:-count]

part_two_answer = ''.join(stack[-1] for stack in stacks)
print(f'Answer to part 2: {part_two_answer}')
