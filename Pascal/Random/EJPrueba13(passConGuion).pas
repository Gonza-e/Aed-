program a;

uses 
	math, crt;  
var 
	i: integer;
	pass, resultado: string;  
begin
	writeln('Ingrese password: ');
	readln(pass);
	resultado:= ' ';

	clrscr;
	for i:= 1 to length(pass) do 
	begin 
		if resultado = ' ' then 
		begin 
			resultado:= pass[i];
		end 
		else 
		begin 
			resultado:= resultado + '-' + pass[i];
		end;  
	end; 

	writeln('La pass es: ',resultado);
	readln;
end. 