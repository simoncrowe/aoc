import Enum

defmodule Reports do 
  def is_safe(levels) do
    [first, second | _] = levels
    cond do
      first < second -> 
        chunk_every(levels, 2, 1, :discard) |> all?(fn [a, b] -> a < b && b-a < 4 end) 
      first > second ->
        chunk_every(levels, 2, 1, :discard) |> all?(fn [a, b] -> a > b && a-b < 4 end) 
      true ->
        false
    end
  end

  def dampened_is_safe(levels) do
    safe = 0..length(levels)
    |> map(fn i -> reject(with_index(levels), fn {_, idx} -> idx == i end) end)
    |> map(fn levels -> map(levels, fn  {level, _} -> level end) end)
    |> map(&Reports.is_safe/1)
    any?(safe) 
  end
end

{:ok, input} = File.read("../input/02-input.txt")
reports = String.trim(input)
|> String.split("\n")
|> map(fn line -> String.split(line, " ") end)
|> map(fn line -> map(line, fn chunk ->  String.to_integer(chunk) end) end)
|> to_list()

safe_count = filter(reports, &Reports.is_safe/1) |> count()
IO.puts("Part 1 solution: #{safe_count}") 
 
dampened_safe_count = filter(reports, &Reports.dampened_is_safe/1) |> count()
IO.puts("Part 2 solution: #{dampened_safe_count}") 
