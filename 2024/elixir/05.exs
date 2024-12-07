defmodule PrintQueue do
  def violations(row, rules) do
    rules
    |> Enum.filter(fn [first, second] ->
      first_idx = row |> Enum.find_index(fn x -> x == first end)
      second_idx = row |> Enum.find_index(fn y -> y == second end)
      case {first_idx, second_idx} do
          {nil, nil} -> false
          {nil, _} -> false
          {_, nil} -> false
          {x, y} when x > y -> true
          {x, y} when x < y -> false
          _ -> nil
      end
    end)
  end
  
  def order(row, rules) do
    relevant_rules = rules
    |> Enum.filter(fn [first, second] -> first in row && second in row end)
    
    graph = :digraph.new
    Enum.each(relevant_rules, fn [first, second] -> 
      :digraph.add_vertex(graph, first)
      :digraph.add_vertex(graph, second)
      :digraph.add_edge(graph, first, second)
    end)

    :digraph_utils.topsort(graph)
  end
end

{:ok, input} = File.read("../input/05-input.txt")

[rules_raw, pages_raw] = String.split(input, "\n\n")

rules = rules_raw
|> String.split("\n")
|> Enum.map(fn row -> row
  |> String.split("|")
  |> Enum.map(fn num -> String.to_integer(num) end)
end)

pages = pages_raw
|> String.trim
|> String.split("\n")
|> Enum.map(fn row -> row 
  |> String.split(",") 
  |> Enum.map(fn num -> String.to_integer(num) end)
end)

correct_count = pages
|> Enum.filter(fn row -> PrintQueue.violations(row, rules) |> Enum.empty? end)
|> Enum.map(fn row -> Enum.at(row, div(Enum.count(row), 2)) end)
|> Enum.sum()

IO.puts("Part 1 result: #{correct_count}")

corrected = pages
|> Enum.filter(fn row -> 
  empty = PrintQueue.violations(row, rules) |> Enum.empty?
  not empty 
end)
|> Enum.map(fn row -> PrintQueue.order(row, rules) end)
IO.inspect(corrected, charlists: :as_lists)

corrected_count = corrected
|> Enum.map(fn row -> Enum.at(row, div(Enum.count(row), 2)) end)
|> Enum.sum()
IO.puts("Part 2 result: #{corrected_count}")
