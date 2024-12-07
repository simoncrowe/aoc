defmodule Search do
  def all_slices(flat_grid, width, height) do 
    slices = [
      Search.horizontal_slices(flat_grid, width),
      Search.vertical_slices(flat_grid, width, height),
      Search.diagonal_slices_left(flat_grid, width, height),
      Search.diagonal_slices_right(flat_grid, width, height),
    ] |> Enum.concat()
    reversed = slices
    |> Enum.map(&String.reverse/1)
    [slices, reversed] |> Enum.concat
  end
  
  def all_square_chunks(flat_grid, width, height, chunk_size) do
    chunk_offset = div(chunk_size, 2)
    results = chunk_offset..height-(chunk_offset+1)
    |> Enum.map(fn center_y -> 
      chunk_offset..width-(chunk_offset+1)
      |> Enum.map(fn center_x ->
          0..chunk_size-1
          |> Enum.map(fn chunk_num ->
            start_idx = (((center_y+(chunk_num-1))*width) + center_x)-1
            end_idx = start_idx + (chunk_size-1)
            Enum.slice(flat_grid, start_idx..end_idx)
        end)
      end)
    end)
    |> Enum.concat
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

rows = input |> String.trim() |> String.split("\n")
{width, height} = {rows |> Enum.at(0) |> String.length, rows |> Enum.count}
flat = rows |> Enum.join("") |> String.graphemes

occurances = Search.all_slices(flat, width, height) 
|> Enum.map(fn slice -> Regex.scan(~r/XMAS/, slice) |> Enum.concat end) 
|> Enum.concat
|> Enum.count()
IO.puts("Part 1 result #{occurances}")

cross_count = Search.all_square_chunks(flat, width, height, 3) 
|> Enum.filter(fn square ->
  case square do
    [["M", _n_, "M"],
     [_w_, "A", _e_], 
     ["S", _s_, "S"]] -> true
    [["S", _n_, "M"],
     [_w_, "A", _e_], 
     ["S", _s_, "M"]] -> true
    [["M", _n_, "S"],
     [_w_, "A", _e_], 
     ["M", _s_, "S"]] -> true
    [["S", _n_, "S"],
     [_w_, "A", _e_], 
     ["M", _s_, "M"]] -> true
    _ -> false 
  end
end)
|> Enum.count()
IO.puts("Part 2 result #{cross_count}")
