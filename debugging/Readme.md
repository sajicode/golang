# Debugging

<strong>
  Setting the timestamp precision for a logger
</strong>

1. Ldate controls printing the date.<br/>
2. Ltime prints the timestamp.<br/>
3. Lmicrosends adds microsecond precision to the time. This automatically results in the time being printed, even if Ltime isn’t set.</br>
4. LstdFlags turns on both Ldate and Ltime.<br/>


<i>Then a pair of flags deals with the location information:</i><br/>

5. Llongfile shows a full file path and then the line number: /foo/bar/baz.go:123.<br/>
6. Lshortifle shows just the filename and the line number: baz.go:123.<br/>