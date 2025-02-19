Program a; 

{33. Escribir un algoritmo que permita imprimir la siguiente sucesión. Considere que N es un número par, que se ingresa.
2 4 6 8 ... N
4 6 8 10 ... N
6 8 10 12 ... N
....
....
N}

uses 
	crt;
var
	i,j,num: integer; 
begin
	writeln('Ingrese numero par: ');
	readln(num);

	clrscr;
	for i:= 2 to num do 
	begin 
		if (i mod 2) = 0 then  
		begin
			for j:= i to num do 
			begin
				if (j mod 2) = 0 then 
				begin  
					write(j, ' ');
				end;
			end; 
		end;
		writeln();
	end;

	readln;
end. 