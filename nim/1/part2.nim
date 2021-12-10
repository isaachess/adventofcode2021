import std/os
import std/strutils

func sum(x: seq[int]): int =
  for i in x:
    result.inc i

var
  slide: seq[int]
  incs = 0
for line in open(paramStr(1)).lines:
  slide.add line.parseInt()
  if slide.len > 4:
    slide.delete(0)
  if slide.len == 4:
    if slide[1..^1].sum > slide[0..2].sum:
      inc incs
echo incs