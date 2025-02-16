Program Ejercicio10; 

{Hacer un algoritmo que detecte si un numero ingresado es primo o no}

uses 
	math, crt, sysutils;
var 
	numero, raiz, i: integer;
	esprimo: Boolean;  
begin
	esprimo:= True;
	writeln('Ingrese numero: '); 
	readln(numero); 
	raiz:= trunc(sqrt(numero));

	while (esprimo) and (raiz <> i) and (numero <> 1) do 
	begin  
		for i:= 2 to (raiz) do 
		begin
			if ((numero mod i) = 0) then 
			begin 
				esprimo:= False;
			end;
		end; 
	end;

	if numero = 1 then  
	begin 
		esprimo:= False;
	end;

	if esprimo then  
	begin
		writeln('El numero es primo'); 
	end 
	else  
	begin 
		writeln('El numero no es primo');
	end;


	readln;
end. 
