defmodule Search do
  def scan_multi(text_grid, pattern) do 
    rows = text_grid |> String.trim() |> String.split("\n")
    {width, height} = {rows |> Enum.at(0) |> String.length, rows |> Enum.count}
    IO.puts("Width: #{width}; height: #{height}")
    flat = rows |> Enum.join("") |> String.graphemes
    flat_reversed = flat |> Enum.reverse
    [
      Search.horizontal_scan(flat, pattern, width),
      Search.horizontal_scan(flat_reversed, pattern, width),
      Search.vertical_scan(flat, pattern, width, height),
      Search.vertical_scan(flat_reversed, pattern, width, height),
      Search.diagonal_scan_left(flat, pattern, width, height),
      Search.diagonal_scan_left(flat_reversed, pattern, width, height),
      Search.diagonal_scan_right(flat, pattern, width, height),
      Search.diagonal_scan_right(flat_reversed, pattern, width, height),
    ] |> Enum.concat()
  end
  
  def vertical_scan(flat_grid, pattern, offset, height) do
    results = 0..height-1
    |> Enum.map(fn column -> 
      Enum.slice(flat_grid, column..(offset*height)-(column-1)//height)
      |> Enum.join("")
    end)
    |> Enum.map(fn slice -> Regex.scan(pattern, slice) |> Enum.concat end)
    |> Enum.concat
  end

  def horizontal_scan(flat_grid, pattern, width) do
    results = flat_grid
    |> Enum.chunk_every(width)
    |> Enum.map(fn slice -> Enum.join(slice, "") end)
    |> Enum.map(fn slice -> Regex.scan(pattern, slice) |> Enum.concat end)
    |> Enum.concat
  end
  
  def diagonal_scan_left(flat_grid, pattern, width, height) do
    vertical_starts = 0..width-1
    horizontal_starts = (width*2)-1..(width*height)-1//height
    results = Enum.concat([vertical_starts, horizontal_starts])
    |> Enum.map(fn start -> 
      slice = Enum.slice(flat_grid, start..(width*start)//width-1)
      |> Enum.join("")
    end) 
    |> Enum.map(fn slice -> Regex.scan(pattern, slice) |> Enum.concat end) 
    |> Enum.concat
  end
  
  def diagonal_scan_right(flat_grid, pattern, width, height) do
    vertical_starts = width-1..0 
    |> Enum.with_index
    horizontal_starts = width..width*(height-1)//height 
    |> Enum.reverse
    |> Enum.with_index
    IO.inspect(Enum.concat([vertical_starts, horizontal_starts]), limit: :infinity)
     
    results = Enum.concat([vertical_starts, horizontal_starts])
    |> Enum.map(fn {start, idx} -> 
      slice = Enum.slice(flat_grid, start..start+(width*(idx+1))//width+1)
      |> Enum.join("")
    end) 
    |> Enum.map(fn slice -> Regex.scan(pattern, slice) |> Enum.concat end) 
    |> Enum.concat
  end
end

{:ok, input} = File.read("../input/04-input.txt")

occurances = Search.scan_multi(input, ~r/XMAS/) 
|> Enum.count()
IO.puts("Part 1 result #{occurances}")

