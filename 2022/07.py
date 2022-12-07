from collections import defaultdict
from os import path


fs = defaultdict(set)
working_dir = list()

with open('07_input.txt') as file_obj:
    for line in file_obj:
        match line.strip().split(' '):
            case ['$', 'cd', '..']:
                working_dir = working_dir[:-1]
            case ['$', 'cd', dir_path]:
                working_dir.append(dir_path)
            case ['$', 'ls']:
                pass
            case ['dir', name]:
                fs[path.join(*working_dir)].add(
                    ('dir', path.join(*working_dir, name))
                )
            case [size, name]:
                fs[path.join(*working_dir)].add((int(size), name))


def total_sizes(file_system: dict, size_pred):
    for root_path in file_system.keys():
        total_size = 0
        dir_paths = [root_path]
        while dir_paths:
            dir_path = dir_paths.pop()
            for item in file_system[dir_path]:
                match item:
                    case ['dir', subdir_path]:
                        dir_paths.append(subdir_path)
                    case [int(size), _]:
                        total_size += size
        if size_pred(total_size):
            yield total_size


total_size_small_files = sum(total_sizes(fs, lambda s: s <= 100_000))
print(f'Part 1 answer: {total_size_small_files}')

FS_TOTAL_SIZE = 70_000_000
FREE_SPACE_REQUIRED = 30_000_000
fs_used = sum(s for d in fs.values() for s, _ in d if isinstance(s, int))
more_space_needed = abs(FS_TOTAL_SIZE - (fs_used + FREE_SPACE_REQUIRED))
deletion_candidate_size = min(
    total_sizes(fs, lambda s: s >= more_space_needed)
)
print(f'Part 2 answer: {deletion_candidate_size}')
