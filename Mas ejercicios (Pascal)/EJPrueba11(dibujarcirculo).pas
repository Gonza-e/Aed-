
program DibujarCirculo;
uses 
	math, crt;

const
  	radio = 10;  { Define el radio del círculo }
  	escalaX = 2; { Ajusta la escala para compensar la distorsión de la consola }

var
  	x, y: integer;

begin
  	clrscr;
  
	for y := -radio to radio do
	begin
		for x := -radio * escalaX to radio * escalaX do
	    begin
	      if ((x**2) div (escalaX**2) + (y**2) <= (radio**2)) and ((x**2) div (escalaX**2) + (y**2) >= (radio-1)**2) then
	        write('*')
	      else
	        write(' ');
	    end;
	    writeln;
	end;
	
	readln;
end.