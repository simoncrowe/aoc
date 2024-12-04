import Enum

{:ok, input} = File.read("../input/01-input.txt")
[left_raw, right_raw] = String.trim(input) 
|> String.split("\n") 
|> map(fn l -> String.split(l, "   ") end) 
|> zip() 
|> map(fn t -> Tuple.to_list(t) end)
left = map(left_raw, fn n -> String.to_integer(n) end)
right = map(right_raw, fn n -> String.to_integer(n) end)

difference = zip(sort(left), sort(right)) 
|> map(fn {l, r} -> abs(l-r) end) 
|> sum()
IO.puts("Part 1 result: #{difference}")

similarity = map(left, fn l -> count(filter(right, fn r -> l == r end)) * l end)
|> sum()
IO.puts("Part 2 result #{similarity}")
