Program EJ17; 

{1) Calcular el IMC (índice de masa corporal) para esto se debe tener en cuenta los
siguientes datos: Altura, Edad y Peso.
El cálculo del IMC es: IMC = Peso / (altura ^2)
Con el Peso en Kg y la altura en metros.
La altura puede ser de 10 cm a 240 cm, la edad puede ser de 4 a 140 años y el Peso puede ser
de 10 kg a 300 kg.
Según el valor que da el calculo podemos tener:
- Para los mayores o iguales a 20 años, el IMC menor a 19 es bajo peso, entre 19 y
24 es normal, entre 25 y 29 es sobrepeso, y más de 30 es obeso.
- Para menores de 20 años el IMC es: menos de 14 es bajo peso, entre entre 14 y 19 es
normal, entre 20 y 24 es sobrepeso y más de 25 es obeso.}

uses 
	math;
var 
	peso, edad: integer;
	altura, IMC: real;

	procedure mayores20();
	begin
		if IMC < 19 then 
		begin
			writeln('Bajo peso');
		end 
		else if (IMC > 19) and (IMC < 24) then 
		begin
			writeln('Normal');
		end 
		else if (IMC > 25) and (IMC < 29) then
		begin 
			writeln('Sobrepeso');
		end 
		else if IMC > 30 then 
		begin 
			writeln('Obeso'); 
		end; 
	end;

	procedure menores20();
	begin
		if IMC < 14 then 
		begin
			writeln('Bajo peso');
		end 
		else if (IMC > 14) and (IMC < 19) then 
		begin
			writeln('Normal');
		end 
		else if (IMC > 20) and (IMC < 24) then
		begin 
			writeln('Sobrepeso');
		end 
		else if IMC > 25 then 
		begin 
			writeln('Obeso'); 
		end; 
	end;
	
begin
	writeln('Ingrese el peso en kg: ');
	readln(peso); 
	writeln('Ingrese la altura en cm: ');
	readln(altura); 
	writeln('Ingrese la edad: ');
	readln(edad); 

	altura:= (altura/100);

	IMC:= (peso/altura**2);

	if edad < 20 then 
	begin 
		menores20();
	end 
	else 
	begin
		mayores20();
	end;

	readln;
end.
