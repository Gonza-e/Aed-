Program prueba1; 

var
	niveles, i, j: integer;
begin 
	writeln('Ingrese numero (este sera la cantidad de niveles del triangulo): ');
	readln(niveles);
	writeln();
	
	for i:= 1 to niveles do 
	begin 
		for j:= 1 to i do 
		begin 
			write('*');
		end;
		writeln(); 
	end; 

	readln;
end.