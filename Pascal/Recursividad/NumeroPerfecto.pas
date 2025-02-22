program perfecto;  

uses 
	crt;
var 
	numero: integer;

	function perfecto(num1,num2,suma:integer):boolean;
	begin
		if num2 = 0 then 
		begin 
			if num1 = suma then 
			begin
				perfecto:= true;
			end
			else 
			begin 
				perfecto:= false; 
			end;
		end 
		else 
		begin 
			if (num1 mod num2) = 0 then 
			begin 
				perfecto:= perfecto(num1,num2-1,suma + num2); // Si es divisor se acumula en la variable 
			end 
			else 
			begin
				perfecto:= perfecto(num1,num2-1,suma); // Si no es divisor solo se decrementa el numero
			end;
		end; 
	end; 
begin
	writeln('Ingrese un numero: ');
	readln(numero); 

	clrscr; 

	if perfecto(numero,numero-1,0) then // El numero por el cual vamos a dividir el numero ingresado es el mismo numero -1 (para no contarlo como divisor)
	begin
		writeln('El numero es perfecto');
	end 
	else 
	begin 
		writeln('El numero no es perfecto');
	end;

	readln;
end. 