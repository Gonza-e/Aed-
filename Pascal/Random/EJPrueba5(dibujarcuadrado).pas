Program Prueba2; 

var
	i, j, lado: integer;
begin

	writeln('Ingrese lado');
	readln(lado);

	for i:= 1 to lado do 
	begin 
		for j:= 1 to lado do
		begin 
			if (i=1) or (j=1) or (i=lado) or (j=lado) then  
			begin 
				write('*');
			end 
			else 
			begin 
				write(' ');
			end;
		end;
		writeln();
	end; 

	readln;
end. 