defmodule Search do
  def all_slices(text_grid) do 
    rows = text_grid |> String.trim() |> String.split("\n")
    {width, height} = {rows |> Enum.at(0) |> String.length, rows |> Enum.count}
    IO.puts("Width: #{width}; height: #{height}")
    flat = rows |> Enum.join("") |> String.graphemes
    
    slices = [
      Search.horizontal_slices(flat, width),
      Search.vertical_slices(flat, width, height),
      Search.diagonal_slices_left(flat, width, height),
      Search.diagonal_slices_right(flat, width, height),
    ] |> Enum.concat()
    reversed = slices

    |> Enum.map(&String.reverse/1)
    
    [slices, reversed] |> Enum.concat
  end
  
  def vertical_slices(flat_grid, width, height) do
    results = 0..height-1
    |> Enum.map(fn column -> 
      Enum.slice(flat_grid, column..(width*height)//width)
      |> Enum.join("")
    end)
  end

  def horizontal_slices(flat_grid, width) do
    slices = flat_grid
    |> Enum.chunk_every(width)
    |> Enum.map(fn slice -> Enum.join(slice, "") end)
  end
  
  def diagonal_slices_left(flat_grid, width, height) do
    vertical_starts = 0..width-1
    horizontal_starts = (width*2)-1..(width*height)-1//height
    slices = Enum.concat([vertical_starts, horizontal_starts])
    |> Enum.map(fn start -> 
      slice = Enum.slice(flat_grid, start..(width*start)//width-1)
      |> Enum.join("")
    end) 
  end
  
  def diagonal_slices_right(flat_grid, width, height) do
    vertical_starts = width-1..0 
    |> Enum.with_index
    horizontal_starts = width..width*(height-1)//height 
    |> Enum.reverse
    |> Enum.with_index
     
    slices = Enum.concat([vertical_starts, horizontal_starts])
    |> Enum.map(fn {start, idx} -> 
      slice = Enum.slice(flat_grid, start..start+(width*(idx+1))//width+1)
      |> Enum.join("")
    end) 
  end
end

{:ok, input} = File.read("../input/04-input.txt")

occurances = Search.all_slices(input) 
|> Enum.map(fn slice -> Regex.scan(~r/XMAS/, slice) |> Enum.concat end) 
|> Enum.concat
|> Enum.count()
IO.inspect(occurances, charlists: :as_lists)
IO.puts("Part 1 result #{occurances}")
