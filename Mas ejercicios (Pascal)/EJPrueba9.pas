 Program a; 

{36. Todo número cuya suma de sus dígitos sea múltiplo de 3 lo es también.
Ej: 117 → 1+1+7 = 9 que es múltiplo de 3 , entonces 117 es múltiplo de 3
Realizar un algoritmo que determine si un número es múltiplo de 3 en función de la afirmación antes realizada}

uses 
	crt, sysutils;
var 
	num, num_res: integer; 
begin
	writeln('Ingrese el numero:');
	readln(num);

	clrscr;  
	num_res:= 0;
	while num > 0 do 
	begin
		num_res:= num_res + (num mod 10); 
		num:= num div 10;
	end;  

	if (num_res mod 3) = 0 then 
	begin 
		writeln('El numero es multiplo de 3');
	end
	else 
	begin 
		writeln('El numero no es multiplo de 3'); 
	end;

	readln;
end. 