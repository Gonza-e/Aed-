Program a; 

uses 
	math;
var 
	x, y: LongInt; 
begin
	writeln('Ingrese el valor de X: ');
	readln(x); 

	while (x <> 20) do 
	begin
		y:= (x**2) + x;

		writeln('(',x,',',y,')');

		x:= x + 2; 
	end; 

	readln;

end.  