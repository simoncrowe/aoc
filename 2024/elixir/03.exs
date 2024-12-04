import Enum

defmodule Mull do
  def sum_products(input) do
    Regex.scan(~r/mul\((\d+),(\d+)\)/, input, capture: :all_but_first)
    |> filter(fn [a, b] -> String.length(a) <= 3 && String.length(b) <= 3 end)
    |> map(fn [a, b] -> [String.to_integer(a), String.to_integer(b)] end)
    |> map(fn [a, b] -> a*b end)
    |> sum()
  end
  
  def valid_calls(input) do 
    calls = Regex.scan(~r/mul\(\d+,\d+\)|do\(\)|don\'t\(\)/, input)
    |> map(fn [call |_] -> call end)
    dispositions = calls |> with_index 
    |> filter(fn {call, _idx} -> call in ["do()", "don't()"] end)
    |> map(fn {call, idx} ->   
      case call do
          call when call == "do()" -> {idx, true}
          call when call == "don't()" -> {idx, false}
          _ -> :error 
      end
    end)
    |> into(%{})
    dispositions = Map.put(dispositions, 0, true)  # starting disposition is positive
    calls |> with_index |> filter(fn {call, idx} ->
        String.starts_with?(call, "mul") && dispositions[find(idx..-1, fn 
          i -> Map.has_key?(dispositions, i)
        end)]
      end
    ) |> map(fn {call, _idx} -> call end)
  end
end

{:ok, input} = File.read("../input/03-input.txt")

sum_of_products = Mull.sum_products(input)
IO.puts("Part 1 result #{sum_of_products}")

sum_of_valid_products = Mull.valid_calls(input) |> join(" ") |> Mull.sum_products()
IO.puts("Part 2 result #{sum_of_valid_products}")
