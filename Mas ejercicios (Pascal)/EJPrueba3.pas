Program a; 

var
	edad, edad, anio, anio_ac, mes, mes_ac, dia, dia_ac: integer; 
begin
	writeln('Ingrese anio: ');
	readln(anio); 

	writeln('Ingrese mes: ');
	readln(mes); 

	writeln('Ingrese dia: ');
	readln(dia); 

	writeln('Ingrese anio actual: ');
	readln(anio_ac); 

	writeln('Ingrese mes actual: ');
	readln(mes_ac); 

	writeln('Ingrese dia actual: ');
	readln(dia_ac); 

	edad:= anio_ac - anio; 

	if (mes <= mes_ac) and (dia < dia_ac) then 
	begin 
		edad:= edad - 1; 
	end; 

	readln;
end. 